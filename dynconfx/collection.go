package dynconfx

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/qor5/kx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CollectionSchema says where a Collection keeps its entries: one row per
// entry in Table, found by KeyColumn, with the entry as JSON (sensitive fields
// removed) in DataColumn and those fields encrypted in CiphertextColumn.
//
// The rows belong to the application: a Collection reads and writes only
// these two columns of rows that already exist, and never inserts or deletes
// one. A row whose DataColumn is NULL has no entry yet. KeyColumn must be
// unique within Scope. Collection is written for PostgreSQL.
type CollectionSchema struct {
	Table            string
	KeyColumn        string
	DataColumn       string
	CiphertextColumn string

	// Scope narrows the rows that belong to the collection, e.g.
	// `func(db *gorm.DB) *gorm.DB { return db.Where("protocol = ? AND deleted_at IS NULL", "SAML") }`.
	// Optional.
	Scope func(*gorm.DB) *gorm.DB

	// OrderBy orders List, e.g. "created_at, code". Defaults to KeyColumn.
	// Unlike the column names above, which are quoted as identifiers, it is
	// SQL taken as is, so it must come from the application's code.
	OrderBy string

	// EncryptionContext returns the encryption context of the entry under
	// key. It must name the key, so that a ciphertext copied onto another row
	// does not decrypt there. Defaults to the table, the ciphertext column and
	// the key.
	EncryptionContext func(key string) map[string]string
}

// Collection is a configuration with one entry per row of an application's
// table — a provider, a tenant — each entry encrypted on its own.
//
// Like Document it has no cache of its own: every read goes to the rows, and
// a manager from NewDecryptCachingManager spares the KMS call for a
// ciphertext already decrypted. Get reads and decrypts one row, so an entry
// that cannot be read affects that entry only; writes touch one row, so
// writers of different entries never wait for or overwrite each other.
type Collection[T any] struct {
	db        *gorm.DB
	kxManager *kx.Manager
	schema    CollectionSchema
}

// Entry is one entry of a Collection.
type Entry[T any] struct {
	Key   string
	Value T
}

// NewCollection returns the Collection schema describes in db. T must be
// registered with kxManager's registry, naming its sensitive fields.
func NewCollection[T any](db *gorm.DB, kxManager *kx.Manager, schema CollectionSchema) (*Collection[T], error) {
	for name, v := range map[string]string{
		"Table":            schema.Table,
		"KeyColumn":        schema.KeyColumn,
		"DataColumn":       schema.DataColumn,
		"CiphertextColumn": schema.CiphertextColumn,
	} {
		if v == "" {
			return nil, errors.Errorf("CollectionSchema.%s is required", name)
		}
	}
	if schema.EncryptionContext == nil {
		table, column := schema.Table, schema.CiphertextColumn
		schema.EncryptionContext = func(key string) map[string]string {
			return map[string]string{"dynconfx:table": table, "dynconfx:column": column, "dynconfx:key": key}
		}
	}
	return &Collection[T]{db: db, kxManager: kxManager, schema: schema}, nil
}

// WithDB returns the same Collection reading and writing through db,
// typically a transaction.
func (c *Collection[T]) WithDB(db *gorm.DB) *Collection[T] {
	clone := *c
	clone.db = db
	return &clone
}

type collectionRow struct {
	Key        string  `gorm:"column:dynconfx_key"`
	Data       []byte  `gorm:"column:dynconfx_data"`
	Ciphertext *string `gorm:"column:dynconfx_ciphertext"`
}

func (c *Collection[T]) rows(ctx context.Context, db *gorm.DB) *gorm.DB {
	query := db.WithContext(ctx).Table(c.schema.Table)
	if c.schema.Scope != nil {
		query = c.schema.Scope(query)
	}
	return query
}

// selectRows selects the schema's columns under the names collectionRow
// reads, quoted as identifiers.
func (c *Collection[T]) selectRows(query *gorm.DB) *gorm.DB {
	return query.Select("? AS dynconfx_key, ? AS dynconfx_data, ? AS dynconfx_ciphertext",
		clause.Column{Name: c.schema.KeyColumn},
		clause.Column{Name: c.schema.DataColumn},
		clause.Column{Name: c.schema.CiphertextColumn})
}

// keyIs matches the row under key.
func (c *Collection[T]) keyIs(key string) clause.Eq {
	return clause.Eq{Column: clause.Column{Name: c.schema.KeyColumn}, Value: key}
}

// Get returns the entry under key, sensitive fields decrypted, or ErrNotFound
// when there is no such row or it has no entry.
func (c *Collection[T]) Get(ctx context.Context, key string) (T, error) {
	return c.get(ctx, c.db, key)
}

func (c *Collection[T]) get(ctx context.Context, db *gorm.DB, key string) (T, error) {
	var zero T
	var rows []collectionRow
	if err := c.selectRows(c.rows(ctx, db)).Where(c.keyIs(key)).
		Limit(1).Find(&rows).Error; err != nil {
		return zero, errors.Wrapf(err, "failed to load entry %q of %s", key, c.schema.Table)
	}
	if len(rows) == 0 || rows[0].Data == nil {
		return zero, errors.Wrapf(ErrNotFound, "entry %q of %s", key, c.schema.Table)
	}
	return c.decode(ctx, rows[0], true)
}

