package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/qor5/x/v3/oss"
	"github.com/samber/lo"
)

// Client S3 storage
type Client struct {
	S3     *s3.Client
	Config *Config
}

// Config S3 client config
type Config struct {
	AccessID         string
	AccessKey        string
	Region           string
	Bucket           string
	SessionToken     string
	ACL              string
	Endpoint         string
	S3Endpoint       string
	S3ForcePathStyle bool
	CacheControl     string

	AWSConfig *aws.Config

	RoleARN string
}

// New initialize S3 storage
func New(config *Config) *Client {
	if config.ACL == "" {
		config.ACL = "public-read"
	}

	client := &Client{Config: config}

	var optFns []func(o *s3.Options)

	optFns = append(optFns, func(o *s3.Options) {
		o.Region = config.Region
		o.UsePathStyle = config.S3ForcePathStyle
	})
	if config.S3Endpoint != "" {
		optFns = append(optFns, func(o *s3.Options) {
			o.BaseEndpoint = lo.ToPtr(config.S3Endpoint)
		})
	}

	if config.RoleARN != "" {
		awscfg := lo.Must(awsconfig.LoadDefaultConfig(context.Background()))
		creds := stscreds.NewAssumeRoleProvider(sts.NewFromConfig(awscfg), config.RoleARN)
		optFns = append(optFns, func(o *s3.Options) {
			o.Credentials = creds
		})

		client.S3 = s3.NewFromConfig(awscfg, optFns...)
		return client
	}

	if config.AWSConfig != nil {
		client.S3 = s3.NewFromConfig(*config.AWSConfig, optFns...)
	} else if config.AccessID == "" && config.AccessKey == "" {
		// use aws default Credentials
		awsCfg := lo.Must(awsconfig.LoadDefaultConfig(context.Background()))
		client.S3 = s3.NewFromConfig(awsCfg, optFns...)
	} else {
		creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(config.AccessID, config.AccessKey, config.SessionToken))
		if _, err := creds.Retrieve(context.Background()); err == nil {
			optFns = append(optFns, func(o *s3.Options) {
				o.Credentials = creds
			})
			awsCfg := lo.Must(awsconfig.LoadDefaultConfig(context.Background()))
			client.S3 = s3.NewFromConfig(awsCfg, optFns...)
		}
	}

	return client
}

// Get receive file with given path
func (client Client) Get(ctx context.Context, path string) (file *os.File, err error) {
	readCloser, err := client.GetStream(ctx, path)

	ext := filepath.Ext(path)
	pattern := fmt.Sprintf("s3*%s", ext)

	if err == nil {
		if file, err = os.CreateTemp("/tmp", pattern); err == nil {
			defer readCloser.Close()
			_, err = io.Copy(file, readCloser)
			file.Seek(0, 0)
		}
	}

	return file, err
}

// GetStream get file as stream
func (client Client) GetStream(ctx context.Context, path string) (io.ReadCloser, error) {
	getResponse, err := client.S3.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(client.Config.Bucket),
		Key:    aws.String(client.ToS3Key(path)),
	})
	if err != nil {
		return nil, err
	}

	return getResponse.Body, err
}

// Put store a reader into given path
func (client Client) Put(ctx context.Context, urlPath string, reader io.Reader) (*oss.Object, error) {
	if seeker, ok := reader.(io.ReadSeeker); ok {
		seeker.Seek(0, 0)
	}

	urlPath = client.ToS3Key(urlPath)
	buffer, err := io.ReadAll(reader)

	fileType := mime.TypeByExtension(path.Ext(urlPath))
	if fileType == "" {
		fileType = http.DetectContentType(buffer)
	}

	params := &s3.PutObjectInput{
		Bucket:        aws.String(client.Config.Bucket), // required
		Key:           aws.String(urlPath),              // required
		ACL:           types.ObjectCannedACL(client.Config.ACL),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(int64(len(buffer))),
		ContentType:   aws.String(fileType),
	}
	if client.Config.CacheControl != "" {
		params.CacheControl = aws.String(client.Config.CacheControl)
	}

	_, err = client.S3.PutObject(ctx, params)

	now := time.Now()
	return &oss.Object{
		Path:             urlPath,
		Name:             filepath.Base(urlPath),
		LastModified:     &now,
		StorageInterface: client,
	}, err
}

