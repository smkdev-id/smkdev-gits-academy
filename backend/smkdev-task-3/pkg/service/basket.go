package service

import (
	"BookStore/pkg/models"
	"BookStore/pkg/repository"
	"errors"

	"github.com/google/uuid"
)

type BasketService interface {
	CreateBasket(customerID uuid.UUID) (*models.Basket, error)
	GetBasket(customerID uuid.UUID) (*models.Basket, error)
	AddToBasket(customerID uuid.UUID, bookID uuid.UUID, quantity int) error
	UpdateBasketItem(customerID uuid.UUID, bookID uuid.UUID, quantity int) error
	RemoveFromBasket(customerID uuid.UUID, bookID uuid.UUID) error
	ClearBasket(customerID uuid.UUID) error
}

type basketService struct {
	basketRepo repository.BasketRepository
}

func NewBasketService(basketRepo repository.BasketRepository) BasketService {
	return &basketService{basketRepo: basketRepo}
}

func (s *basketService) CreateBasket(customerID uuid.UUID) (*models.Basket, error) {
	basket := &models.Basket{ID: uuid.New(), CustomerID: customerID}
	err := s.basketRepo.CreateBasket(basket)
	if err != nil {
		return nil, err
	}
	return basket, nil
}

func (s *basketService) GetBasket(customerID uuid.UUID) (*models.Basket, error) {
	return s.basketRepo.GetBasketByCustomerID(customerID)
}

func (s *basketService) AddToBasket(customerID uuid.UUID, bookID uuid.UUID, quantity int) error {
	basket, err := s.basketRepo.GetBasketByCustomerID(customerID)
	if err != nil {
		return err
	}
	if basket == nil {
		basket, err = s.CreateBasket(customerID)
		if err != nil {
			return err
		}
	}

	for i, item := range basket.Items {
		if item.BookID == bookID {
			basket.Items[i].Quantity += quantity
			return s.basketRepo.UpdateBasketItem(&basket.Items[i])
		}
	}

	item := &models.BasketItem{ID: uuid.New(), BasketID: basket.ID, BookID: bookID, Quantity: quantity}
	return s.basketRepo.AddToBasket(item)
}

func (s *basketService) UpdateBasketItem(customerID uuid.UUID, bookID uuid.UUID, quantity int) error {
	basket, err := s.basketRepo.GetBasketByCustomerID(customerID)
	if err != nil {
		return err
	}
	if basket == nil {
		return errors.New("basket not found")
	}

	for i, item := range basket.Items {
		if item.BookID == bookID {
			basket.Items[i].Quantity = quantity
			return s.basketRepo.UpdateBasketItem(&basket.Items[i])
		}
	}

	return errors.New("item not found in basket")
}

func (s *basketService) RemoveFromBasket(customerID uuid.UUID, bookID uuid.UUID) error {
	basket, err := s.basketRepo.GetBasketByCustomerID(customerID)
	if err != nil {
		return err
	}
	if basket == nil {
		return errors.New("basket not found")
	}
	return s.basketRepo.RemoveFromBasket(basket.ID, bookID)
}

func (s *basketService) ClearBasket(customerID uuid.UUID) error {
	basket, err := s.basketRepo.GetBasketByCustomerID(customerID)
	if err != nil {
		return err
	}
	if basket == nil {
		return errors.New("basket not found")
	}
	return s.basketRepo.ClearBasket(basket.ID)
}
