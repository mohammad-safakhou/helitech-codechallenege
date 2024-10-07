package dto

import "time"

type StorageUploadResponse struct {
	FileName string `json:"file_name"`
}

type AddTodoItemRequest struct {
	Description string    `json:"description" validate:"required"`
	DueDate     time.Time `json:"due_date" validate:"required"`
	FileID      string    `json:"file_id"`
}

type AddTodoItemResponse struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	FileID      string    `json:"file_id"`
}
