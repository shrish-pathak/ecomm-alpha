package models

import "gorm.io/gorm"

// Cart struct
type Cart struct {
	gorm.Model
	Product   Product
	ProductID uint `gorm:"not null" json:"productID"`
	Buyer     Buyer
	BuyerID   uint `gorm:"not null" json:"buyerID"`
}
