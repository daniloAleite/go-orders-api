package main

import (
	"log"
	"net/http"

	"github.com/daniloAleite/go-orders-api/internal/server"
)

func main() {

	log.Printf("Initializing server settings!")
	srv := server.New()

	addr := ":8080"
	log.Printf("server is listening on %s", addr)

	if err := http.ListenAndServe(addr, srv.Handler()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