// Delete delete file
func (client Client) Delete(ctx context.Context, path string) error {
	_, err := client.S3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(client.Config.Bucket),
		Key:    aws.String(client.ToS3Key(path)),
	})
	return err
}

// DeleteObjects delete files in bulk
func (client Client) DeleteObjects(ctx context.Context, paths []string) (err error) {
	var objs []types.ObjectIdentifier
	for _, v := range paths {
		var obj types.ObjectIdentifier
		obj.Key = aws.String(strings.TrimPrefix(client.ToS3Key(v), "/"))
		objs = append(objs, obj)
	}
	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(client.Config.Bucket),
		Delete: &types.Delete{
			Objects: objs,
		},
	}

	_, err = client.S3.DeleteObjects(ctx, input)
	if err != nil {
		return
	}
	return
}

// List list all objects under current path
func (client Client) List(ctx context.Context, path string) ([]*oss.Object, error) {
	var objects []*oss.Object
	var prefix string

	if path != "" {
		prefix = strings.Trim(path, "/") + "/"
	}

	listObjectsResponse, err := client.S3.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(client.Config.Bucket),
		Prefix: aws.String(prefix),
	})

	if err == nil {
		for _, content := range listObjectsResponse.Contents {
			objects = append(objects, &oss.Object{
				Path:             "/" + client.ToS3Key(*content.Key),
				Name:             filepath.Base(*content.Key),
				LastModified:     content.LastModified,
				StorageInterface: client,
			})
		}
	}

	return objects, err
}

// GetEndpoint get endpoint, FileSystem's endpoint is /
func (client Client) GetEndpoint(ctx context.Context) string {
	if client.Config.Endpoint != "" {
		return client.Config.Endpoint
	}

	endpoint := client.getS3Endpoint(ctx)
	for _, prefix := range []string{"https://", "http://"} {
		endpoint = strings.TrimPrefix(endpoint, prefix)
	}

	return client.Config.Bucket + "." + endpoint
}

var urlRegexp = regexp.MustCompile(`(https?:)?//((\w+).)+(\w+)/`)

// ToS3Key process path to s3 key
func (client Client) ToS3Key(urlPath string) string {
	if urlRegexp.MatchString(urlPath) {
		if u, err := url.Parse(urlPath); err == nil {
			if client.Config.S3ForcePathStyle { // First part of path will be bucket name
				return strings.TrimPrefix(strings.TrimPrefix(u.Path, "/"+client.Config.Bucket), "/")
			}
			return strings.TrimPrefix(u.Path, "/")
		}
	}

	if client.Config.S3ForcePathStyle { // First part of path will be bucket name
		return strings.TrimPrefix(urlPath, "/"+client.Config.Bucket+"/")
	}
	return strings.TrimPrefix(urlPath, "/")
}

// GetURL get public accessible URL
func (client Client) GetURL(ctx context.Context, path string) (url string, err error) {
	if client.Config.S3Endpoint == "" {
		if client.Config.ACL == "private" || client.Config.ACL == "authenticated-read" {
			presignReq, err := s3.NewPresignClient(client.S3).PresignGetObject(ctx, &s3.GetObjectInput{
				Bucket:                     aws.String(client.Config.Bucket),
				Key:                        aws.String(client.ToS3Key(path)),
				ResponseContentDisposition: aws.String(fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(path))),
			}, func(po *s3.PresignOptions) {
				po.Expires = 1 * time.Hour
			})
			if err != nil {
				return "", err
			}
			return presignReq.URL, nil
		}
	}

	return path, nil
}

// Copy copy s3 file from "from" to "to"
func (client Client) Copy(ctx context.Context, from, to string) (err error) {
	_, err = client.S3.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(client.Config.Bucket),
		CopySource: aws.String(from),
		Key:        aws.String(to),
	})
	return
}

func (client Client) getS3Endpoint(ctx context.Context) string {
	endpoint := client.Config.S3Endpoint
	if endpoint == "" {
		re := lo.Must(client.S3.Options().EndpointResolverV2.ResolveEndpoint(ctx, s3.EndpointParameters{
			Region: lo.ToPtr(client.Config.Region),
		}))
		endpoint = re.URI.String()
	}
	return endpoint
}
