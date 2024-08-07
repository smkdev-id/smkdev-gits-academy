package manager

import (
	"EkoEdyPurwanto/task-3/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	UseCaseManager interface {
		BookUseCase() usecase.BookUseCase
	}
	useCaseManager struct {
		RepositoryM RepositoryManager
		Validate    *validator.Validate
		Logger      *logrus.Logger
	}
)

func (u *useCaseManager) BookUseCase() usecase.BookUseCase {
	return usecase.NewBookUseCase(u.RepositoryM.BookRepository(), u.Validate, u.Logger)
}

// NewUseCaseManager Constructor
func NewUseCaseManager(repositoryM RepositoryManager, validate *validator.Validate, logger *logrus.Logger) UseCaseManager {
	return &useCaseManager{
		RepositoryM: repositoryM,
		Validate:    validate,
		Logger:      logger,
	}
}
