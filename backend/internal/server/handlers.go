package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"terrariadle-backend/internal/services"
	"terrariadle-backend/internal/utils"
)

// Used for check-guess post body
type GuessRequest struct {
	UserID string `json:"userId"`
	Guess  int    `json:"guess"`
}

// Handler for checking the health of the backend
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// Handler for getting the data for the user to start playing
func InitializeGame(w http.ResponseWriter, r *http.Request) {
	mode := r.PathValue("mode")
	userId := r.PathValue("userId")

	switch mode {
	case "daily-slash":
		puzzleData, guesses, err := services.InitializeDailySlashGame(userId)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, map[string]any{"puzzleData": puzzleData, "guesses": guesses})
		return
	case "connections":
	case "guess-the-npc":
	case "hangman":
	default:
		writeJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found", mode),
		})
		return
	}
}

// Handler for searching for items and returning a set ammount of options
func Search(w http.ResponseWriter, r *http.Request) {
	mode := r.PathValue("mode")
	q := r.URL.Query().Get("q")

	switch mode {
	case "daily-slash":
		data, err := services.DailySlashSearch(q)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, data)
		return
	case "guess-the-npc":
	default:
		writeJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found", mode),
		})
		return
	}
}

// Handler for checking the user guesses and saving the guess in the database
func CheckGuess(w http.ResponseWriter, r *http.Request) {
	mode := r.PathValue("mode")

	// Decodes request body into json
	var req GuessRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	switch mode {
	case "daily-slash":
		won, guess, err := services.CheckDailySlashGuess(req.UserID, req.Guess)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, map[string]any{"won": won, "guess": guess})
		return
	case "connections":
	case "guess-the-npc":
	case "hangman":
	default:
		writeJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found", mode),
		})
		return
	}
}

// Handler for getting the 'winning' position of the user
func GetWinningData(w http.ResponseWriter, r *http.Request) {
	mode := r.PathValue("mode")
	userId := r.PathValue("userId")

	switch mode {
	case "daily-slash":
		pos, count, err := services.GetDailySlashWinningData(userId)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, map[string]any{"pos": pos, "count": count})
		return
	case "connections":
	case "guess-the-npc":
	case "hangman":
	default:
		writeJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found", mode),
		})
		return
	}
}

// Handler for getting the remaining time in the day in seconds
func GetRemainingTime(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, int64(utils.TimeUntilNextMidnightFromNow().Seconds()))
}

// Helper method for writing a response
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
