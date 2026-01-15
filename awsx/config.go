package awsx

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/pkg/errors"
)

// Config defines AWS configuration that supports both confx and environment variables
type Config struct {
	// Region specifies the AWS region
	Region string `confx:"region" usage:"AWS region (e.g., ap-northeast-1)"`

	// AccessKeyID specifies the AWS access key ID
	AccessKeyID string `confx:"accessKeyID" usage:"AWS access key ID"`

	// SecretAccessKey specifies the AWS secret access key
	SecretAccessKey string `confx:"secretAccessKey" usage:"AWS secret access key"`

	// SessionToken specifies the AWS session token for temporary credentials
	SessionToken string `confx:"sessionToken" usage:"AWS session token for temporary credentials"`
}

func SetupAWSConfig(ctx context.Context, conf *Config) (*aws.Config, error) {
	// Set environment variables from confx config if provided
	if conf.Region != "" {
		os.Setenv("AWS_REGION", conf.Region)
	}
	if conf.AccessKeyID != "" {
		os.Setenv("AWS_ACCESS_KEY_ID", conf.AccessKeyID)
	}
	if conf.SecretAccessKey != "" {
		os.Setenv("AWS_SECRET_ACCESS_KEY", conf.SecretAccessKey)
	}
	if conf.SessionToken != "" {
		os.Setenv("AWS_SESSION_TOKEN", conf.SessionToken)
	}

	// Load AWS configuration using standard AWS SDK methods
	// This will automatically pick up the environment variables we just set
	awsConfig, err := LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &awsConfig, nil
}

// DefaultCredsExpiryWindow is the default expiry window for AWS credentials cache.
// DefaultCredsExpiryWindowJitterFrac is the jitter fraction applied to the expiry window.
//
// These values are configured so that:
//
//	DefaultCredsExpiryWindow * DefaultCredsExpiryWindowJitterFrac = 15 minutes (minimum refresh window)
//
// Rationale:
//   - AWS RDS IAM authentication tokens have a maximum validity of 15 minutes.
//   - If the underlying AWS credentials (e.g., from IRSA, AssumeRole, OIDC) have less than
//     15 minutes remaining, the generated IAM token's effective validity will be shortened
//     to match the credential expiration.
//   - By setting a 15-minute minimum refresh window for credentials, we ensure that
//     credentials are always refreshed before they could cause IAM token generation to
//     produce short-lived or already-expired tokens.
//   - The jitter (0.75) randomizes the actual refresh time between 5-20 minutes before
//     expiration, preventing thundering herd issues in distributed systems.
//
// AWS temporary credentials typical validity periods:
//   - IRSA (EKS): 1 hour (default), up to 12 hours
//   - AssumeRole: 1 hour (default), up to 12 hours
//   - OIDC: varies by configuration, typically 1 hour
//   - EC2 Instance Profile: typically 6 hours
var (
	DefaultCredsExpiryWindow           = 20 * time.Minute
	DefaultCredsExpiryWindowJitterFrac = 0.75
)

// LoadDefaultConfig loads the default AWS configuration using the AWS SDK
func LoadDefaultConfig(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
	optFns = append([]func(*config.LoadOptions) error{config.WithCredentialsCacheOptions(func(o *aws.CredentialsCacheOptions) {
		o.ExpiryWindow = DefaultCredsExpiryWindow
		o.ExpiryWindowJitterFrac = DefaultCredsExpiryWindowJitterFrac
	})}, optFns...)
	conf, err := config.LoadDefaultConfig(ctx, optFns...)
	if err != nil {
		return aws.Config{}, errors.Wrap(err, "failed to load AWS configuration")
	}
	return conf, nil
}
