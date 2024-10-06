package service

import (
	"codechallenge/internal/repository"
	"context"
	"errors"
	"io"
	"strings"
)

// Storage is the interface that defines the methods that the Storage service must implement
type storageService struct {
	storageRepository   repository.StorageRepository
	maxUploadLength     int64
	validFileExtensions []string
}

// NewStorageService creates a new Storage service
func NewStorageService(storageRepository repository.StorageRepository, maxUploadLength int64, validFileExtensions []string) Storage {
	return &storageService{
		storageRepository:   storageRepository,
		maxUploadLength:     maxUploadLength,
		validFileExtensions: validFileExtensions,
	}
}

// Upload uploads a file
func (s *storageService) Upload(ctx context.Context, file io.ReadCloser, fileName string) (string, error) {
	err := s.validateFileName(fileName)
	if err != nil {
		return "", err
	}

	err = s.validateFile(file)
	if err != nil {
		return "", err
	}

	return s.storageRepository.Upload(ctx, file, fileName)
}

func (s *storageService) validateFileName(fileName string) error {
	if fileName == "" {
		return errors.New("file name is required")
	}

	// with regex validate file types of png and jpeg
	flag := false
	for _, ext := range s.validFileExtensions {
		if strings.HasSuffix(fileName, ext) {
			flag = true
			break
		}
	}

	if !flag {
		return errors.New("file type is not supported")
	}

	return nil
}

func (s *storageService) validateFile(file io.ReadCloser) error {
	length, err := io.Copy(io.Discard, file)
	if err != nil {
		return err
	}

	if length > s.maxUploadLength {
		return errors.New("file size is too large")
	}

	return nil
}
