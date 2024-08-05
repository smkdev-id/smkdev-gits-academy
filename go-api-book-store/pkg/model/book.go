package models

type Book struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
	Year        int    `json:"year"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}
