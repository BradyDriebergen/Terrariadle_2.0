package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"terrariadle-backend/internal/services"
	"terrariadle-backend/internal/utils"

	"github.com/go-chi/chi/v5"
)

// Used for check-guess post body
type GuessRequest struct {
	UserID string `json:"userId"`
	Guess  int    `json:"guess"`
}

// Handler for checking the health of the backend
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// Handler for getting the daily puzzle data
func GetPuzzleData(w http.ResponseWriter, r *http.Request) {
	mode := chi.URLParam(r, "mode")

	switch mode {
	case "daily-slash":
		data, err := services.GetDailySlashPuzzleData()
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		utils.WriteJSON(w, http.StatusOK, data)
		return
	case "connections":
	case "guess-the-npc":
	case "hangman":
	default:
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found", mode),
		})
		return
	}
}

// Handler for getting the user guesses
func GetUserGuesses(w http.ResponseWriter, r *http.Request) {
	mode := chi.URLParam(r, "mode")
	userId := chi.URLParam(r, "userId")

	switch mode {
	case "daily-slash":
		data, err := services.GetDailySlashUserGuesses(userId)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		utils.WriteJSON(w, http.StatusOK, data)
		return
	case "connections":
	case "guess-the-npc":
	case "hangman":
	default:
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found", mode),
		})
		return
	}
}

// Handler for checking the user guesses and saving the guess in the database
func CheckGuess(w http.ResponseWriter, r *http.Request) {
	mode := chi.URLParam(r, "mode")
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
			utils.WriteJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		utils.WriteJSON(w, http.StatusOK, map[string]any{"won": won, "guess": guess})
		return
	case "connections":
	case "guess-the-npc":
	case "hangman":
	default:
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found", mode),
		})
		return
	}
}

// Handler for getting the 'winning' position of the user
func GetUserPosition(w http.ResponseWriter, r *http.Request) {
	mode := chi.URLParam(r, "mode")
	userId := chi.URLParam(r, "userId")

	switch mode {
	case "daily-slash":
		pos, err := services.GetDailySlashPlayerPosition(userId)
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		utils.WriteJSON(w, http.StatusOK, pos)
		return
	case "connections":
	case "guess-the-npc":
	case "hangman":
	default:
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found", mode),
		})
		return
	}
}

// Handler for getting the total number of players that guessed correctly
func GetTotalPlayersGuessed(w http.ResponseWriter, r *http.Request) {
	mode := chi.URLParam(r, "mode")

	switch mode {
	case "daily-slash":
		data, err := services.GetDailySlashPlayersGuessed()
		if err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		utils.WriteJSON(w, http.StatusOK, data)
		return
	case "connections":
	case "guess-the-npc":
	case "hangman":
	default:
		utils.WriteJSON(w, http.StatusNotFound, map[string]any{
			"error": fmt.Sprintf("Gamemode %s not found", mode),
		})
		return
	}
}

// Handler for getting the remaining time in the day in seconds
func GetRemainingTime(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, int64(utils.TimeUntilNextMidnightFromNow().Seconds()))
}
