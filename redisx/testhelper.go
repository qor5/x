package redisx

import (
	"context"
	"strings"
	"time"

	redis "github.com/redis/go-redis/v9"
	testredis "github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/theplant/inject"
	"github.com/theplant/inject/lifecycle"
)

// ContainerConfig defines the configuration for starting a Redis test container.
type ContainerConfig struct {
	Image string
}

// DefaultContainerConfig returns the default configuration for a Redis test container.
var DefaultContainerConfig = func() *ContainerConfig {
	return &ContainerConfig{
		Image: "redis:7-alpine",
	}
}

// Container wraps the started Redis test container and exposes the address for connecting.
type Container struct {
	container *testredis.RedisContainer
	addr      string
}

// Addr returns the Redis address (host:port) for the test container.
func (c *Container) Addr() string {
	return c.addr
}

// Terminate terminates the Redis container.
func (c *Container) Terminate(ctx context.Context) error {
	return c.container.Terminate(ctx)
}

// DefaultContainerFailCleanupTimeout is used only for abnormal cleanup when container startup fails.
var DefaultContainerFailCleanupTimeout = 10 * time.Second

// OpenContainer starts a Redis test container and returns a Container with a ready-to-use address.
func OpenContainer(ctx context.Context, conf *ContainerConfig) (_ *Container, xerr error) {
	if conf == nil {
		conf = DefaultContainerConfig()
	}
	c, err := testredis.Run(ctx, conf.Image)
	if err != nil {
		return nil, err
	}
	defer func() {
		if xerr != nil {
			cleanupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), DefaultContainerFailCleanupTimeout)
			defer cancel()
			_ = c.Terminate(cleanupCtx)
		}
	}()
	connStr, err := c.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}
	addr := strings.TrimPrefix(connStr, "redis://")
	return &Container{container: c, addr: addr}, nil
}

// SetupContainer starts a Redis test container and registers a lifecycle cleanup actor.
func SetupContainer(ctx context.Context, lc *lifecycle.Lifecycle, conf *ContainerConfig) (*Container, error) {
	c, err := OpenContainer(ctx, conf)
	if err != nil {
		return nil, err
	}
	lc.Add(lifecycle.NewFuncActor(nil, func(ctx context.Context) error {
		return c.Terminate(ctx)
	}).WithName("redis-container"))
	return c, nil
}

// TestSuite provides a test environment with a Redis container and client for integration tests.
type TestSuite struct {
	*lifecycle.Lifecycle
	*Container
	client *redis.Client
}

// Client returns the Redis client for the test suite.
func (s *TestSuite) Client() *redis.Client {
	return s.client
}

// Stop stops the test suite.
func (s *TestSuite) Stop(ctx context.Context) error {
	return s.Lifecycle.Stop(ctx)
}

// TestSuiteOption configures TestSuite creation.
type TestSuiteOption func(*testSuiteOptions)

type testSuiteOptions struct {
	containerConfig *ContainerConfig
}

// WithContainerConfig sets a custom container configuration.
func WithContainerConfig(config *ContainerConfig) TestSuiteOption {
	return func(opts *testSuiteOptions) {
		opts.containerConfig = config
	}
}

// StartTestSuite creates and starts a new test suite with a Redis container.
func StartTestSuite(ctx context.Context, opts ...TestSuiteOption) (*TestSuite, error) {
	options := &testSuiteOptions{
		containerConfig: DefaultContainerConfig(),
	}
	for _, opt := range opts {
		opt(options)
	}

	lc, err := lifecycle.Start(ctx,
		lifecycle.SetupSignal,
		func() *ContainerConfig { return options.containerConfig },
		SetupContainer,
		func(c *Container) *redis.Client {
			return redis.NewClient(&redis.Options{Addr: c.Addr()})
		},
	)
	if err != nil {
		return nil, err
	}

	client := inject.MustResolve[*redis.Client](lc)
	container := inject.MustResolve[*Container](lc)
	return &TestSuite{Lifecycle: lc, Container: container, client: client}, nil
}

// MustStartTestSuite creates and starts a new test suite, panicking on error.
func MustStartTestSuite(ctx context.Context, opts ...TestSuiteOption) *TestSuite {
	suite, err := StartTestSuite(ctx, opts...)
	if err != nil {
		panic(err)
	}
	return suite
}
