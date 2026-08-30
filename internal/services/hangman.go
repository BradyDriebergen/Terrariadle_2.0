package services

import (
	"context"
	"slices"
	"strings"
	"terrariadle/internal/domain"
	"terrariadle/internal/store"
)

type HangmanService interface {
	InitializeGame(ctx context.Context, userId string) (HangmanInitData, error)
	CheckGuess(ctx context.Context, userId string, guess string) (HangmanCheckData, error)
	GetWinningData(ctx context.Context, userId string) (HangmanWinningData, error)
}

type Hangman struct {
	answerCache     store.AnswerStore
	guessCountCache store.GuessCountsStore
	catalogCache    store.CatalogStore
	userCache       store.UserStore
}

func NewHangmanGame(
	answerCache store.AnswerStore,
	guessCountCache store.GuessCountsStore,
	catalogCache store.CatalogStore,
	userCache store.UserStore,
) *Hangman {

	return &Hangman{
		answerCache:     answerCache,
		guessCountCache: guessCountCache,
		catalogCache:    catalogCache,
		userCache:       userCache,
	}
}

func (g *Hangman) InitializeGame(ctx context.Context, userId string) (HangmanInitData, error) {
	user, err := g.userCache.GetOrCreateUser(ctx, userId)
	if err != nil {
		return HangmanInitData{}, domain.UserNotFound("Error creating user", err)
	}

	enemyAnswer := g.answerCache.GetAnswers().Hangman

	phrase, guessedLetters := buildHangmanPhrase(enemyAnswer.Enemy.Name, user.Hangman.Game.Guesses)

	return HangmanInitData{
		Attempts: user.Hangman.Attempts,
		Finished: user.Hangman.Game.Finished,
		Phrase:   phrase,
		Guesses:  guessedLetters,
	}, nil
}

func (g *Hangman) CheckGuess(ctx context.Context, userId string, guess string) (HangmanCheckData, error) {
	err := validateGuessedLetter(guess)
	if err != nil {
		return HangmanCheckData{}, err
	}

	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return HangmanCheckData{}, domain.UserNotFound("User not found", err)
	}

	if user.Hangman.Game.Finished {
		return HangmanCheckData{}, domain.Conflict("Game already finished", nil)
	}

	enemyAnswer := g.answerCache.GetAnswers().Hangman
	phrase := strings.Split(enemyAnswer.Enemy.Name, "")
	guessedLetters := []string{}

	// Hydrate int guesses into characters
	for _, c := range user.Hangman.Game.Guesses {
		guessedLetters = append(guessedLetters, string(rune(c)))
	}

	if slices.Contains(guessedLetters, guess) {
		return HangmanCheckData{}, domain.Conflict("User already guessed this letter", err)
	}

	guessedLetters = append(guessedLetters, guess)

	isCorrect := slices.Contains(phrase, guess)
	if !isCorrect {
		user.Hangman.Attempts--
		if user.Hangman.Attempts <= 0 {
			user.Hangman.Game.Finished = true
		}
	}

	for i, c := range phrase {
		// Removes any unguessed letters
		if !slices.Contains(guessedLetters, c) && !slices.Contains(specialChars, c) {
			phrase[i] = "_"
		}
	}

	if !slices.Contains(phrase, "_") {
		position, err := g.guessCountCache.IncrementHangmanCount(ctx)
		if err != nil {
			return HangmanCheckData{}, domain.Internal("An error occurred updating user's position", err)
		}

		user.Hangman.Game.Finished = true
		user.Hangman.Game.Position = position
	}

	user.Hangman.Game.Guesses = append(user.Hangman.Game.Guesses, int(guess[0]))

	err = g.userCache.UpsertUser(ctx, user)
	if err != nil {
		return HangmanCheckData{}, domain.Internal("An error occurred updating user's guess", err)
	}

	return HangmanCheckData{
		Phrase:   phrase,
		Finished: user.Hangman.Game.Finished,
		Attempts: user.Hangman.Attempts,
		Guess: HangmanGuess{
			Letter:  guess,
			Correct: isCorrect,
		},
	}, nil
}

func (g *Hangman) GetWinningData(ctx context.Context, userId string) (HangmanWinningData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return HangmanWinningData{}, domain.UserNotFound("User not found", err)
	}

	if !user.Hangman.Game.Finished {
		return HangmanWinningData{}, domain.Conflict("User isn't finished guessing", nil)
	}

	enemyAnswer := g.answerCache.GetAnswers().Hangman

	return HangmanWinningData{
		Position:  user.Hangman.Game.Position,
		EnemyName: enemyAnswer.Enemy.Name,
		EnemyPath: enemyAnswer.Enemy.ImagePath,
	}, nil
}
