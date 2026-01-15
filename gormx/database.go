package gormx

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	"github.com/cenkalti/backoff/v4"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
	"github.com/qor5/confx"
	"github.com/qor5/x/v3/awsx"
	"github.com/qor5/x/v3/gormx/postgresx"
	"github.com/theplant/cachex"
	"github.com/theplant/inject"
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

type IAMDialectorConfig struct {
	TokenEndpoint string      `confx:"tokenEndpoint" usage:"RDS endpoint for IAM token generation (optional, defaults to DSN host:port)"`
	AWSConfig     *aws.Config `confx:"-" `
}

type DatabaseConfig struct {
	DSN             string             `confx:"dsn" usage:"Database connection string" validate:"required"`
	Debug           bool               `confx:"debug" usage:"Enable debug mode"`
	Tracing         TracingConfig      `confx:"tracing" usage:"Tracing configuration"`
	MaxIdleConns    int                `confx:"maxIdleConns" usage:"Maximum number of idle connections" validate:"ltefield=MaxOpenConns"`
	MaxOpenConns    int                `confx:"maxOpenConns" usage:"Maximum number of open connections"`
	ConnMaxLifetime time.Duration      `confx:"connMaxLifetime" usage:"Maximum connection lifetime"`
	ConnMaxIdleTime time.Duration      `confx:"connMaxIdleTime" usage:"Maximum idle time for connections" validate:"ltefield=ConnMaxLifetime"`
	AuthMethod      AuthMethod         `confx:"authMethod" usage:"Authentication method: 'password' or 'iam'" validate:"required,oneof=password iam"`
	IAM             IAMDialectorConfig `confx:"iam" validate:"skip_nested_unless=AuthMethod iam" usage:"IAM configuration"`
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
		if conf.AuthMethod == AuthMethodIAM && conf.IAM.AWSConfig == nil {
			var awsCfg *aws.Config
			if err := lc.ResolveContext(ctx, &awsCfg); err != nil && !errors.Is(err, inject.ErrTypeNotProvided) {
				return nil, err
			}
			conf.IAM.AWSConfig = awsCfg
		}

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
		awsConfig := conf.IAM.AWSConfig
		if awsConfig == nil {
			awsCfg, err := awsx.LoadDefaultConfig(ctx)
			if err != nil {
				return nil, nil, errors.Wrap(err, "failed to load AWS config")
			}
			awsConfig = &awsCfg
		}
		dialector, err = NewIAMDialector(conf.DSN, awsConfig.Region, awsConfig.Credentials, WithIAMTokenEndpoint(conf.IAM.TokenEndpoint))
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
				// QueryFields explicitly lists columns in SELECT queries instead of using SELECT *.
				// Required to prevent "cached plan must not change result type" errors when pgx's
				// prepared statement cache encounters schema changes during rolling deployments.
				// Trade-off: Field removal will fail-fast (error immediately) rather than silently
				// returning zero values, requiring proper multi-phase deployment for breaking changes.
				QueryFields: true,
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

var (
	// IAMTokenFreshTTL is the duration for which the IAM token is considered fresh in the local cache.
	//
	// Rationale:
	//   - AWS RDS IAM authentication tokens are valid for up to 15 minutes, according to the official
	//     documentation.
	//   - We choose 10 minutes as the fresh window to ensure that any token served as "fresh" is well
	//     within the 15-minute validity period.
	//   - This reduces the frequency of calling auth.BuildAuthToken while still staying safely below
	//     the server-side expiration time.
	IAMTokenFreshTTL = 10 * time.Minute
	// IAMTokenStaleTTL is the additional duration during which a cached token can be served as stale
	// when the upstream token generation is temporarily failing.
	//
	// This is used together with cachex.WithServeStale(true):
	//   - When a token age is between IAMTokenFreshTTL and IAMTokenFreshTTL+IAMTokenStaleTTL, cachex
	//     considers it stale and will trigger an asynchronous refresh via auth.BuildAuthToken while
	//     still returning the stale token to callers during this window.
	//   - With IAMTokenFreshTTL=10m and IAMTokenStaleTTL=3m, the maximum age of a served token is
	//     13 minutes, which is still less than the AWS RDS IAM 15-minute validity period, so the
	//     database continues to accept the stale token.
	IAMTokenStaleTTL = 3 * time.Minute
	// IAMTokenFetchConcurrency is the number of concurrent IAM token fetches allowed when refreshing
	// tokens for the same endpoint, to avoid thundering herd under high concurrency.
	IAMTokenFetchConcurrency = 2
)

func newIAMTokenClient(region string, dbUser string, credProvider aws.CredentialsProvider) *cachex.Client[*cachex.Entry[string]] {
	upstream := cachex.UpstreamFunc[*cachex.Entry[string]](
		func(ctx context.Context, endpoint string) (*cachex.Entry[string], error) {
			var token string
			if err := backoff.RetryNotify(
				func() error {
					slog.InfoContext(ctx, "Generating new IAM token ...")
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
				return nil, errors.Wrap(err, "failed to get IAM auth token after retries")
			}

			slog.InfoContext(ctx, "Successfully got IAM Token")
			return &cachex.Entry[string]{
				Data:     token,
				CachedAt: time.Now(),
			}, nil
		},
	)

	return cachex.NewClient(
		cachex.NewSyncMap[*cachex.Entry[string]](),
		upstream,
		cachex.EntryWithTTL[string](IAMTokenFreshTTL, IAMTokenStaleTTL),
		cachex.WithServeStale[*cachex.Entry[string]](true),
		cachex.WithFetchConcurrency[*cachex.Entry[string]](IAMTokenFetchConcurrency),
		cachex.WithDoubleCheck[*cachex.Entry[string]](cachex.DoubleCheckEnabled),
	)
}

type IAMDialectorOption func(*iamDialectorOptions)

type iamDialectorOptions struct {
	tokenEndpoint    string
	connectorWrapper func(driver.Connector) driver.Connector // for testing
}

func WithIAMTokenEndpoint(endpoint string) IAMDialectorOption {
	return func(o *iamDialectorOptions) {
		o.tokenEndpoint = endpoint
	}
}

// withConnectorWrapper is an internal option for testing that wraps the underlying connector.
// This allows tests to intercept and modify connection behavior.
func withConnectorWrapper(wrapper func(driver.Connector) driver.Connector) IAMDialectorOption {
	return func(o *iamDialectorOptions) {
		o.connectorWrapper = wrapper
	}
}

func NewIAMDialector(dsn string, region string, credProvider aws.CredentialsProvider, opts ...IAMDialectorOption) (gorm.Dialector, error) {
	var options iamDialectorOptions
	for _, opt := range opts {
		opt(&options)
	}

	if dsn == "" {
		return nil, errors.New("dsn is required")
	}
	// Parse DSN first to validate format before modifying it.
	conf, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}
	// AWS RDS IAM authentication requires SSL/TLS connection.
	// Force sslmode to "require" if not explicitly set.
	// This prevents intermittent "pg_hba.conf rejects connection ... no encryption" errors
	// that can occur when sslmode is "prefer" (default) and SSL negotiation occasionally fails.
	// Note: pgx defaults to sslmode=prefer, so TLSConfig is non-nil even without explicit sslmode.
	// We must check the DSN string directly to detect if sslmode was explicitly specified.
	if !strings.Contains(strings.ToLower(dsn), "sslmode=") {
		dsn = appendSSLMode(dsn, "require")
		var err error
		conf, err = pgx.ParseConfig(dsn)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse dsn after appending sslmode")
		}
		slog.Info("IAM authentication: sslmode not specified, defaulting to 'require'")
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

	tokenEndpoint := options.tokenEndpoint
	if tokenEndpoint == "" {
		tokenEndpoint = fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	}

	tokenClient := newIAMTokenClient(region, conf.User, credProvider)
	optBeforeConnect := stdlib.OptionBeforeConnect(func(ctx context.Context, conf *pgx.ConnConfig) error {
		entry, err := tokenClient.Get(ctx, tokenEndpoint)
		if err != nil {
			return err
		}
		conf.Password = entry.Data
		return nil
	})
	tzOpts := postgresx.ConfigureTimezone(conf)
	var baseConnector driver.Connector = stdlib.GetConnector(*conf, append(tzOpts, optBeforeConnect)...)

	// Apply test wrapper if provided
	if options.connectorWrapper != nil {
		baseConnector = options.connectorWrapper(baseConnector)
	}

	// Wrap the connector with retry logic for PAM authentication failures
	retryConnector := &iamRetryConnector{
		connector:     baseConnector,
		tokenClient:   tokenClient,
		tokenEndpoint: tokenEndpoint,
	}

	conn := sql.OpenDB(retryConnector)
	return postgresx.New(postgres.Config{
		Conn: conn,
		// We are using pgx as postgresql's database/sql driver, it enables prepared statement cache by default (extended protocol)
		// PreferSimpleProtocol: true, // disables implicit prepared statement usage.
	}), nil
}

// iamRetryConnector wraps a driver.Connector to retry on PAM authentication failures
// by invalidating the IAM token cache and retrying once.
type iamRetryConnector struct {
	connector     driver.Connector
	tokenClient   *cachex.Client[*cachex.Entry[string]]
	tokenEndpoint string
}

// isPAMAuthError checks if the error is a PAM authentication failure
func isPAMAuthError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return strings.Contains(errStr, "PAM authentication failed") ||
		strings.Contains(errStr, "SQLSTATE 28000")
}

