package routes

import (
	"terrariadle-backend/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", handlers.HealthHandler)

	return r
}
