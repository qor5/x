package dynconfx_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"maps"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/qor5/kx"
	"github.com/qor5/kx/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/qor5/x/v3/dynconfx"
	"github.com/qor5/x/v3/gormx"
)

var suite *gormx.TestSuite

func TestMain(m *testing.M) {
	ctx := context.Background()
	suite = gormx.MustStartTestSuite(ctx)
	defer func() { _ = suite.Stop(ctx) }()
	m.Run()
}

// fakeCipher stands in for KMS: the "ciphertext" is the plaintext and the
// encryption context, so Decrypt refuses a different context as KMS does. It
// counts Decrypt calls, and fails the next failNext of them.
type fakeCipher struct {
	decrypts atomic.Int32
	failNext atomic.Int32
}

type fakeEnvelope struct {
	Context   map[string]string `json:"context"`
	Plaintext []byte            `json:"plaintext"`
}

func (f *fakeCipher) NewEncrypter() api.Encrypter { return f }
func (f *fakeCipher) NewDecrypter() api.Decrypter { return f }

func (f *fakeCipher) Encrypt(_ context.Context, plaintext []byte, encryptionContext map[string]string) ([]byte, error) {
	return json.Marshal(fakeEnvelope{Context: encryptionContext, Plaintext: plaintext})
}

func (f *fakeCipher) Decrypt(_ context.Context, ciphertext []byte, encryptionContext map[string]string) ([]byte, error) {
	f.decrypts.Add(1)
	if f.failNext.Load() > 0 {
		f.failNext.Add(-1)
		return nil, errors.New("fake KMS unavailable")
	}
	var env fakeEnvelope
	if err := json.Unmarshal(ciphertext, &env); err != nil {
		return nil, err
	}
	if !maps.Equal(env.Context, encryptionContext) && !(len(env.Context) == 0 && len(encryptionContext) == 0) {
		return nil, errors.New("encryption context mismatch")
	}
	return env.Plaintext, nil
}

type settings struct {
	Name    string            `json:"name"`
	Secret  string            `json:"secret"`
	Counter int               `json:"counter"`
	Tags    map[string]string `json:"tags,omitempty"`
}

func newManagers(t *testing.T) (*fakeCipher, *kx.Manager) {
	t.Helper()
	return newManagersWithCacheSize(t, 0)
}

func newManagersWithCacheSize(t *testing.T, size int) (*fakeCipher, *kx.Manager) {
	t.Helper()
	registry, err := kx.NewRegistry()
	require.NoError(t, err)
	require.NoError(t, registry.RegisterStruct(settings{}, kx.WithRegularField("Secret", false)))
	hashFactory, err := kx.NewHashFactory(base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{7}, 32)))
	require.NoError(t, err)
	cipher := &fakeCipher{}
	plain, err := kx.NewManager(cipher, hashFactory, registry)
	require.NoError(t, err)
	cached, err := dynconfx.NewDecryptCachingManager(plain, size)
	require.NoError(t, err)
	return cipher, cached
}

func resetConfigs(t *testing.T) *gorm.DB {
	t.Helper()
	db := suite.DB()
	require.NoError(t, db.AutoMigrate(&dynconfx.ConfigRecord{}))
	require.NoError(t, db.Exec(`DELETE FROM configs`).Error)
	return db
}

func TestDocument_GetMissing(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)

	_, err := dynconfx.NewDocument[settings](db, m, "app", nil).Get(ctx)
	require.ErrorIs(t, err, dynconfx.ErrNotFound)

	got, err := dynconfx.NewDocument(db, m, "app", nil, dynconfx.WithDocumentDefault(settings{Name: "default"})).Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, "default", got.Name)
}

func TestDocument_StoresSecretEncrypted(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)
	doc := dynconfx.NewDocument[settings](db, m, "app", map[string]string{"app": "test"})

	require.NoError(t, doc.Save(ctx, settings{Name: "a", Secret: "s3cret"}))

	var record dynconfx.ConfigRecord
	require.NoError(t, db.Where("key = ?", "app").Take(&record).Error)
	assert.NotContains(t, string(record.Data), "s3cret")
	assert.NotEmpty(t, record.Ciphertext)

	got, err := doc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, settings{Name: "a", Secret: "s3cret"}, got)
}

