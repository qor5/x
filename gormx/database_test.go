package gormx_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/qor5/confx"
	"github.com/qor5/x/v3/awsx"
	"github.com/qor5/x/v3/gormx"
	"github.com/qor5/x/v3/gormx/postgresx"
	"github.com/stretchr/testify/require"
	"github.com/theplant/inject/lifecycle"
	"gorm.io/gorm"
)

func TestConfig(t *testing.T) {
	suite := confx.NewValidationSuite(t)
	suite.RunTests([]confx.ExpectedValidation{
		{
			Name: "valid config with password auth",
			Config: &gormx.DatabaseConfig{
				DSN:             "postgres://user:pass@localhost:5432/db",
				Debug:           true,
				Tracing:         gormx.TracingConfig{},
				MaxIdleConns:    5,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 10 * time.Minute,
				ConnMaxLifetime: 30 * time.Minute,
				AuthMethod:      gormx.AuthMethodPassword,
			},
			ExpectedErrors: nil,
		},
		{
			Name: "valid config with iam auth",
			Config: &gormx.DatabaseConfig{
				DSN:             "postgres://user@localhost:5432/db",
				Debug:           false,
				Tracing:         gormx.TracingConfig{},
				MaxIdleConns:    10,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 30 * time.Minute,
				ConnMaxLifetime: 30 * time.Minute,
				AuthMethod:      gormx.AuthMethodIAM,
			},
			ExpectedErrors: nil,
		},
		{
			Name: "valid config with iam auth and custom token endpoint",
			Config: &gormx.DatabaseConfig{
				DSN:             "postgres://user@proxy-host:5432/db",
				Debug:           false,
				Tracing:         gormx.TracingConfig{},
				MaxIdleConns:    10,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 30 * time.Minute,
				ConnMaxLifetime: 30 * time.Minute,
				AuthMethod:      gormx.AuthMethodIAM,
				IAM: gormx.IAMDialectorConfig{
					TokenEndpoint: "rds-cluster.xxx.ap-northeast-1.rds.amazonaws.com:5432",
				},
			},
			ExpectedErrors: nil,
		},
		{
			Name: "invalid config - missing required fields",
			Config: &gormx.DatabaseConfig{
				Debug:           true,
				Tracing:         gormx.TracingConfig{},
				MaxIdleConns:    10,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 10 * time.Minute,
				ConnMaxLifetime: 30 * time.Minute,
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "DSN", Tag: "required"},
				{Path: "AuthMethod", Tag: "required"},
			},
		},
		{
			Name: "invalid config - connection constraints",
			Config: &gormx.DatabaseConfig{
				DSN:             "postgres://user:pass@localhost:5432/db",
				Debug:           true,
				Tracing:         gormx.TracingConfig{},
				MaxIdleConns:    11,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 30 * time.Minute,
				ConnMaxLifetime: 10 * time.Minute,
				AuthMethod:      gormx.AuthMethodPassword,
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "MaxIdleConns", Tag: "ltefield"},
				{Path: "ConnMaxIdleTime", Tag: "ltefield"},
			},
		},
		{
			Name: "invalid config - auth method",
			Config: &gormx.DatabaseConfig{
				DSN:             "postgres://user:pass@localhost:5432/db",
				Debug:           true,
				Tracing:         gormx.TracingConfig{},
				MaxIdleConns:    5,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 10 * time.Minute,
				ConnMaxLifetime: 30 * time.Minute,
				AuthMethod:      gormx.AuthMethod("invalid"),
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "AuthMethod", Tag: "oneof"},
			},
		},
		{
			Name: "invalid config",
			Config: &gormx.DatabaseConfig{
				DSN:             "", // empty dsn
				Debug:           true,
				Tracing:         gormx.TracingConfig{},
				MaxIdleConns:    11, // maxIdleConns > maxOpenConns
				MaxOpenConns:    10,
				ConnMaxIdleTime: 30 * time.Minute, // maxIdleTime > maxLifetime
				ConnMaxLifetime: 10 * time.Minute,
				AuthMethod:      gormx.AuthMethod("not password or iam"), // invalid auth method
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "DSN", Tag: "required"},
				{Path: "AuthMethod", Tag: "oneof"},
				{Path: "MaxIdleConns", Tag: "ltefield"},
				{Path: "ConnMaxIdleTime", Tag: "ltefield"},
			},
		},
	})
}

