package models

import "gorm.io/gorm"

// Addresses struct
type Addresses struct {
	gorm.Model   `json:"-"`
	UserID       string `gorm:"not null" json:"userID"`
	Mobile_No    int    `gorm:"not null" json:"mobileNo"`
	City         string `gorm:"not null" json:"city"`
	State        string `gorm:"not null" json:"state"`
	Country_Code int    `gorm:"not null" json:"countryCode"`
	Country      string `gorm:"not null" json:"country"`
	Address      string `gorm:"not null" json:"address"`
}
