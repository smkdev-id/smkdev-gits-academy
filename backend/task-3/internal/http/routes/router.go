package routes

import (
	"net/http"

	"github.com/alwiirfan/internal/http/controllers"
	"github.com/labstack/echo/v4"
)

type Route struct {
	Path    string
	Method  string
	Handler echo.HandlerFunc
}

func PublicRoutes(bookContrller *controllers.BookController) []*Route {
	return []*Route{
		{
			Path:    "/api/v1/books",
			Method:  http.MethodPost,
			Handler: bookContrller.Create,
		},
	}
}

func PrivateRoutes(bookController *controllers.BookController) []*Route {
	return []*Route{}
}
