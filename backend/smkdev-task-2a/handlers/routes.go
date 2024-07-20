package handlers

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

// URL Routes
func Route(e *echo.Echo, db *sql.DB) {
	e.GET("/", getAllTodoHandler(db))
	e.POST("/", createTodoHandler(db))
	e.GET("/:id", getTodoHandler(db))
	e.PUT("/:id", updateTodoHandler(db))
	e.DELETE("/:id", deleteTodoHandler(db))
}

