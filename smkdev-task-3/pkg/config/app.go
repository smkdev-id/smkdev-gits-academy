package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "smkdev-task-3/pkg/models"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("pkg/config/bookstore.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrasi schema
    DB.AutoMigrate(&models.Book{})
}
