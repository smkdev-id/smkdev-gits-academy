package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

// Fungsi untuk menjalankan server dalam goroutine terpisah
func RunServer(srv *Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

// Fungsi untuk menangani shutdown server dengan aman
func WaitForShutdown(srv *Server) {
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
