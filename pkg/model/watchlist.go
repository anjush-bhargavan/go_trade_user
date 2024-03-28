package model

import "gorm.io/gorm"

type Wathlist struct {
	gorm.Model
	CategoryID uint `gorm:"not null"`
	UserID     uint `gorm:"not null"`
}
