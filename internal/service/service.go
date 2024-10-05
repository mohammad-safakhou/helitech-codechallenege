package service

import (
	"codechallenge/internal/service/service_models"
	"context"
	"database/sql"
)

// Todo is the interface that defines the methods that the Todo service must implement
type Todo interface {
	Create(ctx context.Context, todoItem service_models.TodoItem) (service_models.TodoItem, error)
	Get(ctx context.Context, id string) (service_models.TodoItem, error)

	GetWithTX(tx *sql.Tx) Todo
}
