package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) initializeNpcGame(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "missing user_id")
		return
	}

	result, err := s.guessTheNpc.InitializeGame(r.Context(), userID)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) getNpcSearchItems(w http.ResponseWriter, r *http.Request) {
	result := s.guessTheNpc.GetSearchableNpcs()

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) checkNpcGuess(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID string `json:"user_id"`
		Guess  int    `json:"guess"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body for request")
		return
	}

	result, err := s.guessTheNpc.CheckGuess(r.Context(), payload.UserID, payload.Guess)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) getNpcWinningData(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "missing user_id")
		return
	}

	result, err := s.guessTheNpc.GetWinningData(r.Context(), userID)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) checkNpcNameGuess(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID string `json:"user_id"`
		Guess  string `json:"guess"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body for request")
		return
	}

	result, err := s.guessTheNpc.CheckName(r.Context(), payload.UserID, payload.Guess)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}
