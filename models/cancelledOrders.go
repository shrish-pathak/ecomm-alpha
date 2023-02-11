package models

import "gorm.io/gorm"

// Cancelled Order struct
type CancelledOrder struct {
	gorm.Model
	Order   *Order `json:"order"`
	OrderId uint   `gorm:"not null" json:"orderID"`
}
