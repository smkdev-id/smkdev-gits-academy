package models

import "time"

type Todo struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Todo string    `json:"todo" gorm:"type:text"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}