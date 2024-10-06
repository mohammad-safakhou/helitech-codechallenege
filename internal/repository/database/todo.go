package database

import (
	"codechallenge/internal/repository"
	models "codechallenge/internal/repository/database/boiler"
	"codechallenge/internal/service/service_models"
	"codechallenge/utils"
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type todoRepository struct {
	dbRead  *sql.DB
	dbWrite *sql.DB
}

func NewTodoRepository(dbRead, dbWrite *sql.DB) repository.TodoRepository {
	return &todoRepository{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}

func (s *todoRepository) CreateWithTX(ctx context.Context, todoItem service_models.TodoItem) (dbFunc utils.DbTransaction, item service_models.TodoItem, err error) {
	tx, err := s.dbWrite.Begin()
	if err != nil {
		return nil, service_models.TodoItem{}, err
	}

	dbModel := toTodoBoilerModel(todoItem)
	err = dbModel.Insert(ctx, tx, boil.Infer())
	if err != nil {
		return nil, service_models.TodoItem{}, err
	}

	item = toTodoServiceModel(dbModel)

	return tx, item, nil
}

func (s *todoRepository) Get(ctx context.Context, id string) (service_models.TodoItem, error) {
	todo, err := models.TodoItems(models.TodoItemWhere.ID.EQ(id)).One(ctx, s.dbRead)
	if err != nil {
		return service_models.TodoItem{}, err
	}
	return toTodoServiceModel(*todo), nil
}

func toTodoBoilerModel(s service_models.TodoItem) models.TodoItem {
	return models.TodoItem{
		ID:          s.ID,
		Description: s.Description,
		DueDate:     s.DueDate,
		FileID:      s.FileID,
	}
}

func toTodoServiceModel(s models.TodoItem) service_models.TodoItem {
	return service_models.TodoItem{
		ID:          s.ID,
		Description: s.Description,
		DueDate:     s.DueDate,
		FileID:      s.FileID,
	}
}
