package server

import (
	"net/http"
)

func NewMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", HealthHandler) // Checks backend health

	mux.HandleFunc("GET /api/initialize-game/{userId}", InitializeGame) // Gets user guesses and puzzle data
	mux.HandleFunc("GET /api/search", HealthHandler)                    // Searches the data and returns a reduced list
	mux.HandleFunc("POST /api/check-guess", CheckGuess)                 // Checks user guess
	mux.HandleFunc("GET /api/winning-data/{userId}", HealthHandler)     // Gets users position and total players guessed

	mux.HandleFunc("GET /api/remaining-time", HealthHandler) // Gets the remaining time in the day

	return mux
}
