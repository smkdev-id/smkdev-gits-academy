package controllers

import (
	"gasruk/pkg/config"
	"gasruk/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// CreateProduct: Fungsi untuk membuat produk baru
func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
		})
	}

	// Simpan produk ke database
	if err := config.DB.Create(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error creating product",
		})
	}

	return c.JSON(http.StatusCreated, product)
}

// GetProducts: Fungsi untuk mengambil semua produk
func GetProducts(c echo.Context) error {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving products",
		})
	}
	return c.JSON(http.StatusOK, products)
}

// GetProductByID: Fungsi untuk mengambil produk berdasarkan ID
func GetProductByID(c echo.Context) error {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{
				"message": "Product not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Error retrieving product",
		})
	}
	return c.JSON(http.StatusOK, product)
}

// UpdateProduct: Fungsi untuk mengupdate produk berdasarkan ID
func UpdateProduct(c echo.Context) error {
	var product models.Product
	id := c.Param("id")
	// Temukan produk berdasarkan ID
	if err := config.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve product"})
	}

	// Bind JSON payload ke model Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := config.DB.Save(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to update product"})
	}
	return c.JSON(http.StatusOK, product)
}

// DeleteProduct: Fungsi untuk menghapus produk berdasarkan ID
func DeleteProduct(c echo.Context) error {
	var product models.Product
	id := c.Param("id")
	// Temukan produk berdasarkan ID
	if err := config.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Product not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to retrieve product"})
	}

	// Hapus produk dari database
	if err := config.DB.Delete(&product).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to delete product"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted"})
}
