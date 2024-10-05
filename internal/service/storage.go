package service

import "context"

// Storage is the interface that defines the methods that the Storage service must implement
type storageService struct {
}

// NewStorageService creates a new Storage service
func NewStorageService() Storage {
	return &storageService{}
}

// Upload uploads a file
func (s *storageService) Upload(ctx context.Context, file []byte, fileName string) (string, error) {
	panic("not implemented") // TODO: Implement
}

// Download downloads a file
func (s *storageService) Download(ctx context.Context, fileID string) ([]byte, error) {
	panic("not implemented") // TODO: Implement
}
