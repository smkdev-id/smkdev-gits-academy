package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal terhubung dengan DB:", err)
    }
    log.Println("Berhasil terhubung dengan DB")
}
