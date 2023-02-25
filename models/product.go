package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	ID                uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Store             *Store    `json:"store"`
	StoreID           uuid.UUID `gorm:"not null;type:uuid" json:"storeID"`
	Title             string    `gorm:"not null" json:"title"`
	Description       string    `gorm:"not null" json:"description"`
	Price             float64   `gorm:"not null" json:"price"`
	Discount          float32   `gorm:"not null" json:"discount"` //percent 0-100
	AvailableQuantity uint      `json:"availableQuantity"`
}
