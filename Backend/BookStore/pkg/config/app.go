package config

import (
	"errors"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Struct Config untuk menyimpan konfigurasi aplikasi
type Config struct {
	Port   string `env:"PORT" envDefault:"8080"`         // Port di mana server akan dijalankan
	DbName string `env:"DB_NAME" envDefault:"bookstore"` // Nama database yang akan digunakan
}

// Fungsi NewConfig untuk membuat instance Config baru
func NewConfig(envPath string) (*Config, error) {
	// Memanggil fungsi parseConfig untuk membaca dan mengurai konfigurasi dari file .env
	cfg, err := parseConfig(envPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Fungsi parseConfig untuk membaca file .env dan mengurai variabel-variabel lingkungan
func parseConfig(envPath string) (*Config, error) {
	// Memuat file .env
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, errors.New("failed to load env") // Mengembalikan error jika gagal memuat file .env
	}

	cfg := &Config{}
	// Mengurai variabel-variabel lingkungan ke dalam struct Config
	err = env.Parse(cfg)
	if err != nil {
		return nil, errors.New("failed to parse config") // Mengembalikan error jika gagal mengurai konfigurasi
	}
	return cfg, nil
}
