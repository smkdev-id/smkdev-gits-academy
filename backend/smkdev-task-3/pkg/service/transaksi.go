package service

import (
	"BookStore/pkg/models"
	"BookStore/pkg/repository"
	"errors"

	"github.com/google/uuid"
)

type TransactionService interface {
	CreateTransactionFromBasket(customerID uuid.UUID) (*models.Transaction, error)
	GetTransaction(id uuid.UUID) (*models.Transaction, error)
	GetCustomerTransactions(customerID uuid.UUID) ([]models.Transaction, error)
	UpdateTransactionStatus(id uuid.UUID, status string) error
	GetAllTransactions() ([]models.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	basketService   BasketService
}

func NewTransactionService(transactionRepo repository.TransactionRepository, basketService BasketService) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		basketService:   basketService,
	}
}

func (s *transactionService) CreateTransactionFromBasket(customerID uuid.UUID) (*models.Transaction, error) {
	basket, err := s.basketService.GetBasket(customerID)
	if err != nil {
		return nil, err
	}
	if basket == nil || len(basket.Items) == 0 {
		return nil, errors.New("basket is empty")
	}

	transaction := &models.Transaction{
		ID:         uuid.New(),
		CustomerID: customerID,
		Status:     "pending",
	}

	var totalPrice float64
	for _, item := range basket.Items {
		orderItem := models.OrderItem{
			ID:            uuid.New(),
			TransactionID: transaction.ID,
			BookID:        item.BookID,
			Quantity:      item.Quantity,
			Price:         float64(item.Book.Price),
		}
		transaction.Items = append(transaction.Items, orderItem)
		totalPrice += float64(item.Quantity) * float64(item.Book.Price)
	}
	transaction.TotalPrice = totalPrice

	if err := s.transactionRepo.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	if err := s.basketService.ClearBasket(customerID); err != nil {
		// Log this error, but don't fail the transaction creation
		// Consider implementing a compensation mechanism here
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) GetTransaction(id uuid.UUID) (*models.Transaction, error) {
	return s.transactionRepo.GetTransactionByID(id)
}
func (s *transactionService) GetAllTransactions() ([]models.Transaction, error) {
	return s.transactionRepo.GetAllTransactions()
}
func (s *transactionService) GetCustomerTransactions(customerID uuid.UUID) ([]models.Transaction, error) {
	return s.transactionRepo.GetTransactionsByCustomerID(customerID)
}

func (s *transactionService) UpdateTransactionStatus(id uuid.UUID, status string) error {
	return s.transactionRepo.UpdateTransactionStatus(id, status)
}
