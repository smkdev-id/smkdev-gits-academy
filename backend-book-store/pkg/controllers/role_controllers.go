package controllers

import (
	"net/http"

	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/config"
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/models"
	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var role models.Role

	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := config.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, role)
}

func GetRoles(c *gin.Context) {
	var roles []models.Role

	if err := config.DB.Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)
}