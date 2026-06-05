package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) initializeDailySlashGame(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "missing user_id")
		return
	}

	result, err := s.dailySlash.InitializeGame(r.Context(), userID)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) getDailySlashSearchItems(w http.ResponseWriter, r *http.Request) {
	result := s.dailySlash.GetSearchableWeapons()

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) getDailySlashHint(w http.ResponseWriter, r *http.Request) {
	hint := r.URL.Query().Get("hint")

	hintNumber, err := strconv.Atoi(hint)
	if err != nil {
		writeError(w, http.StatusBadRequest, "hint must be a number")
		return
	}

	result, err := s.dailySlash.GetHint(hintNumber)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) checkDailySlashGuess(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID string `json:"user_id"`
		Guess  int    `json:"guess"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body for request")
		return
	}

	result, err := s.dailySlash.CheckGuess(r.Context(), payload.UserID, payload.Guess)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) getDailySlashWinningData(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "missing user_id")
		return
	}

	result, err := s.dailySlash.GetWinningData(r.Context(), userID)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}
