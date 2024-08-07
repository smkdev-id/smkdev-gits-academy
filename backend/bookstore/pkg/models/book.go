package models

import "time"

type Book struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Author    string    `json:"author"`
    Publisher string    `json:"publisher"`
    Copies    int       `json:"copies"`
    Year      int       `json:"year"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
