package services

import (
	"context"
	"slices"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/store"
)

type DailySlashService interface {
	InitializeGame(ctx context.Context, userId string) (DailySlashInitData, error)
	GetSearchableWeapons() []SearchWeaponData
	GetHint(hintNum int) (string, error)
	CheckGuess(ctx context.Context, userId string, weaponId int) (DailySlashCheckData, error)
	GetWinningData(ctx context.Context, userId string) (DailySlashWinningData, error)
}

type DailySlash struct {
	answerCache     store.AnswerStore
	guessCountCache store.GuessCountsStore
	catalogCache    store.CatalogStore
	userCache       store.UserStore
}

func NewDailySlashGame(
	answerCache store.AnswerStore,
	guessCountCache store.GuessCountsStore,
	catalogCache store.CatalogStore,
	userCache store.UserStore,
) *DailySlash {

	return &DailySlash{
		answerCache:     answerCache,
		guessCountCache: guessCountCache,
		catalogCache:    catalogCache,
		userCache:       userCache,
	}
}

func (g *DailySlash) InitializeGame(ctx context.Context, userId string) (DailySlashInitData, error) {
	user, err := g.userCache.GetOrCreateUser(ctx, userId)
	if err != nil {
		return DailySlashInitData{}, domain.UserNotFound("Error creating user", err)
	}

	checks := make(map[int]domain.WeaponChecks, len(user.DailySlash.Checks))
	for _, c := range user.DailySlash.Checks {
		checks[c.WeaponID] = c
	}

	guesses := make([]WeaponGuess, 0, len(user.DailySlash.Game.Guesses))

	for _, weaponID := range user.DailySlash.Game.Guesses {
		guessedWeapon, ok := g.catalogCache.GetWeapon(weaponID)
		if !ok {
			return DailySlashInitData{}, domain.NotFound("The guessed weapon doesn't exist", nil)
		}

		check, ok := checks[guessedWeapon.ID]
		if !ok {
			return DailySlashInitData{}, domain.Internal("Weapon does not have check associated with it", nil)
		}

		guesses = append(guesses, WeaponGuess{
			Weapon: toWeaponData(guessedWeapon),
			Checks: check,
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

func (g *DailySlash) GetSearchableWeapons() []SearchWeaponData {
	return toSearchableWeapons(g.catalogCache.GetSearchableWeapons())
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
		return "", domain.NotFound("The requested hint does not exist", nil)
	}
}

func (g *DailySlash) CheckGuess(ctx context.Context, userId string, weaponId int) (DailySlashCheckData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return DailySlashCheckData{}, domain.UserNotFound("User not found", err)
	}

	if slices.Contains(user.DailySlash.Game.Guesses, weaponId) {
		return DailySlashCheckData{}, domain.Conflict("User previously guessed this weapon", nil)
	}

	weaponAnswer := g.answerCache.GetAnswers().DailySlash.CurrentWeapon
	guessedWeapon, ok := g.catalogCache.GetWeapon(weaponId)
	if !ok {
		return DailySlashCheckData{}, domain.NotFound("The requested weapon doesn't exist", nil)
	}

	checks := generateWeaponChecks(guessedWeapon, weaponAnswer)
	correct := weaponAnswer.ID == guessedWeapon.ID

	if correct {
		position, err := g.guessCountCache.IncrementDailySlashCount(ctx)
		if err != nil {
			return DailySlashCheckData{}, domain.Internal("An error occurred updating user's position", err)
		}

		user.DailySlash.Game.Finished = true
		user.DailySlash.Game.Position = position
	}

	user.DailySlash.Game.Guesses = append(user.DailySlash.Game.Guesses, guessedWeapon.ID)
	user.DailySlash.Checks = append(user.DailySlash.Checks, checks)

	err = g.userCache.UpsertUser(ctx, user)
	if err != nil {
		return DailySlashCheckData{}, domain.Internal("An error occurred updating user's guess", err)
	}

	return DailySlashCheckData{
		Finished: correct,
		GuessResult: WeaponGuess{
			Weapon: toWeaponData(guessedWeapon),
			Checks: checks,
		},
	}, nil
}

func (g *DailySlash) GetWinningData(ctx context.Context, userId string) (DailySlashWinningData, error) {
	user, err := g.userCache.GetUser(ctx, userId)
	if err != nil {
		return DailySlashWinningData{}, domain.UserNotFound("User not found", err)
	}

	if !user.DailySlash.Game.Finished {
		return DailySlashWinningData{}, domain.Conflict("User isn't finished guessing", err)
	}

	return DailySlashWinningData{
		Position: user.DailySlash.Game.Position,
	}, nil
}
