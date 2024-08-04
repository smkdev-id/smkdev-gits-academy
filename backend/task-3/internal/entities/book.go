package entities

import "time"

type (
	Book struct {
		ID                string     `gorm:"primaryKey;type:varchar(55);column:id;unique;not null" json:"id"`
		ISBN              string     `gorm:"type:varchar(25);column:isbn;unique;not null" json:"isbn"`
		Title             string     `gorm:"type:varchar(255);column:title;not null" json:"book_title"`
		Description       string     `gorm:"type:varchar(555);column:description" json:"book_description"`
		Author            string     `gorm:"type:varchar(255);column:author;not null" json:"author"`
		YearOfManufacture int        `gorm:"type:int;column:year_of_manufacture;not null" json:"year_of_manufacture"`
		Stock             int        `gorm:"type:int;column:stock;not null" json:"stock"`
		Price             int        `gorm:"type:int;column:price;not null" json:"price"`
		IsDisplayed       bool       `gorm:"type:bool;column:is_displayed;default:false" json:"is_displayed"`
		CreatedAt         time.Time  `gorm:"column:created_at;not null" json:"created_at"`
		UpdatedAt         *time.Time `gorm:"autoUpdateTime:false;column:updated_at" json:"updated_at"`
	}
)

type Books []Book

func (Book) TableName() string {
	return "books"
}
