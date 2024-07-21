package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"todolist/models"
	"todolist/services"
	"todolist/utils"

	"github.com/labstack/echo/v4"
)

// GetTodos mengambil semua todo dari service dan mengembalikannya dalam respons JSON.
func GetTodos(c echo.Context) error {
	// Memanggil service untuk mendapatkan semua todo.
	todos, err := services.GetTodos()
	if err != nil {
		// Jika terjadi kesalahan, kembalikan respons error.
		return utils.NewErrorResponse(c, "Gagal mengambil todos", http.StatusInternalServerError)
	}
	// Jika sukses, kembalikan respons sukses dengan data todo.
	return utils.NewSuccessResponse(c, "Sukses mendapatkan todos", todos, http.StatusOK)
}

// CreateTodo membuat todo baru dari data yang diterima dalam request body.
func CreateTodo(c echo.Context) error {
	todo := new(models.Todo)
	// Mengikat data dari request body ke struct Todo.
	if err := c.Bind(todo); err != nil {
		// Jika terjadi kesalahan saat mengikat data, kembalikan respons error.
		return utils.NewErrorResponse(c, "Gagal membaca request body", http.StatusBadRequest)
	}
	// Memanggil service untuk membuat todo baru.
	createTodo, err := services.CreateTodo(*todo)
	if err != nil {
		// Jika terjadi kesalahan saat membuat todo baru, kembalikan respons error.
		return utils.NewErrorResponse(c, "Gagal membuat todo baru", http.StatusInternalServerError)
	}
	// Jika sukses, kembalikan respons sukses dengan data todo yang baru dibuat.
	return utils.NewSuccessResponse(c, "Sukses membuat todo baru", createTodo, http.StatusCreated)
}

// UpdateTodo memperbarui todo yang ada berdasarkan ID yang diterima dalam parameter URL.
func UpdateTodo(c echo.Context) error {
	// Mengonversi parameter ID dari URL menjadi integer.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Jika ID tidak valid, kembalikan respons error.
		return utils.NewErrorResponse(c, "Id tidak Valid", http.StatusBadRequest)
	}
	// Mengikat data dari request body ke struct Todo.
	todo := new(models.Todo)
	if err := c.Bind(todo); err != nil {
		// Jika terjadi kesalahan saat mengikat data, kembalikan respons error.
		return utils.NewErrorResponse(c, "Gagal membaca request body", http.StatusBadRequest)
	}
	// Memanggil service untuk memperbarui todo berdasarkan ID.
	updatedTodo, err := services.UpdateTodo(id, *todo)
	if err == sql.ErrNoRows {
		// Jika todo tidak ditemukan, kembalikan respons error.
		return utils.NewErrorResponse(c, "Todo tidak ditemukan", http.StatusNotFound)
	} else if err != nil {
		// Jika terjadi kesalahan saat memperbarui todo, kembalikan respons error.
		return utils.NewErrorResponse(c, "Gagal memperbarui todo", http.StatusInternalServerError)
	} else if updatedTodo == (models.Todo{}) {
		// Jika tidak ada perubahan yang dilakukan, kembalikan respons sukses tanpa data.
		return utils.NewSuccessResponse(c, "Tidak ada perubahan yang dilakukan", nil, http.StatusOK)
	}
	// Jika sukses, kembalikan respons sukses dengan data todo yang diperbarui.
	return utils.NewSuccessResponse(c, "Sukses memperbarui todo", updatedTodo, http.StatusOK)
}

// DeleteTodo menghapus todo berdasarkan ID yang diterima dalam parameter URL.
func DeleteTodo(c echo.Context) error {
	// Mengonversi parameter ID dari URL menjadi integer.
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Jika ID tidak valid, kembalikan respons error.
		return utils.NewErrorResponse(c, "ID tidak valid", http.StatusBadRequest)
	}
	// Memanggil service untuk menghapus todo berdasarkan ID.
	err = services.DeleteTodo(id)
	if err == sql.ErrNoRows {
		// Jika todo tidak ditemukan, kembalikan respons error.
		return utils.NewErrorResponse(c, "Todo tidak ditemukan", http.StatusNotFound)
	} else if err != nil {
		// Jika terjadi kesalahan saat menghapus todo, kembalikan respons error.
		return utils.NewErrorResponse(c, "Gagal menghapus todo", http.StatusInternalServerError)
	}
	// Jika sukses, kembalikan respons sukses tanpa data.
	return utils.NewSuccessResponse(c, "Todo Berhasil Dihapus", nil, http.StatusNoContent)
}
