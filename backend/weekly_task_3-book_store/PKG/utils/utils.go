package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPErrorHandler is a custom error handler for Echo framework
func HTTPErrorHandler(err error, c echo.Context) {
	var code int
	var message interface{}

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
		if m, ok := he.Message.(string); ok {
			message = map[string]string{"error": m}
		}
	} else {
		code = http.StatusInternalServerError
		message = map[string]string{"error": http.StatusText(code)}
	}

	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, message)
		}
		if err != nil {
			c.Logger().Error(err)
		}
	}
}

func ResponseError(c echo.Context, status int, message string) error {
	return c.JSON(status, map[string]string{"error": message})
}
