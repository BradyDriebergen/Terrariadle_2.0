package server

import (
	"github.com/go-chi/chi/v5"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/health", HealthHandler)

	// Game specific routes
	r.Route("/api/{mode}", gameSpecificRouters)

	r.Get("/api/remaining-time", GetRemainingTime)

	return r
}

func gameSpecificRouters(r chi.Router) {
	r.Get("/puzzle-data", GetPuzzleData)            // Responds puzzle data
	r.Get("/user-guesses/{userId}", GetUserGuesses) // Responds user guesses/attempts

	r.Post("/check", CheckGuess) // Checks user guess

	r.Get("/get-position/{userId}", GetUserPosition)  // Responds user winning position
	r.Get("/players-guessed", GetTotalPlayersGuessed) // Responds total players guessed
}
