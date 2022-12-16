package model

import "gorm.io/gorm"

// Store struct
type Store struct {
	gorm.Model
	Seller      Seller
	SellerID    uint   `gorm:"not null" json:"sellerID"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}