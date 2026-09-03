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

	answers := newAnswerData(1)
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
		newAnswerData(1),
		newAnswerData(2),
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

	guessCounts := newGuessCounts(1)
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
		newGuessCounts(1),
		newGuessCounts(2),
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
func newAnswerData(id int) answerData {
	return answerData{
		DailySlash: weaponData{
			CurrentWeaponID: id,
		},
		Connections: connectionData{
			CategoryIDs: []int{id},
		},
		GuessTheNpc: npcData{
			NpcID: id,
		},
		Hangman: hangmanData{
			EnemyID: id,
		},
		TerraTrivia: terraTriviaData{
			QuestionIDs: []int{id},
		},
		ResetTime:     testutils.TestingTime(),
		NextResetTime: testutils.TestingTime(),
	}
}

// Returns unique guess counts
func newGuessCounts(num int) guessCounts {
	return guessCounts{
		DailySlashCount:  num,
		ConnectionsCount: num,
		GuessTheNpcCount: num,
		HangmanCount:     num,
		TerraTriviaCount: num,
	}
}
