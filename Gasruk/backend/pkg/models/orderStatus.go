package models

import (
	"gorm.io/gorm"
)

type OrderStatus struct {
	gorm.Model
	Name string `gorm:"size:100"` // Nama status seperti "Pending", "Shipped", "Delivered", dll.
}
