package models

import "gorm.io/gorm"

// Cart struct
type CartItem struct {
	gorm.Model `swaggerignore:"true"`
	Product    *Product `json:"product"`

	ProductID uint   `gorm:"not null" json:"productID"`
	Buyer     *Buyer `swaggerignore:"true"`
	BuyerID   uint   `gorm:"not null" json:"buyerID"`
	Quantity  uint   `json:"quantity"`
}
