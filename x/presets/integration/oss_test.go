package integration_test

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"strings"
	"testing"

	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/qor/media"
	"github.com/qor/media/oss"
	"github.com/qor/oss/s3"
)

type TestOSS struct {
	ID    int
	Image oss.OSS
}

func TestOSSDemo(t *testing.T) {
	oss.Storage = s3.New(&s3.Config{
		AccessID:  "access_id",
		AccessKey: "access_key",
		Region:    "region",
		Bucket:    "bucket",
		Endpoint:  "cdn.getqor.com",
		ACL:       awss3.BucketCannedACLPublicRead,
	})

	db := ConnectDB()
	db.AutoMigrate(&TestOSS{})
	media.RegisterCallbacks(db)
	//media_library.RegisterCallbacks(db)

	r := multipart.NewReader(strings.NewReader(`
----123123123
Content-Disposition: form-data; name="TestName"; filename="hello.txt"
Content-Type: application/octet-stream

Hello Content
----123123123
`), "--123123123")

	form, err := r.ReadForm(10 << 20)
	if err != nil {
		panic(err)
	}

	_ = sql.NullBool{}

	img := oss.OSS{}
	img.FileHeader = form.File["TestName"][0]
	img.FileName = "hello.txt"
	fmt.Println("img.FileHeader", img.FileHeader)

	err = db.Create(&TestOSS{
		Image: img,
	}).Error
	if err != nil {
		panic(err)
	}
}
