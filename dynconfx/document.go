package dynconfx

import (
	"context"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/qor5/kx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Document is one configuration kept as a row of the configs table
// (ConfigRecord): the configuration as JSON with its sensitive fields
// removed, and those fields encrypted together in one ciphertext.
//
// It has no cache of its own. Get reads the row every time, so it returns
// what is committed, on every replica, the moment it is committed; give it a
// manager from NewDecryptCachingManager so that reading an unchanged row does
// not call KMS again.
//
// Everything sensitive in a Document shares one ciphertext, so a single KMS
// call can carry at most 4096 bytes of it, and one bad ciphertext makes the
// whole document unreadable. When the configuration is really one entry per
// something (a provider, a tenant), keep it in a Collection instead.
type Document[T any] struct {
	db                *gorm.DB
	kxManager         *kx.Manager
	key               string
	encryptionContext map[string]string
	defaultValue      *T
}

// DocumentOption configures a Document.
type DocumentOption[T any] func(*Document[T])

// WithDocumentDefault makes Get return value when no row is stored, and
// Update start from value instead of the zero value.
func WithDocumentDefault[T any](value T) DocumentOption[T] {
	return func(d *Document[T]) {
		d.defaultValue = &value
	}
}

// NewDocument returns the Document stored in db under key. T must be
// registered with kxManager's registry, naming its sensitive fields.
func NewDocument[T any](db *gorm.DB, kxManager *kx.Manager, key string, encryptionContext map[string]string, opts ...DocumentOption[T]) *Document[T] {
	d := &Document[T]{db: db, kxManager: kxManager, key: key, encryptionContext: encryptionContext}
	for _, opt := range opts {
		opt(d)
	}
	return d
}

var _ Source[struct{}] = (*Document[struct{}])(nil)

// WithDB returns the same Document reading and writing through db, typically
// a transaction.
func (d *Document[T]) WithDB(db *gorm.DB) *Document[T] {
	clone := *d
	clone.db = db
	return &clone
}

// Get returns the stored configuration: the default when no row is stored
// and one was given, ErrNotFound otherwise.
func (d *Document[T]) Get(ctx context.Context) (T, error) {
	return d.read(ctx, d.db, false)
}

// Update locks the row, applies mutate to the stored configuration (or to the
// default, or the zero value, when none is stored) and writes the result, all
// in one transaction (a savepoint when db is already one). Two Updates of the
// same Document therefore run one after the other, the second seeing the
// first's result. When mutate returns an error nothing is written.
//
// The lock needs a row: the very first two Updates, racing before any row
// exists, both start from the default and the later write wins.
func (d *Document[T]) Update(ctx context.Context, mutate func(T) (T, error)) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		current, err := d.read(ctx, tx, true)
		if err != nil {
			return err
		}
		next, err := mutate(current)
		if err != nil {
			return err
		}
		return d.write(ctx, tx, next)
	})
}

// Save replaces the stored configuration with value.
func (d *Document[T]) Save(ctx context.Context, value T) error {
	return d.write(ctx, d.db, value)
}

// read loads the row, locking it when lock is set. A missing row is the
// default, or ErrNotFound without lock; with lock, the zero value, since an
// Update creates the row.
func (d *Document[T]) read(ctx context.Context, db *gorm.DB, lock bool) (T, error) {
	var zero T
	query := db.WithContext(ctx).Where("key = ?", d.key)
	if lock {
		query = query.Clauses(clause.Locking{Strength: "UPDATE"})
	}
	var record ConfigRecord
	err := query.Take(&record).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if d.defaultValue != nil {
			return *d.defaultValue, nil
		}
		if lock {
			return zero, nil
		}
		return zero, errors.Wrapf(ErrNotFound, "config %q", d.key)
	}
	if err != nil {
		return zero, errors.Wrapf(err, "failed to load config %q", d.key)
	}

	var value T
	if err := json.Unmarshal(record.Data, &value); err != nil {
		return zero, errors.Wrapf(err, "failed to unmarshal config %q", d.key)
	}
	if record.Ciphertext == "" {
		return value, nil
	}
	value, err = kx.DecryptStruct(ctx, d.kxManager, value, record.Ciphertext, d.encryptionContext)
	if err != nil {
		return zero, errors.Wrapf(err, "failed to decrypt config %q", d.key)
	}
	return value, nil
}

func (d *Document[T]) write(ctx context.Context, db *gorm.DB, value T) error {
	encrypted, ciphertext, err := kx.EncryptStruct(ctx, d.kxManager, value, d.encryptionContext)
	if err != nil {
		return errors.Wrapf(err, "failed to encrypt config %q", d.key)
	}
	data, err := json.Marshal(encrypted)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal config %q", d.key)
	}
	record := ConfigRecord{Key: d.key, Data: data, Ciphertext: ciphertext, UpdatedAt: time.Now()}
	err = db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"data", "ciphertext", "updated_at"}),
		}).
		Create(&record).Error
	return errors.Wrapf(err, "failed to save config %q", d.key)
}
