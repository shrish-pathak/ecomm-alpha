package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Cancelled Order struct
type CancelledOrder struct {
	gorm.Model
	ID      uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Order   *Order    `json:"order"`
	OrderId uuid.UUID `gorm:"not null;type:uuid" json:"orderID"`
}
