package routes

import (
	"gasruk/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func OrderRoutes(e *echo.Echo) {
	// Routes
	e.POST("/orders", controllers.CreateOrder)       // register data Order
	e.GET("/orders", controllers.GetOrders)          // Tampilkan semua data Order
	e.GET("/orders/:id", controllers.GetOrderByID)   // Select data Order
	e.PUT("/orders/:id", controllers.UpdateOrder)    // Edit data Order
	e.DELETE("/orders/:id", controllers.DeleteOrder) // Delete data Order
}
