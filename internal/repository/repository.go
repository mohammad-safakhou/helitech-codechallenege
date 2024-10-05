package repository

import (
	"codechallenge/internal/service/service_models"
	"context"
	"database/sql"
	"io"
)

type StorageRepository interface {
	Upload(ctx context.Context, file io.ReadCloser, filename string) (string, error)
	Download(ctx context.Context, filename string) (io.ReadCloser, error)
}

type QueueRepository interface {
	PushTodoItem(ctx context.Context, message service_models.TodoItem) error
}

type TodoRepository interface {
	Create(ctx context.Context, todoItem service_models.TodoItem) (service_models.TodoItem, error)
	Get(ctx context.Context, id string) (service_models.TodoItem, error)

	GetWithTX(tx *sql.Tx) TodoRepository
}
