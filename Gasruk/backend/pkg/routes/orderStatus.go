package routes

import (
	"gasruk/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func OrderStatusRoutes(e *echo.Echo) {
	// Routes
	e.POST("/orderStatuses", controllers.CreateOrderStatus)       // register data OrderStatus
	e.GET("/orderStatuses", controllers.GetOrderStatuses)         // Tampilkan semua data OrderStatus
	e.GET("/orderStatuses/:id", controllers.GetOrderStatusByID)   // Select data OrderStatus
	e.PUT("/orderStatuses/:id", controllers.UpdateOrderStatus)    // Edit data OrderStatus
	e.DELETE("/orderStatuses/:id", controllers.DeleteOrderStatus) // Delete data OrderStatus
}
