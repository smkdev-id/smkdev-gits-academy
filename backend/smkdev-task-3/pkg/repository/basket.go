package repository

import (
	"BookStore/pkg/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BasketRepository interface {
	CreateBasket(basket *models.Basket) error
	GetBasketByCustomerID(customerID uuid.UUID) (*models.Basket, error)
	AddToBasket(item *models.BasketItem) error
	UpdateBasketItem(item *models.BasketItem) error
	RemoveFromBasket(basketID, bookID uuid.UUID) error
	ClearBasket(basketID uuid.UUID) error
}

type basketRepo struct {
	db *gorm.DB
}

func NewBasketRepository(db *gorm.DB) BasketRepository {
	return &basketRepo{db: db}
}

func (r *basketRepo) CreateBasket(basket *models.Basket) error {
	return r.db.Create(basket).Error
}

func (r *basketRepo) GetBasketByCustomerID(customerID uuid.UUID) (*models.Basket, error) {
	var basket models.Basket
	if err := r.db.Preload("Items.Book").Where("customer_id = ?", customerID).First(&basket).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &basket, nil
}

func (r *basketRepo) AddToBasket(item *models.BasketItem) error {
	return r.db.Create(item).Error
}

func (r *basketRepo) UpdateBasketItem(item *models.BasketItem) error {
	return r.db.Save(item).Error
}

func (r *basketRepo) RemoveFromBasket(basketID, bookID uuid.UUID) error {
	return r.db.Where("basket_id = ? AND book_id = ?", basketID, bookID).Delete(&models.BasketItem{}).Error
}

func (r *basketRepo) ClearBasket(basketID uuid.UUID) error {
	return r.db.Where("basket_id = ?", basketID).Delete(&models.BasketItem{}).Error
}
