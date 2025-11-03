package server

import (
	"net/http"
	"slices"
)

// withCORS wraps a handler to add CORS headers and handle OPTIONS preflights.
func withCORS(next http.Handler) http.Handler {
	// Allowed dev origins. Add others as needed.
	allowedOrigins := []string{
		"http://localhost:5173", // SvelteKit dev
		"http://127.0.0.1:5173", // sometimes used
		"http://localhost:4173", // Vite/SvelteKit preview
		"http://127.0.0.1:4173",
	}

	allowedMethods := "GET, POST, PUT, PATCH, DELETE, OPTIONS"
	allowedHeaders := "Content-Type, Authorization"

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && slices.Contains(allowedOrigins, origin) {
			// IMPORTANT: echo the requesting origin (not "*") if you might ever use credentials.
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin") // keep caches honest
			// Uncomment if you need cookies/Authorization across origins:
			// w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		// Always advertise what we allow.
		w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)

		// Handle preflight early.
		if r.Method == http.MethodOptions {
			// You can return 204; 200 is fine too.
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", HealthHandler) // Checks backend health

	mux.HandleFunc("GET /api/{mode}/initialize-game/{userId}", InitializeGame) // Gets user guesses and puzzle data
	mux.HandleFunc("GET /api/{mode}/search", Search)                           // Searches the data and returns a reduced list
	mux.HandleFunc("POST /api/check-guess", CheckGuess)                        // Checks user guess
	mux.HandleFunc("GET /api/{mode}/winning-data/{userId}", HealthHandler)     // Gets users position and total players guessed

	mux.HandleFunc("GET /api/remaining-time", HealthHandler) // Gets the remaining time in the day

	return withCORS(mux)
}
