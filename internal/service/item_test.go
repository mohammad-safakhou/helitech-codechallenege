package service

import (
	"codechallenge/internal/service/service_models"
	"codechallenge/mocks"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_CreateAndPushTX(t *testing.T) {
	t.Run("Should return ok", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		mockQueue := mocks.NewMockQueueRepository(ctrl)
		mockTodoRepo := mocks.NewMockTodoRepository(ctrl)
		mockDbTransaction := mocks.NewMockDbTransaction(ctrl)

		ctx := context.Background()
		var mockItem = service_models.TodoItem{}
		mockQueue.EXPECT().PushTodoItem(ctx, gomock.Any()).Return(nil)
		mockTodoRepo.EXPECT().CreateWithTX(ctx, gomock.Any()).Return(mockDbTransaction, mockItem, nil)

		mockDbTransaction.EXPECT().Commit().Return(nil)

		queueServiceTest := NewQueueService(mockQueue)
		todoServiceTest := NewTodoService(queueServiceTest, mockTodoRepo)

		resultItem, err := todoServiceTest.CreateAndPushTX(ctx, mockItem)
		assert.Nil(t, err)

		assert.NotEqual(t, resultItem.ID, uuid.UUID{})
	})
	t.Run("Should return pushing error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		mockQueue := mocks.NewMockQueueRepository(ctrl)
		mockTodoRepo := mocks.NewMockTodoRepository(ctrl)
		mockDbTransaction := mocks.NewMockDbTransaction(ctrl)

		mockDbTransaction.EXPECT().Rollback().Return(nil)

		ctx := context.Background()
		var mockItem = service_models.TodoItem{}
		mockQueue.EXPECT().PushTodoItem(ctx, mockItem).Return(errors.New("error"))
		mockTodoRepo.EXPECT().CreateWithTX(ctx, gomock.Any()).Return(mockDbTransaction, mockItem, nil)

		queueServiceTest := NewQueueService(mockQueue)
		todoServiceTest := NewTodoService(queueServiceTest, mockTodoRepo)

		resultItem, err := todoServiceTest.CreateAndPushTX(ctx, mockItem)
		assert.NotNil(t, err)

		assert.NotEqual(t, resultItem.ID, uuid.UUID{})
	})
}
