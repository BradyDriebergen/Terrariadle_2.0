package repo_test

import (
	"context"
	"terrariadle/internal/db"
	"terrariadle/internal/repo"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

// Tests getting answer data from the database
func TestGetAnswerData(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)
	answerRepo := repo.NewAnswerRepo(testMongo, collection, "")

	answers := generateAnswerData1()

	err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": 1}, answers)
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	got, err := answerRepo.GetAnswerData(ctx)
	if err != nil {
		t.Fatalf("getanswerdata failed: %v", err)
	}

	if diff := cmp.Diff(answers, got); diff != "" {
		t.Errorf("answer mismatch (-want +got):\n%s", diff)
	}
}

// Tests adding and updating answer data from the database
func TestUpsertAnswerData(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)
	answerRepo := repo.NewAnswerRepo(testMongo, collection, "")

	answers := []repo.AnswerData{
		generateAnswerData1(),
		generateAnswerData2(),
	}

	for i := range answers {
		err := answerRepo.UpsertAnswerData(ctx, &answers[i])
		if err != nil {
			t.Fatalf("upsertanswerdata failed: %v", err)
		}

		got, err := db.FindOne[repo.AnswerData](ctx, testMongo, collection, db.Filter{"_id": 1})
		if err != nil {
			t.Fatalf("findone failed: %v", err)
		}

		if diff := cmp.Diff(answers[i], *got); diff != "" {
			t.Errorf("answer mismatch (-want +got):\n%s", diff)
		}
	}

	docs, err := db.GetAll[repo.AnswerData](ctx, testMongo, collection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}

	if len(docs) != 1 {
		t.Errorf("expected 1 document after update, got %d", len(docs))
	}
}

// Tests getting player guess counts from the database
func TestGetGuessCounts(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)
	answerRepo := repo.NewAnswerRepo(testMongo, "", collection)

	guessCounts := generateGuessCounts1()

	err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": 1}, guessCounts)
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	got, err := answerRepo.GetGuessCounts(ctx)
	if err != nil {
		t.Fatalf("getguesscounts failed: %v", err)
	}

	if diff := cmp.Diff(guessCounts, got); diff != "" {
		t.Errorf("guess counts mismatch (-want +got):\n%s", diff)
	}
}

// Tests adding and updating player guess counts to the database
func TestUpsertGuessCounts(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)
	answerRepo := repo.NewAnswerRepo(testMongo, "", collection)

	guessCounts := []repo.PlayerGuessCounts{
		generateGuessCounts1(),
		generateGuessCounts2(),
	}

	for i := range guessCounts {
		err := answerRepo.UpsertGuessCounts(ctx, &guessCounts[i])
		if err != nil {
			t.Fatalf("upsertguesscounts failed: %v", err)
		}

		got, err := db.FindOne[repo.PlayerGuessCounts](ctx, testMongo, collection, db.Filter{"_id": 1})
		if err != nil {
			t.Fatalf("findone failed: %v", err)
		}

		if diff := cmp.Diff(guessCounts[i], *got); diff != "" {
			t.Errorf("guess counts mismatch (-want +got):\n%s", diff)
		}
	}

	all, err := db.GetAll[repo.PlayerGuessCounts](ctx, testMongo, collection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}

	if len(all) != 1 {
		t.Errorf("expected 1 document after update, got %d", len(all))
	}
}

// Returns unique answer data
func generateAnswerData1() repo.AnswerData {
	return repo.AnswerData{
		DailySlash: repo.WeaponData{
			CurrentWeaponID: 42,
			PrevWeaponID:    17,
		},
		Connections: repo.ConnectionData{
			CategoryIDs: []int{1, 2, 3, 4},
			Options: []repo.ConnectionOption{
				{Option: "Zenith", CategoryID: 1},
				{Option: "Terra Blade", CategoryID: 2},
			},
		},
		GuessTheNpc: repo.NpcData{
			NpcID:       22,
			Quote:       "Nurses heal wounds. I heal broken bones.",
			Name:        "Nurse",
			NameOptions: []string{"Nurse", "Guide", "Merchant"},
		},
		Hangman: repo.HangmanData{
			EnemyID: 13,
		},
		TerraTrivia: repo.TerraTriviaData{
			QuestionIDs: []int{101, 102, 103},
		},
		ResetTime:     time.Date(2026, 8, 26, 0, 0, 0, 0, time.UTC),
		NextResetTime: time.Date(2026, 8, 27, 0, 0, 0, 0, time.UTC),
	}
}

// Returns unique answer data
func generateAnswerData2() repo.AnswerData {
	return repo.AnswerData{
		DailySlash: repo.WeaponData{
			CurrentWeaponID: 30,
			PrevWeaponID:    10,
		},
		Connections: repo.ConnectionData{
			CategoryIDs: []int{4, 5, 6, 7},
			Options: []repo.ConnectionOption{
				{Option: "Golden Delight", CategoryID: 3},
				{Option: "Skeleton", CategoryID: 4},
			},
		},
		GuessTheNpc: repo.NpcData{
			NpcID:       22,
			Quote:       "Hunters shoot bows. I shoot guns.",
			Name:        "Arms Dealer",
			NameOptions: []string{"Pirate", "Angler", "Princess"},
		},
		Hangman: repo.HangmanData{
			EnemyID: 15,
		},
		TerraTrivia: repo.TerraTriviaData{
			QuestionIDs: []int{111, 112, 113},
		},
		ResetTime:     time.Date(2026, 9, 26, 0, 0, 0, 0, time.UTC),
		NextResetTime: time.Date(2026, 9, 27, 0, 0, 0, 0, time.UTC),
	}
}

// Returns unique guess counts
func generateGuessCounts1() repo.PlayerGuessCounts {
	return repo.PlayerGuessCounts{
		DailySlashCount:  1,
		ConnectionsCount: 1,
		GuessTheNpcCount: 1,
		HangmanCount:     1,
		TerraTriviaCount: 1,
	}
}

// Returns unique guess counts
func generateGuessCounts2() repo.PlayerGuessCounts {
	return repo.PlayerGuessCounts{
		DailySlashCount:  2,
		ConnectionsCount: 2,
		GuessTheNpcCount: 2,
		HangmanCount:     2,
		TerraTriviaCount: 2,
	}
}
