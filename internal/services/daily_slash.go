package services

import (
	"context"
	"fmt"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
)

type DailySlashService interface {
}

type DailySlash struct {
	answerCache  store.AnswerStore
	catalogCache store.CatalogStore
	userCache    store.UserStore
}

func (g *DailySlash) InitializeGame(ctx context.Context, userId string) (DailySlashInitData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return DailySlashInitData{}, fmt.Errorf("Daily Slash: InitializeGame: %w", err)
	}

	if len(user.DailySlash.Game.Guesses) != len(user.DailySlash.Checks) {
		return DailySlashInitData{}, fmt.Errorf("Daily Slash: InitializeGame: guess/check length mismatch for user %s", userId)
	}

	guesses := make([]WeaponGuess, 0, len(user.DailySlash.Game.Guesses))

	for i, weaponID := range user.DailySlash.Game.Guesses {
		guessedWeapon, ok := g.catalogCache.GetWeapon(weaponID)
		if !ok {
			return DailySlashInitData{}, fmt.Errorf("Daily Slash: InitializeGame: unknown weapon ID %d", weaponID)
		}

		guesses = append(guesses, WeaponGuess{
			Weapon: toWeaponData(guessedWeapon),
			Checks: user.DailySlash.Checks[i],
		})
	}

	previousWeapon := toPreview(g.answerCache.GetAnswers().DailySlash.PrevWeapon)

	return DailySlashInitData{
		PreviousWeapon: previousWeapon,
		GuessedIDs:     user.DailySlash.Game.Guesses,
		Guesses:        guesses,
		HasWon:         user.DailySlash.Game.Finished,
	}, nil
}

func (g *DailySlash) SearchableWeapons(ctx context.Context) ([]domain.SearchWeaponResult, error) {
	return []domain.SearchWeaponResult{}, nil
}
