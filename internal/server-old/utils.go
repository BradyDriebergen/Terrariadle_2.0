package server

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"

	"github.com/google/uuid"
)

// Helper method for writing a response
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Checks if a UUID is valid
func isValidUUID(id string) bool {
	if id == "" {
		return false
	}
	_, err := uuid.Parse(id)
	return err == nil
}

func shuffle[T any](list []T) {
	rand.Shuffle(len(list), func(i, j int) {
		list[i], list[j] = list[j], list[i]
	})
}
