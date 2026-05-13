// Package redisx provides Redis helpers, mirroring gormx for Postgres.
//
// OpenContainer / SetupContainer start a Redis test container backed by
// testcontainers-go's redis module — kept here so consumers (ratelimiter,
// future syncx/admin needs, etc.) do not depend on github.com/theplant/testenv,
// which still pins github.com/docker/docker.
package redisx

import (
	"cmp"
	"context"
	"log/slog"
	"time"

	"github.com/pkg/errors"
	redis "github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	testredis "github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/theplant/inject/lifecycle"
)

// SetupContainer starts a Redis test container and registers a
// lifecycle cleanup actor to terminate the container on shutdown.
func SetupContainer(ctx context.Context, lc *lifecycle.Lifecycle, conf *ContainerConfig) (*Container, error) {
	container, err := OpenContainer(ctx, conf)
	if err != nil {
		return nil, err
	}
	lc.Add(lifecycle.NewFuncActor(nil, func(ctx context.Context) error {
		return container.Close(ctx)
	}).WithName("redis-container"))
	return container, nil
}

// ContainerConfig defines the configuration for starting a Redis test
// container. If a field is left empty, a sensible default will be used.
type ContainerConfig struct {
	Image string `confx:"image" usage:"Redis image to use"`
}

// DefaultContainerConfig is the default configuration for starting a Redis test container.
var DefaultContainerConfig = func() *ContainerConfig {
	return &ContainerConfig{
		Image: "redis:8.0-M04-alpine",
	}
}

// Container wraps the started test container and exposes a ready-to-use
// *redis.Client.
type Container struct {
	testcontainers.Container
	Client *redis.Client
}

// Close closes the underlying redis.Client (if any) and terminates the container.
// Both errors are reported; the container termination error takes precedence.
func (c *Container) Close(ctx context.Context) error {
	var clientErr error
	if c.Client != nil {
		clientErr = c.Client.Close()
	}
	if err := c.Container.Terminate(ctx); err != nil {
		return errors.Wrap(err, "failed to terminate redis container")
	}
	return clientErr
}

// DefaultContainerFailCleanupTimeout is used only for abnormal cleanup when
// container startup fails in OpenContainer. It does not affect normal
// termination flows managed by the lifecycle or callers.
var DefaultContainerFailCleanupTimeout = 10 * time.Second

// OpenContainer starts a Redis test container and returns a Container
// with a ready-to-use *redis.Client. Call Close on the returned Container
// to stop and clean up resources when finished. If startup fails,
// the container is terminated using DefaultContainerFailCleanupTimeout.
func OpenContainer(ctx context.Context, conf *ContainerConfig) (_ *Container, xerr error) {
	conf = cmp.Or(conf, DefaultContainerConfig())
	container, err := testredis.Run(ctx, conf.Image)
	if err != nil {
		return nil, errors.Wrap(err, "fail to start redis container")
	}
	defer func() {
		if xerr != nil {
			ctx, cancel := context.WithTimeout(
				context.WithoutCancel(ctx), DefaultContainerFailCleanupTimeout)
			defer cancel()
			err := container.Terminate(ctx)
			if err != nil {
				slog.ErrorContext(ctx, "failed to terminate redis container", "error", err)
			}
		}
	}()
	endpoint, err := container.ConnectionString(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "fail to get redis endpoint")
	}
	opts, err := redis.ParseURL(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "fail to parse redis endpoint")
	}
	return &Container{
		Container: container,
		Client:    redis.NewClient(opts),
	}, nil
}
