package utils

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

type Logger struct {
	*log.Logger
}

func InitializeLogger() *Logger {
	return &Logger{log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)}
}

func (l *Logger) Info(v ...interface{}) {
	l.SetPrefix("INFO: ")
	l.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.SetPrefix("WARN: ")
	l.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.SetPrefix("ERROR: ")
	l.Println(v...)
}

func RespondWithError(c echo.Context, status int, message string) error {
	return c.JSON(status, map[string]string{"error": message})
}

func NewSuccessResponse(c echo.Context, message string, data interface{}, statusCode int) error {
	response := map[string]interface{}{
		"message": message,
		"status":  statusCode,
		"data":    data,
	}
	return c.JSON(statusCode, response)
}