package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name              string   `gorm:"type:varchar(100);not null"`
	Description       string   `gorm:"type:text"`
	Price             float64  `gorm:"not null"`
	Stock             int      `gorm:"not null"`
	MerchantID        uint     // Foreign key referring to Merchant
	Category          string   `gorm:"type:varchar(50)"` // Could be Digital, Physical, etc.
	SystemRequirement string   `gorm:"type:text"`        // Possible
	GameDetails       string   `gorm:"type:text"`
	Leanguages        string   `gorm:"type:varchar(100)"`
	Orders            []Order  `gorm:"foreignKey:ProductID"` // One-to-many relationship with Order
	Riviews           []Review `gorm:"foreignKey:ProductID"` // One-to-many relationship with Order
}
