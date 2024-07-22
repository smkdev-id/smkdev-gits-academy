package handlers

import (
	"go-todo/models"
	"go-todo/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
    Service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
    return &TodoHandler{Service: service}
}

func (h *TodoHandler) FetchTodos(c echo.Context) error {
    todos, err := h.Service.FetchTodos()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }
   
    return c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodos(c echo.Context) error {
    var input models.Todo
    if err := c.Bind(&input); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }
    if err := h.Service.CreateTodos(&input); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, input)
}

func (h *TodoHandler) UpdateTodos(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    todo, err := h.Service.GetTodoByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Todo not found!"})
    }
    if err := c.Bind(todo); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
    }
    if err := h.Service.UpdateTodos(todo); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) DeleteTodos(c echo.Context) error {
    id, _ := strconv.Atoi(c.Param("id"))
    todo, err := h.Service.GetTodoByID(uint(id))
    if err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Todo not found!"})
    }
    if err := h.Service.DeleteTodos(todo); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, echo.Map{"data": true})
}
