package config

import (
	"BookStore/pkg/models"
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	dbPath := filepath.Join("pkg", "database", "bookstore.db")

	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return nil, err
	}

	// Migrate the schema
	if err := DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Transaction{}, &models.Basket{}, &models.BasketItem{}, &models.Transaction{}, &models.OrderItem{}); err != nil {
		log.Fatalf("Could not migrate database schema: %v", err)
		return nil, err
	}

	return DB, nil
}
