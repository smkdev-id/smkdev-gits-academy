package config

import (
	"gasruk/pkg/models"
	"log"

	_ "github.com/lib/pq" // Import driver PostgreSQL
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase creates the database if it doesn't exist and then connects to it.

func ConnectDatabase() *gorm.DB {
	var err error

	dsn := "host=localhost user=postgres password=cimapag1 dbname=gasruk port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Menjalankan query sederhana untuk mengecek koneksi
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the PostgreSQL database.")

	// Migrasi model
	err = DB.AutoMigrate(&models.User{}, &models.Merchant{}, &models.Order{}, &models.OrderStatus{}, &models.Payment{}, &models.Product{}, &models.Review{}, &models.Shipping{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return DB
}
