package controllers

import (
	"BookStore/pkg/config"
	"BookStore/pkg/dto"
	"BookStore/pkg/models"
	"BookStore/pkg/utils"
	"fmt"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

// GetBooks mengembalikan daftar semua buku.
func GetBooks(c echo.Context) error {
	var books []*models.Book
	var bookRes []*dto.BookResponse
	// Ambil semua buku dari database
	config.DB.Find(&books)
	// Ubah setiap buku menjadi format respons
	for _, book := range books {
		BookResponses := dto.EntityToResponse(book)
		bookRes = append(bookRes, BookResponses)
	}
	fmt.Println(bookRes)
	// Kembalikan respons sukses dengan data buku
	return utils.SuccessResponse(c, http.StatusOK, bookRes)
}

// GetBookByID mengembalikan buku berdasarkan ID.
func GetBookByID(c echo.Context) error {
	var book *models.Book
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Book not found")
	}
	// Cari buku berdasarkan ID
	if err := config.DB.First(&book, intID).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Book not found")
	}
	// Ubah buku menjadi format respons
	bookRes := dto.EntityToResponse(book)
	// Kembalikan respons sukses dengan data buku
	return utils.SuccessResponse(c, http.StatusOK, bookRes)
}

// CreateBook membuat buku baru.
func CreateBook(c echo.Context) error {
	var book *models.Book
	// Bind data dari request ke model buku
	if err := c.Bind(&book); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	// Validasi input
	if err := utils.ValidateStruct(book); err != nil {
		errors := utils.ParseValidationErrors(err)
		return utils.ErrorResponseValidation(c, http.StatusBadRequest, errors)
	}
	// Simpan buku baru ke database
	config.DB.Create(&book)
	// Kembalikan respons sukses tanpa data
	return utils.SuccessResponseWithoutData(c, http.StatusCreated)
}

// UpdateBook memperbarui buku berdasarkan ID.
func UpdateBook(c echo.Context) error {
	var book *models.Book
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Book not found")
	}
	// Cari buku berdasarkan ID
	if err := config.DB.First(&book, intID).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Book not found")
	}
	// Bind data dari request ke model buku
	if err := c.Bind(&book); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	// Validasi input
	if err := utils.ValidateStruct(book); err != nil {
		errors := utils.ParseValidationErrors(err)
		return utils.ErrorResponseValidation(c, http.StatusBadRequest, errors)
	}
	// Simpan perubahan buku ke database
	config.DB.Save(&book)
	// Kembalikan respons sukses tanpa data
	return utils.SuccessResponseWithoutData(c, http.StatusOK)
}

// DeleteBook menghapus buku berdasarkan ID.
func DeleteBook(c echo.Context) error {
	var book models.Book
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Book not found")
	}
	// Cari buku berdasarkan ID
	if err := config.DB.First(&book, intID).Error; err != nil {
		return utils.ErrorResponse(c, http.StatusNotFound, "Book not found")
	}
	// Hapus buku dari database
	config.DB.Delete(&book)
	// Kembalikan respons sukses tanpa data
	return utils.SuccessResponseWithoutData(c, http.StatusOK)
}
