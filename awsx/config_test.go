package awsx

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	ctx := context.Background()

	t.Run("with explicit credentials", func(t *testing.T) {
		// Clear any existing environment variables first
		os.Unsetenv("AWS_REGION")
		os.Unsetenv("AWS_ACCESS_KEY_ID")
		os.Unsetenv("AWS_SECRET_ACCESS_KEY")
		os.Unsetenv("AWS_SESSION_TOKEN")

		conf := &Config{
			Region:          "ap-northeast-1",
			AccessKeyID:     "test-access-key",
			SecretAccessKey: "test-secret-key",
			SessionToken:    "test-session-token",
		}

		awsConfig, err := SetupAWSConfig(ctx, conf)
		require.NoError(t, err)
		assert.Equal(t, "ap-northeast-1", awsConfig.Region)

		// Test credentials
		creds, err := awsConfig.Credentials.Retrieve(ctx)
		require.NoError(t, err)
		assert.Equal(t, "test-access-key", creds.AccessKeyID)
		assert.Equal(t, "test-secret-key", creds.SecretAccessKey)
		assert.Equal(t, "test-session-token", creds.SessionToken)

		// Verify environment variables were set
		assert.Equal(t, "ap-northeast-1", os.Getenv("AWS_REGION"))
		assert.Equal(t, "test-access-key", os.Getenv("AWS_ACCESS_KEY_ID"))
		assert.Equal(t, "test-secret-key", os.Getenv("AWS_SECRET_ACCESS_KEY"))
		assert.Equal(t, "test-session-token", os.Getenv("AWS_SESSION_TOKEN"))
	})

	t.Run("with environment variables", func(t *testing.T) {
		// Set environment variables
		os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
		os.Setenv("AWS_REGION", "us-west-2")
		os.Setenv("AWS_ACCESS_KEY_ID", "env-access-key")
		os.Setenv("AWS_SECRET_ACCESS_KEY", "env-secret-key")
		os.Setenv("AWS_SESSION_TOKEN", "env-session-token")
		defer func() {
			os.Unsetenv("AWS_SDK_LOAD_CONFIG")
			os.Unsetenv("AWS_REGION")
			os.Unsetenv("AWS_ACCESS_KEY_ID")
			os.Unsetenv("AWS_SECRET_ACCESS_KEY")
			os.Unsetenv("AWS_SESSION_TOKEN")
		}()

		conf := &Config{}

		awsConfig, err := SetupAWSConfig(ctx, conf)
		require.NoError(t, err)

		// The region should be loaded from environment
		assert.Equal(t, "us-west-2", awsConfig.Region)
	})

	t.Run("confx config sets environment variables", func(t *testing.T) {
		// Set initial AWS native environment variables
		os.Setenv("AWS_REGION", "us-west-2")
		os.Setenv("AWS_ACCESS_KEY_ID", "original-access-key")
		os.Setenv("AWS_SECRET_ACCESS_KEY", "original-secret-key")
		defer func() {
			os.Unsetenv("AWS_REGION")
			os.Unsetenv("AWS_ACCESS_KEY_ID")
			os.Unsetenv("AWS_SECRET_ACCESS_KEY")
		}()

		// confx config should override environment variables by setting new values
		conf := &Config{
			Region:          "ap-northeast-1", // This should override AWS_REGION
			AccessKeyID:     "config-access-key",
			SecretAccessKey: "config-secret-key",
		}

		awsConfig, err := SetupAWSConfig(ctx, conf)
		require.NoError(t, err)

		// confx config should have set new environment variable values
		assert.Equal(t, "ap-northeast-1", awsConfig.Region)
		assert.Equal(t, "ap-northeast-1", os.Getenv("AWS_REGION"))
		assert.Equal(t, "config-access-key", os.Getenv("AWS_ACCESS_KEY_ID"))
		assert.Equal(t, "config-secret-key", os.Getenv("AWS_SECRET_ACCESS_KEY"))

		creds, err := awsConfig.Credentials.Retrieve(ctx)
		require.NoError(t, err)
		assert.Equal(t, "config-access-key", creds.AccessKeyID)
		assert.Equal(t, "config-secret-key", creds.SecretAccessKey)
	})

	t.Run("AWS native environment variables work when confx config is empty", func(t *testing.T) {
		// Set AWS native environment variables
		os.Setenv("AWS_REGION", "eu-central-1")
		os.Setenv("AWS_ACCESS_KEY_ID", "native-env-access-key")
		os.Setenv("AWS_SECRET_ACCESS_KEY", "native-env-secret-key")
		os.Setenv("AWS_SESSION_TOKEN", "native-env-session-token")
		defer func() {
			os.Unsetenv("AWS_REGION")
			os.Unsetenv("AWS_ACCESS_KEY_ID")
			os.Unsetenv("AWS_SECRET_ACCESS_KEY")
			os.Unsetenv("AWS_SESSION_TOKEN")
		}()

		// Empty confx config should allow AWS native env vars to work
		conf := &Config{}

		awsConfig, err := SetupAWSConfig(ctx, conf)
		require.NoError(t, err)

		// Should use AWS native environment variables
		assert.Equal(t, "eu-central-1", awsConfig.Region)

		creds, err := awsConfig.Credentials.Retrieve(ctx)
		require.NoError(t, err)
		assert.Equal(t, "native-env-access-key", creds.AccessKeyID)
		assert.Equal(t, "native-env-secret-key", creds.SecretAccessKey)
		assert.Equal(t, "native-env-session-token", creds.SessionToken)
	})
}
