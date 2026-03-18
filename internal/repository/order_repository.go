package repository

import (
	"sync"
	"time"

	orderModel "github.com/daniloAleite/go-orders-api/internal/model/order"
)

type OrderRepository struct {
	mu     sync.Mutex
	data   []orderModel.Order
	nextID int64
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		data:   make([]orderModel.Order, 0),
		nextID: 1,
	}
}

func (r *OrderRepository) Create(customer string, amount float64) orderModel.Order {
	r.mu.Lock()
	defer r.mu.Unlock()

	order := orderModel.Order{
		ID:        r.nextID,
		Customer:  customer,
		Amount:    amount,
		Status:    "created",
		CreatedAt: time.Now(),
	}

	r.data = append(r.data, order)
	r.nextID++

	return order
}

func (r *OrderRepository) List() []orderModel.Order {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]orderModel.Order, len(r.data))
	copy(result, r.data)

	return result
}
