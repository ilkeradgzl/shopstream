package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ilkeradgzl/shopstream/services/user-service/internal/api"
	"github.com/ilkeradgzl/shopstream/services/user-service/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Get("/healthz", api.HealthHandler)
	r.Route("/users", func(r chi.Router) {
		r.Get("/{userID}", api.GetUserHandler)
	})

	addr := ":" + cfg.ServerPort
	log.Printf("User Service listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
