package gormx

import (
	"context"
	"database/sql/driver"
	"errors"
	"sync/atomic"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestIsPAMAuthError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "PAM authentication failed error",
			err:      errors.New("server error: FATAL: PAM authentication failed for user \"test_user\" (SQLSTATE 28000)"),
			expected: true,
		},
		{
			name:     "SQLSTATE 28000 error",
			err:      errors.New("connection failed: SQLSTATE 28000"),
			expected: true,
		},
		{
			name:     "other error",
			err:      errors.New("connection refused"),
			expected: false,
		},
		{
			name:     "wrapped PAM error",
			err:      errors.New("failed to connect: PAM authentication failed"),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPAMAuthError(tt.err)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestNewIAMDialector_PAMRetryLogic(t *testing.T) {
	dsn := "host=localhost port=5432 user=test_user dbname=test_db sslmode=require"
	region := "ap-northeast-1"
	credProvider := &mockCredentialsProvider{}

	t.Run("PAM error on first connect triggers retry and succeeds", func(t *testing.T) {
		var connectCount int32

		dialector, err := NewIAMDialector(dsn, region, credProvider,
			withConnectorWrapper(func(c driver.Connector) driver.Connector {
				return &pamErrorThenMockConnector{
					connectCount: &connectCount,
				}
			}),
		)
		require.NoError(t, err)
		require.NotNil(t, dialector)

		// Try to open the database - first attempt fails with PAM error, retry succeeds
		_, err = gorm.Open(dialector, &gorm.Config{})
		require.NoError(t, err)

		require.Equal(t, int32(2), atomic.LoadInt32(&connectCount),
			"should retry once after PAM authentication failure")
	})

	t.Run("non-PAM error does not trigger retry", func(t *testing.T) {
		var connectCount int32

		dialector, err := NewIAMDialector(dsn, region, credProvider,
			withConnectorWrapper(func(c driver.Connector) driver.Connector {
				return &alwaysErrorConnector{
					connectCount:  &connectCount,
					errorToReturn: errors.New("connection refused"),
				}
			}),
		)
		require.NoError(t, err)
		require.NotNil(t, dialector)

		_, err = gorm.Open(dialector, &gorm.Config{})
		require.Contains(t, err.Error(), "connection refused")
		require.Equal(t, int32(1), atomic.LoadInt32(&connectCount),
			"should not retry on non-PAM error")
	})

	t.Run("PAM error on retry still fails", func(t *testing.T) {
		var connectCount int32

		// Create a connector that always returns PAM error
		dialector, err := NewIAMDialector(dsn, region, credProvider,
			withConnectorWrapper(func(c driver.Connector) driver.Connector {
				return &alwaysPAMErrorConnector{
					connectCount: &connectCount,
				}
			}),
		)
		require.NoError(t, err)
		require.NotNil(t, dialector)

		_, err = gorm.Open(dialector, &gorm.Config{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "retry after token cache invalidation failed")
		require.Equal(t, int32(2), atomic.LoadInt32(&connectCount),
			"should retry exactly once even if retry also fails")
	})
}

// mockCredentialsProvider implements aws.CredentialsProvider for testing
type mockCredentialsProvider struct{}

func (m *mockCredentialsProvider) Retrieve(ctx context.Context) (aws.Credentials, error) {
	return aws.Credentials{
		AccessKeyID:     "test-access-key",
		SecretAccessKey: "test-secret-key",
	}, nil
}

// mockConn implements driver.Conn for testing
type mockConn struct{}

func (c *mockConn) Prepare(query string) (driver.Stmt, error) { return nil, nil }
func (c *mockConn) Close() error                              { return nil }
func (c *mockConn) Begin() (driver.Tx, error)                 { return nil, nil }
func (c *mockConn) Ping(ctx context.Context) error            { return nil }

// pamErrorThenMockConnector returns PAM error on first call, then returns a mock connection
type pamErrorThenMockConnector struct {
	connectCount *int32
}

func (c *pamErrorThenMockConnector) Connect(ctx context.Context) (driver.Conn, error) {
	count := atomic.AddInt32(c.connectCount, 1)
	if count == 1 {
		return nil, errors.New("FATAL: PAM authentication failed for user \"test_user\" (SQLSTATE 28000)")
	}
	return &mockConn{}, nil
}

func (c *pamErrorThenMockConnector) Driver() driver.Driver {
	return nil
}

// alwaysErrorConnector always returns the specified error
type alwaysErrorConnector struct {
	connectCount  *int32
	errorToReturn error
}

func (c *alwaysErrorConnector) Connect(ctx context.Context) (driver.Conn, error) {
	atomic.AddInt32(c.connectCount, 1)
	return nil, c.errorToReturn
}

func (c *alwaysErrorConnector) Driver() driver.Driver {
	return nil
}

// alwaysPAMErrorConnector always returns a PAM authentication error
type alwaysPAMErrorConnector struct {
	connectCount *int32
}

func (c *alwaysPAMErrorConnector) Connect(ctx context.Context) (driver.Conn, error) {
	atomic.AddInt32(c.connectCount, 1)
	return nil, errors.New("FATAL: PAM authentication failed for user \"test_user\" (SQLSTATE 28000)")
}

func (c *alwaysPAMErrorConnector) Driver() driver.Driver {
	return nil
}
