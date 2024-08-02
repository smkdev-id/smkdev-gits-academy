package repositories

import (
	"context"

	"github.com/alwiirfan/internal/entities"
	"github.com/alwiirfan/internal/utils/pagination"
	"gorm.io/gorm"
)

type (
	BookRepository interface {
		FindAll(ctx context.Context, page, pageSize int) ([]*entities.Book, int64, error)
		FindByID(ctx context.Context, id string) (*entities.Book, error)
		Create(ctx context.Context, book *entities.Book) error
		Update(ctx context.Context, book *entities.Book) error
		Delete(ctx context.Context, id string) error
	}

	bookRepository struct {
		DB *gorm.DB
	}
)

// Create implements BookRepository.
func (r *bookRepository) Create(ctx context.Context, book *entities.Book) error {
	if err := r.DB.WithContext(ctx).Create(book).Error; err != nil {
		return err
	}

	return nil
}

// Delete implements BookRepository.
func (r *bookRepository) Delete(ctx context.Context, id string) error {
	if err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&entities.Book{}).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements BookRepository.
func (r *bookRepository) FindAll(ctx context.Context, page, pageSize int) ([]*entities.Book, int64, error) {
	var books []*entities.Book
	var totalRecords int64

	if err := r.DB.WithContext(ctx).Model(&entities.Book{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := r.DB.WithContext(ctx).Scopes(pagination.BookPaginate(page, pageSize)).Find(&books).Error; err != nil {
		return nil, 0, err
	}

	return books, totalRecords, nil
}

// FindByID implements BookRepository.
func (r *bookRepository) FindByID(ctx context.Context, id string) (*entities.Book, error) {
	var book *entities.Book

	if err := r.DB.WithContext(ctx).Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

// Update implements BookRepository.
func (r *bookRepository) Update(ctx context.Context, book *entities.Book) error {
	if err := r.DB.WithContext(ctx).Model(&entities.Book{}).Where("id = ?", book.ID).Updates(book).Error; err != nil {
		return err
	}

	return nil
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		DB: db,
	}
}
