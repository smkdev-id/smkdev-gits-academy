package routes

import (
	"BookStore/pkg/controllers"
	"BookStore/pkg/server"
)

// SetupRoutes mengatur rute API untuk server
func SetupRoutes(e *server.Server) {
	g := e.Group("/api")                          // Buat grup rute dengan prefix /api
	g.GET("/book", controllers.GetBooks)          // Rute untuk mendapatkan semua buku
	g.GET("/book/:id", controllers.GetBookByID)   // Rute untuk mendapatkan buku berdasarkan ID
	g.POST("/book", controllers.CreateBook)       // Rute untuk membuat buku baru
	g.PUT("/book/:id", controllers.UpdateBook)    // Rute untuk memperbarui buku berdasarkan ID
	g.DELETE("/book/:id", controllers.DeleteBook) // Rute untuk menghapus buku berdasarkan ID
}
