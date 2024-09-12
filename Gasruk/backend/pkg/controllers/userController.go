package controllers

import (
	"gasruk/pkg/config"
	"gasruk/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid input",
		})
	}

	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error hashing password"})
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error creating user"})
	}

	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error retrieving users"})
	}
	return c.JSON(http.StatusOK, users)
}

func GetUserByID(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "User not found",
		})
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	// Temukan user berdasarkan ID
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user"})
	}

	// Bind JSON payload ke model User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := db.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	var user models.User
	id := c.Param("id")
	// Temukan user berdasarkan ID
	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve user"})
	}

	// Hapus user dari database
	if err := db.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete user"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}
