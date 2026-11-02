package dynconfx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStaticStore_Load(t *testing.T) {
	type Config struct {
		Host string
		Port int
	}

	cfg := Config{Host: "localhost", Port: 8080}
	store := NewStaticStore(cfg)

	ctx := context.Background()
	loaded, err := store.Load(ctx)

	require.NoError(t, err)
	assert.Equal(t, cfg, loaded)
}

func TestStaticStore_Save_ReturnsError(t *testing.T) {
	store := NewStaticStore("test")

	err := store.Save(context.Background(), "new value")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "read-only")
}

func TestStaticStore_Update_ReturnsError(t *testing.T) {
	store := NewStaticStore("test")

	err := store.Update(context.Background(), func(s string) string {
		return "updated"
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "read-only")
}

func TestStaticStore_WithConfigProvider(t *testing.T) {
	type Config struct {
		Name  string
		Debug bool
	}

	cfg := Config{Name: "app", Debug: true}
	store := NewStaticStore(cfg)
	dc := NewConfigProvider(store, WithTTL[Config](0)) // No caching

	ctx := context.Background()

	// Multiple Get calls should return the same config
	for i := 0; i < 3; i++ {
		loaded, err := dc.Get(ctx)
		require.NoError(t, err)
		assert.Equal(t, cfg, loaded)
	}
}
