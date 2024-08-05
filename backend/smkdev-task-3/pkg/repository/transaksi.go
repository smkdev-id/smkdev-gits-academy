package repository

import (
	"BookStore/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionByID(id uuid.UUID) (*models.Transaction, error)
	GetTransactionsByCustomerID(customerID uuid.UUID) ([]models.Transaction, error)
	UpdateTransactionStatus(id uuid.UUID, status string) error
	GetAllTransactions() ([]models.Transaction, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepo) GetTransactionByID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := r.db.Preload("Items.Book").First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
func (r *transactionRepo) GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Preload("Items.Book").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
func (r *transactionRepo) GetTransactionsByCustomerID(customerID uuid.UUID) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Preload("Items.Book").Where("customer_id = ?", customerID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepo) UpdateTransactionStatus(id uuid.UUID, status string) error {
	return r.db.Model(&models.Transaction{}).Where("id = ?", id).Update("status", status).Error
}
