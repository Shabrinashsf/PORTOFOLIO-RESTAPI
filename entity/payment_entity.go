package entity

import (
	"github.com/google/uuid"
)

type Payment struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID        uuid.UUID `json:"user_id"`
	PaymentStatus string    `json:"payment_status"`
	Type          string    `json:"type"`
	InvoiceURL    string    `json:"invoice_url"`
	PaidAmount    int       `json:"paid_amount"`

	User *User `gorm:"foreignKey:UserID"`

	Timestamp
}