// Another process sees a write the moment it commits: there is no per-process
// cache of the configuration to wait out.
func TestDocument_OtherInstanceSeesWriteAtOnce(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m1 := newManagers(t)
	_, m2 := newManagers(t)
	writer := dynconfx.NewDocument[settings](db, m1, "app", nil)
	reader := dynconfx.NewDocument[settings](db, m2, "app", nil)

	require.NoError(t, writer.Save(ctx, settings{Name: "v1", Secret: "one"}))
	got, err := reader.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, "one", got.Secret)

	require.NoError(t, writer.Save(ctx, settings{Name: "v2", Secret: "two"}))
	got, err = reader.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, settings{Name: "v2", Secret: "two"}, got)
}

func TestDocument_RolledBackUpdateLeavesNothing(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)
	doc := dynconfx.NewDocument[settings](db, m, "app", nil)
	require.NoError(t, doc.Save(ctx, settings{Name: "v1", Secret: "one"}))

	errRollback := errors.New("roll back")
	err := db.Transaction(func(tx *gorm.DB) error {
		require.NoError(t, doc.WithDB(tx).Update(ctx, func(s settings) (settings, error) {
			s.Name, s.Secret = "v2", "two"
			return s, nil
		}))
		got, err := doc.WithDB(tx).Get(ctx)
		require.NoError(t, err)
		require.Equal(t, "two", got.Secret, "the transaction sees its own write")
		return errRollback
	})
	require.ErrorIs(t, err, errRollback)

	got, err := doc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, settings{Name: "v1", Secret: "one"}, got)
}

// Update locks the row, so concurrent read-modify-writes all land.
func TestDocument_ConcurrentUpdatesAllLand(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)
	doc := dynconfx.NewDocument[settings](db, m, "app", nil)
	require.NoError(t, doc.Save(ctx, settings{Name: "n"}))

	const writers, each = 4, 10
	var wg sync.WaitGroup
	for range writers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range each {
				assert.NoError(t, doc.Update(ctx, func(s settings) (settings, error) {
					s.Counter++
					return s, nil
				}))
			}
		}()
	}
	wg.Wait()

	got, err := doc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, writers*each, got.Counter)
}

func TestDocument_UpdateStartsFromDefault(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)
	doc := dynconfx.NewDocument(db, m, "app", nil, dynconfx.WithDocumentDefault(settings{Name: "default"}))

	require.NoError(t, doc.Update(ctx, func(s settings) (settings, error) {
		s.Counter = 1
		return s, nil
	}))
	got, err := doc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, settings{Name: "default", Counter: 1}, got)
}

func TestDecryptCache(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	cipher, m := newManagers(t)
	doc := dynconfx.NewDocument[settings](db, m, "app", nil)
	require.NoError(t, doc.Save(ctx, settings{Secret: "one"}))

	t.Run("an unchanged ciphertext is decrypted once", func(t *testing.T) {
		before := cipher.decrypts.Load()
		for range 3 {
			got, err := doc.Get(ctx)
			require.NoError(t, err)
			require.Equal(t, "one", got.Secret)
		}
		assert.Equal(t, before+1, cipher.decrypts.Load())
	})

	t.Run("a new ciphertext is decrypted", func(t *testing.T) {
		require.NoError(t, doc.Save(ctx, settings{Secret: "two"}))
		before := cipher.decrypts.Load()
		got, err := doc.Get(ctx)
		require.NoError(t, err)
		assert.Equal(t, "two", got.Secret)
		assert.Equal(t, before+1, cipher.decrypts.Load())
	})

	t.Run("a failure is not remembered", func(t *testing.T) {
		require.NoError(t, doc.Save(ctx, settings{Secret: "three"}))
		cipher.failNext.Store(1)
		_, err := doc.Get(ctx)
		require.ErrorContains(t, err, "fake KMS unavailable")
		got, err := doc.Get(ctx)
		require.NoError(t, err)
		assert.Equal(t, "three", got.Secret)
	})

	t.Run("the encryption context is part of the key", func(t *testing.T) {
		var record dynconfx.ConfigRecord
		require.NoError(t, db.Where("key = ?", "app").Take(&record).Error)
		require.NoError(t, db.Create(&dynconfx.ConfigRecord{Key: "copy", Data: record.Data, Ciphertext: record.Ciphertext}).Error)
		_, err := dynconfx.NewDocument[settings](db, m, "copy", map[string]string{"other": "context"}).Get(ctx)
		require.ErrorContains(t, err, "encryption context mismatch")
	})
}

