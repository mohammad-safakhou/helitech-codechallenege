package service

import (
	"codechallenge/mocks"
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"strings"
	"testing"
)

func Test_Upload(t *testing.T) {
	t.Run("Should return ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		mockStorage := mocks.NewMockStorageRepository(ctrl)

		ctx := context.Background()
		f := strings.NewReader("test we are here")
		file := io.NopCloser(f)
		fileName := "TestingFileName.txt"
		mockStorage.EXPECT().Upload(ctx, file, fileName).Return("", nil)

		storageServiceTest := NewStorageService(mockStorage, 100, []string{"txt"})

		_, err := storageServiceTest.Upload(ctx, file, fileName)
		assert.Nil(t, err)
	})
	t.Run("Should return max size file exceeded", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		mockStorage := mocks.NewMockStorageRepository(ctrl)

		ctx := context.Background()
		f := strings.NewReader("test we are here")
		file := io.NopCloser(f)
		fileName := "TestingFileName.txt"
		mockStorage.EXPECT().Upload(ctx, file, fileName).Times(0)

		storageServiceTest := NewStorageService(mockStorage, 2, []string{"txt"})

		_, err := storageServiceTest.Upload(ctx, file, fileName)
		assert.NotNil(t, err)
	})
	t.Run("Should return file type mismatch", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		mockStorage := mocks.NewMockStorageRepository(ctrl)

		ctx := context.Background()
		f := strings.NewReader("test we are here")
		file := io.NopCloser(f)
		fileName := "TestingFileName.png"
		mockStorage.EXPECT().Upload(ctx, file, fileName).Times(0)

		storageServiceTest := NewStorageService(mockStorage, 100, []string{"txt"})

		_, err := storageServiceTest.Upload(ctx, file, fileName)
		assert.NotNil(t, err)
	})
}
