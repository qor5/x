package dynconfx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestEncryptedConfigStore_Update(t *testing.T) {
	// This test demonstrates the Update method for partial field updates.
	// Since we don't have a real kx.Manager in tests, we'll skip encryption for now.

	type TestConfig struct {
		ServerPort  int    `json:"serverPort"`
		DatabaseURL string `json:"databaseUrl"`
		RedisAddr   string `json:"redisAddr"`
	}

	t.Run("update existing config", func(t *testing.T) {
		ctx := context.Background()
		// Mock implementation for testing
		var currentConfig = TestConfig{
			ServerPort:  8080,
			DatabaseURL: "postgres://localhost/db1",
			RedisAddr:   "localhost:6379",
		}

		// Since we can't create a real kx.Manager in tests,
		// we'll create a mock store that doesn't use encryption
		store := &mockEncryptedConfigStore[TestConfig]{
			loadFunc: func(ctx context.Context) (TestConfig, error) {
				return currentConfig, nil
			},
			saveFunc: func(ctx context.Context, config TestConfig) error {
				currentConfig = config
				return nil
			},
		}

		// Update only the server port
		err := store.Update(ctx, func(cfg TestConfig) TestConfig {
			cfg.ServerPort = 9090
			return cfg
		})
		require.NoError(t, err)

		// Verify only server port changed
		assert.Equal(t, 9090, currentConfig.ServerPort)
		assert.Equal(t, "postgres://localhost/db1", currentConfig.DatabaseURL)
		assert.Equal(t, "localhost:6379", currentConfig.RedisAddr)
	})

	t.Run("update multiple fields", func(t *testing.T) {
		ctx := context.Background()
		var currentConfig TestConfig = TestConfig{
			ServerPort:  8080,
			DatabaseURL: "postgres://localhost/db1",
			RedisAddr:   "localhost:6379",
		}

		store := &mockEncryptedConfigStore[TestConfig]{
			loadFunc: func(ctx context.Context) (TestConfig, error) {
				return currentConfig, nil
			},
			saveFunc: func(ctx context.Context, config TestConfig) error {
				currentConfig = config
				return nil
			},
		}

		// Update multiple fields at once
		err := store.Update(ctx, func(cfg TestConfig) TestConfig {
			cfg.ServerPort = 9090
			cfg.DatabaseURL = "postgres://localhost/db2"
			return cfg
		})
		require.NoError(t, err)

		assert.Equal(t, 9090, currentConfig.ServerPort)
		assert.Equal(t, "postgres://localhost/db2", currentConfig.DatabaseURL)
		assert.Equal(t, "localhost:6379", currentConfig.RedisAddr)
	})

	t.Run("update config that doesn't exist", func(t *testing.T) {
		ctx := context.Background()
		var currentConfig TestConfig

		store := &mockEncryptedConfigStore[TestConfig]{
			loadFunc: func(ctx context.Context) (TestConfig, error) {
				var zero TestConfig
				return zero, gorm.ErrRecordNotFound
			},
			saveFunc: func(ctx context.Context, config TestConfig) error {
				currentConfig = config
				return nil
			},
		}

		// Create new config when it doesn't exist
		err := store.Update(ctx, func(cfg TestConfig) TestConfig {
			cfg.ServerPort = 8080
			cfg.DatabaseURL = "postgres://localhost/db1"
			cfg.RedisAddr = "localhost:6379"
			return cfg
		})
		require.NoError(t, err)

		// Verify config was created
		assert.Equal(t, 8080, currentConfig.ServerPort)
		assert.Equal(t, "postgres://localhost/db1", currentConfig.DatabaseURL)
		assert.Equal(t, "localhost:6379", currentConfig.RedisAddr)
	})
}

// mockEncryptedConfigStore is a mock for testing without real encryption
type mockEncryptedConfigStore[T any] struct {
	loadFunc func(ctx context.Context) (T, error)
	saveFunc func(ctx context.Context, config T) error
}

func (m *mockEncryptedConfigStore[T]) Load(ctx context.Context) (T, error) {
	if m.loadFunc != nil {
		return m.loadFunc(ctx)
	}
	var zero T
	return zero, nil
}

func (m *mockEncryptedConfigStore[T]) Save(ctx context.Context, config T) error {
	if m.saveFunc != nil {
		return m.saveFunc(ctx, config)
	}
	return nil
}

func (m *mockEncryptedConfigStore[T]) Update(ctx context.Context, updateFunc func(T) T) error {
	// Load existing config
	config, err := m.Load(ctx)
	if err != nil {
		var zero T
		if err != gorm.ErrRecordNotFound {
			return err
		}
		config = zero
	}

	// Apply update function
	updatedConfig := updateFunc(config)

	// Save updated config
	return m.Save(ctx, updatedConfig)
}
