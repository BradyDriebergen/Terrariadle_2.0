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
			ID:          t.ID,
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

func (g *TerraTrivia) CheckGuess(ctx context.Context, userId string, guess string) (TerraTriviaCheckData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return TerraTriviaCheckData{}, domain.UserNotFound("User not found", err)
	}

	if user.TerraTrivia.Game.Finished {
		return TerraTriviaCheckData{}, domain.Conflict("Game already finished", nil)
	}

	triviaQuestions := g.answerCache.GetAnswers().TerraTrivia.Questions
	checkData := TerraTriviaCheckData{} // Assigns the inital check data as false or ""

	for _, q := range triviaQuestions {
		if q.Answer == guess {
			if slices.Contains(user.TerraTrivia.Game.Guesses, q.ID) {
				return TerraTriviaCheckData{}, domain.Conflict("User previously guessed this question", nil)
			}

			checkData.IsCorrect = true
			checkData.GuessResult = TriviaItem{
				ID:          q.ID,
				Clue:        q.Clue,
				LetterCount: len(q.Answer),
				Answer:      q.Answer,
			}

			user.TerraTrivia.Game.Guesses = append(user.TerraTrivia.Game.Guesses, q.ID)

			if len(user.TerraTrivia.Game.Guesses) >= 7 {
				position, err := g.guessCountCache.IncrementTerraTriviaCount(ctx)
				if err != nil {
					return TerraTriviaCheckData{}, domain.Internal("An error occurred updating user's position", err)
				}

				checkData.Finished = true
				user.TerraTrivia.Game.Finished = true
				user.GuessTheNPC.Game.Position = position
			}

			err = g.userCache.UpsertUser(ctx, user)
			if err != nil {
				return TerraTriviaCheckData{}, domain.Internal("An error occurred updating user's guess", err)
			}

			break
		}
	}

	return checkData, nil
}

func (g *TerraTrivia) GetWinningData(ctx context.Context, userId string) (TerraTriviaWinningData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return TerraTriviaWinningData{}, domain.UserNotFound("User not found", err)
	}

	if !user.TerraTrivia.Game.Finished {
		return TerraTriviaWinningData{}, domain.Conflict("User isn't finished guessing", err)
	}

	return TerraTriviaWinningData{
		Position: user.TerraTrivia.Game.Position,
	}, nil
}
