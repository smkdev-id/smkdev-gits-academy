package service

import (
	"BookStore/pkg/models"
	"BookStore/pkg/repository"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AuthService interface untuk layanan autentikasi
type AuthService interface {
	Register(c *gin.Context, request *models.UserRegister) error
	Login(c *gin.Context, request *models.UserLogin) error
}

// authService implementasi dari AuthService
type authService struct {
	// db         *gorm.DB
	repository repository.AuthRepository
}

// NewAuthService membuat instance baru dari authService
func NewAuthService(repository repository.AuthRepository) AuthService {
	return &authService{repository: repository}
}

// Register mendaftarkan pengguna baru ke dalam sistem
func (s *authService) Register(c *gin.Context, request *models.UserRegister) error {
	user := models.User{
		ID:       uuid.New(),
		Username: request.Username,
		Password: request.Password,
		Role:     request.Role,
	}
	s.repository.CreateUser(c, &user)
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	return nil
}

// Login untuk autentikasi pengguna
func (s *authService) Login(c *gin.Context, request *models.UserLogin) error {

	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	// 	return err
	// }
	user, err := s.repository.FindByUsername(c, request.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return err
	}
	if user.Password != request.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
		return nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	jwtKey := []byte("YNTKTS") // Gantilah dengan kunci rahasia yang lebih kuat
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return err
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
	return nil
}
