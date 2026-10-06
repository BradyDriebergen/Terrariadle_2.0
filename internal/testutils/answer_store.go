package testutils

import (
	"context"
	"terrariadle/internal/domain"
)

type FakeAnswerStore struct {
	Answers domain.DailyAnswers
}

func (f *FakeAnswerStore) GetAnswers() domain.DailyAnswers {
	return f.Answers
}

func (f *FakeAnswerStore) UpsertAnswers(ctx context.Context, answer domain.DailyAnswers) error {
	f.Answers = answer
	return nil
}

func GenerateFakeAnswerStore() *FakeAnswerStore {
	categories := GenerateCategories()
	npc := GenerateNpc(0)

	return &FakeAnswerStore{
		Answers: domain.DailyAnswers{
			DailySlash: domain.WeaponAnswer{
				CurrentWeapon: GenerateWeapon(0),
				PrevWeapon:    GenerateWeapon(1),
			},
			Connections: GenerateConnectionsAnswer(categories),
			GuessTheNpc: domain.NpcAnswer{
				NpcID:       npc.ID,
				Npc:         npc.NPC,
				Quote:       npc.Quotes[0],
				Name:        npc.Names[0],
				NameOptions: []string{npc.Names[0], "fakeName1", "fakeName2", "fakeName3"},
			},
			Hangman: domain.HangmanAnswer{
				Enemy: GenerateEnemy(0),
			},
			TerraTrivia: domain.TerraTriviaAnswer{
				Questions: GenerateTriviaQuestions(),
			},
			NextResetTime: TestingTime(),
			ResetTime:     TestingTime(),
		},
	}
}
