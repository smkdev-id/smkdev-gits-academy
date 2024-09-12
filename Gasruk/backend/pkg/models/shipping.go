package models

import (
	"time"

	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	OrderID        uint   // Foreign key referring to Order
	Address        string `gorm:"type:varchar(255);not null"`
	City           string `gorm:"type:varchar(100);not null"`
	PostalCode     string `gorm:"type:varchar(20);not null"`
	ShippingMethod string `gorm:"type:varchar(50);not null"`
	ShippedDate    time.Time
	DeliveredDate  time.Time
}
