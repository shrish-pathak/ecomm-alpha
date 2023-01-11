package models

import "gorm.io/gorm"

// Order History struct
type OrderHistory struct {
	gorm.Model `json:"-"`
	Product    Product `json:"-"`
	ProductID  uint    `gorm:"not null" json:"productID"`
	Buyer      Buyer   `json:"-"`
	BuyerID    uint    `gorm:"not null" json:"buyerID"`
	Address    Address `json:"-"`
	AddressID  string  `gorm:"not null" json:"address"`
	Tax        int     `gorm:"not null" json:"tax"`
	Quantity   int     `gorm:"not null" json:"quantity"`
	Price      int     `gorm:"not null" json:"price"`
	Status     int     `gorm:"not null" json:"status"`
	Discount   int     `gorm:"not null" json:"discount"`
}
