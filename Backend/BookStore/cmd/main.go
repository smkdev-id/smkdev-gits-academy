package main

import (
	"BookStore/pkg/config"
	"BookStore/pkg/models"
	"BookStore/pkg/routes"
	"BookStore/pkg/server"
	"BookStore/pkg/utils"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	// Membaca konfigurasi dari file .env
	cfg, err := config.NewConfig(".env")
	checkError(err)

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
	runServer(e, cfg.Port)

	// Menunggu dan menangani shutdown server dengan baik
	waitForShutdown(e)
}

// Fungsi untuk memeriksa dan menangani error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Fungsi untuk menjalankan server dalam goroutine terpisah
func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

// Fungsi untuk menangani shutdown server dengan aman
func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}