// List returns every entry, in OrderBy order, without its sensitive fields:
// nothing is decrypted, so it neither calls KMS nor fails on an entry whose
// ciphertext cannot be read. Get an entry to read its sensitive fields.
func (c *Collection[T]) List(ctx context.Context) ([]Entry[T], error) {
	var rows []collectionRow
	query := c.selectRows(c.rows(ctx, c.db)).
		Where(clause.Neq{Column: clause.Column{Name: c.schema.DataColumn}, Value: nil})
	if c.schema.OrderBy != "" {
		query = query.Order(c.schema.OrderBy)
	} else {
		query = query.Order(clause.OrderByColumn{Column: clause.Column{Name: c.schema.KeyColumn}})
	}
	if err := query.Find(&rows).Error; err != nil {
		return nil, errors.Wrapf(err, "failed to list the entries of %s", c.schema.Table)
	}
	entries := make([]Entry[T], 0, len(rows))
	for _, row := range rows {
		value, err := c.decode(ctx, row, false)
		if err != nil {
			return nil, err
		}
		entries = append(entries, Entry[T]{Key: row.Key, Value: value})
	}
	return entries, nil
}

func (c *Collection[T]) decode(ctx context.Context, row collectionRow, decrypt bool) (T, error) {
	var zero, value T
	if err := json.Unmarshal(row.Data, &value); err != nil {
		return zero, errors.Wrapf(err, "failed to unmarshal entry %q of %s", row.Key, c.schema.Table)
	}
	if !decrypt || row.Ciphertext == nil || *row.Ciphertext == "" {
		return value, nil
	}
	value, err := kx.DecryptStruct(ctx, c.kxManager, value, *row.Ciphertext, c.schema.EncryptionContext(row.Key))
	if err != nil {
		return zero, errors.Wrapf(err, "failed to decrypt entry %q of %s", row.Key, c.schema.Table)
	}
	return value, nil
}

// Put stores value as the entry under key, or returns ErrNotFound when there
// is no such row.
func (c *Collection[T]) Put(ctx context.Context, key string, value T) error {
	return c.put(ctx, c.db, key, value)
}

// Update locks the row under key, applies mutate to its entry and stores the
// result, in one transaction (a savepoint when db is already one). An Update
// racing another on the same entry waits for it and sees its result. mutate
// gets the entry, or ErrNotFound when the row has none yet, and may return
// either; when it returns an error nothing is written. A missing row is
// ErrNotFound without calling mutate.
func (c *Collection[T]) Update(ctx context.Context, key string, mutate func(value T, err error) (T, error)) error {
	return c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var exists []string
		if err := c.rows(ctx, tx).
			Where(c.keyIs(key)).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Limit(1).
			Pluck(c.schema.KeyColumn, &exists).Error; err != nil {
			return errors.Wrapf(err, "failed to lock entry %q of %s", key, c.schema.Table)
		}
		if len(exists) == 0 {
			return errors.Wrapf(ErrNotFound, "entry %q of %s", key, c.schema.Table)
		}
		current, err := c.get(ctx, tx, key)
		if err != nil && !errors.Is(err, ErrNotFound) {
			return err
		}
		next, err := mutate(current, err)
		if err != nil {
			return err
		}
		return c.put(ctx, tx, key, next)
	})
}

// Clear removes the entry under key, leaving its row without one. Clearing a
// row that has no entry is not an error; a missing row is ErrNotFound.
func (c *Collection[T]) Clear(ctx context.Context, key string) error {
	return c.write(ctx, c.db, key, map[string]any{
		c.schema.DataColumn:       nil,
		c.schema.CiphertextColumn: nil,
	})
}

func (c *Collection[T]) put(ctx context.Context, db *gorm.DB, key string, value T) error {
	encrypted, ciphertext, err := kx.EncryptStruct(ctx, c.kxManager, value, c.schema.EncryptionContext(key))
	if err != nil {
		return errors.Wrapf(err, "failed to encrypt entry %q of %s", key, c.schema.Table)
	}
	data, err := json.Marshal(encrypted)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal entry %q of %s", key, c.schema.Table)
	}
	return c.write(ctx, db, key, map[string]any{
		c.schema.DataColumn:       data,
		c.schema.CiphertextColumn: ciphertext,
	})
}

func (c *Collection[T]) write(ctx context.Context, db *gorm.DB, key string, columns map[string]any) error {
	result := c.rows(ctx, db).Where(c.keyIs(key)).Updates(columns)
	if result.Error != nil {
		return errors.Wrapf(result.Error, "failed to store entry %q of %s", key, c.schema.Table)
	}
	if result.RowsAffected == 0 {
		return errors.Wrapf(ErrNotFound, "entry %q of %s", key, c.schema.Table)
	}
	return nil
}
