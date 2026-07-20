package dynconfx

import (
	"context"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/pkg/errors"
	"golang.org/x/sync/singleflight"
)

// ConfigProvider manages configuration with automatic caching and TTL support.
// It provides thread-safe access to configuration values that can be reloaded at runtime.
type ConfigProvider[T any] struct {
	store        ConfigStore[T]
	cache        T
	cacheSet     bool
	ttl          time.Duration
	expiresAt    time.Time
	mu           sync.RWMutex
	validator    Validator
	defaultValue *T
	sf           singleflight.Group
}

// ConfigProviderOption is a function that configures a ConfigProvider.
type ConfigProviderOption[T any] func(*ConfigProvider[T])

// Validator defines the interface for validating configuration structs.
type Validator interface {
	StructCtx(ctx context.Context, v any) error
}

// WithTTL sets cache TTL. If 0, no caching (always query DB).
func WithTTL[T any](ttl time.Duration) ConfigProviderOption[T] {
	return func(cp *ConfigProvider[T]) {
		cp.ttl = ttl
	}
}

// WithValidator sets config validator.
func WithValidator[T any](v Validator) ConfigProviderOption[T] {
	return func(cp *ConfigProvider[T]) {
		cp.validator = v
	}
}

// WithDefault sets a default value to return when config is not found.
func WithDefault[T any](defaultValue T) ConfigProviderOption[T] {
	return func(cp *ConfigProvider[T]) {
		cp.defaultValue = &defaultValue
	}
}

// NewConfigProvider creates a new ConfigProvider instance with the provided store and options.
func NewConfigProvider[T any](store ConfigStore[T], opts ...ConfigProviderOption[T]) *ConfigProvider[T] {
	cp := &ConfigProvider[T]{
		store: store,
		ttl:   5 * time.Minute, // Default TTL
	}
	for _, opt := range opts {
		opt(cp)
	}
	return cp
}

// NewStaticConfigProvider creates a ConfigProvider that always returns the given config value.
// This is a convenience function that wraps NewStaticStore + NewConfigProvider.
// Useful for testing or when you need a temporary config that doesn't change.
//
// Example:
//
//	tempCfg := oidc.Config{...}
//	tempProvider := dynamic.NewStaticConfigProvider(tempCfg)
//	clonedSvc := oidcSvc.Clone(tempProvider)
func NewStaticConfigProvider[T any](config T) *ConfigProvider[T] {
	return NewConfigProvider(NewStaticStore(config), WithTTL[T](0))
}

// Get returns the config. Returns cached value if valid, otherwise loads from storage.
// It uses singleflight to prevent thundering herd when cache expires.
func (c *ConfigProvider[T]) Get(ctx context.Context) (T, error) {
	// Try to get from cache first
	if val, ok := c.getCached(ctx); ok {
		return val, nil
	}

	// Cache miss or expired, load from storage
	return c.load(ctx)
}

// Reload forces a reload from storage, updating the cache.
func (c *ConfigProvider[T]) Reload(ctx context.Context) (T, error) {
	var zero T

	// Capture the current cached value before invalidating, so a failed
	// reload can still fall back to it. getStaleCache gates on cacheSet,
	// which Invalidate clears, so the stale value must be read first.
	staleVal, hasStale := c.getStaleCache()

	// Invalidate cache first
	c.Invalidate()

	// Load directly from storage without singleflight
	config, err := c.store.Load(ctx)
	if err != nil {
		// If config not found and we have a default value, return it
		if c.defaultValue != nil && isNotFoundError(err) {
			return *c.defaultValue, nil
		}
		// On error, return stale cache if available
		if hasStale {
			return staleVal, nil
		}
		return zero, err
	}

	// Validate config if validator is set
	if c.validator != nil {
		if err := c.validator.StructCtx(ctx, config); err != nil {
			return zero, errors.Wrap(err, "validation failed for config")
		}
	}

	// Update cache
	c.mu.Lock()
	c.cache = config
	c.cacheSet = true
	if c.ttl > 0 {
		c.expiresAt = time.Now().Add(c.ttl)
	}
	c.mu.Unlock()

	return config, nil
}

// Invalidate marks the cache as invalid. Next Get will reload from storage.
func (c *ConfigProvider[T]) Invalidate() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheSet = false
}

// Update applies a partial update to the config via the underlying store,
// then updates the cache with the new value.
func (c *ConfigProvider[T]) Update(ctx context.Context, updateFunc func(T) T) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.updateWithStore(ctx, c.store, updateFunc)
}

// TxUpdater is an interface for stores that support transactions.
type TxUpdater[T any] interface {
	WithTx(tx *gorm.DB) ConfigStore[T]
}

