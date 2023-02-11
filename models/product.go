package models

import "gorm.io/gorm"

// Product struct
type Product struct {
	gorm.Model        `swaggerignore:"true"`
	Store             *Store  `swaggerignore:"true"`
	StoreID           uint    `gorm:"not null" json:"storeID"`
	Title             string  `gorm:"not null" json:"title"`
	Description       string  `gorm:"not null" json:"description"`
	Price             float64 `gorm:"not null" json:"price"`
	Discount          float32 `gorm:"not null" json:"discount"` //percent 0-100
	AvailableQuantity uint    `json:"availableQuantity"`
}
