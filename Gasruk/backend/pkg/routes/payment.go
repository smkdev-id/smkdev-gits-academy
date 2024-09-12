package routes

import (
	"gasruk/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func PaymentRoutes(e *echo.Echo) {
	// Routes
	e.POST("/payments", controllers.CreatePayment)       // register data Payment
	e.GET("/payments", controllers.GetPayments)          // Tampilkan semua data Payment
	e.GET("/payments/:id", controllers.GetPaymentByID)   // Select data Payment
	e.PUT("/payments/:id", controllers.UpdatePayment)    // Edit data Payment
	e.DELETE("/payments/:id", controllers.DeletePayment) // Delete data Payment
}
