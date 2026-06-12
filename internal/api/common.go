package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"terrariadle-backend/internal/domain"
	"time"
)

func (s *Server) CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) GetRemainingTime(w http.ResponseWriter, r *http.Request) {
	remaining := domain.TimeUntilNextMidnight(time.Now())

	writeJSON(w, http.StatusOK, remaining.Seconds())
}

func (s *Server) GuessCountStream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

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

	// TODO: Fix this so it sends the current guess count to the user on initial load
	// currentCount := s.broker.CurrentCount(mode)
	// fmt.Fprintf(w, "data: {\"count\":%d}\n\n", currentCount)
	// flusher.Flush()

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
