package services

import (
	"context"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
)

type CommonService interface {
	GetGuessCount(game domain.GameMode) (int, error)
	GetUserFinishedGames(ctx context.Context, userId string) UserGameStatuses
}

type Common struct {
	userCache       store.UserStore
	guessCountCache store.GuessCountsStore
}

func NewSseStream(
	guessCountCache store.GuessCountsStore,
	userCache store.UserStore,
) *Common {
	return &Common{
		userCache:       userCache,
		guessCountCache: guessCountCache,
	}
}

func (s *Common) GetGuessCount(game domain.GameMode) (int, error) {
	switch game {
	case "daily_slash":
		return s.guessCountCache.GetGuessCounts().DailySlashCount, nil
	case "connections":
		return s.guessCountCache.GetGuessCounts().ConnectionsCount, nil
	case "guess_the_npc":
		return s.guessCountCache.GetGuessCounts().GuessTheNpcCount, nil
	case "hangman":
		return s.guessCountCache.GetGuessCounts().HangmanCount, nil
	case "terratrivia":
		return s.guessCountCache.GetGuessCounts().TerraTriviaCount, nil
	default:
		return 0, domain.NotFound("The requested guess count doesn't exist", nil)
	}
}

func (s *Common) GetUserFinishedGames(ctx context.Context, userId string) UserGameStatuses {
	user, err := s.userCache.GetUser(ctx, userId)
	if err != nil {
		return UserGameStatuses{}
	}

	return UserGameStatuses{
		DaliySlash:  user.DailySlash.Game.Finished,
		Connections: user.Connections.Game.Finished,
		GuessTheNpc: user.GuessTheNPC.Game.Finished,
		Hangman:     user.Hangman.Game.Finished,
		TerraTrivia: user.TerraTrivia.Game.Finished,
	}
}
