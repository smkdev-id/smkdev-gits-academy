package routes

import (
	"gasruk/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	// Routes
	e.POST("/users", controllers.CreateUser)       // register data user
	e.GET("/users", controllers.GetUsers)          // Tampilkan semua data user
	e.GET("/users/:id", controllers.GetUserByID)   // Select data user
	e.PUT("/users/:id", controllers.UpdateUser)    // Edit data user
	e.DELETE("/users/:id", controllers.DeleteUser) // Delete data user
}
