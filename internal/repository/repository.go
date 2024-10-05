package repository

import (
	"codechallenge/internal/service/service_models"
	"context"
	"database/sql"
)

type StorageRepository interface {
	Upload(ctx context.Context, file []byte, filename string) (string, error)
	Download(ctx context.Context, filename string) ([]byte, error)
}

type QueueRepository interface {
	PushTodoItem(ctx context.Context, message service_models.TodoItem) error
}

type TodoRepository interface {
	Create(ctx context.Context, todoItem service_models.TodoItem) (service_models.TodoItem, error)
	Get(ctx context.Context, id string) (service_models.TodoItem, error)

	GetWithTX(tx *sql.Tx) TodoRepository
}
