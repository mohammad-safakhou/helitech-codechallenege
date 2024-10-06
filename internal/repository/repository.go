package repository

import (
	"codechallenge/internal/service/service_models"
	"codechallenge/utils"
	"context"
	"io"
)

type StorageRepository interface {
	Upload(ctx context.Context, file io.ReadCloser, filename string) (string, error)
}

type QueueRepository interface {
	PushTodoItem(ctx context.Context, message service_models.TodoItem) error
}

type TodoRepository interface {
	CreateWithTX(ctx context.Context, todoItem service_models.TodoItem) (dbFunc utils.DbTransaction, item service_models.TodoItem, err error)
	Get(ctx context.Context, id string) (service_models.TodoItem, error)
}
