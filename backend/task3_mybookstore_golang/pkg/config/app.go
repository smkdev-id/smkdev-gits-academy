package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var (
    DB *gorm.DB
)

func Connect() {
    d, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    DB = d
}

func GetDB() *gorm.DB {
    return DB
}
