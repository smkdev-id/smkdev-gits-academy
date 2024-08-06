package models

import "time"

type Book struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	NamaBuku  string    `json:"nama_buku"`
	Penulis   string    `json:"nama_penulis"`
	Penerbit  string    `json:"penerbit"`
	Stock     int       `json:"stock"`
	Harga     int       `json:"harga_buku"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Book) TableName() string {
	return "books"
}
