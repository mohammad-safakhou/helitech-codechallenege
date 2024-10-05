package service

import (
	"codechallenge/internal/repository"
	"codechallenge/internal/service/service_models"
	"codechallenge/utils"
	"context"
	"database/sql"
)

type todoService struct {
	queueService   Queue
	todoRepository repository.TodoRepository

	tx *sql.DB
}

// NewTodoService creates a new Todo service
func NewTodoService(
	queueService Queue,
	todoRepository repository.TodoRepository,
	tx *sql.DB,
) Todo {
	return &todoService{
		queueService:   queueService,
		todoRepository: todoRepository,
		tx:             tx,
	}
}

// CreateAndPushTX creates a new todo item and pushes it to the queue
func (s *todoService) CreateAndPushTX(ctx context.Context, todoItem service_models.TodoItem) (service_models.TodoItem, error) {
	deferFunc, tx, err := utils.GetDbTx(s.tx)
	if err != nil {
		return service_models.TodoItem{}, err
	}
	defer deferFunc(&err)

	self := s.getWithTX(tx)

	todoItem, err = self.create(ctx, todoItem)
	if err != nil {
		return service_models.TodoItem{}, err
	}

	err = s.queueService.pushTodoItem(ctx, todoItem)
	if err != nil {
		return service_models.TodoItem{}, err
	}

	return todoItem, nil
}

// Create creates a new todo item
func (s *todoService) create(ctx context.Context, todoItem service_models.TodoItem) (service_models.TodoItem, error) {
	return s.todoRepository.Create(ctx, todoItem)
}

// Get gets a todo item by its ID
func (s *todoService) get(ctx context.Context, id string) (service_models.TodoItem, error) {
	return s.todoRepository.Get(ctx, id)
}

// GetWithTX returns a new Todo service that uses the provided transaction
func (s *todoService) getWithTX(tx *sql.Tx) Todo {
	return &todoService{
		todoRepository: s.todoRepository.GetWithTX(tx),
		queueService:   s.queueService,
		tx:             s.tx,
	}
}
