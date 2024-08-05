package controllers

import (
	"BookStore/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BasketController struct {
	basketService service.BasketService
}

func NewBasketController(basketService service.BasketService) *BasketController {
	return &BasketController{basketService: basketService}
}

func (bc *BasketController) GetBasket(ctx *gin.Context) {
	customerID, err := bc.getCustomerID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	basket, err := bc.basketService.GetBasket(customerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if basket == nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Basket is empty"})
		return
	}
	ctx.JSON(http.StatusOK, basket)
}

func (bc *BasketController) AddToBasket(ctx *gin.Context) {
	customerID, err := bc.getCustomerID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request struct {
		BookID   uuid.UUID `json:"book_id" binding:"required"`
		Quantity int       `json:"quantity" binding:"required,min=1"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = bc.basketService.AddToBasket(customerID, request.BookID, request.Quantity)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Added to basket"})
}

func (bc *BasketController) UpdateBasketItem(ctx *gin.Context) {
	customerID, err := bc.getCustomerID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request struct {
		BookID   uuid.UUID `json:"book_id" binding:"required"`
		Quantity int       `json:"quantity" binding:"required,min=0"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Quantity == 0 {
		err = bc.basketService.RemoveFromBasket(customerID, request.BookID)
	} else {
		err = bc.basketService.UpdateBasketItem(customerID, request.BookID, request.Quantity)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Basket updated"})
}

func (bc *BasketController) RemoveFromBasket(ctx *gin.Context) {
	customerID, err := bc.getCustomerID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookID, err := uuid.Parse(ctx.Param("book_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := bc.basketService.RemoveFromBasket(customerID, bookID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Removed from basket"})
}

func (bc *BasketController) ClearBasket(ctx *gin.Context) {
	customerID, err := bc.getCustomerID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bc.basketService.ClearBasket(customerID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Basket cleared"})
}

func (bc *BasketController) getCustomerID(ctx *gin.Context) (uuid.UUID, error) {
	userIDStr := ctx.MustGet("user_id").(string)
	return uuid.Parse(userIDStr)
}
