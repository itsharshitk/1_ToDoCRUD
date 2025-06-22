package model

import (
	"time"

	"gorm.io/gorm"
)

// Todo represents a task belonging to a user
type Todo struct {
	ID          uint           `gorm:"primaryKey" json:"id" example:"1"`                                // Unique identifier
	UserId      uint           `json:"user_id" validate:"required" example:"10"`                        // ID of the user who owns the task
	Title       string         `json:"title" validate:"required,min=2,max=255" example:"Buy groceries"` // Task title
	Description string         `json:"description" validate:"max=255" example:"Buy milk and bread"`     // Optional task description
	IsComplete  *bool          `gorm:"default:false" json:"is_complete" example:"false"`                // Completion status
	CreatedAt   time.Time      `json:"created_at" example:"2025-06-23T12:00:00Z"`                       // Task creation timestamp
	UpdatedAt   time.Time      `json:"updated_at" example:"2025-06-23T12:30:00Z"`                       // Task update timestamp
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`                             // Soft delete
}
