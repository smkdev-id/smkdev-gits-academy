package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB menginisialisasi koneksi ke database SQLite3 dan membuat tabel jika belum ada
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal("Gagal membuka koneksi ke database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Gagal memeriksa koneksi ke database:", err)
	}

	log.Println("Berhasil koneksi ke database")

	return db
}
