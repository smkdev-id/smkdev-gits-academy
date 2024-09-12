package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID    uint   // Foreign key referring to User (Buyer)
	ProductID uint   // Foreign key referring to Product
	Rating    int    `gorm:"not null"` // Rating (1-5)
	Comment   string `gorm:"type:text"`
}
