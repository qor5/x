package s3_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/jinzhu/configor"
	"github.com/qor5/x/v3/oss/s3"
	"github.com/qor5/x/v3/oss/tests"
)

type Config struct {
	AccessID  string `env:"AWS_ACCESS_KEY_ID"`
	AccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
	Region    string `env:"AWS_REGION"`
	Bucket    string `env:"AWS_BUCKET"`
	Endpoint  string `env:"AWS_ENDPOINT"`
}

var (
	client *s3.Client
	config = Config{}
)

func init() {
	configor.Load(&config)

	client = s3.New(&s3.Config{AccessID: config.AccessID, AccessKey: config.AccessKey, Region: config.Region, Bucket: config.Bucket, Endpoint: config.Endpoint})
}

func TestAll(t *testing.T) {
	if len(config.AccessID) == 0 {
		fmt.Println("No aws configuration")
		t.Skip(`skip because of no config: `)
	}

	fmt.Println("testing S3 with public ACL")
	tests.TestAll(client, t)

	fmt.Println("testing S3 with private ACL")
	privateClient := s3.New(&s3.Config{AccessID: config.AccessID, AccessKey: config.AccessKey, Region: config.Region, Bucket: config.Bucket, ACL: string(types.BucketCannedACLPrivate), Endpoint: config.Endpoint})
	tests.TestAll(privateClient, t)

	fmt.Println("testing S3 with AuthenticatedRead ACL")
	authenticatedReadClient := s3.New(&s3.Config{AccessID: config.AccessID, AccessKey: config.AccessKey, Region: config.Region, Bucket: config.Bucket, ACL: string(types.BucketCannedACLAuthenticatedRead), Endpoint: config.Endpoint})
	tests.TestAll(authenticatedReadClient, t)
}

func TestToS3Key(t *testing.T) {
	urlMap := map[string]string{
		"https://mybucket.s3.amazonaws.com/myobject.ext": "myobject.ext",
		"https://qor-example.com/myobject.ext":           "myobject.ext",
		"//mybucket.s3.amazonaws.com/myobject.ext":       "myobject.ext",
		"http://mybucket.s3.amazonaws.com/myobject.ext":  "myobject.ext",
		"myobject.ext": "myobject.ext",
	}

	for url, path := range urlMap {
		if client.ToS3Key(url) != path {
			t.Errorf("%v's s3 key should be %v, but got %v", url, path, client.ToS3Key(url))
		}
	}
}

func TestToS3KeyWithS3ForcePathStyle(t *testing.T) {
	urlMap := map[string]string{
		"https://s3.amazonaws.com/mybucket/myobject.ext": "myobject.ext",
		"https://qor-example.com/myobject.ext":           "myobject.ext",
		"//s3.amazonaws.com/mybucket/myobject.ext":       "myobject.ext",
		"http://s3.amazonaws.com/mybucket/myobject.ext":  "myobject.ext",
		"/mybucket/myobject.ext":                         "myobject.ext",
		"myobject.ext":                                   "myobject.ext",
	}

	client := s3.New(&s3.Config{AccessID: config.AccessID, AccessKey: config.AccessKey, Region: config.Region, Bucket: "mybucket", S3ForcePathStyle: true, Endpoint: config.Endpoint})

	for url, path := range urlMap {
		if client.ToS3Key(url) != path {
			t.Errorf("%v's s3 key should be %v, but got %v", url, path, client.ToS3Key(url))
		}
	}
}