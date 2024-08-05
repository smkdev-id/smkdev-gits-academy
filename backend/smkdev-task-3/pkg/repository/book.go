package repository

import (
	"BookStore/pkg/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll(c *gin.Context) ([]models.Book, error)

	FindByID(c *gin.Context, id uuid.UUID) (*models.Book, error)
	CreateBook(c *gin.Context, book *models.Book) error
	UpdateBook(c *gin.Context, book *models.Book) error
	DeleteBook(c *gin.Context, id uuid.UUID) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}
func (r *bookRepository) FindAll(c *gin.Context) ([]models.Book, error) {
	books := make([]models.Book, 0)

	if err := r.db.WithContext(c).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
func (r *bookRepository) FindByID(c *gin.Context, id uuid.UUID) (*models.Book, error) {
	book := new(models.Book)
	if err := r.db.WithContext(c).Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}
func (r *bookRepository) CreateBook(c *gin.Context, book *models.Book) error {
	// book := new(models.Book)
	if err := r.db.WithContext(c).Create(&book).Error; err != nil {
		return err
	}

	return nil
}
func (r *bookRepository) UpdateBook(c *gin.Context, book *models.Book) error {
	if err := r.db.WithContext(c).Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}
func (r *bookRepository) DeleteBook(c *gin.Context, id uuid.UUID) error {

	if err := r.db.WithContext(c).Delete(&models.Book{}, id).Error; err != nil {
		log.Println("GA BISA DELETE", err)

		return err
	}
	return nil
}