func TestDocument_UpdateWithoutDefaultStartsFromZero(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)
	doc := dynconfx.NewDocument[settings](db, m, "app", nil)

	require.NoError(t, doc.Update(ctx, func(s settings) (settings, error) {
		require.Equal(t, settings{}, s)
		s.Counter = 1
		return s, nil
	}))
	got, err := doc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, settings{Counter: 1}, got)
}

// Two Updates racing to create the row both land: the second waits for the
// first's insert instead of starting from the default too.
func TestDocument_ConcurrentFirstUpdatesAllLand(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)
	doc := dynconfx.NewDocument(db, m, "fresh", nil, dynconfx.WithDocumentDefault(settings{Name: "default"}))

	const writers = 8
	start := make(chan struct{})
	var wg sync.WaitGroup
	for range writers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			assert.NoError(t, doc.Update(ctx, func(s settings) (settings, error) {
				// Hold the transaction open so that, without the insert's
				// lock, the other writers read the missing row meanwhile.
				time.Sleep(50 * time.Millisecond)
				s.Counter++
				return s, nil
			}))
		}()
	}
	close(start)
	wg.Wait()

	got, err := doc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, settings{Name: "default", Counter: writers}, got)
}

// Neither Get nor a rolled back Update can change the default through a map
// they were handed.
func TestDocument_DefaultIsNotShared(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)
	doc := dynconfx.NewDocument(db, m, "app", nil, dynconfx.WithDocumentDefault(settings{Tags: map[string]string{"a": "1"}}))

	got, err := doc.Get(ctx)
	require.NoError(t, err)
	got.Tags["from-get"] = "x"

	errRollback := errors.New("roll back")
	err = db.Transaction(func(tx *gorm.DB) error {
		require.NoError(t, doc.WithDB(tx).Update(ctx, func(s settings) (settings, error) {
			s.Tags["from-update"] = "x"
			return s, nil
		}))
		return errRollback
	})
	require.ErrorIs(t, err, errRollback)

	got, err = doc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, map[string]string{"a": "1"}, got.Tags)
}

func TestDocument_UpdateErrorWritesNothing(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)
	doc := dynconfx.NewDocument[settings](db, m, "app", nil)
	require.NoError(t, doc.Save(ctx, settings{Name: "kept"}))

	errMutate := errors.New("refuse")
	err := doc.Update(ctx, func(s settings) (settings, error) {
		s.Name = "changed"
		return s, errMutate
	})
	require.ErrorIs(t, err, errMutate)
	got, err := doc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, "kept", got.Name)
}

func TestDecryptCache_EvictsLeastRecentlyUsed(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	cipher, m := newManagersWithCacheSize(t, 2)
	docs := map[string]*dynconfx.Document[settings]{}
	for _, key := range []string{"a", "b", "c"} {
		docs[key] = dynconfx.NewDocument[settings](db, m, key, nil)
		require.NoError(t, docs[key].Save(ctx, settings{Secret: key}))
	}
	decryptsFor := func(key string) int32 {
		before := cipher.decrypts.Load()
		got, err := docs[key].Get(ctx)
		require.NoError(t, err)
		require.Equal(t, key, got.Secret)
		return cipher.decrypts.Load() - before
	}

	assert.Equal(t, int32(1), decryptsFor("a"))
	assert.Equal(t, int32(1), decryptsFor("b"))
	assert.Equal(t, int32(0), decryptsFor("a"), "a is now the most recently used")
	assert.Equal(t, int32(1), decryptsFor("c"), "c evicts b, the least recently used")
	assert.Equal(t, int32(0), decryptsFor("a"))
	assert.Equal(t, int32(1), decryptsFor("b"), "b was evicted")
}

