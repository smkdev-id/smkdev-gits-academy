package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key"`
	Name     string    `json:"name"`
	Author   string    `json:"author"`
	Publisher string   `json:"publisher"`
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}
