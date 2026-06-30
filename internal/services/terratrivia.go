package services

import (
	"context"
	"slices"
	"terrariadle-backend/internal/domain"
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

func (g *TerraTrivia) InitializeGame(ctx context.Context, userId string) (TerraTriviaInitData, error) {
	user, err := g.userCache.GetOrCreateUser(ctx, userId)
	if err != nil {
		return TerraTriviaInitData{}, domain.UserNotFound("Error creating user", err)
	}

	triviaAnswers := g.answerCache.GetAnswers().TerraTrivia.Questions
	if len(triviaAnswers) != 7 {
		return TerraTriviaInitData{}, domain.Internal("Internal error with trivia answers", nil)
	}

	triviaItems := make([]TriviaItem, 7)
	chunks := []string{}

	for i, t := range triviaAnswers {
		answer := ""

		if slices.Contains(user.TerraTrivia.Game.Guesses, t.ID) {
			answer = t.Answer
		} else {
			chunks = append(chunks, t.Chunks...)
		}

		triviaItems[i] = TriviaItem{
			Clue:        t.Clue,
			LetterCount: len(t.Answer),
			Answer:      answer,
		}
	}

	return TerraTriviaInitData{
		Finished:    user.TerraTrivia.Game.Finished,
		TriviaItems: triviaItems,
		Chunks:      chunks,
	}, nil
}
