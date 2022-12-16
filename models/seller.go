package model

import "gorm.io/gorm"

// Seller struct
type Seller struct {
	gorm.Model
	FullName string `gorm:"not null" json:"fullName"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