func (c *iamRetryConnector) Connect(ctx context.Context) (driver.Conn, error) {
	conn, err := c.connector.Connect(ctx)
	if err != nil && isPAMAuthError(err) {
		slog.WarnContext(ctx, "PAM authentication failed, invalidating token cache and retrying",
			"error", err)

		// Invalidate the cached token
		if delErr := c.tokenClient.Del(ctx, c.tokenEndpoint); delErr != nil {
			slog.WarnContext(ctx, "Failed to invalidate token cache", "error", delErr)
		}

		// Retry the connection once
		conn, err = c.connector.Connect(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "retry after token cache invalidation failed")
		}
		slog.InfoContext(ctx, "Successfully reconnected after token cache invalidation")
	}
	return conn, err
}

func (c *iamRetryConnector) Driver() driver.Driver {
	return c.connector.Driver()
}

// appendSSLMode appends sslmode parameter to a DSN string.
// Handles both URL format (postgres://...) and keyword format (host=... port=...).
func appendSSLMode(dsn, mode string) string {
	dsn = strings.TrimSpace(dsn)
	param := "sslmode=" + mode
	if strings.Contains(dsn, "://") {
		// URL format: postgres://user@host:port/dbname?params
		idx := strings.Index(dsn, "?")
		if idx == -1 {
			// No query params: append ?sslmode=...
			return dsn + "?" + param
		}
		// Has query params: check if it ends with ? or &
		if idx == len(dsn)-1 || dsn[len(dsn)-1] == '&' {
			// Ends with ? or &: append directly
			return dsn + param
		}
		// Has existing params: append &sslmode=...
		return dsn + "&" + param
	}
	// Keyword format: host=... port=... - append with space
	return dsn + " " + param
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
