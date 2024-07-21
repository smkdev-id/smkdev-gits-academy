package routes

import (
	"todolist/controllers"

	"github.com/labstack/echo/v4"
)

// Init menginisialisasi semua route yang tersedia dalam aplikasi
// Fungsi ini menerima instance Echo sebagai parameter dan mendefinisikan route beserta handler yang terkait
func Init(e *echo.Echo) {
	// Route untuk mengambil semua todo list
	e.GET("/todos", controllers.GetTodos)

	// Route untuk membuat todo list baru
	e.POST("/todos", controllers.CreateTodo)

	// Route untuk mengupdate todo list berdasarkan ID
	e.PATCH("/todos/:id", controllers.UpdateTodo)

	// Route untuk menghapus todo list berdasarkan ID
	e.DELETE("/todos/:id", controllers.DeleteTodo)
}
