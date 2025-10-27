package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/health", HealthHandler)            // Checks if backend is running
	r.Route("/api/{mode}", gameSpecificRouters)    // Game specific routes
	r.Get("/api/remaining-time", GetRemainingTime) // Gets the remaining time in the day

	return r
}

func gameSpecificRouters(r chi.Router) {
	r.Get("/puzzle-data", GetPuzzleData)              // Responds puzzle data
	r.Get("/user-guesses/{userId}", GetUserGuesses)   // Responds user guesses/attempts
	r.Post("/check", CheckGuess)                      // Checks user guess
	r.Get("/get-position/{userId}", GetUserPosition)  // Responds user winning position
	r.Get("/players-guessed", GetTotalPlayersGuessed) // Responds total players guessed
}

func NewMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", healthHandler) // Checks backend health

	mux.HandleFunc("GET /api/initialize-game/{userId}", healthHandler) // Gets user guesses and puzzle data
	mux.HandleFunc("POST /api/check-guess", healthHandler)             // Checks user guess
	mux.HandleFunc("GET /api/winning-data/{userId}", healthHandler)    // Gets users position and total players guessed

	mux.HandleFunc("GET /api/remaining-time", healthHandler) // Gets the remaining time in the day

	return mux
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok buddy"}`))
}
