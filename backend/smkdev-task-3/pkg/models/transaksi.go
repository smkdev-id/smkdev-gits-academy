package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID   `json:"id" gorm:"primaryKey"`
	CustomerID uuid.UUID   `json:"customer_id" gorm:"type:uuid;not null"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:TransactionID"`
	TotalPrice float64     `json:"total_price" gorm:"type:decimal(10,2);not null"`
	Status     string      `json:"status" gorm:"type:varchar(50);not null"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
type OrderItem struct {
	ID            uuid.UUID `json:"id" gorm:"primaryKey"`
	TransactionID uuid.UUID `json:"transaction_id" gorm:"type:uuid;not null"`
	BookID        uuid.UUID `json:"book_id" gorm:"type:uuid;not null"`
	Quantity      int       `json:"quantity" gorm:"type:int;not null"`
	Price         float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Book          Book      `json:"book" gorm:"foreignKey:BookID"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
