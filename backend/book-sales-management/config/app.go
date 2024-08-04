package config

import (
	"book-sales-management/models"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database
func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
