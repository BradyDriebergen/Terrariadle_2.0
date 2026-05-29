package services

import (
	"context"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
)

type GuessTheNpcService interface {
}

type GuessTheNpc struct {
	answerCache     store.AnswerStore
	guessCountCache store.GuessCountsStore
	catalogCache    store.CatalogStore
	userCache       store.UserStore
}

func NewGuessTheNpcGame(
	answerCache store.AnswerStore,
	guessCountCache store.GuessCountsStore,
	catalogCache store.CatalogStore,
	userCache store.UserStore,
) *GuessTheNpc {

	return &GuessTheNpc{
		answerCache:     answerCache,
		guessCountCache: guessCountCache,
		catalogCache:    catalogCache,
		userCache:       userCache,
	}
}

func (g *GuessTheNpc) InitializeGame(ctx context.Context, userId string) (GuessTheNpcInitData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return GuessTheNpcInitData{}, domain.UserNotFound("User not found", err)
	}

	guesses := []domain.NpcInfo{}

	for _, id := range user.GuessTheNPC.Game.Guesses {
		npc, ok := g.catalogCache.GetNpc(id)
		if !ok {
			return GuessTheNpcInitData{}, domain.Internal("Failed to find matching npc", nil)
		}

		guesses = append(guesses, domain.NpcInfo{
			Name: npc.NPC,
			Path: npc.NPCPath,
		})
	}

	return GuessTheNpcInitData{
		Quote:    g.answerCache.GetAnswers().GuessTheNpc.Quote,
		Finished: user.GuessTheNPC.Game.Finished,
		Guesses:  guesses,
	}, nil
}
