package models

import (
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type Book struct {
    ID        uint           `gorm:"primaryKey"`
    Title     string         `json:"title" validate:"required"`
    Description  string      `json:"description" gorm:"type:text" validate:"required"`
	PageCount   int       	 `json:"page_count" validate:"required,gte=0"`
    Author    string         `json:"author" validate:"required"`
    Price     int            `json:"price" validate:"required,gt=0"`
    Quantity  int            `json:"quantity" validate:"required,gte=0"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
}

// FormatPrice formats the price as Rupiah
func (b Book) FormatPrice() string {
    p := message.NewPrinter(language.Indonesian)
    return p.Sprintf("Rp %d", b.Price)
}

// FormatDate formats time to Indonesian date and time format
func FormatDate(t time.Time) string {
    return t.Format("02 Januari 2006 15:04")
}

