package service

import (
	"errors"
	"strings"

	orderModel "github.com/daniloAleite/go-orders-api/internal/model/order"
	"github.com/daniloAleite/go-orders-api/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
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

	order := s.repo.Create(customer, req.Amount)
	return order, nil
}

func (s *OrderService) List() []orderModel.Order {
	return s.repo.List()
}
