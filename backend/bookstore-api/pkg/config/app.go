package config

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		log.Println("Database Not Found. Creating one....")
		file, err := os.Create("bookstore.db") // Create SQLite file
		if err != nil {
			log.Println("Failed to create a database")
		}
		file.Close()
	}
	
	return db, nil
}
