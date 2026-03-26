package main

import (
	"log"
	"net/http"

	"github.com/daniloAleite/go-orders-api/internal/config"
	"github.com/daniloAleite/go-orders-api/internal/db"
	"github.com/daniloAleite/go-orders-api/internal/server"
)

func main() {

	log.Printf("initializing server settings")

	cfg := config.Load()

	postgressDB, err := db.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer postgressDB.Close()

	srv := server.New(cfg, postgressDB)

	addr := ":" + cfg.Port
	log.Printf("starting %s on port %s", cfg.AppName, cfg.Port)

	if err := http.ListenAndServe(addr, srv.Handler()); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
