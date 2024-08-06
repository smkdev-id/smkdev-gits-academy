package controllers

import (
	"BookStore/pkg/service"
	"BookStore/pkg/utils"
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
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	transaction, err := tc.transactionService.CreateTransactionFromBasket(customerID)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONResponse(ctx, http.StatusCreated, "Transaction created successfully", transaction)
}

func (tc *TransactionController) GetTransaction(ctx *gin.Context) {
	transactionID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, "Invalid transaction ID", nil)
		return
	}

	transaction, err := tc.transactionService.GetTransaction(transactionID)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONResponse(ctx, http.StatusOK, "Transaction retrieved successfully", transaction)
}

func (tc *TransactionController) GetAllTransactions(ctx *gin.Context) {
	transactions, err := tc.transactionService.GetAllTransactions()
	if err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.JSONResponse(ctx, http.StatusOK, "Transactions retrieved successfully", transactions)
}

func (tc *TransactionController) GetCustomerTransactions(ctx *gin.Context) {
	customerID, err := tc.getCustomerID(ctx)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	transactions, err := tc.transactionService.GetCustomerTransactions(customerID)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONResponse(ctx, http.StatusOK, "Customer transactions retrieved successfully", transactions)
}

func (tc *TransactionController) UpdateTransactionStatus(ctx *gin.Context) {
	transactionID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, "Invalid transaction ID", nil)
		return
	}

	var request struct {
		Status string `json:"status" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = tc.transactionService.UpdateTransactionStatus(transactionID, request.Status)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.JSONResponse(ctx, http.StatusOK, "Transaction status updated successfully", nil)
}

func (tc *TransactionController) getCustomerID(ctx *gin.Context) (uuid.UUID, error) {
	userIDStr := ctx.MustGet("user_id").(string)
	return uuid.Parse(userIDStr)
}
