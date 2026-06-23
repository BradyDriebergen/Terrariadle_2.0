package services

import (
	"context"
	"slices"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
)

type ConnectionsService interface {
	InitializeGame(ctx context.Context, userId string) (ConnectionsInitData, error)
	CheckGuess(ctx context.Context, userId string, guessedOptions []string) (ConnectionsCheckData, error)
	GetWinningData(ctx context.Context, userId string) (ConnectionsWinningData, error)
}

type Connections struct {
	answerCache     store.AnswerStore
	guessCountCache store.GuessCountsStore
	catalogCache    store.CatalogStore
	userCache       store.UserStore
}

func NewConnectionsGame(
	answerCache store.AnswerStore,
	guessCountCache store.GuessCountsStore,
	catalogCache store.CatalogStore,
	userCache store.UserStore,
) *Connections {

	return &Connections{
		answerCache:     answerCache,
		guessCountCache: guessCountCache,
		catalogCache:    catalogCache,
		userCache:       userCache,
	}
}

func (g *Connections) InitializeGame(ctx context.Context, userId string) (ConnectionsInitData, error) {
	user, err := g.userCache.GetOrCreateUser(ctx, userId)
	if err != nil {
		return ConnectionsInitData{}, domain.UserNotFound("Error creating user", err)
	}

	answerOptions := g.answerCache.GetAnswers().Connections.Options
	userGuesses := user.Connections.Game.Guesses

	guessedCategoryMap := make(map[int][]string) // category id -> options slice
	guessedCategories := []SolvedCategory{}
	options := []string{}

	for _, o := range answerOptions {
		if slices.Contains(userGuesses, o.CategoryID) {
			guessedCategoryMap[o.CategoryID] = append(guessedCategoryMap[o.CategoryID], o.Option)
		} else {
			options = append(options, o.Option)
		}
	}

	for catId := range guessedCategoryMap {
		category, ok := g.catalogCache.GetCategory(catId)
		if !ok {
			return ConnectionsInitData{}, domain.Internal("Failed to find matching category", nil)
		}

		guessedCategories = append(guessedCategories, SolvedCategory{
			Name:    category.Category,
			Options: guessedCategoryMap[catId],
		})
	}

	shuffleOptions(options)

	return ConnectionsInitData{
		Attempts:         user.Connections.Attempts,
		Finished:         user.Connections.Game.Finished,
		Options:          options,
		SolvedCategories: guessedCategories,
	}, nil
}

func (g *Connections) CheckGuess(ctx context.Context, userId string, guessedOptions []string) (ConnectionsCheckData, error) {
	if err := validateGuessedOptions(guessedOptions); err != nil {
		return ConnectionsCheckData{}, err
	}

	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return ConnectionsCheckData{}, domain.UserNotFound("User not found", err)
	}

	if user.Connections.Attempts <= 0 {
		return ConnectionsCheckData{}, domain.Conflict("User does not have any more attempts", nil)
	}

	answerOptions := g.answerCache.GetAnswers().Connections.Options
	answerMap := make(map[string]int, 16)
	for _, o := range answerOptions {
		answerMap[o.Option] = o.CategoryID
	}

	oneAway := false
	isCorrect := false
	correctGuess := SolvedCategory{}

	guessedCategoryIdMap := make(map[int]int) // id -> id count

	for _, guess := range guessedOptions {
		guessedID, ok := answerMap[guess]
		if !ok {
			return ConnectionsCheckData{}, domain.NotFound("Connection guess is not an option", nil)
		}

		guessedCategoryIdMap[guessedID] += 1
		switch guessedCategoryIdMap[guessedID] {
		case 3:
			oneAway = true
		case 4:
			if slices.Contains(user.Connections.Game.Guesses, guessedID) {
				return ConnectionsCheckData{}, domain.Conflict("User already guessed this category", nil)
			}

			oneAway = false
			isCorrect = true

			correctCategory, ok := g.catalogCache.GetCategory(guessedID)
			if !ok {
				return ConnectionsCheckData{}, domain.Internal("Catagory does not exist", nil)
			}

			correctGuess = SolvedCategory{
				Name:    correctCategory.Category,
				Options: guessedOptions,
			}

			user.Connections.Game.Guesses = append(user.Connections.Game.Guesses, guessedID)
		}
	}

	if !isCorrect {
		user.Connections.Attempts--

		if user.Connections.Attempts <= 0 {
			user.Connections.Game.Finished = true
		}
	} else if len(user.Connections.Game.Guesses) == 4 {
		position, err := g.guessCountCache.IncrementConnectionsCount(ctx)
		if err != nil {
			return ConnectionsCheckData{}, domain.Internal("An error occurred updating user's position", err)
		}

		user.Connections.Game.Finished = true
		user.Connections.Game.Position = position
	}

	err = g.userCache.UpsertUser(ctx, user)
	if err != nil {
		return ConnectionsCheckData{}, domain.Internal("An error occurred updating user's guess", err)
	}

	return ConnectionsCheckData{
		Attempts:     user.Connections.Attempts,
		IsCorrect:    isCorrect,
		OneAway:      oneAway,
		CorrectGuess: correctGuess,
		Finished:     user.Connections.Game.Finished,
	}, nil
}

func (g *Connections) revealAnswers(ctx context.Context, userId string) (ConnectionsRevealData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return ConnectionsRevealData{}, domain.UserNotFound("User not found", err)
	}

	if !user.Connections.Game.Finished {
		return ConnectionsRevealData{}, domain.Conflict("User isn't finished guessing", err)
	}
}

func (g *Connections) GetWinningData(ctx context.Context, userId string) (ConnectionsWinningData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return ConnectionsWinningData{}, domain.UserNotFound("User not found", err)
	}

	if !user.Connections.Game.Finished {
		return ConnectionsWinningData{}, domain.Conflict("User isn't finished guessing", err)
	}

	return ConnectionsWinningData{
		Position: user.Connections.Game.Position,
	}, nil
}
