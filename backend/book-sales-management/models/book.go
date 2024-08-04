package models

import "gorm.io/gorm"

// Book represents the book model
type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Price       float64 `json:"price"`
	Publisher   string  `json:"publisher"`
	Genre       string  `json:"genre"`
	PageCount   int     `json:"page_count"`
	Language    string  `json:"language"`
}
