package entity

import (
	"github.com/google/uuid"
)

type EventPayment struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	RegistID  uuid.UUID `json:"regist_id"`
	PaymentID uuid.UUID `json:"payment_id"`

	Payment *Payment `gorm:"foreignKey:PaymentID"`

	Timestamp
}
