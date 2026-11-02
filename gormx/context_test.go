package gormx_test

import (
	"context"
	"testing"

	"github.com/qor5/x/v3/gormx"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFromContext(t *testing.T) {
	db := suite.DB()

	t.Run("returns db from context", func(t *testing.T) {
		ctx := gormx.NewContext(context.Background(), db)
		result := gormx.FromContext(ctx, nil)
		assert.NotNil(t, result)
		assert.Equal(t, ctx, result.Statement.Context)
	})

	t.Run("returns fallback when not in context", func(t *testing.T) {
		ctx := context.Background()
		result := gormx.FromContext(ctx, db)
		assert.NotNil(t, result)
		assert.Equal(t, ctx, result.Statement.Context)
	})

	t.Run("returns nil when no db in context and no fallback", func(t *testing.T) {
		ctx := context.Background()
		result := gormx.FromContext(ctx, nil)
		assert.Nil(t, result)
	})

	t.Run("preserves context when db exists in context", func(t *testing.T) {
		originalCtx := context.WithValue(context.Background(), "key", "value")
		ctxWithDB := gormx.NewContext(originalCtx, db)
		result := gormx.FromContext(ctxWithDB, nil)
		assert.NotNil(t, result)
		assert.Equal(t, "value", result.Statement.Context.Value("key"))
	})

	t.Run("preserves context when using fallback", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), "key", "value")
		result := gormx.FromContext(ctx, db)
		assert.NotNil(t, result)
		assert.Equal(t, "value", result.Statement.Context.Value("key"))
	})
}

func TestNewContext(t *testing.T) {
	db := suite.DB()

	t.Run("stores db in context", func(t *testing.T) {
		ctx := gormx.NewContext(context.Background(), db)
		result := gormx.FromContext(ctx, nil)
		assert.NotNil(t, result)
	})

	t.Run("preserves existing context values", func(t *testing.T) {
		originalCtx := context.WithValue(context.Background(), "key", "value")
		ctx := gormx.NewContext(originalCtx, db)
		assert.Equal(t, "value", ctx.Value("key"))
	})

	t.Run("can be overwritten", func(t *testing.T) {
		db2 := db.Session(&gorm.Session{}).Set("key", "value")
		ctx := gormx.NewContext(context.Background(), db)
		ctx = gormx.NewContext(ctx, db2)
		result := gormx.FromContext(ctx, nil)
		assert.NotNil(t, result)
		value, ok := result.Get("key")
		assert.Equal(t, "value", value)
		assert.True(t, ok)
		{
			_, ok := db.Get("key")
			assert.False(t, ok)
		}
	})
}
