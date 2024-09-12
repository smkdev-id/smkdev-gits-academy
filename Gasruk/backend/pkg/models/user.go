package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string `gorm:"type:varchar(100);not null"`
	LastName       string `gorm:"type:varchar(100);not null"`
	Email          string `gorm:"type:varchar(100);unique;not null"`
	Password       string `gorm:"type:varchar(255);not null" json:"-"`
	PhoneNumber    string `gorm:"type:varchar(20)"`
	Gender         string `gorm:"type:varchar(10)"`
	BirthDate      time.Time
	ProfilePicture string     `gorm:"type:varchar(255)"`
	Role           string     `gorm:"type:varchar(20);not null"` // Role: Buyer/Seller
	Merchants      []Merchant `gorm:"foreignKey:UserID"`         // One-to-many relationship with Merchant
	Orders         []Order    `gorm:"foreignKey:UserID"`         // One-to-many relationship with Order (for Buyers)
}
