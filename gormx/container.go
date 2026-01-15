package gormx

import (
	"cmp"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/theplant/inject/lifecycle"
)

// SetupContainer starts a PostgreSQL test container and registers a
// lifecycle cleanup actor to terminate the container on shutdown.
func SetupContainer(ctx context.Context, lc *lifecycle.Lifecycle, conf *ContainerConfig) (*Container, error) {
	container, err := OpenContainer(ctx, conf)
	if err != nil {
		return nil, err
	}
	lc.Add(lifecycle.NewFuncActor(nil, func(ctx context.Context) error {
		return container.Terminate(ctx)
	}).WithName("database-container"))
	return container, nil
}

// ContainerConfig defines the configuration for starting a PostgreSQL
// test container. If a field is left empty, a sensible default will be used.
//
// HostPort binds the container's 5432/tcp to a specific host port when set.
// Leave it empty to let Docker allocate a random available port.
type ContainerConfig struct {
	Image        string `confx:"image" usage:"PostgreSQL image to use"`
	Username     string `confx:"username" usage:"PostgreSQL username"`
	Password     string `confx:"password" usage:"PostgreSQL password"`
	DatabaseName string `confx:"databaseName" usage:"PostgreSQL database name"`
	HostPort     string `confx:"hostPort" usage:"PostgreSQL host port"`
}

// DefaultContainerConfig is the default configuration for starting a PostgreSQL test container.
var DefaultContainerConfig = func() *ContainerConfig {
	return &ContainerConfig{
		Image:        "postgres:17.4-alpine3.21",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "postgres",
		HostPort:     "",
	}
}

// Container wraps the started test container and exposes a DSN for
// connecting with PostgreSQL clients.
type Container struct {
	testcontainers.Container
	DSN string
}

// Terminate terminates the PostgreSQL container.
func (c *Container) Terminate(ctx context.Context, opts ...testcontainers.TerminateOption) error {
	err := c.Container.Terminate(ctx, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to terminate postgres container")
	}
	return nil
}

// DefaultContainerFailCleanupTimeout is used only for abnormal cleanup when
// container startup fails in OpenContainer. It does not affect normal
// termination flows managed by the lifecycle or callers.
var DefaultContainerFailCleanupTimeout = 10 * time.Second

// OpenContainer starts a PostgreSQL test container and returns a
// Container with a ready-to-use DSN. Call Terminate on the underlying
// container to stop and clean up resources when finished. If startup fails,
// the container is terminated using ContainerFailCleanupTimeout.
func OpenContainer(ctx context.Context, conf *ContainerConfig) (_ *Container, xerr error) {
	conf = cmp.Or(conf, DefaultContainerConfig())
	req := testcontainers.ContainerRequest{
		Image:        conf.Image,
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     conf.Username,
			"POSTGRES_PASSWORD": conf.Password,
			"POSTGRES_DB":       conf.DatabaseName,
		},
		Cmd:        []string{"postgres", "-c", "fsync=off"},
		WaitingFor: wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
	}
	if conf.HostPort != "" {
		req.HostConfigModifier = func(hostConfig *container.HostConfig) {
			hostConfig.PortBindings = map[nat.Port][]nat.PortBinding{
				"5432/tcp": {
					{
						HostIP:   "0.0.0.0",
						HostPort: conf.HostPort,
					},
				},
			}
		}
	}
	container, err := testcontainers.GenericContainer(ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "fail to start container")
	}
	defer func() {
		if xerr != nil {
			ctx, cancel := context.WithTimeout(
				context.WithoutCancel(ctx), DefaultContainerFailCleanupTimeout)
			defer cancel()
			err := container.Terminate(ctx)
			if err != nil {
				slog.ErrorContext(ctx, "failed to terminate postgres container", "error", err)
			}
		}
	}()
	endpoint, err := container.Endpoint(ctx, "")
	if err != nil {
		return nil, errors.Wrap(err, "fail to get test container endpoint")
	}
	return &Container{
		Container: container,
		DSN:       fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", conf.Username, conf.Password, endpoint, conf.DatabaseName),
	}, nil
}
