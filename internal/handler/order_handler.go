package handler

import (
	"encoding/json"
	"net/http"

	"github.com/daniloAleite/go-orders-api/internal/httpapi"
	model "github.com/daniloAleite/go-orders-api/internal/model/order"
	"github.com/daniloAleite/go-orders-api/internal/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpapi.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	order, err := h.service.Create(req)
	if err != nil {
		httpapi.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	httpapi.WriteJSON(w, http.StatusCreated, order)
}

func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	orders := h.service.List()
	httpapi.WriteJSON(w, http.StatusOK, orders)
}
