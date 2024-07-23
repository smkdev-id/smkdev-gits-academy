package controller

import (
	"net/http"
	"todo-list/helper"
	"todo-list/service"

	"github.com/labstack/echo/v4"
)

type todoController struct {
	service service.TodoService
}

func NewTodoController(service service.TodoService) *todoController {
	return &todoController{service}
}

func (s todoController) GetAllTodo(c echo.Context) error {
	todos, err := s.service.GetAllTodo()

	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to get all todo", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to get all todo", todos)
	c.JSON(http.StatusOK, response)	
	return nil
}

func (s todoController) GetTodoByID(c echo.Context) error {
	id := c.Param("id")

	todo, err := s.service.GetTodoByID(id)

	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to get todo", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to get todo", todo)
	c.JSON(http.StatusOK, response)	
	return nil
}

func (s todoController) CreateTodo(c echo.Context) error {
	var input helper.Input

	err := c.Bind(&input)

	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to parse JSON data", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	todo, err := s.service.CreateTodo(input)
	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to create todo", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to create todo", todo)
	c.JSON(http.StatusOK, response)	
	return nil
}

func (s todoController) UpdateTodo(c echo.Context) error {
	var input helper.Input
	
	err := c.Bind(&input)
	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to parse JSON data", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	id := c.Param("id")

	todo, err := s.service.UpdateTodo(id, input)

	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to update todo", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to update todo", todo)
	c.JSON(http.StatusOK, response)	
	return nil
}

func (s todoController) DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	err := s.service.DeleteTodo(id)
	if err != nil {
		response := helper.ResponseJSON(http.StatusBadRequest, "Failed to delete todo", nil)
		c.JSON(http.StatusBadRequest, response)
		return err
	}

	response := helper.ResponseJSON(http.StatusOK, "Success to delete todo", nil)
	c.JSON(http.StatusOK, response)	
	return nil
}

