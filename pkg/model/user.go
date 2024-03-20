package model

import "gorm.io/gorm"

// User represents the model of user table.
type User struct {
	gorm.Model

	UserName      string  `gorm:"not null"`
	Email         string  `gorm:"not null;unique"`
	Password      string  `gorm:"not null"`
	Mobile        string  `gorm:"not null"`
	ReferralCode  string  `gorm:"not null"`
	Wallet        float64 `gorm:"default:0"`
	IsBlocked     bool    `gorm:"default:false"`
	RejecedOrders uint    `gorm:"default:0"`
}

// Transaction represents the transaction happen in the users wallet
type Transaction struct {
	gorm.Model
	UserID uint    `gorm:"not null"`
	Name   string  `gorm:"not null"`
	Amount float64 `gorm:"default:0"`
}
