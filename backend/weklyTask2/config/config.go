package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB adalah variabel global untuk koneksi database.
var DB *sql.DB
var err error

// Connect menghubungkan ke database SQLite dan membuat tabel todolist jika belum ada.
func Connect() {
	// Membuka koneksi ke database SQLite dengan file todolist.db.
	DB, err = sql.Open("sqlite3", "./todolist.db")
	if err != nil {
		// Jika terjadi kesalahan saat membuka koneksi, cetak kesalahan dan hentikan program.
		log.Fatal(err)
	}
	// Membuat tabel todolist jika belum ada.
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS todolist (
	
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		todo TEXT,
		completed BOOLEAN
	
	)`)

	if err != nil {
		// Jika terjadi kesalahan saat membuat tabel, cetak kesalahan dan hentikan program.
		log.Fatal(err)
	}
}
