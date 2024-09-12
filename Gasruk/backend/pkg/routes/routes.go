package routes

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	UserRoutes(e)
	MerchantRoutes(e)
	ProductRoutes(e)
	ReviewRoutes(e)
}
