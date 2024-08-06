package controllers

import (
	"BookStore/pkg/service"
	"BookStore/pkg/utils"
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
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	basket, err := bc.basketService.GetBasket(customerID)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if basket == nil {
		utils.JSONResponse(ctx, http.StatusOK, "Basket is empty", nil)
		return
	}
	utils.JSONResponse(ctx, http.StatusOK, "Basket retrieved", basket)
}

func (bc *BasketController) AddToBasket(ctx *gin.Context) {
	customerID, err := bc.getCustomerID(ctx)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var request struct {
		BookID   uuid.UUID `json:"book_id" binding:"required"`
		Quantity int       `json:"quantity" binding:"required,min=1"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = bc.basketService.AddToBasket(customerID, request.BookID, request.Quantity)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.JSONResponse(ctx, http.StatusOK, "Added to basket", nil)
}

func (bc *BasketController) UpdateBasketItem(ctx *gin.Context) {
	customerID, err := bc.getCustomerID(ctx)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var request struct {
		BookID   uuid.UUID `json:"book_id" binding:"required"`
		Quantity int       `json:"quantity" binding:"required,min=0"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if request.Quantity == 0 {
		err = bc.basketService.RemoveFromBasket(customerID, request.BookID)
	} else {
		err = bc.basketService.UpdateBasketItem(customerID, request.BookID, request.Quantity)
	}
	if err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.JSONResponse(ctx, http.StatusOK, "Basket updated", nil)
}

func (bc *BasketController) RemoveFromBasket(ctx *gin.Context) {
	customerID, err := bc.getCustomerID(ctx)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	bookID, err := uuid.Parse(ctx.Param("book_id"))
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, "Invalid book ID", nil)
		return
	}

	if err := bc.basketService.RemoveFromBasket(customerID, bookID); err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.JSONResponse(ctx, http.StatusOK, "Removed from basket", nil)
}

func (bc *BasketController) ClearBasket(ctx *gin.Context) {
	customerID, err := bc.getCustomerID(ctx)
	if err != nil {
		utils.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := bc.basketService.ClearBasket(customerID); err != nil {
		utils.JSONResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.JSONResponse(ctx, http.StatusOK, "Basket cleared", nil)
}

func (bc *BasketController) getCustomerID(ctx *gin.Context) (uuid.UUID, error) {
	userIDStr := ctx.MustGet("user_id").(string)
	return uuid.Parse(userIDStr)
}
