package repo_test

import (
	"context"
	"reflect"
	"terrariadle/internal/db"
	"terrariadle/internal/repo"
	"testing"
)

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

	if !reflect.DeepEqual(got, answers) {
		t.Errorf("got %+v, want %+v", got, answers)
	}
}

func TestUpsertAnswerData(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)
	answerRepo := repo.NewAnswerRepo(testMongo, collection, "")

	answers1 := generateAnswerData1()
	answers2 := generateAnswerData2()

	// Initial test for adding to a collection
	err := answerRepo.UpsertAnswerData(ctx, &answers1)
	if err != nil {
		t.Fatalf("upsertanswerdata failed: %v", err)
	}

	got, err := db.FindOne[repo.AnswerData](ctx, testMongo, collection, db.Filter{"_id": 1})
	if err != nil {
		t.Fatalf("findone failed: %v", err)
	}
	if !reflect.DeepEqual(*got, answers1) {
		t.Errorf("adding new answers: got %+v, want %+v", got, answers1)
	}

	// Test for updating existing value
	err = answerRepo.UpsertAnswerData(ctx, &answers2)
	if err != nil {
		t.Fatalf("upsertanswerdata failed: %v", err)
	}

	got, err = db.FindOne[repo.AnswerData](ctx, testMongo, collection, db.Filter{"_id": 1})
	if err != nil {
		t.Fatalf("findone failed: %v", err)
	}
	if !reflect.DeepEqual(*got, answers2) {
		t.Errorf("updating answers: got %+v, want %+v", got, answers2)
	}

	// Test for checking if there is only one record in the collection
	all, err := db.GetAll[repo.AnswerData](ctx, testMongo, collection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}
	if len(all) != 1 {
		t.Errorf("expected 1 document after update, got %d", len(all))
	}
}

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

	if !reflect.DeepEqual(got, guessCounts) {
		t.Errorf("got %+v, want %+v", got, guessCounts)
	}
}

func TestUpsertGuessCounts(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)
	answerRepo := repo.NewAnswerRepo(testMongo, "", collection)

	guessCounts1 := generateGuessCounts1()
	guessCounts2 := generateGuessCounts2()

	// Initial test for adding to a collection
	err := answerRepo.UpsertGuessCounts(ctx, &guessCounts1)
	if err != nil {
		t.Fatalf("upsertguesscounts failed: %v", err)
	}

	got, err := db.FindOne[repo.PlayerGuessCounts](ctx, testMongo, collection, db.Filter{"_id": 1})
	if err != nil {
		t.Fatalf("findone failed: %v", err)
	}
	if !reflect.DeepEqual(*got, guessCounts1) {
		t.Errorf("adding new guess counts: got %+v, want %+v", got, guessCounts1)
	}

	// Test for updating existing value
	err = answerRepo.UpsertGuessCounts(ctx, &guessCounts2)
	if err != nil {
		t.Fatalf("upsertguesscounts failed: %v", err)
	}

	got, err = db.FindOne[repo.PlayerGuessCounts](ctx, testMongo, collection, db.Filter{"_id": 1})
	if err != nil {
		t.Fatalf("findone failed: %v", err)
	}
	if !reflect.DeepEqual(*got, guessCounts2) {
		t.Errorf("updating guess counts: got %+v, want %+v", got, guessCounts2)
	}

	// Test for checking if there is only one record in the collection
	all, err := db.GetAll[repo.PlayerGuessCounts](ctx, testMongo, collection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}
	if len(all) != 1 {
		t.Errorf("expected 1 document after update, got %d", len(all))
	}
}
