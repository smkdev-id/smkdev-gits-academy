package config

import (
	"go-bookstore/pkg/models"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error

      // Create pkg/config directory if it doesn't exist
	  configDir := "pkg/config"
    if _, err := os.Stat(configDir); os.IsNotExist(err) {
        err = os.Mkdir(configDir, os.ModePerm)
        if err != nil {
            log.Fatal("failed to create config directory")
        }
    }

    // Database file path
    dbPath := filepath.Join(configDir, "bookstore.db")

    DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database")
    }

    DB.AutoMigrate(&models.Book{})
}
