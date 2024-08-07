package main

import (
	"BookStore/pkg/config"
	"BookStore/pkg/models"
	"BookStore/pkg/routes"
	"BookStore/pkg/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	// Memuat variabel lingkungan dari file .env
	utils.LoadEnv()

	// Inisialisasi fungsionalitas logging
	utils.InitLogger()

	// Inisialisasi validator untuk validasi data
	utils.InitValidator()

	// Menghubungkan ke database dan melakukan migrasi otomatis
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Book{})

	// Membuat instance baru dari Echo untuk server HTTP
	e := echo.New()

	// Mengatur middleware untuk instance Echo
	utils.SetupMiddleware(e)

	// Mengatur rute untuk instance Echo
	routes.SetupRoutes(e)

	// Memulai server dan mendengarkan pada port yang ditentukan
	RunApp(e)
}

func RunApp(e *echo.Echo) {
	// Mengambil port dari env
	port := utils.GetEnv("PORT")
	// Memformat port untuk digunakan dalam startup server
	strPort := fmt.Sprintf(":%v", port)
	// Memulai server Echo dan mencatat kesalahan fatal
	e.Logger.Fatal(e.Start(strPort))
}
