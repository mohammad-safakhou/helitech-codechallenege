package service

import (
	"codechallenge/internal/service/service_models"
	"context"
	"database/sql"
)

type todoService struct {
	tx *sql.DB
}

// NewTodoService creates a new Todo service
func NewTodoService() Todo {
	return &todoService{}
}

// Create creates a new todo item
func (s *todoService) Create(ctx context.Context, todoItem service_models.TodoItem) (service_models.TodoItem, error) {
	panic("not implemented") // TODO: Implement
}

// Get gets a todo item by its ID
func (s *todoService) Get(ctx context.Context, id string) (service_models.TodoItem, error) {
	panic("not implemented") // TODO: Implement
}

// GetWithTX returns a new Todo service that uses the provided transaction
func (s *todoService) GetWithTX(tx *sql.Tx) Todo {
	return &todoService{
		tx: s.tx,
	}
}
