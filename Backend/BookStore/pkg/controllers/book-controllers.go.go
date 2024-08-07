package controllers

import (
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
	models.DB.Find(&books) // Ambil semua buku dari database
	for _, book := range books {
		bookRes = append(bookRes, dto.EntityToResponse(book)) // Ubah setiap buku menjadi format respons
	}
	fmt.Println(bookRes)
	return utils.SuccessResponse(c, http.StatusOK, bookRes) // Kembalikan respons sukses dengan data buku
}

// GetBookByID mengembalikan buku berdasarkan ID.
func GetBookByID(c echo.Context) error {
	var book *models.Book
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Book not found")
	}
	if err := models.DB.First(&book, intID).Error; err != nil { // Cari buku berdasarkan ID
		return utils.ErrorResponse(c, http.StatusNotFound, "Book not found")
	}
	bookRes := dto.EntityToResponse(book)                   // Ubah buku menjadi format respons
	return utils.SuccessResponse(c, http.StatusOK, bookRes) // Kembalikan respons sukses dengan data buku
}

// CreateBook membuat buku baru.
func CreateBook(c echo.Context) error {
	var book *models.Book
	if err := c.Bind(&book); err != nil { // Bind data dari request ke model buku
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(book); err != nil { // Validasi input
		errors := utils.ParseValidationErrors(err)
		return utils.ErrorResponseValidation(c, http.StatusBadRequest, errors)
	}
	models.DB.Create(&book)                                        // Simpan buku baru ke database
	return utils.SuccessResponseWithoutData(c, http.StatusCreated) // Kembalikan respons sukses tanpa data
}

// UpdateBook memperbarui buku berdasarkan ID.
func UpdateBook(c echo.Context) error {
	var book *models.Book
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Book not found")
	}
	if err := models.DB.First(&book, intID).Error; err != nil { // Cari buku berdasarkan ID
		return utils.ErrorResponse(c, http.StatusNotFound, "Book not found")
	}
	if err := c.Bind(&book); err != nil { // Bind data dari request ke model buku
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	if err := utils.ValidateStruct(book); err != nil { // Validasi input
		errors := utils.ParseValidationErrors(err)
		return utils.ErrorResponseValidation(c, http.StatusBadRequest, errors)
	}
	models.DB.Save(&book)                                     // Simpan perubahan buku ke database
	return utils.SuccessResponseWithoutData(c, http.StatusOK) // Kembalikan respons sukses tanpa data
}

// DeleteBook menghapus buku berdasarkan ID.
func DeleteBook(c echo.Context) error {
	var book models.Book
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Book not found")
	}
	if err := models.DB.First(&book, intID).Error; err != nil { // Cari buku berdasarkan ID
		return utils.ErrorResponse(c, http.StatusNotFound, "Book not found")
	}
	models.DB.Delete(&book)                                   // Hapus buku dari database
	return utils.SuccessResponseWithoutData(c, http.StatusOK) // Kembalikan respons sukses tanpa data
}
