package models

import "gorm.io/gorm"

// Cart struct
type Cart struct {
	gorm.Model `json:"-"`
	Product    Product `json:"-"`
	ProductID  uint    `gorm:"not null" json:"productID"`
	Buyer      Buyer   `json:"-"`
	BuyerID    uint    `gorm:"not null" json:"buyerID"`
}
