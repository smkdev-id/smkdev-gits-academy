package database

import (
	"database/sql"
	"fmt"

	// Import driver SQLite dari modernc.org/sqlite
	_ "modernc.org/sqlite"
)

// ConnectDB membuat koneksi ke database SQLite dan mengembalikan objek *sql.DB
func ConnectDB() (*sql.DB, error) {
	// Membuka koneksi ke database SQLite yang ada di path "database/todos.db"
	db, err := sql.Open("sqlite", "database/todos.db")
	// Handle error saat membuka koneksi, kembalikan error
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}
	// Return objek *sql.DB dan error (nil jika tidak ada kesalahan)
	return db, nil
}

// CreateTodoTable membuat tabel 'todos' jika belum ada dalam database
func CreateTodoTable(db *sql.DB) error {
	// Eksekusi query SQL untuk membuat tabel 'todos' dengan kolom id, title, description, completed, created_at, dan updated_at
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS todos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            title TEXT,
            description TEXT,
            completed BOOLEAN,
            created_at TIMESTAMP,
            updated_at TIMESTAMP
        )
    `)
	// Handle error saat membuka koneksi, kembalikan error
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}
	// Jika berhasil, kembalikan nil (tidak ada error)
	return nil
}
