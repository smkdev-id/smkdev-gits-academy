package models

import (
	"gorm.io/gorm"
)

// Book structure 
type Book struct {
	gorm.Model
	Title  string
	Author string
	Price  uint
}
