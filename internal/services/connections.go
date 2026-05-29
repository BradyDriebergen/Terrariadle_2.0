package services

import (
	"context"
	"slices"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
)

type ConnectionsService interface {
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
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return ConnectionsInitData{}, domain.UserNotFound("User not found", err)
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
		Attepts:          user.Connections.Attempts,
		Finished:         user.Connections.Game.Finished,
		Options:          options,
		SolvedCategories: guessedCategories,
	}, nil
}
