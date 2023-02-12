package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Address struct
type Address struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID   uuid.UUID `gorm:"not null;type:uuid" json:"userID"`
	MobileNo string    `gorm:"not null" json:"mobileNo"`
	City     string    `gorm:"not null" json:"city"`
	State    string    `json:"state"`
	Zip      string    `gorm:"not null" json:"zip"`
	Country  string    `gorm:"not null" json:"country"`
	Address  string    `gorm:"not null" json:"address"`
}
