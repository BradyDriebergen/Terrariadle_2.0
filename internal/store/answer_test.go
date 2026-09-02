package store

import (
	"context"
	"terrariadle/internal/domain"
	"terrariadle/internal/testutils"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Tests if get answers returns the answers in a domain format
func TestGetAnswers(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeStore := &testutils.FakeCatalogStore{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	fakeRepo := &testutils.FakeAnswerRepo{
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
			ResetTime:     testutils.TestingTime(),
			NextResetTime: testutils.TestingTime(),
		},
	}

	store, err := NewAnswerStore(ctx, fakeRepo, fakeStore)
	if err != nil {
		t.Fatalf("newanswerstore failed: %v", err)
	}

	got := store.GetAnswers()

	want := domain.DailyAnswers{
		DailySlash: domain.WeaponAnswer{
			CurrentWeapon: testutils.FindFromList(t, weapons, 1, func(w domain.Weapon) int { return w.ID }),
		},
	}

	if diff := cmp.Diff(want.DailySlash.CurrentWeapon, got.DailySlash.CurrentWeapon); diff != "" {
		t.Errorf("answer store mismatch (-want +got):\n%s", diff)
	}
}

// Tests if updating the answers creates and updates the answsers
func TestUpsertAnswers(t *testing.T) {
	ctx := context.Background()

	weapons := testutils.GenerateWeapons()
	npcs := testutils.GenerateNpcs()
	enemies := testutils.GenerateEnemies()
	categories := testutils.GenerateCategories()
	triviaQuestions := testutils.GenerateTriviaQuestions()

	fakeStore := &testutils.FakeCatalogStore{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}

	fakeRepo := &testutils.FakeAnswerRepo{
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
			ResetTime:     testutils.TestingTime(),
			NextResetTime: testutils.TestingTime(),
		},
	}

	store, err := NewAnswerStore(ctx, fakeRepo, fakeStore)
	if err != nil {
		t.Fatalf("newanswerstore failed: %v", err)
	}

	err = store.UpsertAnswers(ctx, domain.DailyAnswers{
		DailySlash: domain.WeaponAnswer{
			CurrentWeapon: testutils.GenerateWeapon(0),
		},
	})
	if err != nil {
		t.Fatalf("upsertanswers failed: %v", err)
	}

	got := store.answerCache

	want := domain.DailyAnswers{
		DailySlash: domain.WeaponAnswer{
			CurrentWeapon: testutils.FindFromList(t, weapons, 0, func(w domain.Weapon) int { return w.ID }),
		},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("answer store mismatch (-want +got):\n%s", diff)
	}
}
