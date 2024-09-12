package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	UserID      uint      // Foreign key referring to User
	ProductSold int64
	TotalProduct int64
	Products    []Product `gorm:"foreignKey:MerchantID"` // One-to-many relationship with Product
}
