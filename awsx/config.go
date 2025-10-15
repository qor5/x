package awsx

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/pkg/errors"
)

// Config defines AWS configuration that supports both confx and environment variables
type Config struct {
	// Region specifies the AWS region
	Region string `confx:"region" usage:"AWS region (e.g., ap-northeast-1)"`

	// AccessKeyID specifies the AWS access key ID
	AccessKeyID string `confx:"accessKeyId" usage:"AWS access key ID"`

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
	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load AWS configuration")
	}

	return &awsConfig, nil
}
