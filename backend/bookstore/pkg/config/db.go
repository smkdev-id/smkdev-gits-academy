package config

import (
	"bookstore/pkg/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	return db
}
