package controllers

import (
	"gasruk/pkg/config"
	"gasruk/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreatePayment: Fungsi untuk membuat payment baru
func CreatePayment(c echo.Context) error {
	payment := new(models.Payment)
	if err := c.Bind(payment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
		})
	}

	// Simpan payment ke database
	if err := config.DB.Create(&payment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error creating payment",
		})
	}

	return c.JSON(http.StatusCreated, payment)
}

// GetPayments: Fungsi untuk mengambil semua payment
func GetPayments(c echo.Context) error {
	var payments []models.Payment
	if err := config.DB.Find(&payments).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving payments",
		})
	}
	return c.JSON(http.StatusOK, payments)
}

// GetPaymentByID: Fungsi untuk mengambil payment berdasarkan ID
func GetPaymentByID(c echo.Context) error {
	var payment models.Payment
	id := c.Param("id")
	if err := config.DB.First(&payment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Payment not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving payment",
		})
	}
	return c.JSON(http.StatusOK, payment)
}

// UpdatePayment: Fungsi untuk mengupdate payment berdasarkan ID
func UpdatePayment(c echo.Context) error {
	var payment models.Payment
	id := c.Param("id")
	// Temukan payment berdasarkan ID
	if err := config.DB.First(&payment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Payment not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve payment"})
	}

	// Bind JSON payload ke model Payment
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := config.DB.Save(&payment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update payment"})
	}
	return c.JSON(http.StatusOK, payment)
}

// DeletePayment: Fungsi untuk menghapus payment berdasarkan ID
func DeletePayment(c echo.Context) error {
	var payment models.Payment
	id := c.Param("id")
	// Temukan payment berdasarkan ID
	if err := config.DB.First(&payment, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Payment not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve payment"})
	}

	// Hapus payment dari database
	if err := config.DB.Delete(&payment).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete payment"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Payment deleted"})
}
