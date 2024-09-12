package controllers

import (
	"gasruk/pkg/config"
	"gasruk/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateReview: Fungsi untuk membuat review baru
func CreateReview(c echo.Context) error {
	review := new(models.Review)
	if err := c.Bind(review); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
		})
	}

	// Simpan review ke database
	if err := config.DB.Create(&review).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error creating review",
		})
	}

	return c.JSON(http.StatusCreated, review)
}

// GetReviews: Fungsi untuk mengambil semua review
func GetReviews(c echo.Context) error {
	var reviews []models.Review
	if err := config.DB.Find(&reviews).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving reviews",
		})
	}
	return c.JSON(http.StatusOK, reviews)
}

// GetReviewByID: Fungsi untuk mengambil review berdasarkan ID
func GetReviewByID(c echo.Context) error {
	var review models.Review
	id := c.Param("id")
	if err := config.DB.First(&review, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Review not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving review",
		})
	}
	return c.JSON(http.StatusOK, review)
}

// UpdateReview: Fungsi untuk mengupdate review berdasarkan ID
func UpdateReview(c echo.Context) error {
	var review models.Review
	id := c.Param("id")
	// Temukan review berdasarkan ID
	if err := config.DB.First(&review, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Review not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve review"})
	}

	// Bind JSON payload ke model Review
	if err := c.Bind(&review); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := config.DB.Save(&review).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update review"})
	}
	return c.JSON(http.StatusOK, review)
}

// DeleteReview: Fungsi untuk menghapus review berdasarkan ID
func DeleteReview(c echo.Context) error {
	var review models.Review
	id := c.Param("id")
	// Temukan review berdasarkan ID
	if err := config.DB.First(&review, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Review not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve review"})
	}

	// Hapus review dari database
	if err := config.DB.Delete(&review).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete review"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Review deleted"})
}
