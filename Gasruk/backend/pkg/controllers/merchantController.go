package controllers

import (
	"gasruk/pkg/config"
	"gasruk/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateMerchant(c echo.Context) error {
	merchant := new(models.Merchant)
	if err := c.Bind(merchant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
		})
	}

	// Simpan merchant ke database
	if err := config.DB.Create(&merchant).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error creating merchant",
		})
	}

	return c.JSON(http.StatusCreated, merchant)
}

// GetMerchants: Fungsi untuk mengambil semua merchant
func GetMerchants(c echo.Context) error {
	var merchants []models.Merchant
	if err := config.DB.Find(&merchants).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving merchants",
		})
	}
	return c.JSON(http.StatusOK, merchants)
}

// GetMerchantByID: Fungsi untuk mengambil merchant berdasarkan ID
func GetMerchantByID(c echo.Context) error {
	var merchant models.Merchant
	id := c.Param("id")
	if err := config.DB.First(&merchant, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Merchant not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving merchant",
		})
	}
	return c.JSON(http.StatusOK, merchant)
}

// UpdateMerchant: Fungsi untuk mengupdate merchant berdasarkan ID
func UpdateMerchant(c echo.Context) error {
	var merchant models.Merchant
	id := c.Param("id")
	// Temukan merchant berdasarkan ID
	if err := config.DB.First(&merchant, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Merchant not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve merchant"})
	}

	// Bind JSON payload ke model Merchant
	if err := c.Bind(&merchant); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := config.DB.Save(&merchant).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update merchant"})
	}
	return c.JSON(http.StatusOK, merchant)
}

// DeleteMerchant: Fungsi untuk menghapus merchant berdasarkan ID
func DeleteMerchant(c echo.Context) error {
	var merchant models.Merchant
	id := c.Param("id")
	// Temukan merchant berdasarkan ID
	if err := config.DB.First(&merchant, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Merchant not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve merchant"})
	}

	// Hapus merchant dari database
	if err := config.DB.Delete(&merchant).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete merchant"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Merchant deleted"})
}
