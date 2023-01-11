package models

import "gorm.io/gorm"

// Address struct
type Address struct {
	gorm.Model  `json:"-"`
	UserID      uint   `gorm:"not null" json:"userID"`
	MobileNo    string `gorm:"not null" json:"mobileNo"`
	City        string `gorm:"not null" json:"city"`
	State       string `gorm:"not null" json:"state"`
	CountryCode string `gorm:"not null" json:"countryCode"`
	Country     string `gorm:"not null" json:"country"`
	Address     string `gorm:"not null" json:"address"`
}
