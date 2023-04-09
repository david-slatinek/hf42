package main

import (
	"context"
	"log"
	"main/email"
	"main/env"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func handler(ctx context.Context, s3Event events.S3Event) error {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("failed to load default config: %s", err)
		return err
	}
	s3Client := s3.NewFromConfig(sdkConfig)

	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.URLDecodedKey
		headOutput, err := s3Client.HeadObject(ctx, &s3.HeadObjectInput{
			Bucket: &bucket,
			Key:    &key,
		})
		if err != nil {
			log.Printf("error getting head of object %s/%s: %s", bucket, key, err)
			return err
		}
		log.Printf("successfully retrieved %s/%s of type %s", bucket, key, *headOutput.ContentType)

		err = email.SendEmail("https://" + bucket + ".s3." + os.Getenv("REGION") + ".amazonaws.com/" + key)
		if err != nil {
			log.Printf("error sending email: %s", err)
			return err
		}
		log.Printf("successfully sent email")
	}
	return nil
}

func main() {
	err := env.Load("env/.env")
	if err != nil {
		log.Fatalf("error loading env file: %s", err)
	}
	lambda.Start(handler)
}
