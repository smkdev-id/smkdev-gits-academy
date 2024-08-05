package controllers

import (
	"BookStore/pkg/models"
	"BookStore/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController menangani request terkait autentikasi
type AuthController struct {
	authService service.AuthService
}

// NewAuthController membuat instance baru dari AuthController
func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// RegisterUser mengontrol proses pendaftaran user
func (ctrl *AuthController) RegisterUser(c *gin.Context) {
	var registerRequest models.UserRegister
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.authService.Register(c, &registerRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

// LoginUser mengontrol proses login user
func (ctrl *AuthController) LoginUser(c *gin.Context) {
	var loginRequest models.UserLogin
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.authService.Login(c, &loginRequest); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
