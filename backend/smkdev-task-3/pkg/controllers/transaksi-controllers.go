package controllers

import (
	"BookStore/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

func (tc *TransactionController) CreateTransaction(ctx *gin.Context) {
	customerID, err := tc.getCustomerID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := tc.transactionService.CreateTransactionFromBasket(customerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transaction)
}

func (tc *TransactionController) GetTransaction(ctx *gin.Context) {
	transactionID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := tc.transactionService.GetTransaction(transactionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}
func (tc *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := tc.transactionService.GetAllTransactions()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transactions)
}
func (tc *TransactionController) GetCustomerTransactions(ctx *gin.Context) {
	customerID, err := tc.getCustomerID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactions, err := tc.transactionService.GetCustomerTransactions(customerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) UpdateTransactionStatus(ctx *gin.Context) {
	transactionID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var request struct {
		Status string `json:"status" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = tc.transactionService.UpdateTransactionStatus(transactionID, request.Status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction status updated"})
}

func (tc *TransactionController) getCustomerID(ctx *gin.Context) (uuid.UUID, error) {
	userIDStr := ctx.MustGet("user_id").(string)
	return uuid.Parse(userIDStr)
}
