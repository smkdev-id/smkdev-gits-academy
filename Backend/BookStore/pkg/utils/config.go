package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv memuat variabel lingkungan dari file .env.
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file") // Menghentikan aplikasi jika terjadi kesalahan saat memuat file .env
	}
}

// GetEnv mengambil nilai env berdasarkan key.
func GetEnv(key string) string {
	return os.Getenv(key) // Mengembalikan nilai variabel lingkungan untuk kunci yang diberikan
}
