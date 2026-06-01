package services

import (
	"context"
	"slices"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
)

type GuessTheNpcService interface {
	InitializeGame(ctx context.Context, userId string) (GuessTheNpcInitData, error)
	GetSearchableNpcs() []domain.SearchNpcResult
	CheckGuess(ctx context.Context, userId string, npcId int) (GuessTheNpcCheckData, error)
	GetWinningData(ctx context.Context, userId string) (GuessTheNpcWinningData, error)
	CheckName(ctx context.Context, userId string, name string) (GuessTheNpcMiniGameData, error)
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

func (g *GuessTheNpc) GetSearchableNpcs() []domain.SearchNpcResult {
	return g.catalogCache.GetSearchableNpcs()
}

func (g *GuessTheNpc) CheckGuess(ctx context.Context, userId string, npcId int) (GuessTheNpcCheckData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return GuessTheNpcCheckData{}, domain.UserNotFound("User not found", err)
	}

	if slices.Contains(user.GuessTheNPC.Game.Guesses, npcId) {
		return GuessTheNpcCheckData{}, domain.Conflict("User previously guessed this npc", nil)
	}

	npcAnswer := g.answerCache.GetAnswers().GuessTheNpc
	npcGuess, ok := g.catalogCache.GetNpc(npcId)
	if !ok {
		return GuessTheNpcCheckData{}, domain.NotFound("The requested npc doesn't exist", nil)
	}

	isCorrect := npcAnswer.NpcID == npcId
	if isCorrect {
		position, err := g.guessCountCache.IncrementGuessTheNpcCount(ctx)
		if err != nil {
			return GuessTheNpcCheckData{}, domain.Internal("An error occurred updating user's position", err)
		}

		user.GuessTheNPC.Game.Finished = true
		user.GuessTheNPC.Game.Position = position
	}

	user.GuessTheNPC.Game.Guesses = append(user.GuessTheNPC.Game.Guesses, npcId)

	err = g.userCache.UpsertUser(ctx, user)
	if err != nil {
		return GuessTheNpcCheckData{}, domain.Internal("An error occurred updating user's guess", err)
	}

	return GuessTheNpcCheckData{
		Finished: isCorrect,
		Guess: domain.NpcInfo{
			Name: npcGuess.NPC,
			Path: npcGuess.NPCPath,
		},
	}, nil
}

func (g *GuessTheNpc) GetWinningData(ctx context.Context, userId string) (GuessTheNpcWinningData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return GuessTheNpcWinningData{}, domain.UserNotFound("User not found", err)
	}

	if !user.GuessTheNPC.Game.Finished {
		return GuessTheNpcWinningData{}, domain.Conflict("User isn't finished guessing", nil)
	}

	npcAnswer := g.answerCache.GetAnswers().GuessTheNpc

	correctName := ""
	if user.GuessTheNPC.GuessedName != "" {
		correctName = npcAnswer.Name
	}

	return GuessTheNpcWinningData{
		Position:    user.GuessTheNPC.Game.Position,
		PlayerCount: g.guessCountCache.GetGuessCounts().GuessTheNpcCount,
		Names:       npcAnswer.NameOptions,
		GuessedName: user.GuessTheNPC.GuessedName,
		CorrectName: correctName,
	}, nil
}

func (g *GuessTheNpc) CheckName(ctx context.Context, userId string, name string) (GuessTheNpcMiniGameData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return GuessTheNpcMiniGameData{}, domain.UserNotFound("User not found", err)
	}

	if !user.GuessTheNPC.Game.Finished {
		return GuessTheNpcMiniGameData{}, domain.Conflict("User isn't finished guessing", err)
	}

	if user.GuessTheNPC.GuessedName != "" {
		return GuessTheNpcMiniGameData{}, domain.Conflict("User already guessed a name", nil)
	}

	npcAnswer := g.answerCache.GetAnswers().GuessTheNpc

	if !slices.Contains(npcAnswer.NameOptions, name) {
		return GuessTheNpcMiniGameData{}, domain.Conflict("Name guessed isn't an option", nil)
	}

	user.GuessTheNPC.GuessedName = name

	err = g.userCache.UpsertUser(ctx, user)
	if err != nil {
		return GuessTheNpcMiniGameData{}, domain.Internal("An error occurred updating user's guess", err)
	}

	return GuessTheNpcMiniGameData{
		GuessedName: name,
		CorrectName: npcAnswer.Name,
	}, nil
}
