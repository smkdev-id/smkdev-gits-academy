package config

import (
	"log"

	"Book-Store/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB * gorm.DB

//membuka koneksi ke data base SQLite
func InitDB(){
	var err error
	DB, err = gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to database", err)
	}
	//Migrasi Schema
DB.AutoMigrate(&models.Book{})

}



