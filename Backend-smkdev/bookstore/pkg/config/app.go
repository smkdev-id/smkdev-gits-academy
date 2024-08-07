package config

import (
	"bookstore/pkg/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = database.AutoMigrate(&models.Kategori{}, &models.Book{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
}
