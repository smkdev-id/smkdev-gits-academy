package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Method string  `gorm:"type:varchar(50);not null"` // Payment method (e.g., Credit Card, PayPal)
	Orders []Order `gorm:"foreignKey:PaymentID"`
}
