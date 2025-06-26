package model

import (
	"time"

	"gorm.io/gorm"
)

// Todo represents a task belonging to a user
type Todo struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserId      uint           `json:"user_id" validate:"required"`
	Title       string         `json:"title" validate:"required,min=2,max=255"`
	Description string         `json:"description" validate:"max=255"`
	IsComplete  *bool          `gorm:"default:false" json:"is_complete"`
	CreatedAt   time.Time      `json:"created_at" example:"2025-06-23T12:00:00Z"`
	UpdatedAt   time.Time      `json:"updated_at" example:"2025-06-23T12:30:00Z"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-" swaggerignore:"true"`
}

type AddTaskRequest struct {
	Title       string `json:"title" validate:"required,min=2,max=255" example:"Buy groceries"`
	Description string `json:"description" validate:"max=255" example:"Buy milk and bread"`
}
