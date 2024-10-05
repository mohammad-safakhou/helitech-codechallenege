package service

import (
	"codechallenge/internal/service/service_models"
	"context"
)

// queueService is the struct that implements the Queue interface
type queueService struct {
}

// NewQueueService creates a new Queue service
func NewQueueService() Queue {
	return &queueService{}
}

// PushTodoItem pushes a todo item to the queue
func (s *queueService) PushTodoItem(ctx context.Context, message service_models.TodoItem) error {
	panic("not implemented") // TODO: Implement
}
