package gormx_test

import (
	"context"
	"sync/atomic"
	"testing"

	gklog "github.com/go-kit/log"
	"github.com/qor5/x/v3/gormx"
	"github.com/stretchr/testify/require"
	kitlog "github.com/theplant/appkit/log"
	"github.com/theplant/appkit/logtracing"
)

type TracingTestModel struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func TestWithoutTracing(t *testing.T) {
	ctx := context.Background()
	db := suite.DB()

	err := suite.ResetDB(ctx, &TracingTestModel{})
	require.NoError(t, err)

	err = db.WithContext(ctx).Create(&TracingTestModel{Name: "test"}).Error
	require.NoError(t, err)

	// Create a custom logger to track tracing log calls
	var logCount atomic.Int32
	countingLogger := kitlog.Logger{Logger: gklog.LoggerFunc(func(keyvals ...any) error {
		logCount.Add(1)
		return nil
	})}

	// Set custom logger in context
	tracedCtx, _ := logtracing.StartSpan(ctx, "test-parent")
	tracedCtx = kitlog.Context(tracedCtx, countingLogger)
	defer logtracing.EndSpan(tracedCtx, nil)

	// Normal query: should produce tracing log
	logCount.Store(0)
	var result1 TracingTestModel
	err = db.WithContext(tracedCtx).First(&result1).Error
	require.NoError(t, err)
	require.Greater(t, logCount.Load(), int32(0), "normal query should trigger tracing log")

	// WithoutTracing: should NOT produce tracing log
	logCount.Store(0)
	var result2 TracingTestModel
	err = db.WithContext(tracedCtx).Scopes(gormx.WithoutTracing()).First(&result2).Error
	require.NoError(t, err)
	require.Equal(t, int32(0), logCount.Load(), "WithoutTracing should not trigger tracing log")
}
