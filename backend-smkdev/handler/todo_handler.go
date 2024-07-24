package handler

import (
	"net/http"
	"strconv"
	"todos/dto"
	"todos/helper"
	"todos/service"

	"github.com/labstack/echo/v4"
)

// TodoHandler adalah struct yang menangani request HTTP untuk operasi Todo
type TodoHandler struct {
	service service.TodoServiceInteface
}

// NewTodoHandler membuat instance baru dari TodoHandler
func NewTodoHandler(service service.TodoServiceInteface) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

// GetAll menangani request GET untuk mengambil semua todo
func (h *TodoHandler) GetAll(c echo.Context) error {
	var todosRes []*dto.TodoResponse
	// Memanggil service untuk mendapatkan semua todo
	todos, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}
	// Mengonversi entity todo ke response DTO
	for _, todo := range todos {
		todoResponses := dto.EntityToResponse(todo)
		todosRes = append(todosRes, todoResponses)
	}
	// Membuat response sukses
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success get all todos",
		Data:       todosRes,
	})
	return c.JSON(http.StatusOK, res)
}

// GetByID menangani request GET untuk mengambil todo berdasarkan ID
func (h *TodoHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	// Memanggil service untuk mendapatkan todo berdasarkan ID
	todo, err := h.service.GetByID(c.Request().Context(), idInt)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}
	// Mengonversi entity todo ke response DTO
	todoResponses := dto.EntityToResponse(todo)
	// Membuat response sukses
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success get todo by id",
		Data:       todoResponses,
	})
	return c.JSON(http.StatusOK, res)
}

// Create menangani request POST untuk membuat todo baru
func (h *TodoHandler) Create(c echo.Context) error {
	todoReq := new(dto.TodoRequest)
	// Binding request body ke TodoRequest
	if err := c.Bind(&todoReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// Mengonversi request DTO ke entity todo
	todo := todoReq.ToEntity()
	// Memanggil service untuk membuat todo baru
	if err := h.service.Create(c.Request().Context(), todo); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	// Membuat response sukses
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "success create todo",
	})
	return c.JSON(http.StatusCreated, res)
}

// Update menangani request PUT untuk memperbarui todo berdasarkan ID
func (h *TodoHandler) Update(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	todoReq := new(dto.TodoRequest)
	// Binding request body ke TodoRequest
	if err := c.Bind(&todoReq); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}
	// Mengonversi request DTO ke entity todo
	todo := todoReq.ToEntity()
	// Memanggil service untuk memperbarui todo berdasarkan ID
	if err := h.service.Update(c.Request().Context(), todo, idInt); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	// Membuat response sukses
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success update todo",
	})
	return c.JSON(http.StatusOK, res)
}

// Delete menangani request DELETE untuk menghapus todo berdasarkan ID
func (h *TodoHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	// Memanggil service untuk menghapus todo berdasarkan ID
	if err := h.service.Delete(c.Request().Context(), idInt); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	// Membuat response sukses
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "success delete todo",
	})
	return c.JSON(http.StatusOK, res)
}
