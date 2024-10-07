package rest

import (
	"codechallenge/internal/gateway/dto"
	"codechallenge/internal/service"
	"codechallenge/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type StorageController interface {
	UploadHandler(ctx echo.Context) error
}
type storageControllers struct {
	storageService service.Storage
}

func NewStorageController(storageService service.Storage) StorageController {
	return &storageControllers{
		storageService: storageService,
	}
}

func (sc *storageControllers) UploadHandler(ctx echo.Context) error {
	formFile, err := ctx.FormFile("file")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.StandardHttpResponse{
			Message: "",
			Status:  http.StatusBadRequest,
			Data:    err.Error(),
		})
	}

	file, err := formFile.Open()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.StandardHttpResponse{
			Message: "",
			Status:  http.StatusBadRequest,
			Data:    err.Error(),
		})
	}
	defer file.Close()

	err = sc.storageService.Upload(ctx.Request().Context(), file, formFile.Filename)
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
		Data:    dto.StorageUploadResponse{FileName: formFile.Filename},
	})
}
