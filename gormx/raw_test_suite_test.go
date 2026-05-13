package gormx_test

import (
	"bytes"
	"context"
	"io"
	"os"
	"sync"
	"testing"

	"github.com/qor5/x/v3/gormx"
	"github.com/stretchr/testify/require"
)

type rawTestModel struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

// TestMustStartRawTestSuite verifies that StartRawTestSuite returns a
// working *gorm.DB and lifecycle-managed container without installing
// SetupDatabase's plugins.
func TestMustStartRawTestSuite(t *testing.T) {
	ctx := context.Background()
	s := gormx.MustStartRawTestSuite(ctx)
	defer func() { _ = s.Stop(ctx) }()

	require.NoError(t, s.DB().AutoMigrate(&rawTestModel{}))
	require.NoError(t, s.DB().Create(&rawTestModel{Name: "alice"}).Error)

	var got rawTestModel
	require.NoError(t, s.DB().First(&got).Error)
	require.Equal(t, "alice", got.Name)
}

// TestRawTestSuiteNoTracingOutput ensures the raw suite does not emit
// SQL-trace JSON log lines on stdout/stderr (the regression that drove
// adding MustStartRawTestSuite — Example_* tests rely on byte-exact
// stdout capture).
func TestRawTestSuiteNoTracingOutput(t *testing.T) {
	origStdout, origStderr := os.Stdout, os.Stderr
	rOut, wOut, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()
	os.Stdout, os.Stderr = wOut, wErr
	t.Cleanup(func() {
		os.Stdout, os.Stderr = origStdout, origStderr
	})

	var bufOut, bufErr bytes.Buffer
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); _, _ = io.Copy(&bufOut, rOut) }()
	go func() { defer wg.Done(); _, _ = io.Copy(&bufErr, rErr) }()

	func() {
		ctx := context.Background()
		s := gormx.MustStartRawTestSuite(ctx)
		defer func() { _ = s.Stop(ctx) }()

		require.NoError(t, s.DB().AutoMigrate(&rawTestModel{}))
		require.NoError(t, s.DB().Create(&rawTestModel{Name: "bob"}).Error)
		var got rawTestModel
		require.NoError(t, s.DB().First(&got).Error)
	}()

	_ = wOut.Close()
	_ = wErr.Close()
	wg.Wait()

	// Tracing plugin emits JSON like {"caller":"trace.go:...","msg":"gorm.Create..."}
	combined := bufOut.String() + bufErr.String()
	require.NotContains(t, combined, `"caller":"trace.go`, "raw suite must not emit gormx tracing logs")
	require.NotContains(t, combined, `gorm.Create:`, "raw suite must not emit gorm.Create trace spans")
}
