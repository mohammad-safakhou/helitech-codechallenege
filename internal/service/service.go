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

// Storage is the interface that defines the methods that the Storage service must implement
type Storage interface {
	Upload(ctx context.Context, file []byte, fileName string) (string, error)
	Download(ctx context.Context, fileID string) ([]byte, error)
}

// Queue is the interface that defines the methods that the Queue service must implement
type Queue interface {
	PushTodoItem(ctx context.Context, message service_models.TodoItem) error
}
