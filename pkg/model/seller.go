package model

import "gorm.io/gorm"

// Seller represnts the seller details stored in the database.
type Seller struct {
	gorm.Model
	UserID    uint `gorm:"not null;unique"`
	// Rating    uint
	// Review    string
	SoldCount uint `gorm:"default:0"`
}


