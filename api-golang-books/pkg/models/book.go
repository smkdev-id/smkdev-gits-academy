package models

import "time"

// entity or type struct
type Book struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Price     int    `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
