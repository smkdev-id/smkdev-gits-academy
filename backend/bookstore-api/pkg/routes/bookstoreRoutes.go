package routes

import (
	"bookstore-api/pkg/controllers"

	"github.com/labstack/echo/v4"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func BookRoutes(bookcontroller *controllers.BookController) []*Route {
	return []*Route{
		{
			Method:  "GET",
			Path:    "/",
			Handler: bookcontroller.GetAllBooks,
		},
	}
}
