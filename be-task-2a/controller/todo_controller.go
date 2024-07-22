package controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/HerlyRyan/be-task-2a/model"
	"github.com/HerlyRyan/be-task-2a/repository"
)

type todoController struct {
	repo repository.TodoRepository
}

func NewTodoController(repo repository.TodoRepository) TodoController {
	return &todoController{repo: repo}
}

func (ctrl *todoController) CreateTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.repo.CreateTodo(c.Request.Context(), &todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func (ctrl *todoController) GetTodoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	todo, err := ctrl.repo.GetTodoByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (ctrl *todoController) GetAllTodos(c *gin.Context) {
	todos, err := ctrl.repo.GetAllTodos(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func (ctrl *todoController) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo.ID = id

	if err := ctrl.repo.UpdateTodo(c.Request.Context(), &todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func (ctrl *todoController) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.repo.DeleteTodo(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
