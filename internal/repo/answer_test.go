package repo

import (
	"context"
	"terrariadle/internal/db"
	"terrariadle/internal/testutils"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Tests getting answer data from the database
func TestGetAnswerData(t *testing.T) {
	ctx := context.Background()
	answerRepo := mockAnswerRepo(t)

	answers := generateAnswerDataA()
	want := toAnswerRef(answers)

	err := db.Upsert(ctx, testMongo, answerRepo.answerCollection, db.Filter{"_id": 1}, answers)
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	got, err := answerRepo.GetAnswerData(ctx)
	if err != nil {
		t.Fatalf("getanswerdata failed: %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("answer mismatch (-want +got):\n%s", diff)
	}
}

// Tests adding and updating answer data from the database
func TestUpsertAnswerData(t *testing.T) {
	ctx := context.Background()
	answerRepo := mockAnswerRepo(t)

	answers := []answerData{
		generateAnswerDataA(),
		generateAnswerDataB(),
	}

	for i := range answers {
		ref := toAnswerRef(answers[i])
		err := answerRepo.UpsertAnswerData(ctx, &ref)
		if err != nil {
			t.Fatalf("upsertanswerdata failed: %v", err)
		}

		got, err := db.FindOne[answerData](ctx, testMongo, answerRepo.answerCollection, db.Filter{"_id": 1})
		if err != nil {
			t.Fatalf("findone failed: %v", err)
		}

		if diff := cmp.Diff(answers[i], *got); diff != "" {
			t.Errorf("answer mismatch (-want +got):\n%s", diff)
		}
	}

	docs, err := db.GetAll[answerData](ctx, testMongo, answerRepo.answerCollection)
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
	answerRepo := mockAnswerRepo(t)

	guessCounts := generateGuessCountsA()
	want := toPlayerGuessCounts(guessCounts)

	err := db.Upsert(ctx, testMongo, answerRepo.guessCountCollection, db.Filter{"_id": 1}, guessCounts)
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	got, err := answerRepo.GetGuessCounts(ctx)
	if err != nil {
		t.Fatalf("getguesscounts failed: %v", err)
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("guess counts mismatch (-want +got):\n%s", diff)
	}
}

// Tests adding and updating player guess counts to the database
func TestUpsertGuessCounts(t *testing.T) {
	ctx := context.Background()
	answerRepo := mockAnswerRepo(t)

	counts := []guessCounts{
		generateGuessCountsA(),
		generateGuessCountsB(),
	}

	for i := range counts {
		ref := toPlayerGuessCounts(counts[i])
		err := answerRepo.UpsertGuessCounts(ctx, &ref)
		if err != nil {
			t.Fatalf("upsertguesscounts failed: %v", err)
		}

		got, err := db.FindOne[guessCounts](ctx, testMongo, answerRepo.guessCountCollection, db.Filter{"_id": 1})
		if err != nil {
			t.Fatalf("findone failed: %v", err)
		}

		if diff := cmp.Diff(counts[i], *got); diff != "" {
			t.Errorf("guess counts mismatch (-want +got):\n%s", diff)
		}
	}

	all, err := db.GetAll[guessCounts](ctx, testMongo, answerRepo.guessCountCollection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}

	if len(all) != 1 {
		t.Errorf("expected 1 document after update, got %d", len(all))
	}
}

// Helper method creating collections for answer_repo
func mockAnswerRepo(t *testing.T) *MongoAnswerRepo {
	t.Helper()

	answerCollection := "answer_" + t.Name()
	guessCountCollection := "guess_count_" + t.Name()

	t.Cleanup(func() {
		_ = db.DeleteAll(context.Background(), testMongo, answerCollection)
		_ = db.DeleteAll(context.Background(), testMongo, guessCountCollection)
	})

	return NewAnswerRepo(
		testMongo,
		answerCollection,
		guessCountCollection,
	)
}

// Returns unique answer data
func generateAnswerDataA() answerData {
	return answerData{
		DailySlash: weaponData{
			CurrentWeaponID: 42,
			PrevWeaponID:    17,
		},
		Connections: connectionData{
			CategoryIDs: []int{1, 2, 3, 4},
			Options: []connectionOption{
				{Option: "Zenith", CategoryID: 1},
				{Option: "Terra Blade", CategoryID: 2},
			},
		},
		GuessTheNpc: npcData{
			NpcID:       22,
			Quote:       "Nurses heal wounds. I heal broken bones.",
			Name:        "Nurse",
			NameOptions: []string{"Nurse", "Guide", "Merchant"},
		},
		Hangman: hangmanData{
			EnemyID: 13,
		},
		TerraTrivia: terraTriviaData{
			QuestionIDs: []int{101, 102, 103},
		},
		ResetTime:     testutils.TestingTime(),
		NextResetTime: testutils.TestingTime(),
	}
}

// Returns unique answer data
func generateAnswerDataB() answerData {
	return answerData{
		DailySlash: weaponData{
			CurrentWeaponID: 30,
			PrevWeaponID:    10,
		},
		Connections: connectionData{
			CategoryIDs: []int{4, 5, 6, 7},
			Options: []connectionOption{
				{Option: "Golden Delight", CategoryID: 3},
				{Option: "Skeleton", CategoryID: 4},
			},
		},
		GuessTheNpc: npcData{
			NpcID:       22,
			Quote:       "Hunters shoot bows. I shoot guns.",
			Name:        "Arms Dealer",
			NameOptions: []string{"Pirate", "Angler", "Princess"},
		},
		Hangman: hangmanData{
			EnemyID: 15,
		},
		TerraTrivia: terraTriviaData{
			QuestionIDs: []int{111, 112, 113},
		},
		ResetTime:     testutils.TestingTime(),
		NextResetTime: testutils.TestingTime(),
	}
}

// Returns unique guess counts
func generateGuessCountsA() guessCounts {
	return guessCounts{
		DailySlashCount:  1,
		ConnectionsCount: 1,
		GuessTheNpcCount: 1,
		HangmanCount:     1,
		TerraTriviaCount: 1,
	}
}

// Returns unique guess counts
func generateGuessCountsB() guessCounts {
	return guessCounts{
		DailySlashCount:  2,
		ConnectionsCount: 2,
		GuessTheNpcCount: 2,
		HangmanCount:     2,
		TerraTriviaCount: 2,
	}
}
