package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
	"log"
)

type BucketBasics struct {
	S3Client *s3.Client
}

var BucketBase = BucketBasics{}

func Init() {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == "ru-central1" {
			return aws.Endpoint{
				PartitionID:   "yc",
				URL:           "https://storage.yandexcloud.net",
				SigningRegion: "ru-central1",
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	cfg, err := cfg.LoadDefaultConfig(context.TODO(), cfg.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		log.Fatal(111, err)
	}

	client := s3.NewFromConfig(cfg)
	BucketBase = BucketBasics{S3Client: client}
}

func (basics *BucketBasics) UploadFile(bucketName string, objectKey string, file io.Reader) (string, error) {
	_, err := basics.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		log.Printf("Couldn't upload file to %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
		return "", err
	}
	return fmt.Sprintf("https://storage.yandexcloud.net/%s/%s", bucketName, objectKey), nil
}

func (basics BucketBasics) DeleteObjects(bucketName string, objectKeys []string) error {
	var objectIds []types.ObjectIdentifier
	for _, key := range objectKeys {
		objectIds = append(objectIds, types.ObjectIdentifier{Key: aws.String(key)})
	}
	_, err := basics.S3Client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{Objects: objectIds},
	})
	if err != nil {
		log.Printf("Couldn't delete objects from bucket %v. Here's why: %v\n", bucketName, err)
	}
	return err
}
