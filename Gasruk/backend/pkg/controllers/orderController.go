package controllers

import (
	"gasruk/pkg/config"
	"gasruk/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateOrder(c echo.Context) error {
	order := new(models.Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
		})
	}

	// Simpan order ke database
	if err := config.DB.Create(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error creating order",
		})
	}

	return c.JSON(http.StatusCreated, order)
}

func GetOrders(c echo.Context) error {
	var orders []models.Order
	if err := config.DB.Preload("Payment").Preload("OrderStatus").Preload("Shipping").Find(&orders).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving orders",
		})
	}
	return c.JSON(http.StatusOK, orders)
}

func GetOrderByID(c echo.Context) error {
	var order models.Order
	id := c.Param("id")
	if err := config.DB.Preload("Payment").Preload("OrderStatus").Preload("Shipping").First(&order, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Order not found",
		})
	}
	return c.JSON(http.StatusOK, order)
}

func UpdateOrder(c echo.Context) error {
	var order models.Order
	id := c.Param("id")

	// Temukan order berdasarkan ID
	if err := config.DB.First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Order not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve order",
		})
	}

	// Bind JSON payload ke model Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Simpan perubahan ke database
	if err := config.DB.Save(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update order",
		})
	}
	return c.JSON(http.StatusOK, order)
}

func DeleteOrder(c echo.Context) error {
	var order models.Order
	id := c.Param("id")

	// Temukan order berdasarkan ID
	if err := config.DB.First(&order, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Order not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve order",
		})
	}

	// Hapus order dari database
	if err := config.DB.Delete(&order).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete order",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Order deleted",
	})
}
