package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"terrariadle-backend/internal/domain"
	"time"
)

func (s *Server) checkHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) getRemainingTime(w http.ResponseWriter, r *http.Request) {
	remaining := int(domain.TimeUntilNextMidnight(time.Now()).Seconds())

	writeJSON(w, http.StatusOK, remaining)
}

func (s *Server) guessCountStream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	mode := domain.GameMode(r.URL.Query().Get("mode"))
	if !validGameMode(mode) {
		http.Error(w, "invalid game mode", http.StatusBadRequest)
		return
	}

	ch := s.broker.Subscribe()
	defer s.broker.Unsubscribe(ch)

	// Ignore error because it's already accounted for
	initialCount, _ := s.common.GetGuessCount(mode)
	initialEvent := domain.GuessCountEvent{GameMode: mode, Count: initialCount}
	data, _ := json.Marshal(initialEvent)
	fmt.Fprintf(w, "data: %s\n\n", data)
	flusher.Flush()

	for {
		select {
		case <-r.Context().Done():
			return
		case event := <-ch:
			if event.GameMode != mode {
				continue // not our game mode, skip it
			}
			data, _ := json.Marshal(event)
			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		}
	}
}

func (s *Server) getUserGameStatuses(w http.ResponseWriter, r *http.Request) {
	userID, ok := getUserID(r)
	if !ok {
		writeError(w, http.StatusBadRequest, "missing user_id")
		return
	}

	result := s.common.GetUserFinishedGames(r.Context(), userID)

	if err := writeJSON(w, http.StatusOK, result); err != nil {
		// Add logger later
	}
}
