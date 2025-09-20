package gormx

import (
	"context"
	_ "embed"

	"github.com/theplant/inject"
	"github.com/theplant/inject/lifecycle"
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

// StartTestSuiteWithConfig creates a test suite with custom container configuration.
// If containerConfig is nil, default configuration will be used.
func StartTestSuiteWithConfig(ctx context.Context, containerConfig *ContainerConfig, ctors ...any) (*TestSuite, error) {
	configProvider := DefaultContainerConfig
	if containerConfig != nil {
		configProvider = func() *ContainerConfig { return containerConfig }
	}

	lc, err := lifecycle.Start(ctx,
		lifecycle.SetupSignal,
		configProvider,
		SetupContainer,
		func(c *Container) *DatabaseConfig {
			conf := DefaultDatabaseConfig()
			conf.DSN = c.DSN
			return conf
		},
		SetupDatabase,
		ctors,
	)
	if err != nil {
		return nil, err
	}

	db := inject.MustResolve[*gorm.DB](lc)
	container := inject.MustResolve[*Container](lc)
	return &TestSuite{Lifecycle: lc, Container: container, db: db}, nil
}

// StartTestSuite creates and starts a new test suite with PostgreSQL container.
// Additional constructors can be provided to extend the dependency injection setup.
func StartTestSuite(ctx context.Context, ctors ...any) (*TestSuite, error) {
	return StartTestSuiteWithConfig(ctx, nil, ctors...)
}

// MustStartTestSuite creates and starts a new test suite, panicking on error.
// This is a convenience wrapper around StartTestSuite for test code that prefers panics.
func MustStartTestSuite(ctx context.Context, ctors ...any) *TestSuite {
	suite, err := StartTestSuite(ctx, ctors...)
	if err != nil {
		panic(err)
	}
	return suite
}

// SetupTestSuiteFactory creates a factory function for creating test suites.
func SetupTestSuiteFactory(ctors ...any) func(ctx context.Context, lc *lifecycle.Lifecycle) (*TestSuite, error) {
	return func(ctx context.Context, lc *lifecycle.Lifecycle) (*TestSuite, error) {
		suite, err := StartTestSuite(ctx, ctors...)
		if err != nil {
			return nil, err
		}
		lc.Add(suite)
		return suite, nil
	}
}
