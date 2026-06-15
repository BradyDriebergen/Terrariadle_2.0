package services

import (
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
)

type SseStreamService interface {
	GetGuessCount(game domain.GameMode) (int, error)
}

type SseStream struct {
	guessCountCache store.GuessCountsStore
}

func NewSseStream(guessCountCache store.GuessCountsStore) *SseStream {
	return &SseStream{
		guessCountCache: guessCountCache,
	}
}

func (s *SseStream) GetGuessCount(game domain.GameMode) (int, error) {
	switch game {
	case "daily_slash":
		return s.guessCountCache.GetGuessCounts().DailySlashCount, nil
	case "connections":
		return s.guessCountCache.GetGuessCounts().ConnectionsCount, nil
	case "guess_the_npc":
		return s.guessCountCache.GetGuessCounts().GuessTheNpcCount, nil
	case "hangman":
		return s.guessCountCache.GetGuessCounts().HangmanCount, nil
	default:
		return 0, domain.NotFound("The requested guess count doesn't exist", nil)
	}
}
