package services

import (
	"context"
	"terrariadle-backend/internal/store"
)

type TerraTriviaService interface {
	InitializeGame(ctx context.Context, userId string) (HangmanInitData, error)
	CheckGuess(ctx context.Context, userId string, guess string) (HangmanCheckData, error)
	GetWinningData(ctx context.Context, userId string) (HangmanWinningData, error)
}

type TerraTrivia struct {
	answerCache     store.AnswerStore
	guessCountCache store.GuessCountsStore
	catalogCache    store.CatalogStore
	userCache       store.UserStore
}

func NewTerraTriviaGame(
	answerCache store.AnswerStore,
	guessCountCache store.GuessCountsStore,
	catalogCache store.CatalogStore,
	userCache store.UserStore,
) *TerraTrivia {

	return &TerraTrivia{
		answerCache:     answerCache,
		guessCountCache: guessCountCache,
		catalogCache:    catalogCache,
		userCache:       userCache,
	}
}
