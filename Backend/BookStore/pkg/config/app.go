package config

import (
	"BookStore/pkg/utils"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase menghubungkan ke database SQLite.
func ConnectDatabase() {
	var err error
	// Ambil nama database dari env
	dbName := utils.GetEnv("DB_NAME")
	// Bentuk string format di simpannya database
	strDB := fmt.Sprintf("pkg/models/%v.db", dbName)
	// Buka koneksi ke database
	DB, err = gorm.Open(sqlite.Open(strDB), &gorm.Config{})
	if err != nil {
		// Log error jika gagal
		log.Fatal("Gagal menghubungkan ke database!", err)
	}
}
