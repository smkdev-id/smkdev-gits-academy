package controllers

import (
	"gasruk/pkg/config"
	"gasruk/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateShipping(c echo.Context) error {
	shipping := new(models.Shipping)
	if err := c.Bind(shipping); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
		})
	}

	// Simpan shipping ke database
	if err := config.DB.Create(&shipping).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error creating shipping",
		})
	}

	return c.JSON(http.StatusCreated, shipping)
}

func GetShippings(c echo.Context) error {
	var shippings []models.Shipping
	if err := config.DB.Find(&shippings).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving shippings",
		})
	}
	return c.JSON(http.StatusOK, shippings)
}

func GetShippingByID(c echo.Context) error {
	var shipping models.Shipping
	id := c.Param("id")
	if err := config.DB.First(&shipping, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Shipping not found",
		})
	}
	return c.JSON(http.StatusOK, shipping)
}
func UpdateShipping(c echo.Context) error {
	var shipping models.Shipping
	id := c.Param("id")

	// Temukan shipping berdasarkan ID
	if err := config.DB.First(&shipping, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Shipping not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve shipping",
		})
	}

	// Bind JSON payload ke model Shipping
	if err := c.Bind(&shipping); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Simpan perubahan ke database
	if err := config.DB.Save(&shipping).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update shipping",
		})
	}
	return c.JSON(http.StatusOK, shipping)
}
func DeleteShipping(c echo.Context) error {
	var shipping models.Shipping
	id := c.Param("id")

	// Temukan shipping berdasarkan ID
	if err := config.DB.First(&shipping, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Shipping not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve shipping",
		})
	}

	// Hapus shipping dari database
	if err := config.DB.Delete(&shipping).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete shipping",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Shipping deleted",
	})
}
