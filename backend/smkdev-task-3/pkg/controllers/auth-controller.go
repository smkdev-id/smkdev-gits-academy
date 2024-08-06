package controllers

import (
	"BookStore/pkg/models"
	"BookStore/pkg/service"
	"BookStore/pkg/utils"
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
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	if err := ctrl.authService.Register(c, &registerRequest); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Registration successful", nil)
}

// LoginUser mengontrol proses login user
func (ctrl *AuthController) LoginUser(c *gin.Context) {
	var loginRequest models.UserLogin
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	if err := ctrl.authService.Login(c, &loginRequest); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "Login successful", nil)
}