func TestNewDecryptCachingManager_RejectsNegativeSize(t *testing.T) {
	_, m := newManagers(t)
	_, err := dynconfx.NewDecryptCachingManager(m, -1)
	require.Error(t, err)
}

type tenant struct {
	Code       string `gorm:"primaryKey"`
	Kind       string
	Config     []byte `gorm:"type:jsonb"`
	Ciphertext *string
}

func resetTenants(t *testing.T, codes ...string) *gorm.DB {
	t.Helper()
	db := suite.DB()
	require.NoError(t, db.AutoMigrate(&tenant{}))
	require.NoError(t, db.Exec(`DELETE FROM tenants`).Error)
	for _, code := range codes {
		require.NoError(t, db.Create(&tenant{Code: code, Kind: "a"}).Error)
	}
	return db
}

func newTenantCollection(t *testing.T, db *gorm.DB, m *kx.Manager) *dynconfx.Collection[settings] {
	t.Helper()
	c, err := dynconfx.NewCollection[settings](db, m, dynconfx.CollectionSchema{
		Table:            "tenants",
		KeyColumn:        "code",
		DataColumn:       "config",
		CiphertextColumn: "ciphertext",
		Scope:            func(db *gorm.DB) *gorm.DB { return db.Where("kind = ?", "a") },
	})
	require.NoError(t, err)
	return c
}

func TestNewCollection_RequiresColumns(t *testing.T) {
	_, m := newManagers(t)
	_, err := dynconfx.NewCollection[settings](suite.DB(), m, dynconfx.CollectionSchema{Table: "tenants"})
	require.Error(t, err)
}

func TestCollection_GetMissing(t *testing.T) {
	ctx := context.Background()
	db := resetTenants(t, "draft")
	_, m := newManagers(t)
	c := newTenantCollection(t, db, m)

	_, err := c.Get(ctx, "draft")
	require.ErrorIs(t, err, dynconfx.ErrNotFound, "a row without an entry")
	_, err = c.Get(ctx, "nobody")
	require.ErrorIs(t, err, dynconfx.ErrNotFound, "no row")
	require.ErrorIs(t, c.Put(ctx, "nobody", settings{}), dynconfx.ErrNotFound, "Put never inserts")
}

func TestCollection_PutGetEachRowEncryptedOnItsOwn(t *testing.T) {
	ctx := context.Background()
	db := resetTenants(t, "x", "y")
	_, m := newManagers(t)
	c := newTenantCollection(t, db, m)

	require.NoError(t, c.Put(ctx, "x", settings{Name: "X", Secret: "x-secret"}))
	require.NoError(t, c.Put(ctx, "y", settings{Name: "Y", Secret: "y-secret"}))

	var row tenant
	require.NoError(t, db.Take(&row, "code = ?", "x").Error)
	assert.NotContains(t, string(row.Config), "x-secret")
	require.NotNil(t, row.Ciphertext)

	got, err := c.Get(ctx, "x")
	require.NoError(t, err)
	assert.Equal(t, settings{Name: "X", Secret: "x-secret"}, got)

	// x's ciphertext copied onto y does not decrypt there.
	require.NoError(t, db.Model(&tenant{}).Where("code = ?", "y").Update("ciphertext", *row.Ciphertext).Error)
	_, err = c.Get(ctx, "y")
	require.ErrorContains(t, err, "encryption context mismatch")
}

