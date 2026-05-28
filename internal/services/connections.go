package services

import (
	"context"
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

	options := g.answerCache.GetAnswers().Connections.Options
	solvedCategories := []SolvedCategory{}

	// filter options down

	return ConnectionsInitData{
		Attepts:          user.Connections.Attempts,
		Finished:         user.Connections.Game.Finished,
		Options:          options,
		SolvedCategories: solvedCategories,
	}, nil
}
