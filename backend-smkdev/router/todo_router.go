package router

import (
	"database/sql"
	"todos/handler"
	"todos/repository"
	"todos/service"

	"github.com/labstack/echo/v4"
)

func TodoRouter(e *echo.Echo, db *sql.DB) {
	repo := repository.NewTodoRepository(db)
	service := service.NewTodoService(repo)
	handlers := handler.NewTodoHandler(service)
	g := e.Group("/api/v1")
	g.GET("/todo", handlers.GetAll)
	g.GET("/todo/:id", handlers.GetByID)
	g.POST("/todo", handlers.Create)
	g.PUT("/todo/:id", handlers.Update)
	g.DELETE("/todo/:id", handlers.Delete)
}
