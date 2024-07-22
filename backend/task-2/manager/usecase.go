package manager

import (
	"EkoEdyPurwanto/task-2/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	UseCaseManager interface {
		TodoListUseCase() usecase.TodoListUseCase
	}
	useCaseManager struct {
		RepositoryM RepositoryManager
		Validate    *validator.Validate
		Logger      *logrus.Logger
	}
)

func (u *useCaseManager) TodoListUseCase() usecase.TodoListUseCase {
	return usecase.NewTodoListUseCase(u.RepositoryM.TodoListRepository(), u.Validate, u.Logger)
}

// NewUseCaseManager Constructor
func NewUseCaseManager(repositoryM RepositoryManager, validate *validator.Validate, logger *logrus.Logger) UseCaseManager {
	return &useCaseManager{
		RepositoryM: repositoryM,
		Validate:    validate,
		Logger:      logger,
	}
}
