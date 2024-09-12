package routes

import (
	"gasruk/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo) {
	// Routes
	e.POST("/products", controllers.CreateProduct)       // register data Product
	e.GET("/products", controllers.GetProducts)          // Tampilkan semua data Product
	e.GET("/products/:id", controllers.GetProductByID)   // Select data Product
	e.PUT("/products/:id", controllers.UpdateProduct)    // Edit data Product
	e.DELETE("/products/:id", controllers.DeleteProduct) // Delete data user
}
