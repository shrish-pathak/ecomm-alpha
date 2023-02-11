package models

import "gorm.io/gorm"

// Rating struct
type Rating struct {
	gorm.Model
	Product     *Product `json:"product"`
	ProductID   uint     `gorm:"not null" json:"productID"`
	Buyer       *Buyer   `json:"buyer"`
	BuyerID     uint     `gorm:"not null" json:"buyerID"`
	Stars       uint8    `gorm:"not null" json:"stars"`
	Description string   `json:"description"`
}
