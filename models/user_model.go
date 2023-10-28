package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primary_key;autoIncrement" json:"id"`
	UserID       string         `json:"user_id"`
	FirstName    string         `gorm:"not null" json:"first_name"`
	LastName     string         `gorm:"not null" json:"last_name"`
	Email        string         `gorm:"unique, not null" json:"email"`
	Password     string         `gorm:"not null" json:"password"`
	UserType     string         `gorm:"default:'user'" json:"user_type"`
	Token        string         `json:"token"`
	RefreshToken string         `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
