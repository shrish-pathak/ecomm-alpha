package models

import "gorm.io/gorm"

// Address struct
type Address struct {
	gorm.Model
	UserID   uint   `gorm:"not null" json:"userID"`
	MobileNo string `gorm:"not null" json:"mobileNo"`
	City     string `gorm:"not null" json:"city"`
	State    string `json:"state"`
	Zip      string `gorm:"not null" json:"zip"`
	Country  string `gorm:"not null" json:"country"`
	Address  string `gorm:"not null" json:"address"`
}
