package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `gorm:"unique"`
	Author string
	Price  float64
}
