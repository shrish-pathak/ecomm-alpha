package model

import "gorm.io/gorm"

// Rating struct
type Rating struct {
	gorm.Model
	Product     Product
	ProductID   uint `gorm:"not null" json:"productID"`
	Buyer       Buyer
	BuyerID     uint   `gorm:"not null" json:"buyerID"`
	Stars       uint8  `gorm:"not null" json:"stars"`
	Description string `json:"description"`
}
