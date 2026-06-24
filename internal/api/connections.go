package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) initializeConnectionsGame(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "missing user_id")
		return
	}

	result, err := s.connections.InitializeGame(r.Context(), userID)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) checkConnectionsGuess(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID string   `json:"user_id"`
		Guess  []string `json:"guess"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body for request")
		return
	}

	result, err := s.connections.CheckGuess(r.Context(), payload.UserID, payload.Guess)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) revealConnectionAnswers(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		UserID string `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid body for request")
		return
	}

	result, err := s.connections.RevealAnswers(r.Context(), payload.UserID)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}

func (s *Server) getConnectionsWinningData(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "missing user_id")
		return
	}

	result, err := s.connections.GetWinningData(r.Context(), userID)
	if err != nil {
		handleError(w, err)
		return
	}

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}
