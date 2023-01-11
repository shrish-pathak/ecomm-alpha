package models

import "gorm.io/gorm"

// Address struct
type Address struct {
	gorm.Model   `json:"-"`
	UserID       uint   `gorm:"not null" json:"userID"`
	Mobile_No    string `gorm:"not null" json:"mobileNo"`
	City         string `gorm:"not null" json:"city"`
	State        string `gorm:"not null" json:"state"`
	Country_Code string `gorm:"not null" json:"countryCode"`
	Country      string `gorm:"not null" json:"country"`
	Address      string `gorm:"not null" json:"address"`
}