func TestNewIAMDialector(t *testing.T) {
	type args struct {
		dsn    string
		region string
	}
	tests := []struct {
		name        string
		args        args
		errContains string
	}{
		{
			name: "empty dsn",
			args: args{
				dsn: "",
			},
			errContains: "dsn is required",
		},
		{
			name: "empty region",
			args: args{
				dsn:    "postgres://postgres@localhost:5432/dbname?sslmode=disable",
				region: "",
			},
			errContains: "region is required",
		},
		{
			name: "invalid dsn",
			args: args{
				dsn:    "invalid-dsn",
				region: "us-west-2",
			},
			errContains: "failed to parse dsn",
		},
		{
			name: "invalid port",
			args: args{
				dsn:    "postgres://user:pass@hostname:invalid/dbname",
				region: "us-west-2",
			},
			errContains: "invalid port",
		},
		{
			name: "missing user",
			args: args{
				dsn:    "postgres://:pass@hostname:5432/dbname",
				region: "us-west-2",
			},
			errContains: "user is required",
		},
		{
			name: "password provided",
			args: args{
				dsn:    "postgres://user:pass@hostname:5432/dbname",
				region: "us-west-2",
			},
			errContains: "password should not be provided",
		},
		{
			name: "credentials provider is required",
			args: args{
				dsn:    "postgres://user@hostname:5432/dbname",
				region: "us-west-2",
			},
			errContains: "credentials provider is required",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dialector, err := gormx.NewIAMDialector(tt.args.dsn, tt.args.region, nil)
			if tt.errContains != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.errContains)
				require.Nil(t, dialector)
			} else {
				require.NoError(t, err)
				require.NotNil(t, dialector)
			}
		})
	}
}

func TestNewDefaultDialector(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name: "empty dsn",
			args: args{
				dsn: "",
			},
			wantErr: true,
			errMsg:  "dsn is required",
		},
		{
			name: "valid dsn",
			args: args{
				dsn: "postgres://user:password@localhost:5432/dbname?sslmode=disable",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gormx.NewDefaultDialector(tt.args.dsn)
			if tt.wantErr {
				require.Error(t, err)
				if tt.errMsg != "" {
					require.Contains(t, err.Error(), tt.errMsg)
				}
				require.Nil(t, got)
			} else {
				require.NoError(t, err)
				require.NotNil(t, got)

				// Additional checks for the returned dialector.
				dialector, ok := got.(*postgresx.Dialector)
				require.True(t, ok, "Expected *postgresx.Dialector")
				require.Equal(t, tt.args.dsn, dialector.Config.DSN)
			}
		})
	}
}

func TestAuthMethodPassword(t *testing.T) {
	var result int
	require.NoError(t, suite.DB().Raw("SELECT 1").Scan(&result).Error)
	require.Equal(t, 1, result)
}

/*
eval $(oidc2aws -login -alias xxx-test --env) && \
TEST_IAM_DSN="host=xxx-test-pgproxy port=5432 user=xxx_developer dbname=xxx_test sslmode=require" \
TEST_IAM_TOKEN_ENDPOINT="xxx-test.cluster-xxx.ap-northeast-1.rds.amazonaws.com:5432" \
go test -v -run TestAuthMethodIAM ./gormx/
*/
func TestAuthMethodIAM(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping IAM auth test in CI environment")
	}

	dsn := os.Getenv("TEST_IAM_DSN")
	if dsn == "" {
		t.Skip("TEST_IAM_DSN environment variable not set")
	}

	tokenEndpoint := os.Getenv("TEST_IAM_TOKEN_ENDPOINT")
	ctx := context.Background()

	baseProviders := []any{
		func(ctx context.Context, lc *lifecycle.Lifecycle) (*gorm.DB, error) {
			conf := &gormx.DatabaseConfig{
				DSN:        dsn,
				AuthMethod: gormx.AuthMethodIAM,
				IAM: gormx.IAMDialectorConfig{
					TokenEndpoint: tokenEndpoint,
				},
			}
			return gormx.SetupDatabaseFactory("database")(ctx, lc, conf)
		},
	}

	t.Run("without awsx providers", func(t *testing.T) {
		lc := lifecycle.New()
		require.NoError(t, lc.Provide(baseProviders...))

		var db *gorm.DB
		require.NoError(t, lc.ResolveContext(ctx, &db))

		var result int
		require.NoError(t, db.Raw("SELECT 1").Scan(&result).Error)
		require.Equal(t, 1, result)
	})

	t.Run("with awsx providers", func(t *testing.T) {
		lc := lifecycle.New()
		require.NoError(t, lc.Provide(baseProviders...))
		require.NoError(t, lc.Provide(
			awsx.SetupAWSConfig,
			func() *awsx.Config {
				return &awsx.Config{Region: "ap-northeast-1"}
			},
		))

		var db *gorm.DB
		require.NoError(t, lc.ResolveContext(ctx, &db))

		var result int
		require.NoError(t, db.Raw("SELECT 1").Scan(&result).Error)
		require.Equal(t, 1, result)
	})

	t.Run("with invalid aws credentials", func(t *testing.T) {
		lc := lifecycle.New()
		require.NoError(t, lc.Provide(baseProviders...))
		require.NoError(t, lc.Provide(
			awsx.SetupAWSConfig,
			func() *awsx.Config {
				return &awsx.Config{
					Region:          "ap-northeast-1",
					AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
					SecretAccessKey: "invalid-secret-key",
				}
			},
		))

		var db *gorm.DB
		err := lc.ResolveContext(ctx, &db)
		require.Error(t, err)
		t.Logf("Expected error with invalid credentials: %v", err)
	})
}
