package dynconfx

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"github.com/qor5/kx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// EncryptedConfigStore implements ConfigStore with encryption support using kx library.
// It stores configs in a database with sensitive fields encrypted.
type EncryptedConfigStore[T any] struct {
	db                *gorm.DB
	kxManager         *kx.Manager
	configKey         string
	encryptionContext map[string]string
}

// ConfigRecord is the database model for storing encrypted configs.
type ConfigRecord struct {
	Key        string    `gorm:"primaryKey"`
	Data       []byte    `gorm:"type:jsonb"` // Config with sensitive fields removed
	Ciphertext string    `gorm:"type:text"`  // Encrypted sensitive fields
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

// TableName specifies the table name for ConfigRecord.
func (cr *ConfigRecord) TableName() string {
	return "configs"
}

// Scan implements sql.Scanner interface for ConfigRecord.
func (cr *ConfigRecord) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	return json.Unmarshal(bytes, cr)
}

// Value implements driver.Valuer interface for ConfigRecord.
func (cr ConfigRecord) Value() (driver.Value, error) {
	return json.Marshal(cr)
}

// NewEncryptedConfigStore creates a new EncryptedConfigStore instance.
//
// Parameters:
//   - db: GORM database connection
//   - kxManager: kx.Manager instance for encryption/decryption
//   - configKey: Unique identifier for this config (used as primary key in DB)
//   - encryptionContext: Additional context for encryption (e.g., {"app": "myapp"})
//
// Note: You need to register your config struct with kx.Registry before using this store.
// Example:
//
//	import "github.com/qor5/kx"
//
//	registry, _ := kx.NewRegistry()
//	registry.MustRegisterStruct(&AppConfig{},
//	    kx.WithRegularField("OAuth2ClientSecret", false),
//	)
//
//	kxManager, _ := kx.NewManagerByConfig(&kx.Config{
//	    KMSKeyID: "arn:aws:kms:...",
//	    HashKey:  "base64-encoded-key",
//	}, registry)
//
//	store := dynamic.NewEncryptedConfigStore[*AppConfig](
//	    db,
//	    kxManager,
//	    "app_config",
//	    map[string]string{"app": "myapp"},
//	)
func NewEncryptedConfigStore[T any](
	db *gorm.DB,
	kxManager *kx.Manager,
	configKey string,
	encryptionContext map[string]string,
) *EncryptedConfigStore[T] {
	return &EncryptedConfigStore[T]{
		db:                db,
		kxManager:         kxManager,
		configKey:         configKey,
		encryptionContext: encryptionContext,
	}
}

// Load loads and decrypts config from the database.
func (s *EncryptedConfigStore[T]) Load(ctx context.Context) (T, error) {
	var zero T

	// Query database for config record
	var record ConfigRecord
	result := s.db.WithContext(ctx).Where("key = ?", s.configKey).First(&record)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return zero, errors.Wrapf(result.Error, "config not found for key: %s", s.configKey)
		}
		return zero, errors.Wrap(result.Error, "failed to load config from database")
	}

	// Unmarshal data to config struct
	var config T
	if err := json.Unmarshal(record.Data, &config); err != nil {
		return zero, errors.Wrap(err, "failed to unmarshal config data")
	}

	// Decrypt sensitive fields if ciphertext exists
	if record.Ciphertext != "" {
		decrypted, err := kx.DecryptStruct(ctx, s.kxManager, config, record.Ciphertext, s.encryptionContext)
		if err != nil {
			return zero, errors.Wrap(err, "failed to decrypt config")
		}
		config = decrypted
	}

	return config, nil
}

// Save encrypts and saves config to the database.
// This replaces the entire config.
func (s *EncryptedConfigStore[T]) Save(ctx context.Context, config T) error {
	// Encrypt sensitive fields
	encryptedObj, ciphertext, err := kx.EncryptStruct(ctx, s.kxManager, config, s.encryptionContext)
	if err != nil {
		return errors.Wrap(err, "failed to encrypt config")
	}

	// Marshal encrypted config to JSON
	data, err := json.Marshal(encryptedObj)
	if err != nil {
		return errors.Wrap(err, "failed to marshal config to JSON")
	}

	// Prepare database record
	record := ConfigRecord{
		Key:        s.configKey,
		Data:       data,
		Ciphertext: ciphertext,
	}

	// Save or update record
	result := s.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "key"}},
			DoUpdates: clause.AssignmentColumns([]string{"data", "ciphertext", "updated_at"}),
		}).
		Create(&record)

	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to save config to database")
	}

	return nil
}

// WithTx returns a new EncryptedConfigStore that uses the given transaction.
// This allows the store to participate in an external transaction.
// Returns ConfigStore[T] to satisfy the TxUpdater interface.
//
// Example:
//
//	err := db.Transaction(func(tx *gorm.DB) error {
//	    txStore := store.WithTx(tx)
//	    if err := txStore.Update(ctx, updateFunc); err != nil {
//	        return err
//	    }
//	    // Other operations in the same transaction...
//	    return nil
//	})
func (s *EncryptedConfigStore[T]) WithTx(tx *gorm.DB) ConfigStore[T] {
	return &EncryptedConfigStore[T]{
		db:                tx,
		kxManager:         s.kxManager,
		configKey:         s.configKey,
		encryptionContext: s.encryptionContext,
	}
}

// Update applies a partial update to the config.
// The update function receives the current config and returns the updated config.
// This is useful for updating only specific fields without replacing the entire config.
//
// Example:
//
//	// Update only the server port
//	err := store.Update(ctx, func(cfg *AppConfig) *AppConfig {
//	    cfg.ServerPort = 9090
//	    return cfg
//	})
func (s *EncryptedConfigStore[T]) Update(ctx context.Context, updateFunc func(T) T) error {
	// Load existing config
	config, err := s.Load(ctx)
	if err != nil {
		// If config not found, start with zero value
		var zero T
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Wrap(err, "failed to load config for update")
		}
		config = zero
	}

	// Apply update function
	updatedConfig := updateFunc(config)

	// Save updated config
	return s.Save(ctx, updatedConfig)
}
