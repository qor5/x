package gormx

import (
	"context"

	"gorm.io/gorm"
)

type ctxKeyDB struct{}

// FromContext returns a *gorm.DB from the context, or a fallback if not found.
// The returned *gorm.DB is ensured to have the same context as the provided context.
func FromContext(ctx context.Context, fallback *gorm.DB) *gorm.DB {
	if db, ok := ctx.Value(ctxKeyDB{}).(*gorm.DB); ok {
		return db.WithContext(ctx)
	}
	if fallback == nil {
		return nil
	}
	return fallback.WithContext(ctx)
}

// NewContext returns a new context with the provided *gorm.DB.
func NewContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, ctxKeyDB{}, db)
}
