package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
