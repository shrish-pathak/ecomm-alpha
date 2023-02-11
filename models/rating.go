package models

import "gorm.io/gorm"

// Rating struct
type Rating struct {
	gorm.Model  `swaggerignore:"true"`
	Product     *Product `swaggerignore:"true"`
	ProductID   uint     `gorm:"not null" json:"productID"`
	Buyer       *Buyer   `swaggerignore:"true"`
	BuyerID     uint     `gorm:"not null" json:"buyerID"`
	Stars       uint8    `gorm:"not null" json:"stars"`
	Description string   `json:"description"`
}
