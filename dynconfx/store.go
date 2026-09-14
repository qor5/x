package dynconfx

import (
	"context"
)

// ConfigStore defines the interface for config persistence.
// Implementations can store configs in databases, files, or other backends.
type ConfigStore[T any] interface {
	// Load loads config from storage.
	// If config doesn't exist, it should return an error.
	Load(ctx context.Context) (T, error)

	// Save saves config to storage.
	// This replaces the entire config.
	Save(ctx context.Context, config T) error

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
	Update(ctx context.Context, updateFunc func(T) T) error
}
