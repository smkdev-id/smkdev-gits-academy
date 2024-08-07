package config

import (
	"bookstore/pkg/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }

    DB.AutoMigrate(&models.Book{})
}
