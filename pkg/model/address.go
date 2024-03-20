package model

import "gorm.io/gorm"

// Address represents the table of address in the database schema.
type Address struct {
	gorm.Model
	House  string `gorm:"not null"`
	Street string `gorm:"not null"`
	City   string `gorm:"not null"`
	ZIP    uint   `gorm:"not null"`
	State  string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
