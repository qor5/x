package gormx

import (
	"testing"
	"time"

	"github.com/qor5/confx"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
)

func TestConfig(t *testing.T) {
	suite := confx.NewValidationSuite(t)
	suite.RunTests([]confx.ExpectedValidation{
		{
			Name: "valid config with password auth",
			Config: &DatabaseConfig{
				DSN:             "postgres://user:pass@localhost:5432/db",
				Debug:           true,
				Tracing:         TracingConfig{},
				MaxIdleConns:    5,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 10 * time.Minute,
				ConnMaxLifetime: 30 * time.Minute,
				AuthMethod:      AuthMethodPassword,
			},
			ExpectedErrors: nil,
		},
		{
			Name: "valid config with iam auth",
			Config: &DatabaseConfig{
				DSN:             "postgres://user:pass@localhost:5432/db",
				Debug:           false,
				Tracing:         TracingConfig{},
				MaxIdleConns:    10,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 30 * time.Minute,
				ConnMaxLifetime: 30 * time.Minute,
				AuthMethod:      AuthMethodIAM,
			},
			ExpectedErrors: nil,
		},
		{
			Name: "invalid config - missing required fields",
			Config: &DatabaseConfig{
				Debug:           true,
				Tracing:         TracingConfig{},
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
			Config: &DatabaseConfig{
				DSN:             "postgres://user:pass@localhost:5432/db",
				Debug:           true,
				Tracing:         TracingConfig{},
				MaxIdleConns:    11,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 30 * time.Minute,
				ConnMaxLifetime: 10 * time.Minute,
				AuthMethod:      AuthMethodPassword,
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "MaxIdleConns", Tag: "ltefield"},
				{Path: "ConnMaxIdleTime", Tag: "ltefield"},
			},
		},
		{
			Name: "invalid config - auth method",
			Config: &DatabaseConfig{
				DSN:             "postgres://user:pass@localhost:5432/db",
				Debug:           true,
				Tracing:         TracingConfig{},
				MaxIdleConns:    5,
				MaxOpenConns:    10,
				ConnMaxIdleTime: 10 * time.Minute,
				ConnMaxLifetime: 30 * time.Minute,
				AuthMethod:      AuthMethod("invalid"),
			},
			ExpectedErrors: []confx.ExpectedValidationError{
				{Path: "AuthMethod", Tag: "oneof"},
			},
		},
		{
			Name: "invalid config",
			Config: &DatabaseConfig{
				DSN:             "", // empty dsn
				Debug:           true,
				Tracing:         TracingConfig{},
				MaxIdleConns:    11, // maxIdleConns > maxOpenConns
				MaxOpenConns:    10,
				ConnMaxIdleTime: 30 * time.Minute, // maxIdleTime > maxLifetime
				ConnMaxLifetime: 10 * time.Minute,
				AuthMethod:      AuthMethod("not password or iam"), // invalid auth method
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
			dialector, err := NewIAMDialector(tt.args.dsn, tt.args.region, nil)
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
			got, err := NewDefaultDialector(tt.args.dsn)
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
				dialector, ok := got.(*postgres.Dialector)
				require.True(t, ok, "Expected *postgres.Dialector")
				require.Equal(t, tt.args.dsn, dialector.Config.DSN)
			}
		})
	}
}
