package service_models

import (
	"github.com/google/uuid"
	"time"
)

// TodoItem represents a database item
type TodoItem struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	FileID      string    `json:"file_id"`
}
