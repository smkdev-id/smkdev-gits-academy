package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID        uint        // Foreign key referring to User (Buyer)
	ProductID     uint        // Foreign key referring to Product
	PaymentID     uint        // Foreign key referring to Payment
	ShippingID    uint        // Foreign key referring to Shipping
	OrderStatusID uint        // Foreign key referring to OrderStatus
	Quantity      int         `gorm:"not null"`
	TotalPrice    float64     `gorm:"not null"`
	Status        string      `gorm:"type:varchar(50);not null"` // E.g., Pending, Shipped, Delivered
	Payment       Payment     `gorm:"foreignKey:PaymentID"`      // Relationship with Payment
	OrderStatus   OrderStatus `gorm:"foreignKey:OrderStatusID"`  // Relationship with OrderStatus
	Shipping      Shipping    `gorm:"foreignKey:ShippingID"`     // Relationship with Shippingz
}
