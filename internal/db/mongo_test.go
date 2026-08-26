package db_test

import (
	"context"
	"fmt"
	"os"
	"terrariadle-backend/internal/db"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/log"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

var testMongo *db.MongoDB

type Puzzle struct {
	ID   string `bson:"_id"`
	Mode string `bson:"mode"`
}

// Initial setup of test container for integration tests
func TestMain(m *testing.M) {
	ctx := context.Background()

	container, err := mongodb.Run(ctx, "mongo:7",
		testcontainers.WithLogger(log.NewNoopLogger()),
	)
	if err != nil {
		fmt.Printf("failed to start mongodb container: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			fmt.Printf("failed to terminate container: %v\n", err)
		}
	}()

	connStr, err := container.ConnectionString(ctx)
	if err != nil {
		fmt.Printf("failed to get connection string: %v\n", err)
		os.Exit(1)
	}

	testMongo, err = db.Connect(connStr, "terrariadle_test")
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(ctx, testMongo)

	os.Exit(m.Run())
}

// Helper method that creates a new connection name for each test
func freshCollection(t *testing.T) string {
	t.Helper()
	name := "puzzles_" + t.Name()
	t.Cleanup(func() {
		_ = db.DeleteAll(context.Background(), testMongo, name)
	})
	return name
}

// Checks if Upserting a new item adds it to the database
func TestUpsertAndFindOne(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)

	puzzle := Puzzle{ID: "1", Mode: "Connections"}

	err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": puzzle.ID}, puzzle)
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	got, err := db.FindOne[Puzzle](ctx, testMongo, collection, db.Filter{"_id": puzzle.ID})
	if err != nil {
		t.Fatalf("findone failed: %v", err)
	}

	if got.Mode != puzzle.Mode {
		t.Errorf("got mode %q, want %q", got.Mode, puzzle.Mode)
	}
}

// Checks if not finding an item returns custom error message
func TestFindOne_NotFound(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)

	_, err := db.FindOne[Puzzle](ctx, testMongo, collection, db.Filter{"_id": "does-not-exist"})
	if err != db.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

// Checks if Upsert updates an item
func TestUpsert_UpdatesExisting(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)

	id := "1"

	// Initially adds the puzzle
	if err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": id}, Puzzle{ID: id, Mode: "hangman"}); err != nil {
		t.Fatalf("initial upsert failed: %v", err)
	}

	// Updates the puzzle
	if err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": id}, Puzzle{ID: id, Mode: "trivia"}); err != nil {
		t.Fatalf("update upsert failed: %v", err)
	}

	got, err := db.FindOne[Puzzle](ctx, testMongo, collection, db.Filter{"_id": id})
	if err != nil {
		t.Fatalf("findone failed: %v", err)
	}

	if got.Mode != "trivia" {
		t.Errorf("got mode %q, want %q", got.Mode, "trivia")
	}
}

// Checks if getting all items returns all items
func TestGetAll(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)

	puzzles := []Puzzle{
		{ID: "1", Mode: "daily-slash"},
		{ID: "2", Mode: "connections"},
		{ID: "3", Mode: "hangman"},
	}

	for _, p := range puzzles {
		if err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": p.ID}, p); err != nil {
			t.Fatalf("upsert failed: %v", err)
		}
	}

	got, err := db.GetAll[Puzzle](ctx, testMongo, collection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}

	if len(got) != len(puzzles) {
		t.Errorf("got %d puzzles, want %d", len(got), len(puzzles))
	}
}

// Checks if DeleteAll deletes all items from a collection.
func TestDeleteAll(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)

	puzzle := Puzzle{ID: "1", Mode: "connections"}

	if err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": puzzle.ID}, puzzle); err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	if err := db.DeleteAll(ctx, testMongo, collection); err != nil {
		t.Fatalf("deleteall failed: %v", err)
	}

	got, err := db.GetAll[Puzzle](ctx, testMongo, collection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}

	if len(got) != 0 {
		t.Errorf("expected 0 puzzles after delete, got %d", len(got))
	}
}
