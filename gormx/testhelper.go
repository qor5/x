package gormx

import (
	"context"
	_ "embed"

	"github.com/pkg/errors"
	"github.com/theplant/inject"
	"github.com/theplant/inject/lifecycle"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TestSuite provides a test environment with a PostgreSQL container,
// database connection, and lifecycle management for integration tests.
type TestSuite struct {
	*lifecycle.Lifecycle
	*Container
	db *gorm.DB
}

// DSN returns the database connection string for the test suite.
func (s *TestSuite) DSN() string {
	return s.Container.DSN
}

// DB returns the GORM database connection for the test suite.
func (s *TestSuite) DB() *gorm.DB {
	return s.db
}

// GetName returns the name of the test suite.
func (s *TestSuite) GetName() string {
	return "gormx-test-suite"
}

// Stop stops the test suite.
func (s *TestSuite) Stop(ctx context.Context) error {
	return s.Lifecycle.Stop(ctx)
}

// ResetDB drops and recreates all tables for the provided models.
// This is useful for cleaning up test data between test runs.
func (s *TestSuite) ResetDB(ctx context.Context, models ...any) error {
	if len(models) == 0 {
		return nil
	}

	db := s.db.WithContext(ctx)

	if err := db.Migrator().DropTable(models...); err != nil {
		return err
	}

	return db.AutoMigrate(models...)
}

// TestSuiteOption configures TestSuite creation
type TestSuiteOption func(*testSuiteOptions)

type testSuiteOptions struct {
	containerConfig *ContainerConfig
	providers       []any
}

// WithContainerConfig sets a custom container configuration
func WithContainerConfig(config *ContainerConfig) TestSuiteOption {
	return func(opts *testSuiteOptions) {
		opts.containerConfig = config
	}
}

// WithProviders adds additional dependency injection providers
func WithProviders(providers ...any) TestSuiteOption {
	return func(opts *testSuiteOptions) {
		opts.providers = append(opts.providers, providers...)
	}
}

// StartTestSuite creates and starts a new test suite with PostgreSQL container.
// Configuration can be customized using TestSuiteOption functions.
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
		func(c *Container) *DatabaseConfig {
			conf := DefaultDatabaseConfig()
			conf.DSN = c.DSN
			return conf
		},
		SetupDatabase,
		options.providers,
	)
	if err != nil {
		return nil, err
	}

	db := inject.MustResolve[*gorm.DB](lc)
	container := inject.MustResolve[*Container](lc)
	return &TestSuite{Lifecycle: lc, Container: container, db: db}, nil
}

// MustStartTestSuite creates and starts a new test suite, panicking on error.
// This is a convenience wrapper around StartTestSuite for test code that prefers panics.
func MustStartTestSuite(ctx context.Context, opts ...TestSuiteOption) *TestSuite {
	suite, err := StartTestSuite(ctx, opts...)
	if err != nil {
		panic(err)
	}
	return suite
}

// SetupTestSuiteFactory creates a factory function for creating test suites.
func SetupTestSuiteFactory(opts ...TestSuiteOption) func(ctx context.Context, lc *lifecycle.Lifecycle) (*TestSuite, error) {
	return func(ctx context.Context, lc *lifecycle.Lifecycle) (*TestSuite, error) {
		suite, err := StartTestSuite(ctx, opts...)
		if err != nil {
			return nil, err
		}
		lc.Add(suite)
		return suite, nil
	}
}

// StartRawTestSuite starts a PostgreSQL test container and opens a *gorm.DB
// against it using plain `gorm.Open(postgres.Open(dsn), &gorm.Config{})` —
// without the production plugins that SetupDatabase installs
// (OmitAssociationsPlugin, SoftDeleteUpdatedAtPlugin, TracingPlugin).
//
// Use this when you want the testenv-style "just give me a *gorm.DB and a
// throwaway database" behavior — e.g. tests that exercise GORM associations
// directly (where OmitAssociationsPlugin would silently change semantics),
// or `Example_*` tests that match exact stdout (where TracingPlugin would
// pollute output with JSON log lines).
//
// Use StartTestSuite / MustStartTestSuite when you want the full production
// data path (tracing, IAM auth retry, association omission).
//
// The Container is registered with the suite's Lifecycle, so suite.Stop()
// cleans up. WithProviders is supported (additional DI providers are added);
// WithContainerConfig is supported.
func StartRawTestSuite(ctx context.Context, opts ...TestSuiteOption) (*TestSuite, error) {
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
		func(ctx context.Context, lc *lifecycle.Lifecycle, c *Container) (*gorm.DB, error) {
			db, err := gorm.Open(postgres.Open(c.DSN), &gorm.Config{})
			if err != nil {
				return nil, errors.Wrap(err, "failed to open gorm")
			}
			lc.Add(lifecycle.NewFuncActor(nil, func(_ context.Context) error {
				sqlDB, err := db.DB()
				if err != nil {
					return err
				}
				return sqlDB.Close()
			}).WithName("raw-test-db"))
			return db, nil
		},
		options.providers,
	)
	if err != nil {
		return nil, err
	}

	db := inject.MustResolve[*gorm.DB](lc)
	container := inject.MustResolve[*Container](lc)
	return &TestSuite{Lifecycle: lc, Container: container, db: db}, nil
}

// MustStartRawTestSuite is StartRawTestSuite that panics on error.
// See StartRawTestSuite for when to choose Raw over the regular variant.
func MustStartRawTestSuite(ctx context.Context, opts ...TestSuiteOption) *TestSuite {
	suite, err := StartRawTestSuite(ctx, opts...)
	if err != nil {
		panic(err)
	}
	return suite
}
