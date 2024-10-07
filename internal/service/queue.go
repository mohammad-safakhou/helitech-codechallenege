package service

import (
	"codechallenge/internal/repository"
	"codechallenge/internal/service/service_models"
	"context"
)

// queueService is the struct that implements the Queue interface
type queueService struct {
	queueRepository repository.QueueRepository
}

// NewQueueService creates a new Queue service
func NewQueueService(queueRepository repository.QueueRepository) Queue {
	return &queueService{
		queueRepository: queueRepository,
	}
}

// PushTodoItem pushes a database item to the queue
func (s *queueService) pushTodoItem(ctx context.Context, message service_models.TodoItem) error {
	return s.queueRepository.PushTodoItem(ctx, message)
}
