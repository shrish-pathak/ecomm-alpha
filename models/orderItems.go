package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Order Item struct
type OrderItem struct {
	gorm.Model
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Product   *Product  `json:"product"`
	ProductID uuid.UUID `gorm:"not null;type:uuid" json:"productID"`
	Order     *Order    `json:"order"`
	OrderId   uuid.UUID `gorm:"not null;type:uuid" json:"orderID"`
	Quantity  uint      `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"not null" json:"price"`
	Discount  float32   `gorm:"not null" json:"discount"` //percent value 0-100
}
