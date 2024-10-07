package rest

import (
	"codechallenge/internal/gateway/dto"
	"codechallenge/internal/service"
	"codechallenge/internal/service/service_models"
	"codechallenge/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TodoController interface {
	AddItem(ctx echo.Context) error
}
type todoControllers struct {
	todoService service.Todo
}

func NewTodoController(todoService service.Todo) TodoController {
	return &todoControllers{
		todoService: todoService,
	}
}

func (sc *todoControllers) AddItem(ctx echo.Context) error {
	req := dto.AddTodoItemRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.StandardHttpResponse{
			Message: "",
			Status:  http.StatusBadRequest,
			Data:    err.Error(),
		})
	}

	err := validate.Struct(req)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.StandardHttpResponse{
			Message: "",
			Status:  http.StatusBadRequest,
			Data:    err.Error(),
		})
	}
	item, err := sc.todoService.CreateAndPushTX(ctx.Request().Context(), service_models.TodoItem{
		Description: req.Description,
		DueDate:     req.DueDate,
		FileID:      req.FileID,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, utils.StandardHttpResponse{
			Message: "upload failed",
			Status:  http.StatusInternalServerError,
			Data:    err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, utils.StandardHttpResponse{
		Message: "file successfully uploaded",
		Status:  http.StatusOK,
		Data: dto.AddTodoItemResponse{
			ID:          item.ID,
			Description: item.Description,
			DueDate:     item.DueDate,
			FileID:      item.FileID,
		},
	})
}
