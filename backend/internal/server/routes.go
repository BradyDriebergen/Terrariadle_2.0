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

	registerCommonRoutes(mux)
	registerDailySlashRoutes(mux)
	registerConnectionsRoutes(mux)
	registerGuessTheNpcRoutes(mux)
	// registerHangmanRoutes(mux) ...

	return withCORS(mux)
}

func registerCommonRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/health", healthHandler)
	mux.HandleFunc("GET /api/remaining-time", getRemainingTime)
}

func registerDailySlashRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/daily-slash/initialize-game/{userId}", getDailySlashInitData)
	mux.HandleFunc("GET /api/daily-slash/search-items", getDailySlashSearchItems)
	mux.HandleFunc("GET /api/daily-slash/hint/{hintNum}", getDailySlashHint)
	mux.HandleFunc("POST /api/daily-slash/check-guess", postDailySlashGuess)
	mux.HandleFunc("GET /api/daily-slash/winning-data/{userId}", getDailySlashWinningData)
}

func registerConnectionsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/connections/initialize-game/{userId}", getConnectionInitData)
	mux.HandleFunc("POST /api/connections/check-guess", postConnectionsGuess)
	mux.HandleFunc("GET /api/connections/winning-data/{userId}", getConnectionsWinningData)
}

func registerGuessTheNpcRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/guess-the-npc/search-items", getNpcSearchItems)
	mux.HandleFunc("GET /api/guess-the-npc/initialize-game/{userId}", getNpcInitGame)
	mux.HandleFunc("POST /api/guess-the-npc/check-guess", postNpcGuess)
	mux.HandleFunc("GET /api/guess-the-npc/winning-data/{userId}", getNpcWinningData)
}
