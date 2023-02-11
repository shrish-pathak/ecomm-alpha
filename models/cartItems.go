package models

import "gorm.io/gorm"

// Cart struct
type CartItem struct {
	gorm.Model
	Product   *Product `json:"product"`
	ProductID uint     `gorm:"not null" json:"productID"`
	Buyer     *Buyer   `json:"buyer"`
	BuyerID   uint     `gorm:"not null" json:"buyerID"`
	Quantity  uint     `json:"quantity"`
}
