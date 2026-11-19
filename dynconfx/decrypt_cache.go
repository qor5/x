package dynconfx

import (
	"bytes"
	"container/list"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"slices"
	"sync"

	"github.com/pkg/errors"
	"github.com/qor5/kx"
	"github.com/qor5/kx/api"
)

// DefaultDecryptCacheSize is how many plaintexts NewDecryptCachingManager
// keeps when given a size of zero.
const DefaultDecryptCacheSize = 1024

// NewDecryptCachingManager returns a kx.Manager sharing m's encrypter, hash
// factory and registry, whose decrypter remembers the plaintext of each
// ciphertext it has decrypted, keyed by the ciphertext and the encryption
// context together.
//
// This is the only cache Document and Collection need. They read their rows
// on every Get, so what they return is always what is committed (or, inside a
// transaction, what that transaction sees), on every replica; what is
// expensive is the KMS call, and a ciphertext decrypts to the same plaintext
// forever. A changed configuration has a new ciphertext, so nothing has to be
// invalidated, and a rolled back write leaves nothing behind.
//
// Failures are not remembered: a ciphertext that does not decrypt is tried
// again on the next read. size bounds the entries kept, least recently used
// first out; zero means DefaultDecryptCacheSize.
func NewDecryptCachingManager(m *kx.Manager, size int) (*kx.Manager, error) {
	if size < 0 {
		return nil, errors.Errorf("decrypt cache size must not be negative, got %d", size)
	}
	if size == 0 {
		size = DefaultDecryptCacheSize
	}
	factory := &decryptCachingFactory{
		next:  m.GetCipherFactory(),
		cache: newPlaintextCache(size),
	}
	return kx.NewManager(factory, m.GetHashFactory(), m.GetRegistry())
}

type decryptCachingFactory struct {
	next  api.CipherFactory
	cache *plaintextCache
}

func (f *decryptCachingFactory) NewEncrypter() api.Encrypter {
	return f.next.NewEncrypter()
}

func (f *decryptCachingFactory) NewDecrypter() api.Decrypter {
	return &decryptCachingDecrypter{next: f.next.NewDecrypter(), cache: f.cache}
}

type decryptCachingDecrypter struct {
	next  api.Decrypter
	cache *plaintextCache
}

func (d *decryptCachingDecrypter) Decrypt(ctx context.Context, ciphertext []byte, encryptionContext map[string]string) ([]byte, error) {
	key := plaintextCacheKey(ciphertext, encryptionContext)
	if plaintext, ok := d.cache.get(key); ok {
		return plaintext, nil
	}
	plaintext, err := d.next.Decrypt(ctx, ciphertext, encryptionContext)
	if err != nil {
		return nil, err
	}
	d.cache.add(key, plaintext)
	return bytes.Clone(plaintext), nil
}

// plaintextCacheKey hashes the ciphertext and the encryption context, each
// part length-prefixed so that no two different inputs share a key.
func plaintextCacheKey(ciphertext []byte, encryptionContext map[string]string) [sha256.Size]byte {
	h := sha256.New()
	write := func(b []byte) {
		var n [8]byte
		binary.BigEndian.PutUint64(n[:], uint64(len(b)))
		h.Write(n[:])
		h.Write(b)
	}
	write(ciphertext)
	keys := make([]string, 0, len(encryptionContext))
	for k := range encryptionContext {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, k := range keys {
		write([]byte(k))
		write([]byte(encryptionContext[k]))
	}
	var key [sha256.Size]byte
	copy(key[:], h.Sum(nil))
	return key
}

type plaintextCache struct {
	mu      sync.Mutex
	size    int
	order   *list.List
	entries map[[sha256.Size]byte]*list.Element
}

type plaintextEntry struct {
	key       [sha256.Size]byte
	plaintext []byte
}

func newPlaintextCache(size int) *plaintextCache {
	return &plaintextCache{size: size, order: list.New(), entries: map[[sha256.Size]byte]*list.Element{}}
}

// get returns a copy, so that a caller writing into it cannot change what the
// next caller gets.
func (c *plaintextCache) get(key [sha256.Size]byte) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	el, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	c.order.MoveToFront(el)
	return bytes.Clone(el.Value.(*plaintextEntry).plaintext), true
}

func (c *plaintextCache) add(key [sha256.Size]byte, plaintext []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.entries[key]; ok {
		c.order.MoveToFront(el)
		return
	}
	c.entries[key] = c.order.PushFront(&plaintextEntry{key: key, plaintext: bytes.Clone(plaintext)})
	for c.order.Len() > c.size {
		oldest := c.order.Back()
		c.order.Remove(oldest)
		delete(c.entries, oldest.Value.(*plaintextEntry).key)
	}
}
