package dynconfx

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockStore is a mock implementation of ConfigStore for testing.
type mockStore[T any] struct {
	loadFunc  func(ctx context.Context) (T, error)
	saveFunc  func(ctx context.Context, config T) error
	updateFunc func(ctx context.Context, updateFunc func(T) T) error
	loadCount int
	saveCount int
	updateCount int
}

func (m *mockStore[T]) Load(ctx context.Context) (T, error) {
	m.loadCount++
	if m.loadFunc != nil {
		return m.loadFunc(ctx)
	}
	var zero T
	return zero, nil
}

func (m *mockStore[T]) Save(ctx context.Context, config T) error {
	m.saveCount++
	if m.saveFunc != nil {
		return m.saveFunc(ctx, config)
	}
	return nil
}

func (m *mockStore[T]) Update(ctx context.Context, updateFunc func(T) T) error {
	m.updateCount++
	if m.updateFunc != nil {
		return m.updateFunc(ctx, updateFunc)
	}
	// Default implementation: load, update, save
	config, err := m.Load(ctx)
	if err != nil {
		return err
	}
	updated := updateFunc(config)
	return m.Save(ctx, updated)
}


// mockValidator is a mock implementation of Validator for testing.
type mockValidator struct {
	validateFunc func(ctx context.Context, v any) error
}

func (m *mockValidator) StructCtx(ctx context.Context, v any) error {
	if m.validateFunc != nil {
		return m.validateFunc(ctx, v)
	}
	return nil
}

func TestConfigProvider_Get_CacheHit(t *testing.T) {
	type Config struct {
		Port int
	}

	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			return Config{Port: 8080}, nil
		},
	}

	validator := &mockValidator{}
	dc := NewConfigProvider[Config](
		store,
		WithTTL[Config](5*time.Minute),
		WithValidator[Config](validator),
	)

	ctx := context.Background()

	// First call should load from store
	config, err := dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, 1, store.loadCount)

	// Second call should use cache
	config, err = dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, 1, store.loadCount) // No additional load
}

func TestConfigProvider_Get_CacheExpired(t *testing.T) {
	type Config struct {
		Port int
	}

	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			return Config{Port: 8080}, nil
		},
	}

	dc := NewConfigProvider[Config](
		store,
		WithTTL[Config](100*time.Millisecond), // Short TTL for testing
	)

	ctx := context.Background()

	// First call
	config, err := dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, 1, store.loadCount)

	// Wait for cache to expire
	time.Sleep(150 * time.Millisecond)

	// Second call should reload from store
	config, err = dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, 2, store.loadCount)
}

func TestConfigProvider_Get_NoCaching(t *testing.T) {
	type Config struct {
		Port int
	}

	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			return Config{Port: 8080}, nil
		},
	}

	dc := NewConfigProvider[Config](
		store,
		WithTTL[Config](0), // No caching
	)

	ctx := context.Background()

	// Each call should load from store
	config, err := dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, 1, store.loadCount)

	config, err = dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, 2, store.loadCount)
}

func TestConfigProvider_Reload(t *testing.T) {
	type Config struct {
		Port int
	}

	port := 8080
	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			port++
			return Config{Port: port}, nil
		},
	}

	dc := NewConfigProvider[Config](
		store,
		WithTTL[Config](5*time.Minute),
	)

	ctx := context.Background()

	// Initial load
	config, err := dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8081, config.Port)
	assert.Equal(t, 1, store.loadCount)

	// Get should use cache
	config, err = dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8081, config.Port)
	assert.Equal(t, 1, store.loadCount)

	// Force reload
	config, err = dc.Reload(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8082, config.Port)
	assert.Equal(t, 2, store.loadCount)
}

func TestConfigProvider_Reload_ErrorFallsBackToStaleCache(t *testing.T) {
	type Config struct {
		Port int
	}

	loadErr := errors.New("db unavailable")
	failing := false
	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			if failing {
				var zero Config
				return zero, loadErr
			}
			return Config{Port: 8080}, nil
		},
	}

	dc := NewConfigProvider[Config](
		store,
		WithTTL[Config](5*time.Minute),
	)

	ctx := context.Background()

	// Prime the cache with a successful load.
	config, err := dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8080, config.Port)

	// Store now fails: Reload should return the previously cached value
	// instead of surfacing the error.
	failing = true
	config, err = dc.Reload(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8080, config.Port)
}

func TestConfigProvider_Reload_ErrorWithoutCacheReturnsError(t *testing.T) {
	type Config struct {
		Port int
	}

	loadErr := errors.New("db unavailable")
	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			var zero Config
			return zero, loadErr
		},
	}

	dc := NewConfigProvider[Config](
		store,
		WithTTL[Config](5*time.Minute),
	)

	ctx := context.Background()

	// No prior successful load means no stale cache to fall back to,
	// so the load error propagates to the caller.
	_, err := dc.Reload(ctx)
	require.Error(t, err)
	assert.ErrorIs(t, err, loadErr)
}

func TestConfigProvider_Invalidate(t *testing.T) {
	type Config struct {
		Port int
	}

	port := 8080
	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			port++
			return Config{Port: port}, nil
		},
	}

	dc := NewConfigProvider[Config](
		store,
		WithTTL[Config](5*time.Minute),
	)

	ctx := context.Background()

	// Initial load
	config, err := dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8081, config.Port)
	assert.Equal(t, 1, store.loadCount)

	// Invalidate cache
	dc.Invalidate()

	// Next Get should reload
	config, err = dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 8082, config.Port)
	assert.Equal(t, 2, store.loadCount)
}

func TestConfigProvider_ValidationError(t *testing.T) {
	type Config struct {
		Port int
	}

	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			return Config{Port: 8080}, nil
		},
	}

	validator := &mockValidator{
		validateFunc: func(ctx context.Context, v any) error {
			return errors.New("validation failed")
		},
	}

	dc := NewConfigProvider[Config](
		store,
		WithTTL[Config](5*time.Minute),
		WithValidator[Config](validator),
	)

	ctx := context.Background()

	// Load should fail validation
	_, err := dc.Get(ctx)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
}

func TestConfigProvider_WithDefault(t *testing.T) {
	type Config struct {
		Port int
	}

	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			var zero Config
			return zero, errors.New("not found")
		},
	}

	dc := NewConfigProvider[Config](
		store,
		WithDefault[Config](Config{Port: 9090}),
	)

	ctx := context.Background()

	// Should return default value when config not found
	config, err := dc.Get(ctx)
	require.NoError(t, err)
	assert.Equal(t, 9090, config.Port)
}

func TestConfigProvider_NewFixedConfig(t *testing.T) {
	type Config struct {
		Port int
	}

	fixed := Config{Port: 8080}
	dc := NewFixedConfig(fixed)

	ctx := context.Background()

	// Should always return fixed value
	for i := 0; i < 5; i++ {
		config, err := dc.Get(ctx)
		require.NoError(t, err)
		assert.Equal(t, 8080, config.Port)
	}
}

func TestConfigProvider_ConcurrentAccess(t *testing.T) {
	type Config struct {
		Port int
	}

	store := &mockStore[Config]{
		loadFunc: func(ctx context.Context) (Config, error) {
			return Config{Port: 8080}, nil
		},
	}

	dc := NewConfigProvider[Config](
		store,
		WithTTL[Config](5*time.Minute),
	)

	ctx := context.Background()

	// Concurrent reads
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			config, err := dc.Get(ctx)
			require.NoError(t, err)
			assert.Equal(t, 8080, config.Port)
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Should have loaded only once due to singleflight and cache
	assert.Equal(t, 1, store.loadCount)
}
