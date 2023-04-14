package file

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func UploadFile(orderID string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:                        aws.String(os.Getenv("REGION")),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		return err
	}

	f, err := os.Open("invoice.pdf")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Printf("error with closing file: %s\n", err)
		}

		if err := os.Remove("invoice.pdf"); err != nil {
			fmt.Printf("error with removing file: %s\n", err)
		}
	}(f)

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(orderID + ".pdf"),
		ACL:    aws.String("public-read"),
		Body:   f,
	})
	return err
}
