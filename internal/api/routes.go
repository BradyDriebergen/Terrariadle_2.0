package api

import "net/http"

func (s *Server) newMux() http.Handler {
	mux := http.NewServeMux()

	s.registerCommonRoutes(mux)
	s.registerDailySlashRoutes(mux)
	s.registerConnectionsRoutes(mux)
	s.registerGuessTheNpcRoutes(mux)
	s.registerHangmanRoutes(mux)

	return withCORS(mux)
}

func (s *Server) registerCommonRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/health", CheckHealth)
	mux.HandleFunc("GET /api/remaining-time", GetRemainingTime)
}

func (s *Server) registerDailySlashRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/daily-slash/initialize-game/", s.initializeDailySlashGame)
	mux.HandleFunc("GET /api/daily-slash/search-items", s.getDailySlashSearchItems)
	mux.HandleFunc("GET /api/daily-slash/hint/", s.getDailySlashHint)
	mux.HandleFunc("POST /api/daily-slash/check-guess", s.checkDailySlashGuess)
	mux.HandleFunc("GET /api/daily-slash/winning-data/", s.getDailySlashWinningData)
}

func (s *Server) registerConnectionsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/connections/initialize-game/", InitializeConnectionsGame)
	mux.HandleFunc("POST /api/connections/check-guess", CheckConnectionsGuess)
	mux.HandleFunc("GET /api/connections/winning-data/", GetConnectionsWinningData)
}

func (s *Server) registerGuessTheNpcRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/guess-the-npc/initialize-game/", InitializeNpcGame)
	mux.HandleFunc("GET /api/guess-the-npc/search-items", GetNpcSearchItems)
	mux.HandleFunc("POST /api/guess-the-npc/check-guess", CheckNpcGuess)
	mux.HandleFunc("GET /api/guess-the-npc/winning-data/", GetNpcWinningData)
	mux.HandleFunc("POST /api/guess-the-npc/check-name-guess", CheckNpcNameGuess)
}

func (s *Server) registerHangmanRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/hangman/initialize-game/", InitializeHangmanGame)
	mux.HandleFunc("POST /api/hangman/check-guess", CheckHangmanGuess)
	mux.HandleFunc("GET /api/hangman/winning-data/", GetHangmanWinningData)
}
