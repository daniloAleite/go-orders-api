package server

import (
	"net/http"

	"github.com/daniloAleite/go-orders-api/internal/config"
	"github.com/daniloAleite/go-orders-api/internal/handler"
	"github.com/daniloAleite/go-orders-api/internal/repository"
	"github.com/daniloAleite/go-orders-api/internal/service"
)

type Server struct {
	mux *http.ServeMux
}

func New(cfg *config.Config) *Server {
	mux := http.NewServeMux()

	healthHandler := handler.NewHealthHandler(cfg.AppName, cfg.AppVersion)

	//order
	orderRepo := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	// Criação dos endpoints
	mux.HandleFunc("GET /health", healthHandler.GetStatus)
	mux.HandleFunc("GET /orders", orderHandler.List)
	mux.HandleFunc("POST /orders", orderHandler.Create)

	return &Server{
		mux: mux,
	}
}

func (s *Server) Handler() http.Handler {
	return s.mux
}
