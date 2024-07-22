package controller

import (
	"github.com/gin-gonic/gin"
)

type TodoController interface {
	CreateTodo(c *gin.Context)
	GetTodoByID(c *gin.Context)
	GetAllTodos(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}
