package main

import (
	"log"
	"net/http"

	"github.com/daniloAleite/go-orders-api/internal/config"
	"github.com/daniloAleite/go-orders-api/internal/server"
)

func main() {

	log.Printf("initializing server settings")

	cfg := config.Load()

	srv := server.New(cfg)

	addr := ":" + cfg.Port
	log.Printf("starting %s on port %s", cfg.AppName, cfg.Port)

	if err := http.ListenAndServe(addr, srv.Handler()); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
