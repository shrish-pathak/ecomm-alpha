package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Seller struct
type Seller struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	FullName string    `gorm:"not null" json:"fullName" example:"harry potter"`
	Email    string    `gorm:"not null" json:"email" example:"harry@gmail.com"`
	Password string    `gorm:"not null" json:"password" example:"har!@#ryp#$otter123!@#"`
}
