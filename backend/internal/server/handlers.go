package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"terrariadle-backend/internal/services"
	"terrariadle-backend/internal/utils"
)

// Used for check-guess post body daily slash
type DailySlashCheckRequest struct {
	UserID string `json:"userId"`
	Guess  int    `json:"guess"`
}

// Used for check-guess post body connections
type ConnectionsCheckRequest struct {
	UserID string   `json:"userId"`
	Guess  []string `json:"guess"`
}

// Handler for checking the health of the backend
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func GetHint(w http.ResponseWriter, r *http.Request) {
	hintNum, err := strconv.Atoi(r.PathValue("hintNum"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
	}

	searchItems, err := services.GetDailySlashHint(hintNum)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
	}
	writeJSON(w, http.StatusOK, searchItems)
}

func GetSearchItems(w http.ResponseWriter, r *http.Request) {
	mode := r.PathValue("mode")

	switch mode {
	case "daily-slash":
		searchItems, err := services.GetDailySlashSearchItems()
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, searchItems)
		return
	case "guess-the-npc":
		searchItems, err := services.GetNpcSearchItems()
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, searchItems)
		return
	default:
		writeJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found for searching items", mode),
		})
		return
	}
}

// Handler for getting the data for the user to start playing
func InitializeGame(w http.ResponseWriter, r *http.Request) {
	mode := r.PathValue("mode")
	userId := r.PathValue("userId")

	switch mode {
	case "daily-slash":
		previousWeapon, guesses, checks, won, err := services.InitializeDailySlashGame(userId)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, map[string]any{"previousWeapon": previousWeapon, "guesses": guesses, "checks": checks, "won": won})
		return
	case "connections":
		attempts, options, guesses, won, err := services.InitializeConnectionsGame(userId)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, map[string]any{"attempts": attempts, "options": options, "guesses": guesses, "finished": won})
		return
	case "guess-the-npc":
		quote, guesses, won, err := services.InitializeNpcGame(userId)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, map[string]any{"quote": quote, "guesses": guesses, "won": won})
		return
	case "hangman":
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

	switch mode {
	case "daily-slash":
		// Decodes request body into json
		var req DailySlashCheckRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		won, guess, check, err := services.CheckDailySlashGuess(req.UserID, req.Guess)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, map[string]any{"won": won, "guess": guess, "check": check})
		return
	case "connections":
		// Decodes request body into json
		var req ConnectionsCheckRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		guess, oneAway, won, err := services.CheckConnectionsGuess(req.UserID, req.Guess)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, map[string]any{"guess": guess, "oneAway": oneAway, "finished": won})
		return
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
		pos, count, err := services.GetConnectionsWinningData(userId)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		writeJSON(w, http.StatusOK, map[string]any{"pos": pos, "count": count})
		return
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
	data := map[string]int64{
		"remaining": int64(utils.TimeUntilNextMidnightFromNow().Seconds()),
	}

	writeJSON(w, http.StatusOK, data)
}

// Helper method for writing a response
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
