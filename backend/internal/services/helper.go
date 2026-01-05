package services

import (
	"math/rand/v2"

	"github.com/google/uuid"
)

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
