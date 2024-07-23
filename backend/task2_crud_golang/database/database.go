package database

import (
    "task2_crud_golang/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    DB.AutoMigrate(&models.Todo{})
}
