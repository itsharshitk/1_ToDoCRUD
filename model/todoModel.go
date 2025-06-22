package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserId      uint           `json:"user_id" validate:"required"`
	Title       string         `json:"title" validate:"required,min=2,max=255"`
	Description string         `json:"description" validate:"max=255"`
	IsComplete  *bool          `gorm:"default:false" json:"is_complete"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
