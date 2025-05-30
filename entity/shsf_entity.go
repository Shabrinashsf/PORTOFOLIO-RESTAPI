package entity

import "github.com/google/uuid"

const (
	ADMIN_STATUS_PENDING  = "PENDING"
	ADMIN_STATUS_REJECTED = "REJECTED"
	ADMIN_STATUS_REVISION = "REVISION"
	ADMIN_STATUS_ACCEPTED = "ACCEPTED"
	ADMIN_STATUS_REVISED  = "REVISED"
)

type SHSF struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	SubeventID   uuid.UUID `json:"subevent_id"`
	SubeventName string    `json:"subevent_name"`

	// form data
	Name   string `json:"name"`
	NoTelp string `json:"no_telp"`
	Email  string `json:"email"`

	AdminStatus string `json:"admin_status"`

	User     *User     `gorm:"foreignKey:UserID"`
	Subevent *Subevent `gorm:"foreignKey:SubeventID"`

	Timestamp
}
