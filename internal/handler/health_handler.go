package handler

import (
	"net/http"
	"time"

	"github.com/daniloAleite/go-orders-api/internal/httpapi"
	model "github.com/daniloAleite/go-orders-api/internal/model/health"
)

type HealthHandler struct {
	appName    string
	AppVersion string
}

func NewHealthHandler(appName, appVersion string) *HealthHandler {
	return &HealthHandler{
		appName:    appName,
		AppVersion: appVersion,
	}
}

func (h *HealthHandler) GetStatus(w http.ResponseWriter, r *http.Request) {

	health := model.HealthResponse{

		Status:    "stable",
		Version:   h.AppVersion,
		Timestamp: time.Now(),
	}

	httpapi.WriteJSON(w, http.StatusOK, health)
}
