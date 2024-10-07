package service_models

import (
	"time"
)

// TodoItem represents a database item
type TodoItem struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	FileID      string    `json:"file_id"`
}
