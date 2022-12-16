package model

import "gorm.io/gorm"

// Buyer struct
type Buyer struct {
	gorm.Model
	FullName string `gorm:"not null" json:"fullName"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
