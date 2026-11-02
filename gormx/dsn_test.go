package gormx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendSSLMode(t *testing.T) {
	tests := []struct {
		name     string
		dsn      string
		mode     string
		expected string
	}{
		// URL format tests
		{
			name:     "URL format without query params",
			dsn:      "postgres://user@host:5432/dbname",
			mode:     "require",
			expected: "postgres://user@host:5432/dbname?sslmode=require",
		},
		{
			name:     "URL format with existing query params",
			dsn:      "postgres://user@host:5432/dbname?connect_timeout=10",
			mode:     "require",
			expected: "postgres://user@host:5432/dbname?connect_timeout=10&sslmode=require",
		},
		{
			name:     "URL format with trailing ?",
			dsn:      "postgres://user@host:5432/dbname?",
			mode:     "require",
			expected: "postgres://user@host:5432/dbname?sslmode=require",
		},
		{
			name:     "URL format with trailing &",
			dsn:      "postgres://user@host:5432/dbname?foo=bar&",
			mode:     "require",
			expected: "postgres://user@host:5432/dbname?foo=bar&sslmode=require",
		},
		{
			name:     "URL format with multiple params",
			dsn:      "postgres://user@host:5432/dbname?connect_timeout=10&application_name=myapp",
			mode:     "verify-full",
			expected: "postgres://user@host:5432/dbname?connect_timeout=10&application_name=myapp&sslmode=verify-full",
		},
		// Keyword format tests
		{
			name:     "keyword format basic",
			dsn:      "host=localhost port=5432 user=test dbname=testdb",
			mode:     "require",
			expected: "host=localhost port=5432 user=test dbname=testdb sslmode=require",
		},
		{
			name:     "keyword format with existing params",
			dsn:      "host=localhost port=5432 user=test dbname=testdb connect_timeout=10",
			mode:     "require",
			expected: "host=localhost port=5432 user=test dbname=testdb connect_timeout=10 sslmode=require",
		},
		{
			name:     "keyword format minimal",
			dsn:      "host=localhost",
			mode:     "require",
			expected: "host=localhost sslmode=require",
		},
		// TrimSpace tests
		{
			name:     "URL format with leading space",
			dsn:      "  postgres://user@host:5432/dbname",
			mode:     "require",
			expected: "postgres://user@host:5432/dbname?sslmode=require",
		},
		{
			name:     "URL format with trailing space",
			dsn:      "postgres://user@host:5432/dbname  ",
			mode:     "require",
			expected: "postgres://user@host:5432/dbname?sslmode=require",
		},
		{
			name:     "keyword format with leading and trailing spaces",
			dsn:      "  host=localhost port=5432  ",
			mode:     "require",
			expected: "host=localhost port=5432 sslmode=require",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := appendSSLMode(tt.dsn, tt.mode)
			assert.Equal(t, tt.expected, result)
		})
	}
}