// An entry that cannot be decrypted affects that entry only: the others still
// Get, and List, which decrypts nothing, still lists them all.
func TestCollection_BadEntryAffectsOnlyItself(t *testing.T) {
	ctx := context.Background()
	db := resetTenants(t, "good", "bad", "draft")
	cipher, m := newManagers(t)
	c := newTenantCollection(t, db, m)
	require.NoError(t, c.Put(ctx, "good", settings{Name: "G", Secret: "g"}))
	require.NoError(t, c.Put(ctx, "bad", settings{Name: "B", Secret: "b"}))
	require.NoError(t, db.Model(&tenant{}).Where("code = ?", "bad").Update("ciphertext", "not base64!").Error)

	_, err := c.Get(ctx, "bad")
	require.Error(t, err)
	got, err := c.Get(ctx, "good")
	require.NoError(t, err)
	assert.Equal(t, "g", got.Secret)

	before := cipher.decrypts.Load()
	entries, err := c.List(ctx)
	require.NoError(t, err)
	assert.Equal(t, before, cipher.decrypts.Load(), "List decrypts nothing")
	require.Len(t, entries, 2, "a row without an entry is not listed")
	for _, e := range entries {
		assert.Empty(t, e.Value.Secret)
	}
	assert.Equal(t, []string{"bad", "good"}, []string{entries[0].Key, entries[1].Key})
}

func TestCollection_ScopeAndClear(t *testing.T) {
	ctx := context.Background()
	db := resetTenants(t, "in")
	require.NoError(t, db.Create(&tenant{Code: "out", Kind: "b"}).Error)
	_, m := newManagers(t)
	c := newTenantCollection(t, db, m)

	require.ErrorIs(t, c.Put(ctx, "out", settings{}), dynconfx.ErrNotFound, "a row outside the scope")

	require.NoError(t, c.Put(ctx, "in", settings{Secret: "s"}))
	require.NoError(t, c.Clear(ctx, "in"))
	_, err := c.Get(ctx, "in")
	require.ErrorIs(t, err, dynconfx.ErrNotFound)
	var row tenant
	require.NoError(t, db.Take(&row, "code = ?", "in").Error)
	assert.Nil(t, row.Config)
	assert.Nil(t, row.Ciphertext)
}

func TestCollection_Update(t *testing.T) {
	ctx := context.Background()
	db := resetTenants(t, "t")
	_, m := newManagers(t)
	c := newTenantCollection(t, db, m)

	t.Run("a row without an entry starts from ErrNotFound", func(t *testing.T) {
		require.NoError(t, c.Update(ctx, "t", func(s settings, err error) (settings, error) {
			require.ErrorIs(t, err, dynconfx.ErrNotFound)
			s.Secret = "first"
			return s, nil
		}))
		got, err := c.Get(ctx, "t")
		require.NoError(t, err)
		assert.Equal(t, "first", got.Secret)
	})

	t.Run("a missing row is ErrNotFound", func(t *testing.T) {
		err := c.Update(ctx, "nobody", func(settings, error) (settings, error) {
			t.Fatal("mutate must not run")
			return settings{}, nil
		})
		require.ErrorIs(t, err, dynconfx.ErrNotFound)
	})

	t.Run("concurrent updates of one entry all land", func(t *testing.T) {
		require.NoError(t, c.Put(ctx, "t", settings{Secret: "first"}))
		const writers, each = 4, 10
		var wg sync.WaitGroup
		for range writers {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for range each {
					assert.NoError(t, c.Update(ctx, "t", func(s settings, err error) (settings, error) {
						s.Counter++
						return s, err
					}))
				}
			}()
		}
		wg.Wait()
		got, err := c.Get(ctx, "t")
		require.NoError(t, err)
		assert.Equal(t, writers*each, got.Counter)
		assert.Equal(t, "first", got.Secret)
	})
}

