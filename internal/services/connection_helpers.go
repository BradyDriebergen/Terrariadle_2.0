package services

import (
	"math/rand/v2"
	"terrariadle-backend/internal/domain"
)

func shuffleOptions(options []string) {
	rand.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})
}

func validateGuessedOptions(options []string) error {
	if len(options) != 4 {
		return domain.InvalidInput("Guess must contain 4 options", nil)
	}
	seen := make(map[string]struct{}, 4)
	for _, opt := range options {
		if _, exists := seen[opt]; exists {
			return domain.InvalidInput("Duplicate guess options", nil)
		}
		seen[opt] = struct{}{}
	}
	return nil
}