// UpdateWithTx applies a partial update to the config using the given transaction.
// The underlying store must implement TxUpdater interface (e.g., EncryptedConfigStore).
// This allows the config update to participate in an external transaction.
//
// Example:
//
//	err := db.Transaction(func(tx *gorm.DB) error {
//	    if err := provider.UpdateWithTx(ctx, tx, updateFunc); err != nil {
//	        return err
//	    }
//	    // Other operations in the same transaction...
//	    return nil
//	})
func (c *ConfigProvider[T]) UpdateWithTx(ctx context.Context, tx *gorm.DB, updateFunc func(T) T) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if store supports transactions
	txStore, ok := c.store.(TxUpdater[T])
	if !ok {
		return errors.New("underlying store does not support transactions")
	}

	return c.updateWithStore(ctx, txStore.WithTx(tx), updateFunc)
}

// updateWithStore is the common implementation for Update and UpdateWithTx.
// Caller must hold c.mu lock.
func (c *ConfigProvider[T]) updateWithStore(ctx context.Context, store ConfigStore[T], updateFunc func(T) T) error {
	if err := store.Update(ctx, updateFunc); err != nil {
		return err
	}

	// Reload from store to get the updated config
	config, err := store.Load(ctx)
	if err != nil {
		// Invalidate cache on error so next Get will reload
		c.cacheSet = false
		return nil
	}

	// Update cache
	c.cache = config
	c.cacheSet = true
	if c.ttl > 0 {
		c.expiresAt = time.Now().Add(c.ttl)
	}

	return nil
}

// getCached returns cached config if it's valid (set and not expired).
// Returns the cached value and true if cache is valid.
func (c *ConfigProvider[T]) getCached(ctx context.Context) (T, bool) {
	c.mu.RLock()
	valid := c.cacheSet && (c.ttl == 0 || time.Now().Before(c.expiresAt))
	var zero T
	if valid {
		cached := c.cache
		c.mu.RUnlock()
		return cached, true
	}
	c.mu.RUnlock()
	return zero, false
}

// load loads config from storage, validates it, and updates cache.
// Uses singleflight to prevent concurrent loads when caching is enabled.
func (c *ConfigProvider[T]) load(ctx context.Context) (T, error) {
	// When TTL is 0 (no caching), bypass singleflight to always reload
	if c.ttl == 0 {
		return c.loadWithoutCache(ctx)
	}

	result, err, shared := c.sf.Do("load", func() (interface{}, error) {
		// Double-check after acquiring singleflight
		if val, ok := c.getCached(ctx); ok {
			return val, nil
		}

		config, err := c.store.Load(ctx)
		if err != nil {
			// If config not found and we have a default value, return it
			if c.defaultValue != nil && isNotFoundError(err) {
				return *c.defaultValue, nil
			}
			// On error, return stale cache if available
			if val, ok := c.getStaleCache(); ok {
				return val, nil
			}
			return config, err
		}

		// Validate config if validator is set
		if c.validator != nil {
			if err := c.validator.StructCtx(ctx, config); err != nil {
				return config, errors.Wrap(err, "validation failed for config")
			}
		}

		// Update cache
		c.mu.Lock()
		c.cache = config
		c.cacheSet = true
		if c.ttl > 0 {
			c.expiresAt = time.Now().Add(c.ttl)
		}
		c.mu.Unlock()

		return config, nil
	})

	if err != nil {
		var zero T
		return zero, err
	}

	if shared {
		// Another goroutine already loaded, update our cache from the result
		c.mu.Lock()
		c.cache = result.(T)
		c.cacheSet = true
		if c.ttl > 0 {
			c.expiresAt = time.Now().Add(c.ttl)
		}
		c.mu.Unlock()
	}

	return result.(T), nil
}

// loadWithoutCache loads config from storage without using singleflight or cache.
// Used when TTL is 0 to force fresh loads every time.
func (c *ConfigProvider[T]) loadWithoutCache(ctx context.Context) (T, error) {
	var zero T

	config, err := c.store.Load(ctx)
	if err != nil {
		// If config not found and we have a default value, return it
		if c.defaultValue != nil && isNotFoundError(err) {
			return *c.defaultValue, nil
		}
		return zero, err
	}

	// Validate config if validator is set
	if c.validator != nil {
		if err := c.validator.StructCtx(ctx, config); err != nil {
			return zero, errors.Wrap(err, "validation failed for config")
		}
	}

	return config, nil
}

// getStaleCache returns the cached value even if expired, for fallback on errors.
func (c *ConfigProvider[T]) getStaleCache() (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if !c.cacheSet {
		var zero T
		return zero, false
	}
	return c.cache, true
}

// isNotFoundError checks if the error indicates config not found.
// This is a placeholder - implementations should handle this appropriately.
func isNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	// Check for common not found error patterns
	errMsg := err.Error()
	return errMsg == "record not found" ||
		errMsg == "sql: no rows in result set" ||
		errMsg == "not found"
}

// NewFixedConfig creates a ConfigProvider that always returns a fixed value.
// This is useful for testing.
func NewFixedConfig[T any](value T) *ConfigProvider[T] {
	return &ConfigProvider[T]{
		cache:    value,
		cacheSet: true,
		ttl:      0, // Never expire
	}
}
