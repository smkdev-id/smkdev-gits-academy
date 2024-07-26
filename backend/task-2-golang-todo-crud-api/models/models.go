package models

// Item mewakili data item dalam tabel
type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
