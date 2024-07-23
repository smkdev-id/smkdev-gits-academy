package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Inisialisasi Echo
	e := echo.New()

	// Buka koneksi ke database SQLite
	db, err := sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// Handler untuk route utama
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Handler untuk mengambil data dari database
	e.GET("/users", func(c echo.Context) error {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			return err
		}
		defer rows.Close()

		var users []map[string]interface{}
		for rows.Next() {
			var id int
			var name string
			err := rows.Scan(&id, &name)
			if err != nil {
				return err
			}
			user := map[string]interface{}{
				"id":   id,
				"name": name,
			}
			users = append(users, user)
		}
		return c.JSON(http.StatusOK, users)
	})

	// Jalankan server pada port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
