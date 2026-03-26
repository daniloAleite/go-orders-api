package service

import (
	"errors"
	"strings"

	orderModel "github.com/daniloAleite/go-orders-api/internal/model/order"
)

type OrderRepository interface {
	Create(customer string, amount float64) (orderModel.Order, error)
	List() ([]orderModel.Order, error)
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (s *OrderService) Create(req orderModel.CreateOrderRequest) (orderModel.Order, error) {
	customer := strings.TrimSpace(req.Customer)

	if customer == "" {
		return orderModel.Order{}, errors.New("customer is required")
	}

	if req.Amount <= 0 {
		return orderModel.Order{}, errors.New("amount must be greater than zero")
	}

	order, err := s.repo.Create(customer, req.Amount)
	if err != nil {
		return orderModel.Order{}, err
	}
	return order, nil
}

func (s *OrderService) List() ([]orderModel.Order, error) {

	orders, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
