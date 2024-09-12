package routes

import (
	"gasruk/pkg/controllers"

	"github.com/labstack/echo/v4"
)

func MerchantRoutes(e *echo.Echo) {
	e.POST("/merchants", controllers.CreateMerchant)
	e.GET("/merchants", controllers.GetMerchants)
	e.GET("/merchants/:id", controllers.GetMerchantByID)
	e.PUT("/merchants/:id", controllers.UpdateMerchant)
	e.DELETE("/merchants/:id", controllers.DeleteMerchant)
}
