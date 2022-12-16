package model

import "gorm.io/gorm"

// Product struct
type Product struct {
	gorm.Model
	Store       Store
	StoreID     uint   `gorm:"not null" json:"storeID"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Price       int    `gorm:"not null" json:"price"`
	Discount    int    `gorm:"not null" json:"discount"`
}
