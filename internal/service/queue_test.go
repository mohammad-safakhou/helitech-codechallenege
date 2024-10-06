package service

import (
	"codechallenge/internal/service/service_models"
	"codechallenge/mocks"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func Test_PushTodoItem(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockQueue := mocks.NewMockQueueRepository(ctrl)

	ctx := context.Background()
	mockItem := service_models.TodoItem{
		ID:          uuid.UUID{},
		Description: "",
		DueDate:     time.Time{},
		FileID:      "",
	}
	mockQueue.EXPECT().PushTodoItem(ctx, mockItem).Return(nil)

	queueServiceTest := NewQueueService(mockQueue)

	err := queueServiceTest.pushTodoItem(ctx, mockItem)
	assert.Nil(t, err)
}
