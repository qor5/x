package gormx

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/cenkalti/backoff/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/qor5/confx"
	"github.com/qor5/x/v3/gormx/postgresx"
	"github.com/theplant/inject/lifecycle"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "embed"
)

var DefaultCreateBatchSize = 100

type AuthMethod string

const (
	AuthMethodPassword AuthMethod = "password"
	AuthMethodIAM      AuthMethod = "iam"
)

type DatabaseConfig struct {
	DSN             string        `confx:"dsn" usage:"Database connection string" validate:"required"`
	Debug           bool          `confx:"debug" usage:"Enable debug mode"`
	Tracing         TracingConfig `confx:"tracing" usage:"Tracing configuration"`
	MaxIdleConns    int           `confx:"maxIdleConns" usage:"Maximum number of idle connections" validate:"ltefield=MaxOpenConns"`
	MaxOpenConns    int           `confx:"maxOpenConns" usage:"Maximum number of open connections"`
	ConnMaxLifetime time.Duration `confx:"connMaxLifetime" usage:"Maximum connection lifetime"`
	ConnMaxIdleTime time.Duration `confx:"connMaxIdleTime" usage:"Maximum idle time for connections" validate:"ltefield=ConnMaxLifetime"`
	AuthMethod      AuthMethod    `confx:"authMethod" usage:"Authentication method: 'password' or 'iam'" validate:"required,oneof=password iam"`
}

//go:embed embed/default-database-config.yaml
var defaultDatabaseConfigYAML string

var DefaultDatabaseConfig = func() *DatabaseConfig {
	def, err := confx.Read[*struct {
		Database DatabaseConfig `confx:"database"`
	}]("yaml", strings.NewReader(defaultDatabaseConfigYAML))
	if err != nil {
		panic(err)
	}
	return &def.Database
}

var SetupDatabase = SetupDatabaseFactory("database")

func SetupDatabaseFactory(name string, opts ...gorm.Option) func(ctx context.Context, lc *lifecycle.Lifecycle, conf *DatabaseConfig) (*gorm.DB, error) {
	return func(ctx context.Context, lc *lifecycle.Lifecycle, conf *DatabaseConfig) (*gorm.DB, error) {
		db, closer, err := Open(ctx, conf, opts...)
		if err != nil {
			return nil, err
		}

		lc.Add(lifecycle.NewFuncActor(nil, func(_ context.Context) error {
			return closer.Close()
		}).WithName(name))

		return db, nil
	}
}

type dbCloserWrapper struct {
	*sql.DB
}

func (c *dbCloserWrapper) Close() error {
	if err := c.DB.Close(); err != nil {
		return errors.Wrap(err, "failed to close database connection")
	}
	return nil
}

func Open(ctx context.Context, conf *DatabaseConfig, opts ...gorm.Option) (*gorm.DB, io.Closer, error) {
	var (
		dialector gorm.Dialector
		err       error
	)
	switch conf.AuthMethod {
	case AuthMethodIAM:
		slog.InfoContext(ctx, "Using IAM authentication to connect to database")
		awsConfig, err := config.LoadDefaultConfig(ctx)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to load AWS config")
		}
		dialector, err = NewIAMDialector(conf.DSN, awsConfig.Region, awsConfig.Credentials)
		if err != nil {
			return nil, nil, err
		}
	case AuthMethodPassword:
		slog.InfoContext(ctx, "Using password authentication to connect to database")
		dialector, err = NewDefaultDialector(conf.DSN)
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errors.Errorf("unsupported auth method: %s", conf.AuthMethod)
	}

	db, err := gorm.Open(
		dialector,
		append([]gorm.Option{
			&gorm.Config{
				// We don't want to use any foreign key constraint.
				DisableForeignKeyConstraintWhenMigrating: true,
				// CreateBatchSize is set to DefaultCreateBatchSize to improve performance by reducing the number of round trips to the database.
				CreateBatchSize: DefaultCreateBatchSize,
				// PrepareStmt is disabled because pgx driver already enables prepared statement cache
				// by default (extended protocol). Enabling this would create redundant statement caching
				// at both GORM and driver levels, potentially causing memory overhead without performance gain.
				PrepareStmt: false,
				// Enable GORM error translation to convert database-specific errors into standardized GORM errors.
				// Benefits:
				// 1. Better error handling: Enables type-based error handling (errors.Is(err, gorm.ErrDuplicatedKey))
				//    instead of fragile string matching against raw database error messages.
				//    Example: PostgreSQL "ERROR: duplicate key value violates unique constraint" -> gorm.ErrDuplicatedKey
				// 2. Clean abstraction: Business logic remains database-agnostic
				TranslateError: true,
				// QueryFields:                              true,
			},
		}, opts...)...,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to open database connection")
	}

	if err := db.Use(OmitAssociationsPlugin); err != nil {
		return nil, nil, errors.Wrap(err, "failed to setup omit associations plugin")
	}

	if err := db.Use(NewTracingPlugin(&conf.Tracing)); err != nil {
		return nil, nil, errors.Wrap(err, "failed to setup database tracing plugin")
	}

	if conf.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get database connection")
	}
	if conf.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	}
	if conf.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	}
	if conf.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(conf.ConnMaxLifetime)
	}
	if conf.ConnMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(conf.ConnMaxIdleTime)
	}

	return db, &dbCloserWrapper{sqlDB}, nil
}

