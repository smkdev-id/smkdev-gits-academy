package handler

import (
	"net/http"
	"strconv"
	"todos/dto"
	"todos/helper"
	"todos/service"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	service service.TodoServiceInteface
}

func NewTodoHandler(service service.TodoServiceInteface) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

func (h *TodoHandler) GetAll(c echo.Context) error {
	var todosRes []*dto.TodoResponse
	todos, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}
	for _, todo := range todos {
		todoResponses := dto.EntityToResponse(todo)
		todosRes = append(todosRes, todoResponses)
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success get all todos",
		Data:       todosRes,
	})
	return c.JSON(http.StatusOK, res)
}

func (h *TodoHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	todo, err := h.service.GetByID(c.Request().Context(), idInt)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}
	todoResponses := dto.EntityToResponse(todo)
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success get todo by id",
		Data:       todoResponses,
	})
	return c.JSON(http.StatusOK, res)
}

func (h *TodoHandler) Create(c echo.Context) error {
	todoReq := new(dto.TodoRequest)
	if err := c.Bind(&todoReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	todo := todoReq.ToEntity()
	if err := h.service.Create(c.Request().Context(), todo); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success Create Todo",
	})
	return c.JSON(http.StatusOK, res)
}

func (h *TodoHandler) Update(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	todoReq := new(dto.TodoRequest)
	if err := c.Bind(&todoReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	todo := todoReq.ToEntity()
	if err := h.service.Update(c.Request().Context(), todo, idInt); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success update todo",
	})
	return c.JSON(http.StatusOK, res)
}

func (h *TodoHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	if err := h.service.Delete(c.Request().Context(), idInt); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success delete todo",
	})
	return c.JSON(http.StatusOK, res)
}
