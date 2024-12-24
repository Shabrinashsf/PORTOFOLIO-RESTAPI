package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	NoTelp     string `json:"no_telp"`
	Role       string `json:"role"`
	IsVerified bool   `json:"is_verified"`
}
