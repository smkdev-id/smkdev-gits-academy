package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "bookstore/pkg/models"
)

func SetupDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    // Auto Migrate the schema
    db.AutoMigrate(&models.Book{})
    
    return db, nil
}