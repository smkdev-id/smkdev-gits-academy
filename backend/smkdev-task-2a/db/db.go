package db

import (
	"go-todo/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connect to database", err)
	}

	database.AutoMigrate(&models.Todo{})

	DB = database
}