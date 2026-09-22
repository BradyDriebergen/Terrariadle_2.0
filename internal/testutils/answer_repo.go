package testutils

import (
	"context"
	"terrariadle/internal/domain"
)

// struct and methods for mocking answer_repo
type FakeAnswerRepo struct {
	AnswerData      domain.AnswerRefs
	GuessCounts     domain.PlayerGuessCounts
	GetAnswerErr    error
	UpsertAnswerErr error
	GetGuessErr     error
	UpsertGuessErr  error
}

func (f *FakeAnswerRepo) GetAnswerData(ctx context.Context) (domain.AnswerRefs, error) {
	if f.GetAnswerErr != nil {
		return domain.AnswerRefs{}, f.GetAnswerErr
	}
	return f.AnswerData, nil
}

func (f *FakeAnswerRepo) UpsertAnswerData(ctx context.Context, answerData *domain.AnswerRefs) error {
	if f.UpsertAnswerErr != nil {
		return f.UpsertAnswerErr
	}
	f.AnswerData = *answerData
	return nil
}

func (f *FakeAnswerRepo) GetGuessCounts(ctx context.Context) (domain.PlayerGuessCounts, error) {
	if f.GetGuessErr != nil {
		return domain.PlayerGuessCounts{}, f.GetGuessErr
	}
	return f.GuessCounts, nil
}

func (f *FakeAnswerRepo) UpsertGuessCounts(ctx context.Context, guessCounts *domain.PlayerGuessCounts) error {
	if f.UpsertGuessErr != nil {
		return f.UpsertGuessErr
	}
	f.GuessCounts = *guessCounts
	return nil
}

func GenerateFakeAnswerRepo() FakeAnswerRepo {
	return FakeAnswerRepo{
		AnswerData: domain.AnswerRefs{
			DailySlash: domain.WeaponRef{
				CurrentWeaponID: 1,
				PrevWeaponID:    2,
			},
			Connections: domain.ConnectionRef{
				CategoryIDs: []int{4, 5, 6, 7},
				Options: []domain.ConnectionOption{
					{Option: "Golden Delight", CategoryID: 3},
					{Option: "Skeleton", CategoryID: 4},
				},
			},
			GuessTheNpc: domain.NpcRef{
				NpcID:       1,
				Quote:       "Hunters shoot bows. I shoot guns.",
				Name:        "Arms Dealer",
				NameOptions: []string{"Pirate", "Angler", "Princess"},
			},
			Hangman: domain.HangmanRef{
				EnemyID: 1,
			},
			TerraTrivia: domain.TerraTriviaRef{
				QuestionIDs: []int{0, 1, 2},
			},
			ResetTime:     TestingTime(),
			NextResetTime: TestingTime(),
		},
		GuessCounts: domain.PlayerGuessCounts{
			DailySlashCount:  1,
			ConnectionsCount: 1,
			GuessTheNpcCount: 1,
			HangmanCount:     1,
			TerraTriviaCount: 1,
		},
	}
}
