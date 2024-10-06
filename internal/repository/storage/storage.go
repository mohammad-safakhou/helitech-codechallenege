package storage

import (
	"codechallenge/config"
	"codechallenge/internal/repository"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

type storageRepository struct {
	Client *s3.Client
}

func NewStorageRepository(client *s3.Client) repository.StorageRepository {
	return &storageRepository{
		Client: client,
	}
}

func (s *storageRepository) Upload(ctx context.Context, file io.ReadCloser, filename string) error {
	_, err := s.Client.PutObject(ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(config.AppConfig.Storage.Bucket),
			Key:    aws.String(filename),
			Body:   file,
		},
	)
	return err
}
