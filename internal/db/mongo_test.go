package db_test

import (
	"context"
	"fmt"
	"os"
	"terrariadle/internal/db"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	os.Exit(runTests(m))
}

func runTests(m *testing.M) int {
	ctx := context.Background()

	container, err := mongodb.Run(ctx, "mongo:7", testcontainers.WithLogger(log.NewNoopLogger()))
	if err != nil {
		fmt.Printf("failed to start mongodb container: %v\n", err)
		return 1
	}
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			fmt.Printf("failed to terminate container: %v\n", err)
		}
	}()

	connStr, err := container.ConnectionString(ctx)
	if err != nil {
		fmt.Printf("failed to get connection string: %v\n", err)
		return 1
	}

	testMongo, err = db.Connect(connStr, "terrariadle_test")
	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
		return 1
	}
	defer db.Close(ctx, testMongo)

	return m.Run()
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

// Checks if FindOne returns one item from a collection where the
// filter modifier is true.
func TestFindOne(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name    string
		seed    *Puzzle
		id      string
		wantErr error
	}{
		{name: "existing puzzle", seed: &Puzzle{ID: "1", Mode: "Connections"}, id: "1"},
		{name: "nonexistent puzzle", seed: nil, id: "does-not-exist", wantErr: db.ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collection := freshCollection(t)

			if tt.seed != nil {
				if err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": tt.seed.ID}, *tt.seed); err != nil {
					t.Fatalf("upsert failed: %v", err)
				}
			}

			got, err := db.FindOne[Puzzle](ctx, testMongo, collection, db.Filter{"_id": tt.id})
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Fatalf("got err %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("findone failed: %v", err)
			}

			if got.Mode != tt.seed.Mode {
				t.Errorf("got mode %q, want %q", got.Mode, tt.seed.Mode)
			}
		})
	}
}

// Checks if Upsert adds an item if it doesn't exist, and updates existing
// items based on the filter.
func TestUpsert(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)

	id := "1"
	puzzles := []Puzzle{
		{ID: id, Mode: "hangman"},
		{ID: id, Mode: "trivia"},
		{ID: id, Mode: "dailyslash"},
	}

	for i := range puzzles {
		err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": puzzles[i].ID}, puzzles[i])
		if err != nil {
			t.Fatalf("update upsert failed: %v", err)
		}

		got, err := db.FindOne[Puzzle](ctx, testMongo, collection, db.Filter{"_id": puzzles[i].ID})
		if err != nil {
			t.Fatalf("findone failed: %v", err)
		}

		if got.Mode != puzzles[i].Mode {
			t.Errorf("got mode %q, want %q", got.Mode, puzzles[i].Mode)
		}
	}

	got, err := db.GetAll[Puzzle](ctx, testMongo, collection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}

	if len(got) != 1 {
		t.Errorf("expected 1 puzzle after updating, got %d", len(got))
	}
}

// Checks if InsertMany adds multiple documents to a colletion, and
// GetAll returns everything from one collection.
func TestGetAllAndInsertMany(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)

	puzzles := []Puzzle{
		{ID: "1", Mode: "dailyslash"},
		{ID: "2", Mode: "connections"},
		{ID: "3", Mode: "hangman"},
	}

	err := db.InsertMany(ctx, testMongo, collection, puzzles)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	got, err := db.GetAll[Puzzle](ctx, testMongo, collection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}

	if len(got) != len(puzzles) {
		t.Errorf("got %d puzzles, want %d", len(got), len(puzzles))
	}

	if diff := cmp.Diff(puzzles, got); diff != "" {
		t.Errorf("Puzzle mismatch (-want +got):\n%s", diff)
	}
}

// Checks if DeleteAll deletes all items from a collection.
func TestDeleteAll(t *testing.T) {
	ctx := context.Background()
	collection := freshCollection(t)

	puzzle := Puzzle{ID: "1", Mode: "connections"}

	err := db.Upsert(ctx, testMongo, collection, db.Filter{"_id": puzzle.ID}, puzzle)
	if err != nil {
		t.Fatalf("upsert failed: %v", err)
	}

	err = db.DeleteAll(ctx, testMongo, collection)
	if err != nil {
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
