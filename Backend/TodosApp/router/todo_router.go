package router

import (
	"database/sql"
	"todos/handler"
	"todos/repository"
	"todos/service"

	"github.com/labstack/echo/v4"
)

// TodoRouter menginisialisasi rute Todo dan menghubungkannya dengan handler
func TodoRouter(e *echo.Echo, db *sql.DB) {
	// Membuat instance baru dari TodoRepository yang membutuhkan database
	repo := repository.NewTodoRepository(db)
	// Membuat instance baru dari TodoService yang membutuhkan repository
	service := service.NewTodoService(repo)
	// Membuat instance baru dari TodoHandler yang membutuhkan service
	handlers := handler.NewTodoHandler(service)

	// Membuat group untuk endpoint Todo
	g := e.Group("/api/v1")

	// Menetapkan rute-rute Todo ke handler yang sesuai
	g.GET("/todo", handlers.GetAll)        // Mendapatkan semua todo
	g.GET("/todo/:id", handlers.GetByID)   // Mendapatkan todo berdasarkan ID
	g.POST("/todo", handlers.Create)       // Membuat todo baru
	g.PUT("/todo/:id", handlers.Update)    // Memperbarui todo berdasarkan ID
	g.DELETE("/todo/:id", handlers.Delete) // Menghapus todo berdasarkan ID
}
