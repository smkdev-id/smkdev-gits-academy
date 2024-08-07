package controllers

import (
	"bookstore/pkg/config"
	"bookstore/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKategoris(c *gin.Context) {
	var kategoris []models.Kategori
	config.DB.Find(&kategoris)
	c.JSON(http.StatusOK, kategoris)
}

func GetKategoriByID(c *gin.Context) {
	id := c.Param("id")
	var kategori models.Kategori
	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found!"})
		return
	}
	c.JSON(http.StatusOK, kategori)
}

func CreateKategori(c *gin.Context) {
	var kategori models.Kategori
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&kategori)
	c.JSON(http.StatusCreated, kategori)
}

func UpdateKategori(c *gin.Context) {
	id := c.Param("id")
	var kategori models.Kategori
	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found!"})
		return
	}
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&kategori)
	c.JSON(http.StatusOK, kategori)
}

func DeleteKategori(c *gin.Context) {
	id := c.Param("id")
	var kategori models.Kategori
	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori not found!"})
		return
	}
	config.DB.Delete(&kategori)
	c.Status(http.StatusNoContent)
}
