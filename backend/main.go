package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	nama   string
	alamat string
}

func main() {
	// Inisialisasi Echo
	e := echo.New()

	// Buka koneksi ke database SQLite
	db, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrasi schema
	db.AutoMigrate(&User{})

	// Handler untuk route utama
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Handler untuk mengambil data dari database
	e.GET("/users", func(c echo.Context) error {
		var users []User
		if result := db.Find(&users); result.Error != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
		}
		return c.JSON(http.StatusOK, users)
	})

	// Jalankan server pada port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
