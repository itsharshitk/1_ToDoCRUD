package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name" validate:"required,min=2,max=100"`
	Email     string         `gorm:"uniqueIndex" json:"email" validate:"required,email"`
	Password  string         `json:"password" validate:"required,customPassVal"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Todo      []Todo         `gorm:"foreignKey:UserId"`
}

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message,omitempty"`
	Token   string `json:"token,omitempty"`
}

type CustomClaims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"harshit@yopmail.com"`
	Password string `json:"password" validate:"required" example:"Admin@123"`
}

type SignupRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100" example:"Harshit Katiyar"`
	Email    string `json:"email" validate:"required,email" example:"harshit@yopmail.com"`
	Password string `json:"password" validate:"required" example:"Admin@123"`
}
