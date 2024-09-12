package controllers

import (
	"gasruk/pkg/config"
	"gasruk/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateOrderStatus(c echo.Context) error {
	orderStatus := new(models.OrderStatus)
	if err := c.Bind(orderStatus); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
		})
	}

	// Simpan OrderStatus ke database
	if err := config.DB.Create(&orderStatus).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error creating order status",
		})
	}

	return c.JSON(http.StatusCreated, orderStatus)
}
func GetOrderStatuses(c echo.Context) error {
	var orderStatuses []models.OrderStatus
	if err := config.DB.Find(&orderStatuses).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving order statuses",
		})
	}
	return c.JSON(http.StatusOK, orderStatuses)
}
func GetOrderStatusByID(c echo.Context) error {
	var orderStatus models.OrderStatus
	id := c.Param("id")
	if err := config.DB.First(&orderStatus, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Order status not found",
		})
	}
	return c.JSON(http.StatusOK, orderStatus)
}
func UpdateOrderStatus(c echo.Context) error {
	var orderStatus models.OrderStatus
	id := c.Param("id")

	// Temukan OrderStatus berdasarkan ID
	if err := config.DB.First(&orderStatus, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Order status not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve order status",
		})
	}

	// Bind JSON payload ke model OrderStatus
	if err := c.Bind(&orderStatus); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Simpan perubahan ke database
	if err := config.DB.Save(&orderStatus).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update order status",
		})
	}
	return c.JSON(http.StatusOK, orderStatus)
}

func DeleteOrderStatus(c echo.Context) error {
	var orderStatus models.OrderStatus
	id := c.Param("id")

	// Temukan OrderStatus berdasarkan ID
	if err := config.DB.First(&orderStatus, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Order status not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve order status",
		})
	}

	// Hapus OrderStatus dari database
	if err := config.DB.Delete(&orderStatus).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete order status",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Order status deleted",
	})
}