func buildIAMAuthToken(ctx context.Context, endpoint, region, dbUser string, credProvider aws.CredentialsProvider) (string, error) {
	var token string
	if err := backoff.RetryNotify(
		func() error {
			t, err := auth.BuildAuthToken(ctx, endpoint, region, dbUser, credProvider)
			if err != nil {
				return errors.Wrap(err, "failed to build IAM auth token")
			}
			token = t
			return nil
		},
		backoff.WithContext(&backoff.ExponentialBackOff{
			InitialInterval:     100 * time.Millisecond,
			RandomizationFactor: 0.5,
			Multiplier:          2,
			MaxInterval:         time.Second,
			MaxElapsedTime:      0,
			Stop:                backoff.Stop,
			Clock:               backoff.SystemClock,
		}, ctx),
		func(err error, next time.Duration) {
			slog.ErrorContext(ctx, "Failed to get IAM Token, will retry", "retryIn", next.String(), "error", err)
		},
	); err != nil {
		return "", errors.Wrap(err, "failed to get IAM auth token after retries")
	}

	slog.InfoContext(ctx, "Successfully got IAM Token")
	return token, nil
}

func NewIAMDialector(dsn string, region string, credProvider aws.CredentialsProvider) (gorm.Dialector, error) {
	if dsn == "" {
		return nil, errors.New("dsn is required")
	}
	conf, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}
	if conf.Host == "" {
		return nil, errors.New("host is required")
	}
	if conf.Port == 0 {
		return nil, errors.New("port is required")
	}
	if conf.User == "" {
		return nil, errors.New("user is required")
	}
	if conf.Password != "" {
		return nil, errors.New("password should not be provided when connecting to database using IAM authentication")
	}
	if region == "" {
		return nil, errors.New("region is required")
	}
	if credProvider == nil {
		return nil, errors.New("credentials provider is required")
	}

	optBeforeConnect := stdlib.OptionBeforeConnect(func(ctx context.Context, conf *pgx.ConnConfig) error {
		endpoint := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
		token, err := buildIAMAuthToken(ctx, endpoint, region, conf.User, credProvider)
		if err != nil {
			return err
		}
		conf.Password = token
		return nil
	})
	conn := stdlib.OpenDB(*conf, optBeforeConnect)
	return postgresx.New(postgres.Config{
		Conn: conn,
		// We are using pgx as postgresql's database/sql driver, it enables prepared statement cache by default (extended protocol)
		// PreferSimpleProtocol: true, // disables implicit prepared statement usage.
	}), nil
}

func NewDefaultDialector(dsn string) (gorm.Dialector, error) {
	if dsn == "" {
		return nil, errors.New("dsn is required")
	}
	return postgresx.New(postgres.Config{
		DSN: dsn,
		// We are using pgx as postgresql's database/sql driver, it enables prepared statement cache by default (extended protocol)
		// PreferSimpleProtocol: true, // disables implicit prepared statement usage.
	}), nil
}
