package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connection database
func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")

	var dialector gorm.Dialector
	switch dbDriver {
	case "sqlite":
		dialector = sqlite.Open(dbName)
	default:
		log.Fatal("Unsupported DB_DRIVER")
	}

	database, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
}
