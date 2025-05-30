package entity

import (
	"time"

	"github.com/google/uuid"
)

type Subevent struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EventID  uuid.UUID `json:"event_id"`
	Name     string    `json:"name"`
	StartsAt time.Time `gorm:"type:timestamp with time zone" json:"starts_at"`
	EndsAt   time.Time `gorm:"type:timestamp with time zone" json:"ends_at"`
	Price    int       `json:"price"`

	Event *Event `gorm:"foreignKey:EventID"`

	Timestamp
}
