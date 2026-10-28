package dynconfx

import (
	"context"

	"github.com/pkg/errors"
)

// StaticStore implements ConfigStore for static (read-only) configuration.
// It wraps a pre-loaded config value and returns it on every Load call.
// Save and Update operations return an error since static config is immutable.
//
// This is useful when you want to use ConfigProvider's unified interface
// but the config is actually loaded once at startup (e.g., from file or env).
//
// Example:
//
//	// Load config from file at startup
//	cfg, err := confx.Initialize(defaultConfig)
//	config, err := cfg(ctx, "config.yaml")
//
//	// Wrap in StaticStore for use with ConfigProvider
//	store := dynamic.NewStaticStore(config)
//	dc := dynamic.NewConfigProvider(store)
//
//	// Now dc.Get() always returns the same config
//	// This allows code to use ConfigProvider interface uniformly
type StaticStore[T any] struct {
	config T
}

// NewStaticStore creates a new StaticStore with the given config value.
func NewStaticStore[T any](config T) *StaticStore[T] {
	return &StaticStore[T]{
		config: config,
	}
}

// Load returns the static config value.
func (s *StaticStore[T]) Load(ctx context.Context) (T, error) {
	return s.config, nil
}

// Save returns an error because static config is read-only.
func (s *StaticStore[T]) Save(ctx context.Context, config T) error {
	return errors.New("static config store is read-only")
}

// Update returns an error because static config is read-only.
func (s *StaticStore[T]) Update(ctx context.Context, updateFunc func(T) T) error {
	return errors.New("static config store is read-only")
}
