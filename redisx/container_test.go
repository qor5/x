package redisx_test

import (
	"context"
	"testing"

	"github.com/qor5/x/v3/redisx"
	"github.com/stretchr/testify/require"
)

func TestOpenContainer(t *testing.T) {
	ctx := context.Background()
	container, err := redisx.OpenContainer(ctx, nil)
	require.NoError(t, err)
	defer func() { _ = container.Close(ctx) }()

	require.NotNil(t, container.Client)
	require.NoError(t, container.Client.Ping(ctx).Err())

	require.NoError(t, container.Client.Set(ctx, "k", "v", 0).Err())
	got, err := container.Client.Get(ctx, "k").Result()
	require.NoError(t, err)
	require.Equal(t, "v", got)
}
