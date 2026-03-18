package handler

import (
	"encoding/json"
	"net/http"
	"time"

	model "github.com/daniloAleite/go-orders-api/internal/model/health"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GetStatus(w http.ResponseWriter, r *http.Request) {

	now := time.Now()

	health := model.HealthResponse{

		Status:    "stable",
		Version:   "1.0.0",
		Timestamp: now.Format("02/01/2006 03:04 PM"),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(health); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *OrderHandler) List(w http.ResponseWriter, r *http.Request) {
	orders := h.service.List()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(orders); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
