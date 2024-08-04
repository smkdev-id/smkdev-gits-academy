package models

import "time"

// Book represents a book entry in the database.
type Book struct {
	ID              string `gorm:"type:varchar;primaryKey" json:"id"`
	Title           string `gorm:"type:varchar(255);not null" json:"title"`
	Author          string `gorm:"type:varchar(255);not null" json:"author"`
	PublicationYear string `gorm:"type:varchar" json:"publication_year"`
	Price           int    `gorm:"type:int" json:"price"`
	Stock           int    `gorm:"type:int" json:"stock"`
	ISBN            string `gorm:"type:varchar" json:"isbn"`
	Synopsis        string `gorm:"type:varchar" json:"synopsis"`
	CreatedAt       time.Time `gorm:"type:datetime;default:current_timestamp" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"type:datetime;default:current_timestamp" json:"updatedAt"`
}

// TableName in database
func (Book) TableName() string {
	return "books"
}
