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

// UpdateUserWallet method on Transaction to update the user's wallet
func (t *Transaction) UpdateUserWallet(db *gorm.DB) error {
    var user User
    if err := db.First(&user, t.UserID).Error; err != nil {
        return err
    }

    user.Wallet += t.Amount
    if err := db.Save(&user).Error; err != nil {
        return err
    }

    return nil
}

// AfterCreate method helps whenever a new transaction is created, update the user's wallet
func (t *Transaction) AfterCreate(tx *gorm.DB) (err error) {
    return t.UpdateUserWallet(tx)
}