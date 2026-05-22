package handlers

// import "net/http"

// func NewMux() http.Handler {
// 	mux := http.NewServeMux()

// 	registerCommonRoutes(mux)
// 	registerDailySlashRoutes(mux)
// 	registerConnectionsRoutes(mux)
// 	registerGuessTheNpcRoutes(mux)
// 	registerHangmanRoutes(mux)

// 	return mux
// }

// func registerCommonRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("GET /api/health", CheckHealth)
// 	mux.HandleFunc("GET /api/remaining-time", GetRemainingTime)
// }

// func registerDailySlashRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("GET /api/daily-slash/initialize-game/", InitializeDailySlashGame)
// 	mux.HandleFunc("GET /api/daily-slash/search-items", GetDailySlashSearchItems)
// 	mux.HandleFunc("GET /api/daily-slash/hint/", GetDailySlashHint)
// 	mux.HandleFunc("POST /api/daily-slash/check-guess", CheckDailySlashGuess)
// 	mux.HandleFunc("GET /api/daily-slash/winning-data/", GetDailySlashWinningData)
// }

// func registerConnectionsRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("GET /api/connections/initialize-game/", InitializeConnectionsGame)
// 	mux.HandleFunc("POST /api/connections/check-guess", CheckConnectionsGuess)
// 	mux.HandleFunc("GET /api/connections/winning-data/", GetConnectionsWinningData)
// }

// func registerGuessTheNpcRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("GET /api/guess-the-npc/initialize-game/", InitializeNpcGame)
// 	mux.HandleFunc("GET /api/guess-the-npc/search-items", GetNpcSearchItems)
// 	mux.HandleFunc("POST /api/guess-the-npc/check-guess", CheckNpcGuess)
// 	mux.HandleFunc("GET /api/guess-the-npc/winning-data/", GetNpcWinningData)
// 	mux.HandleFunc("POST /api/guess-the-npc/check-name-guess", CheckNpcNameGuess)
// }

// func registerHangmanRoutes(mux *http.ServeMux) {
// 	mux.HandleFunc("GET /api/hangman/initialize-game/", InitializeHangmanGame)
// 	mux.HandleFunc("POST /api/hangman/check-guess", CheckHangmanGuess)
// 	mux.HandleFunc("GET /api/hangman/winning-data/", GetHangmanWinningData)
// }
