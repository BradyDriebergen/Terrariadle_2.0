package api

import "net/http"

func (s *Server) newMux() http.Handler {
	mux := http.NewServeMux()

	s.registerCommonRoutes(mux)
	s.registerDailySlashRoutes(mux)
	s.registerConnectionsRoutes(mux)
	s.registerGuessTheNpcRoutes(mux)
	s.registerHangmanRoutes(mux)
	s.registerTerraTriviaRoutes(mux)

	return withCORS(mux)
}

func (s *Server) registerCommonRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/health", s.checkHealth)
	mux.HandleFunc("GET /api/remaining-time", s.getRemainingTime)
	mux.HandleFunc("GET /api/guess-count", s.guessCountStream)
	mux.HandleFunc("GET /api/finished-games", s.getUserGameStatuses)
}

func (s *Server) registerDailySlashRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/daily-slash/initialize-game", s.initializeDailySlashGame)
	mux.HandleFunc("GET /api/daily-slash/search-items", s.getDailySlashSearchItems)
	mux.HandleFunc("GET /api/daily-slash/hint", s.getDailySlashHint)
	mux.HandleFunc("POST /api/daily-slash/check-guess", s.checkDailySlashGuess)
	mux.HandleFunc("GET /api/daily-slash/winning-data", s.getDailySlashWinningData)
}

func (s *Server) registerConnectionsRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/connections/initialize-game", s.initializeConnectionsGame)
	mux.HandleFunc("POST /api/connections/check-guess", s.checkConnectionsGuess)
	mux.HandleFunc("POST /api/connections/reveal-answers", s.revealConnectionAnswers)
	mux.HandleFunc("GET /api/connections/winning-data", s.getConnectionsWinningData)
}

func (s *Server) registerGuessTheNpcRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/guess-the-npc/initialize-game", s.initializeNpcGame)
	mux.HandleFunc("GET /api/guess-the-npc/search-items", s.getNpcSearchItems)
	mux.HandleFunc("POST /api/guess-the-npc/check-guess", s.checkNpcGuess)
	mux.HandleFunc("GET /api/guess-the-npc/winning-data", s.getNpcWinningData)
	mux.HandleFunc("POST /api/guess-the-npc/check-name-guess", s.checkNpcNameGuess)
}

func (s *Server) registerHangmanRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/hangman/initialize-game", s.initializeHangmanGame)
	mux.HandleFunc("POST /api/hangman/check-guess", s.checkHangmanGuess)
	mux.HandleFunc("GET /api/hangman/winning-data", s.getHangmanWinningData)
}

func (s *Server) registerTerraTriviaRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/terratrivia/initialize-game", s.initializeTerraTriviaGame)
	mux.HandleFunc("POST /api/terratrivia/check-guess", s.checkTerraTriviaGuess)
	mux.HandleFunc("GET /api/terratrivia/winning-data", s.getTerraTriviaWinningData)
}
