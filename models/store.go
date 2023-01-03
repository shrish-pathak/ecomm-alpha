package models

import "gorm.io/gorm"

// Store struct
type Store struct {
	gorm.Model  `json:"-"`
	Seller      Seller `json:"-"`
	SellerID    uint   `gorm:"not null" json:"sellerID"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
}
