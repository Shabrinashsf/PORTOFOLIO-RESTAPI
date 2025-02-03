package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	NoTelp           string    `json:"no_telp"`
	Role             string    `json:"role"`
	IsVerified       bool      `json:"is_verified"`
	VerificationCode string    `json:"verification_code"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}
