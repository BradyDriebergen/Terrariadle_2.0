package server

import (
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
