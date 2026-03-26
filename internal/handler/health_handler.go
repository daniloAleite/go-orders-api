package handler

import (
	"net/http"
	"time"

	"github.com/daniloAleite/go-orders-api/internal/httpapi"
	model "github.com/daniloAleite/go-orders-api/internal/model/health"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GetStatus(w http.ResponseWriter, r *http.Request) {

	health := model.HealthResponse{

		Status:    "stable",
		Version:   "1.0.0",
		Timestamp: time.Now(),
	}

	httpapi.WriteJSON(w, http.StatusOK, health)
}
