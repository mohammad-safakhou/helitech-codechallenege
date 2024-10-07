package service

import (
	"codechallenge/internal/repository"
	"codechallenge/internal/service/service_models"
	"codechallenge/logger"
	"codechallenge/utils"
	"context"
	"github.com/google/uuid"
)

type todoService struct {
	queueService   Queue
	todoRepository repository.TodoRepository
}

// NewTodoService creates a new Todo service
func NewTodoService(
	queueService Queue,
	todoRepository repository.TodoRepository,
) Todo {
	return &todoService{
		queueService:   queueService,
		todoRepository: todoRepository,
	}
}

// CreateAndPushTX creates a new database item and pushes it to the queue
func (s *todoService) CreateAndPushTX(ctx context.Context, todoItem service_models.TodoItem) (service_models.TodoItem, error) {
	todoItem.ID = uuid.New().String()
	dbFunc, todoItem, err := s.create(ctx, todoItem)
	if err != nil {
		return service_models.TodoItem{}, err
	}

	err = s.queueService.pushTodoItem(ctx, todoItem)
	if err != nil {
		rollBackRrr := dbFunc.Rollback()
		if rollBackRrr != nil {
			logger.Logger.Error("Error rolling back transaction", rollBackRrr)
		}
		return service_models.TodoItem{}, err
	}

	// We considered that pushing has more priority than committing, so we push first and then commit
	err = dbFunc.Commit()
	if err != nil {
		return service_models.TodoItem{}, err
	}

	return todoItem, nil
}

// Create creates a new database item
func (s *todoService) create(ctx context.Context, todoItem service_models.TodoItem) (utils.DbTransaction, service_models.TodoItem, error) {
	return s.todoRepository.CreateWithTX(ctx, todoItem)
}

// Get gets a database item by its ID
func (s *todoService) get(ctx context.Context, id string) (service_models.TodoItem, error) {
	return s.todoRepository.Get(ctx, id)
}
