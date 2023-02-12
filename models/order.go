package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order struct
type Order struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Buyer       *Buyer    `json:"buyer"`
	BuyerID     uuid.UUID `gorm:"not null;type:uuid" json:"buyerID"`
	Address     *Address  `json:"address"`
	AddressID   uuid.UUID `gorm:"not null" json:"addressID"`
	Tax         float32   `gorm:"not null" json:"tax"` //percent value 0-100
	TotalAmount float64   `gorm:"not null" json:"amount"`
	Status      uint      `gorm:"not null" json:"status"`
}
