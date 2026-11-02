package gormx_test

import (
	"context"
	"testing"

	"github.com/qor5/x/v3/gormx"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TestOpenContainer_Args pins both halves of the contract: the caller's
// arguments reach the postmaster, and they are appended to the built-in
// "-c fsync=off" rather than replacing it.
func TestOpenContainer_Args(t *testing.T) {
	ctx := context.Background()

	conf := gormx.DefaultContainerConfig()
	conf.Args = []string{"-c", "max_connections=123"}

	container, err := gormx.OpenContainer(ctx, conf)
	require.NoError(t, err)
	t.Cleanup(func() {
		_ = container.Terminate(context.WithoutCancel(ctx))
	})

	db, err := gorm.Open(postgres.Open(container.DSN), &gorm.Config{})
	require.NoError(t, err)

	var maxConnections string
	require.NoError(t, db.Raw("SHOW max_connections").Scan(&maxConnections).Error)
	require.Equal(t, "123", maxConnections)

	var fsync string
	require.NoError(t, db.Raw("SHOW fsync").Scan(&fsync).Error)
	require.Equal(t, "off", fsync, "Args must append to the default arguments, not replace them")
}
