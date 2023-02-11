package models

import "gorm.io/gorm"

// Order struct
type Order struct {
	gorm.Model  `swaggerignore:"true"`
	Buyer       *Buyer   `swaggerignore:"true"`
	BuyerID     uint     `gorm:"not null" json:"buyerID"`
	Address     *Address `swaggerignore:"true"`
	AddressID   string   `gorm:"not null" json:"address"`
	Tax         float32  `gorm:"not null" json:"tax"` //percent value 0-100
	TotalAmount float64  `gorm:"not null" json:"amount"`
	Status      uint     `gorm:"not null" json:"status"`
}
