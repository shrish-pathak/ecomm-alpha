package models

import "gorm.io/gorm"

// Cancelled Order struct
type CancelledOrder struct {
	gorm.Model `swaggerignore:"true"`
	Order      Order `swaggerignore:"true"`
	OrderId    uint  `gorm:"not null" json:"orderID"`
}
