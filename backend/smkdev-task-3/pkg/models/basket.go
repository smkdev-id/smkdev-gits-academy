package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Basket struct {
	ID         uuid.UUID      `json:"id" gorm:"primaryKey"`
	CustomerID uuid.UUID      `json:"customer_id" gorm:"type:uuid;not null"`
	Items      []BasketItem   `json:"items" gorm:"foreignKey:BasketID"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type BasketItem struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	BasketID  uuid.UUID `json:"basket_id" gorm:"type:uuid;not null"`
	BookID    uuid.UUID `json:"book_id" gorm:"type:uuid;not null"`
	Quantity  int       `json:"quantity" gorm:"type:int;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Book      Book      `json:"book" gorm:"foreignKey:BookID"`
}
