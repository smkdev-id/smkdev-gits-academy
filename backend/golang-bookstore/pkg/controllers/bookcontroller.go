package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-bookstore/pkg/models"
	"golang-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	createdBook, err := utils.CreateBook(&book)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondSuccess(c, http.StatusOK, "Book created successfully", map[string]interface{}{
		"book": createdBook,
	})
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := utils.GetBookByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Book not found")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, "Book retrieved successfully", map[string]interface{}{
		"book": book,
	})
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	book.ID = uint(id)
	updatedBook, err := utils.UpdateBook(&book)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondSuccess(c, http.StatusOK, "Book updated successfully", map[string]interface{}{
		"book": updatedBook,
	})
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := utils.DeleteBook(uint(id)); err != nil {
		utils.RespondError(c, http.StatusNotFound, "Book not found")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, "Book deleted successfully", nil)
}

func GetAllBooks(c *gin.Context) {
	books, err := utils.GetAllBooks()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondSuccess(c, http.StatusOK, "Books retrieved successfully", map[string]interface{}{
		"books": books,
	})
}