func TestCollection_UpdateErrorWritesNothing(t *testing.T) {
	ctx := context.Background()
	db := resetTenants(t, "t")
	_, m := newManagers(t)
	c := newTenantCollection(t, db, m)
	require.NoError(t, c.Put(ctx, "t", settings{Name: "kept"}))

	errMutate := errors.New("refuse")
	err := c.Update(ctx, "t", func(s settings, err error) (settings, error) {
		s.Name = "changed"
		return s, errMutate
	})
	require.ErrorIs(t, err, errMutate)
	got, err := c.Get(ctx, "t")
	require.NoError(t, err)
	assert.Equal(t, "kept", got.Name)
}

func TestCollection_WithDBWritesInTheTransaction(t *testing.T) {
	ctx := context.Background()
	db := resetTenants(t, "t")
	_, m := newManagers(t)
	c := newTenantCollection(t, db, m)

	errRollback := errors.New("roll back")
	err := db.Transaction(func(tx *gorm.DB) error {
		require.NoError(t, c.WithDB(tx).Put(ctx, "t", settings{Secret: "inside"}))
		got, err := c.WithDB(tx).Get(ctx, "t")
		require.NoError(t, err)
		require.Equal(t, "inside", got.Secret, "the transaction sees its own write")
		return errRollback
	})
	require.ErrorIs(t, err, errRollback)
	_, err = c.Get(ctx, "t")
	require.ErrorIs(t, err, dynconfx.ErrNotFound, "the rolled back write left nothing")
}

func TestStatic(t *testing.T) {
	src := dynconfx.Static(settings{Name: "fixed", Tags: map[string]string{"a": "1"}})
	got, err := src.Get(context.Background())
	require.NoError(t, err)
	got.Tags["changed"] = "x"
	got, err = src.Get(context.Background())
	require.NoError(t, err)
	assert.Equal(t, settings{Name: "fixed", Tags: map[string]string{"a": "1"}}, got, "each Get returns its own copy")
}

func TestCollection_CustomEncryptionContextAndOrder(t *testing.T) {
	ctx := context.Background()
	db := resetTenants(t, "a", "b")
	_, m := newManagers(t)
	var contexts []map[string]string
	c, err := dynconfx.NewCollection[settings](db, m, dynconfx.CollectionSchema{
		Table:            "tenants",
		KeyColumn:        "code",
		DataColumn:       "config",
		CiphertextColumn: "ciphertext",
		OrderBy:          "code DESC",
		EncryptionContext: func(key string) map[string]string {
			ec := map[string]string{"tenant": key}
			contexts = append(contexts, ec)
			return ec
		},
	})
	require.NoError(t, err)
	require.NoError(t, c.Put(ctx, "a", settings{Secret: "s"}))
	require.NoError(t, c.Put(ctx, "b", settings{Secret: "s"}))

	assert.Contains(t, contexts, map[string]string{"tenant": "a"}, "Put encrypts under the schema's context")
	got, err := c.Get(ctx, "a")
	require.NoError(t, err)
	assert.Equal(t, "s", got.Secret)

	entries, err := c.List(ctx)
	require.NoError(t, err)
	require.Len(t, entries, 2)
	assert.Equal(t, []string{"b", "a"}, []string{entries[0].Key, entries[1].Key}, "List follows OrderBy")
}

// UpdateWithTx must not leave a rolled back value in the cache.
func TestConfigProvider_UpdateWithTxRollbackNotCached(t *testing.T) {
	ctx := context.Background()
	db := resetConfigs(t)
	_, m := newManagers(t)
	store := dynconfx.NewEncryptedConfigStore[settings](db, m, "app", nil)
	require.NoError(t, store.Save(ctx, settings{Name: "v1"}))
	provider := dynconfx.NewConfigProvider[settings](store)
	got, err := provider.Get(ctx)
	require.NoError(t, err)
	require.Equal(t, "v1", got.Name)

	errRollback := errors.New("roll back")
	err = db.Transaction(func(tx *gorm.DB) error {
		require.NoError(t, provider.UpdateWithTx(ctx, tx, func(s settings) settings {
			s.Name = "v2"
			return s
		}))
		return errRollback
	})
	require.ErrorIs(t, err, errRollback)

	got, err = provider.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, "v1", got.Name)
}
