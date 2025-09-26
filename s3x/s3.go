package s3x

import (
	"cmp"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/qor5/x/v3/oss"
	"github.com/qor5/x/v3/oss/s3"
)

type Config struct {
	AccessID  string `confx:"accessID" usage:"AWS access ID, if not provided, will use AWS default config"`
	AccessKey string `confx:"accessKey" usage:"AWS access key, if not provided, will use AWS default config"`
	Region    string `confx:"region" usage:"AWS region" validate:"required"`
	Bucket    string `confx:"bucket" usage:"AWS bucket" validate:"required"`
	Endpoint  string `confx:"endpoint" usage:"AWS endpoint" validate:"required"`
	ACL       string `confx:"-"`
}

func SetupClient(conf *Config) oss.StorageInterface {
	return s3.New(&s3.Config{
		AccessID:  conf.AccessID,
		AccessKey: conf.AccessKey,
		Bucket:    conf.Bucket,
		Region:    conf.Region,
		ACL:       cmp.Or(conf.ACL, string(types.ObjectCannedACLBucketOwnerFullControl)),
		Endpoint:  conf.Endpoint,
	})
}
