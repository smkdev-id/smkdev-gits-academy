package repositories

import (
	"context"

	"github.com/alwiirfan/internal/entities"
	"gorm.io/gorm"
)

type (
	BookRepository interface {
		FindAll(ctx context.Context) ([]*entities.Book, error)
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
	panic("unimplemented")
}

// FindAll implements BookRepository.
func (r *bookRepository) FindAll(ctx context.Context) ([]*entities.Book, error) {
	panic("unimplemented")
}

// FindByID implements BookRepository.
func (r *bookRepository) FindByID(ctx context.Context, id string) (*entities.Book, error) {
	panic("unimplemented")
}

// Update implements BookRepository.
func (r *bookRepository) Update(ctx context.Context, book *entities.Book) error {
	panic("unimplemented")
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{
		DB: db,
	}
}
