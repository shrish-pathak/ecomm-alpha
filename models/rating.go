package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Rating struct
type Rating struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Product     *Product  `json:"product"`
	ProductID   uuid.UUID `gorm:"not null;type:uuid" json:"productID"`
	Buyer       *Buyer    `json:"buyer"`
	BuyerID     uuid.UUID `gorm:"not null;type:uuid" json:"buyerID"`
	Stars       uint8     `gorm:"not null" json:"stars"`
	Description string    `json:"description"`
}
