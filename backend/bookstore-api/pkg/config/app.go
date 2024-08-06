package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database/bookstore.db"), &gorm.Config{})
	if err != nil {
		log.Println("Database Not Found. Creating one....")
		file, err := os.Create("database/bookstore.db")
		if err != nil {
			log.Println("Failed to create a database")
		}
		file.Close()
	}

	return db, nil
}
