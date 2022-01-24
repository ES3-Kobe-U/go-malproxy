package service

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var AWS_S3_REGION = "ap-northeast-1"
var AWS_S3_BUCKET = "go-malproxy"
var AWS_S3_ACCESS_KEY = "********"
var AWS_S3_SECRET_ACCESS_KEY = "********"
var KEY = "********"

func TestUploadFolder(t *testing.T) {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(AWS_S3_ACCESS_KEY, AWS_S3_SECRET_ACCESS_KEY, ""),
		Region:      aws.String(AWS_S3_REGION),
	}))

	uploader := s3manager.NewUploader(sess)
	f, err := os.Open("./example/example.pdf") // ここはフロントから飛ばしてきたデータにする。
	if err != nil {
		log.Fatal(err)
	}

	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(KEY),
		Body:   f,
	})

	if err != nil {
		fmt.Println(res)
		if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
			log.Fatal(err)
		} else {
			log.Fatal(err)
		}
	}
}
