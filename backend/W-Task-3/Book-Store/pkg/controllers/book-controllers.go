// pkg/controllers/bookControllers.go

package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Book-Store/pkg/config"
	"Book-Store/pkg/models"
	"Book-Store/pkg/utils"

	"github.com/gorilla/mux"
)

// GetBooks mengembalikan semua buku dalam format JSON
func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	config.DB.Find(&books)
	utils.RespondJSON(w, books, http.StatusOK)
}

// GetBook mengembalikan buku berdasarkan ID dalam format JSON
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondJSON(w, map[string]string{"error": "Invalid ID"}, http.StatusBadRequest)
		return
	}

	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		utils.RespondJSON(w, map[string]string{"error": "Book not found"}, http.StatusNotFound)
		return
	}

	utils.RespondJSON(w, book, http.StatusOK)
}

// CreateBook membuat buku baru dan mengembalikan buku yang baru dibuat dalam format JSON
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.RespondJSON(w, map[string]string{"error": "Invalid input"}, http.StatusBadRequest)
		return
	}

	config.DB.Create(&book)
	utils.RespondJSON(w, book, http.StatusCreated)
}

// UpdateBook memperbarui buku yang ada berdasarkan ID dan mengembalikan buku yang diperbarui dalam format JSON
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondJSON(w, map[string]string{"error": "Invalid ID"}, http.StatusBadRequest)
		return
	}

	var book models.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		utils.RespondJSON(w, map[string]string{"error": "Book not found"}, http.StatusNotFound)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.RespondJSON(w, map[string]string{"error": "Invalid input"}, http.StatusBadRequest)
		return
	}

	config.DB.Save(&book)
	utils.RespondJSON(w, book, http.StatusOK)
}

// DeleteBook menghapus buku berdasarkan ID dan mengembalikan pesan sukses dalam format JSON
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondJSON(w, map[string]string{"error": "Invalid ID"}, http.StatusBadRequest)
		return
	}

	var book models.Book
	result := config.DB.Delete(&book, id)
	if result.RowsAffected == 0 {
		utils.RespondJSON(w, map[string]string{"error": "Book not found"}, http.StatusNotFound)
		return
	}

	utils.RespondJSON(w, map[string]string{"message": "Book deleted successfully"}, http.StatusOK)
}
