package main

import (
	"todos/database"
	"todos/router"

	"github.com/labstack/echo/v4"
)

func main() {
	// Inisialisasi instance baru dari Echo framework
	e := echo.New()

	// Menghubungkan ke database SQLite
	db, err := database.ConnectDB()
	// Jika gagal menghubungkan ke database, hentikan aplikasi dengan pesan kesalahan
	if err != nil {
		panic("failed to connect database")
	}
	// Menutup koneksi database ketika aplikasi berhenti
	defer db.Close()

	// Membuat tabel 'todos' jika belum ada
	err = database.CreateTodoTable(db)
	// Jika gagal membuat tabel, hentikan aplikasi dengan pesan kesalahan
	if err != nil {
		panic("failed to create table")
	}

	// Mengatur router dan endpoint untuk aplikasi Todo
	router.TodoRouter(e, db)

	// Mulai server pada port 8080 dan log kesalahan jika terjadi
	e.Logger.Fatal(e.Start(":8080"))
}
