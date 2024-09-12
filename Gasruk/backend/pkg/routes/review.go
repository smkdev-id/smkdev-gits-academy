package routes

import (
	"gasruk/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func ReviewRoutes(e *echo.Echo) {
	// Routes
	e.POST("/reviews", controllers.CreateReview)       // register data Review
	e.GET("/reviews", controllers.GetReviews)          // Tampilkan semua data Review
	e.GET("/reviews/:id", controllers.GetReviewByID)   // Select data Review
	e.PUT("/reviews/:id", controllers.UpdateReview)    // Edit data Review
	e.DELETE("/reviews/:id", controllers.DeleteReview) // Delete data Review
}
