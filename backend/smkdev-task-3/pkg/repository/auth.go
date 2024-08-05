package repository

import (
	"BookStore/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByUsername(c *gin.Context, username string) (*models.User, error)
	CreateUser(c *gin.Context, register *models.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindByUsername(c *gin.Context, username string) (*models.User, error) {
	user := new(models.User)

	if err := r.db.WithContext(c).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) CreateUser(c *gin.Context, register *models.User) error {

	if err := r.db.WithContext(c).Create(&register).Error; err != nil {
		return err
	}
	return nil
}
