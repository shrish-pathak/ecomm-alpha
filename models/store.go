package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Store struct
type Store struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Seller      *Seller   `json:"seller"`
	SellerID    uuid.UUID `gorm:"not null;type:uuid" json:"sellerID"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `gorm:"not null" json:"description"`
}
