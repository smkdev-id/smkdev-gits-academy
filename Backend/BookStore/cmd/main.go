package main

import (
	"BookStore/pkg/config"
	"BookStore/pkg/models"
	"BookStore/pkg/routes"
	"BookStore/pkg/server"
	"BookStore/pkg/utils"
)

func main() {

	// Membaca konfigurasi dari file .env
	cfg, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}
	// Inisialisasi logger
	utils.InitLogger()

	// Inisialisasi validator
	utils.InitValidator()

	// Koneksi ke database dan melakukan migrasi untuk model Book
	models.ConnectDatabase(cfg.DbName)
	models.DB.AutoMigrate(&models.Book{})

	// Membuat instance server baru
	e := server.NewServer()

	// Mengatur rute-rute API
	routes.SetupRoutes(e)

	// Menjalankan server
	server.RunServer(e, cfg.Port)

	// Menunggu dan menangani shutdown server dengan baik
	server.WaitForShutdown(e)
}
