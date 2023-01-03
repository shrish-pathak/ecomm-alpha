package models

import "gorm.io/gorm"

// Seller struct
type Seller struct {
	gorm.Model `json:"-"`
	FullName   string `gorm:"not null" json:"fullName" example:"harry potter"`
	Email      string `gorm:"not null" json:"email" example:"harry@gmail.com"`
	Password   string `gorm:"not null" json:"password" example:"har!@#ryp#$otter123!@#"`
}
