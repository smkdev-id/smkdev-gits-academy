package usecase

import (
	"EkoEdyPurwanto/task-3/internal/model"
	"EkoEdyPurwanto/task-3/internal/model/dto/req"
	"EkoEdyPurwanto/task-3/internal/repository"
	"EkoEdyPurwanto/task-3/utility/common"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	BookUseCase interface {
		Create(payload req.CreateRequest) error
		GetAll() (model.Books, error)
		GetById(id string) (model.Books, error)
		UpdateStatus(payload req.UpdateRequest) error
		Delete(payload req.DeleteRequest) error
	}
	bookUseCase struct {
		Repository repository.BookRepository
		Validate   *validator.Validate
		Logger     *logrus.Logger
	}
)

func NewBookUseCase(repository repository.BookRepository, validate *validator.Validate, logger *logrus.Logger) BookUseCase {
	return &bookUseCase{
		Repository: repository,
		Validate:   validate,
		Logger:     logger,
	}
}

// Create implements BookUseCase.
func (b *bookUseCase) Create(payload req.CreateRequest) error {
	// struct validation
	err := b.Validate.Struct(payload)
	if err != nil {
		b.Logger.Warnf("Invalid request body : %+v", err)
		return err
	}

	// Define the start and end dates for the random date generation
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)

	// Generate a random PublishedDate between start and end
	randomPublishedDate := common.RandomGeneratedDate(start, end)

	// fill value of book
	book := model.Book{
		Id:            common.GenerateID(),
		Title:         payload.Title,
		ISBN:          payload.ISBN,
		Author:        payload.Author,
		PublishedDate: randomPublishedDate,
		Genre:         payload.Genre,
		Price:         payload.Price,
		Quantity:      payload.Quantity,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Time{},
		DeletedAt:     time.Time{},
	}

	// save todolist
	err = b.Repository.Save(book)
	if err != nil {
		b.Logger.Warnf("failed save book : %+v", err)
		return err
	}
	return nil
}

// GetAll implements BookUseCase.
func (b *bookUseCase) GetAll() (model.Books, error) {
	books, err := b.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

// GetById implements BookUseCase.
func (b *bookUseCase) GetById(id string) (model.Books, error) {
	// validation
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	books, err := b.Repository.FindById(id)
	if err != nil {
		return nil, err
	}

	return books, nil
}

// UpdateStatus implements BookUseCase.
func (b *bookUseCase) UpdateStatus(payload req.UpdateRequest) error {
	// struct validation
	err := b.Validate.Struct(payload)
	if err != nil {
		b.Logger.Warnf("Invalid request body : %+v", err)
		return err
	}

	_, err = b.GetById(payload.Id)
	if err != nil {
		b.Logger.Warnf("TodoList not found : %+v", err)
		return err
	}

	payload.UpdatedAt = time.Now()

	if err := b.Repository.Modify(payload); err != nil {
		b.Logger.Warnf("Failed to update TodoList : %+v", err)
		return err
	}
	return nil
}

// Delete implements BookUseCase.
func (b *bookUseCase) Delete(payload req.DeleteRequest) error {
	// struct validation
	err := b.Validate.Struct(payload)
	if err != nil {
		b.Logger.Warnf("Invalid request body : %+v", err)
		return err
	}

	if err := b.Repository.Delete(payload); err != nil {
		return err
	}
	return nil
}
