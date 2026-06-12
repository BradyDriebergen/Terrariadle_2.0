package services

import "terrariadle-backend/internal/store"

type SseStreamService interface {
}

type SseStream struct {
	guessCountCache store.GuessCountsStore
}

func NewSseStream()
