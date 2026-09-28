package testutils

import (
	"context"
	"terrariadle/internal/domain"
	"time"
)

type FakeUserRepo struct {
	User domain.User
}

func (f *FakeUserRepo) GetUser(ctx context.Context, userId string) (domain.User, error) {
	if f.User.UserID == "" {
		return domain.User{}, domain.MongoErrNotFound
	}
	return f.User, nil
}

func (f *FakeUserRepo) UpsertUserData(ctx context.Context, user domain.User) error {
	f.User = user
	return nil
}

func (f *FakeUserRepo) DropAllUserData(ctx context.Context) error {
	f.User = domain.User{}
	return nil
}

func GenerateFakeUserRepo() *FakeUserRepo {
	return &FakeUserRepo{
		User: domain.User{},
	}
}

func GenerateUser(id string) domain.User {
	emptyGame := domain.Game{
		Guesses:  []int{},
		Finished: false,
		Position: 0,
	}

	return domain.User{
		UserID: id,
		DailySlash: domain.DailySlashGame{
			Game:   emptyGame,
			Checks: []domain.WeaponChecks{},
		},
		Connections: domain.ConnectionGame{
			Game:     emptyGame,
			Attempts: 4,
		},
		GuessTheNPC: domain.GuessTheNpcGame{
			Game:        emptyGame,
			GuessedName: "",
		},
		Hangman: domain.HangmanGame{
			Game:     emptyGame,
			Attempts: 6,
		},
		TerraTrivia: domain.TerraTriviaGame{
			Game: emptyGame,
		},
		LastSeen: time.Now(),
		Dirty:    true,
	}
}
