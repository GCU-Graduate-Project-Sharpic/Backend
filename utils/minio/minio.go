package minio

import (
	"bytes"
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	MinioClient *minio.Client
}

func Dial(accessHost string, accessID string, accessPW string, useSSL bool) (*Client, error) {
	minioClient, err := minio.New(accessHost, &minio.Options{
		Creds:  credentials.NewStaticV4(accessID, accessPW, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		MinioClient: minioClient,
	}, nil
}

func (c *Client) MakeBucketIfNotExists(bucketName string) error {
	ctx := context.Background()

	err := c.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := c.MinioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Println("We already own " + bucketName)
			return nil
		}
		log.Println(err)
		return err
	}
	log.Println("Successfully created " + bucketName)
	return nil
}

func (c *Client) Upload(bucketName string, objectName string, file []byte) error {
	ctx := context.Background()
	_, err := c.MinioClient.PutObject(ctx, bucketName, objectName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Download(bucketName string, objectName string) ([]byte, error) {
	ctx := context.Background()
	obj, err := c.MinioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(obj)
	return buf.Bytes(), nil
}
