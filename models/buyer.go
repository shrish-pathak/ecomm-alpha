package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Buyer struct
type Buyer struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	FullName string    `gorm:"not null" json:"fullName"`
	Email    string    `gorm:"not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
}
