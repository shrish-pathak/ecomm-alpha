package models

import "gorm.io/gorm"

// Order Item struct
type OrderItem struct {
	gorm.Model
	Product   *Product `json:"product"`
	ProductID uint     `gorm:"not null" json:"productID"`
	Order     *Order   `json:"order"`
	OrderId   uint     `gorm:"not null" json:"orderID"`
	Quantity  uint     `gorm:"not null" json:"quantity"`
	Price     float64  `gorm:"not null" json:"price"`
	Discount  float32  `gorm:"not null" json:"discount"` //percent value 0-100
}
