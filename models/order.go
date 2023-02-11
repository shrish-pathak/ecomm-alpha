package models

import "gorm.io/gorm"

// Order struct
type Order struct {
	gorm.Model
	Buyer       *Buyer   `json:"buyer"`
	BuyerID     uint     `gorm:"not null" json:"buyerID"`
	Address     *Address `json:"address"`
	AddressID   string   `gorm:"not null" json:"addressID"`
	Tax         float32  `gorm:"not null" json:"tax"` //percent value 0-100
	TotalAmount float64  `gorm:"not null" json:"amount"`
	Status      uint     `gorm:"not null" json:"status"`
}
