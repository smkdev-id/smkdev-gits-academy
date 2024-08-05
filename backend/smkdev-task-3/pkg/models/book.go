package models

import (
	"time"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

// type ServiceID struct {
// 	UUID uuid.UUID
// }

type Book struct {
	ID        uuid.UUID      `json:"id" gorm:"primaryKey;"`
	Title     string         `json:"title" gorm:"type:varchar(255);not null"`
	Publisher string         `json:"publisher" gorm:"type:text;not null"`
	Author    string         `json:"author" gorm:"type:text;not null"`
	Price     int            `json:"price" gorm:"type:int;not null"`
	Stock     int            `json:"stock" gorm:"type:int;not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type BookRequest struct {
	ID string `uri:"id" binding:"required"`
	// Title     string `json:"title" binding:"required"`
	// Publisher string `json:"publisher" binding:"required"`
	// Author    string `json:"author"  binding:"required"`
	// Price     int    `json:"price" binding:"required"`
	// Stock     int    `json:"stock" binding:"required"`
}

type CreateBookRequest struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" gorm:"type:varchar(255);not null"`
	Publisher string    `json:"publisher" gorm:"type:text;not null"`
	Author    string    `json:"author" gorm:"type:text;not null"`
	Price     int       `json:"price" gorm:"type:int;not null"`
	Stock     int       `json:"stock" gorm:"type:int;not null"`
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time      `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UpdateBookRequest struct {
	ID        string `uri:"id" binding:"required"`
	Title     string `json:"title" gorm:"type:varchar(255);not null"`
	Publisher string `json:"publisher" gorm:"type:text;not null"`
	Author    string `json:"author" gorm:"type:text;not null"`
	Price     int    `json:"price" gorm:"type:int;not null"`
	Stock     int    `json:"stock" gorm:"type:int;not null"`
	// CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type DeleteBookRequest struct {
	ID string `uri:"id" binding:"required"`
}
