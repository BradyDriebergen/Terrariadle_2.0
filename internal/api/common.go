package api

import (
	"encoding/json"
	"net/http"
	"terrariadle-backend/internal/domain"
	"time"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetRemainingTime(w http.ResponseWriter, r *http.Request) {
	remaining := domain.TimeUntilNextMidnight(time.Now())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"remainingSeconds": int(remaining.Seconds()),
	})
}

// func (s *Server) GuessCountStream(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/event-stream")
// 	w.Header().Set("Cache-Control", "no-cache")
// 	w.Header().Set("Connection", "keep-alive")

// 	flusher, ok := w.(http.Flusher)
// 	if !ok {
// 		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
// 		return
// 	}

// 	mode := domain.GameMode(r.URL.Query().Get("mode"))
// 	// optionally validate mode here

// 	ch := s.broker.Subscribe()
// 	defer s.broker.Unsubscribe(ch)

// 	for {
// 		select {
// 		case <-r.Context().Done():
// 			return
// 		case event := <-ch:
// 			if event.GameMode != mode {
// 				continue // not our game mode, skip it
// 			}
// 			data, _ := json.Marshal(event)
// 			fmt.Fprintf(w, "data: %s\n\n", data)
// 			flusher.Flush()
// 		}
// 	}
// }
