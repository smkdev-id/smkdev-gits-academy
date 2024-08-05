package models

type Book struct {
	ID  		uint 	`gorm:"primaryKey" json:"id"`
	Title 		string 	`json:"title"`
	Author 		string  `json:"author"`
	Publisher	string	`json:"publisher"`
	Price		float64 `json:"price"`
	Description string 	`json:"description"`
}