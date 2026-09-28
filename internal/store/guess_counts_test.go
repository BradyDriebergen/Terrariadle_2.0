package store

import (
	"context"
	"terrariadle/internal/domain"
	"terrariadle/internal/testutils"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Tests if get guess counts returns the guess counts in a domain format
func TestGetGuessCounts(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeAnswerRepo()

	store, err := NewGuessCountStore(ctx, fakeRepo, &domain.Broker{})
	if err != nil {
		t.Fatalf("newanswerstore failed: %v", err)
	}

	got := store.GetGuessCounts()

	want := fakeRepo.GuessCounts

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("guess count store mismatch (-want +got):\n%s", diff)
	}
}

// Tests if reseting guess counts sets everything to 0
func TestResetGuessCounts(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeAnswerRepo()

	store, err := NewGuessCountStore(ctx, fakeRepo, &domain.Broker{})
	if err != nil {
		t.Fatalf("newanswerstore failed: %v", err)
	}

	err = store.ResetGuessCounts(ctx)
	if err != nil {
		t.Fatalf("resetguesscounts failed: %v", err)
	}

	got := store.guessCountsCache

	// Reseted guess counts
	want := domain.PlayerGuessCounts{
		DailySlashCount:  0,
		ConnectionsCount: 0,
		GuessTheNpcCount: 0,
		HangmanCount:     0,
		TerraTriviaCount: 0,
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("guess count store mismatch (-want +got):\n%s", diff)
	}
}

// Table driven test that checks if incrementing every game mode's
// increment guess count increases the guess count by 1
func TestIncrementGuessCounts(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeAnswerRepo()

	store, err := NewGuessCountStore(ctx, fakeRepo, &domain.Broker{})
	if err != nil {
		t.Fatalf("newanswerstore failed: %v", err)
	}

	tests := []struct {
		name      string
		count     func() int
		increment func(ctx context.Context) (int, error)
	}{
		{
			name: "DailySlash",
			count: func() int {
				return store.GetGuessCounts().DailySlashCount
			},
			increment: func(ctx context.Context) (int, error) {
				return store.IncrementDailySlashCount(ctx)
			},
		},
		{
			name: "Connections",
			count: func() int {
				return store.GetGuessCounts().ConnectionsCount
			},
			increment: func(ctx context.Context) (int, error) {
				return store.IncrementConnectionsCount(ctx)
			},
		},
		{
			name: "GuessTheNpc",
			count: func() int {
				return store.GetGuessCounts().GuessTheNpcCount
			},
			increment: func(ctx context.Context) (int, error) {
				return store.IncrementGuessTheNpcCount(ctx)
			},
		},
		{
			name: "Hangman",
			count: func() int {
				return store.GetGuessCounts().HangmanCount
			},
			increment: func(ctx context.Context) (int, error) {
				return store.IncrementHangmanCount(ctx)
			},
		},
		{
			name: "TerraTrivia",
			count: func() int {
				return store.GetGuessCounts().TerraTriviaCount
			},
			increment: func(ctx context.Context) (int, error) {
				return store.IncrementTerraTriviaCount(ctx)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.count() + 1

			got, err := tt.increment(ctx)
			if err != nil {
				t.Fatalf("increment failed: %v", err)
			}
			if got != want {
				t.Errorf("Expected returned value %d, got: %d", want, got)
			}
			if got != tt.count() {
				t.Errorf("Expected actual value %d, got: %d", want, got)
			}
		})
	}
}
