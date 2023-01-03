package models

import "gorm.io/gorm"

// Rating struct
type Rating struct {
	gorm.Model  `json:"-"`
	Product     Product `json:"-"`
	ProductID   uint    `gorm:"not null" json:"productID"`
	Buyer       Buyer   `json:"-"`
	BuyerID     uint    `gorm:"not null" json:"buyerID"`
	Stars       uint8   `gorm:"not null" json:"stars"`
	Description string  `json:"description"`
}
