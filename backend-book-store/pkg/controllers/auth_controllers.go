package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/config"
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/models"
	"github.com/ItsKevinRafaell/go-books-store-crud/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RoleAssignmentRequest struct {
    RoleIDs []uint `json:"role_ids"`
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	user.Password = string(hashedPassword)

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	var defaultRole models.Role
	if err := config.DB.Where("name = ?", "user").First(&defaultRole).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Default role not found"})
		return
	}

	if err := config.DB.Model(&user).Association("Roles").Append(&defaultRole); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign default role"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "User created successfully", "user" : gin.H{
		"ID" : user.ID,
		"username" : user.Username,
	}, "token" : token})
}

func AssignRole(c *gin.Context) {
	userID := c.Param("userID")

	var user models.User

	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "User not found"})
		return
	} 

	var req RoleAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	var roles []models.Role
	if err := config.DB.Where("id IN (?)", req.RoleIDs).Find(&roles).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	var newRoles []models.Role
    existingRoleMap := make(map[uint]bool)
    for _, role := range user.Roles {
        existingRoleMap[role.ID] = true
    }

    for _, role := range roles {
        if !existingRoleMap[role.ID] {
            newRoles = append(newRoles, role)
        }
    }

	if err := config.DB.Model(&user).Association("Roles").Append(newRoles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Role assigned to user"})
}

func Login(c *gin.Context) {
	var user models.User
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		 c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}

	if err := config.DB.Where("username = ?", inputUser.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid Credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Invalid Credentials"})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token" : token})
}

func Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : "Token not found"})
		return
	}
	if len(token) > 7 && token[:7] == "Bearer" {
		token = token[7:]
	}

	ctx := context.Background()
	err := config.RedisClient.Set(ctx, "blacklist:"+token, "true", 24*time.Hour).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message" : "Logout Success"})
}