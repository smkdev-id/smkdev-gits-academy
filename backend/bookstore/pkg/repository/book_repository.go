package repository

import (
	"bookstore/pkg/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetByID(id string) (models.Book, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(id string) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository{
	return &bookRepository{db}
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book

	err := r.db.Find(&books).Error
	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *bookRepository) GetByID(id string) (models.Book, error) {
	var book models.Book

	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Create(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Update(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *bookRepository) Delete(id string) error {
	var book models.Book

	err := r.db.Where("id = ?", id).Delete(&book).Error

	if err != nil {
		return err
	}

	return nil
}