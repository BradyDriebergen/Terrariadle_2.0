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
	answerCache     store.AnswerStore
	guessCountCache store.GuessCountsStore
	catalogCache    store.CatalogStore
	userCache       store.UserStore
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
		Finished:       user.DailySlash.Game.Finished,
	}, nil
}

func (g *DailySlash) GetSearchableWeapons() []domain.SearchWeaponResult {
	return g.catalogCache.GetSearchableWeapons()
}

func (g *DailySlash) GetHint(hintNum int) (string, error) {
	weapon := g.answerCache.GetAnswers().DailySlash.CurrentWeapon

	switch hintNum {
	case 1:
		return weapon.ModeObtained, nil
	case 2:
		return weapon.WeaponType, nil
	case 3:
		return weapon.Info.ImagePath, nil
	default:
		return "", fmt.Errorf("requested hint does not exist")
	}
}

func (g *DailySlash) CheckGuess(ctx context.Context, userId string, weaponId int) (DailySlashCheckData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return DailySlashCheckData{}, fmt.Errorf("Daily Slash: CheckGuess: %w", err)
	}

	weaponAnswer := g.answerCache.GetAnswers().DailySlash.CurrentWeapon
	guessedWeapon, ok := g.catalogCache.GetWeapon(weaponId)
	if !ok {
		return DailySlashCheckData{}, fmt.Errorf("Daily Slash: CheckGuess: guessed weapon id does not exist %w", err)
	}
	checks := generateWeaponChecks(guessedWeapon, weaponAnswer)

	user.DailySlash.Game.Guesses = append(user.DailySlash.Game.Guesses, guessedWeapon.ID)
	user.DailySlash.Checks = append(user.DailySlash.Checks, checks)

	correct := weaponAnswer.ID == guessedWeapon.ID

	if correct {
		user.DailySlash.Game.Finished = true
		user.DailySlash.Game.Position, err = g.guessCountCache.IncrementDailySlashCount(ctx)
		if err != nil {
			return DailySlashCheckData{}, fmt.Errorf("Daily Slash: CheckGuess: error updating user position %w", err)
		}
	}

	err = g.userCache.UpsertUser(ctx, user)
	if err != nil {
		return DailySlashCheckData{}, fmt.Errorf("Daily Slash: CheckGuess: error updating user %w", err)
	}

	return DailySlashCheckData{
		Finished: correct,
		GuessResult: WeaponGuess{
			Weapon: toWeaponData(guessedWeapon),
			Checks: checks,
		},
	}, nil
}
