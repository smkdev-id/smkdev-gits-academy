package config

import (
	"log"
	"todo-list/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	return db
}
