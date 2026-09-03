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

	fakeRepo := &testutils.FakeAnswerRepo{
		GuessCounts: domain.PlayerGuessCounts{
			DailySlashCount:  1,
			ConnectionsCount: 1,
			GuessTheNpcCount: 1,
			HangmanCount:     1,
			TerraTriviaCount: 1,
		},
	}

	store, err := NewGuessCountStore(ctx, fakeRepo, &domain.Broker{})
	if err != nil {
		t.Fatalf("newanswerstore failed: %v", err)
	}

	got := store.GetGuessCounts()

	want := domain.PlayerGuessCounts{
		DailySlashCount:  1,
		ConnectionsCount: 1,
		GuessTheNpcCount: 1,
		HangmanCount:     1,
		TerraTriviaCount: 1,
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("guess count store mismatch (-want +got):\n%s", diff)
	}
}

// Tests if reseting guess counts sets everything to 0
func TestResetGuessCounts(t *testing.T) {
	ctx := context.Background()

	fakeRepo := &testutils.FakeAnswerRepo{
		GuessCounts: domain.PlayerGuessCounts{
			DailySlashCount:  1,
			ConnectionsCount: 1,
			GuessTheNpcCount: 1,
			HangmanCount:     1,
			TerraTriviaCount: 1,
		},
	}

	store, err := NewGuessCountStore(ctx, fakeRepo, &domain.Broker{})
	if err != nil {
		t.Fatalf("newanswerstore failed: %v", err)
	}

	err = store.ResetGuessCounts(ctx)
	if err != nil {
		t.Fatalf("resetguesscounts failed: %v", err)
	}

	got := store.guessCountsCache
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

	tests := []struct {
		name      string
		increment func(store *CachedGuessCountsStore, ctx context.Context) (int, error)
	}{
		{
			name: "DailySlash",
			increment: func(s *CachedGuessCountsStore, ctx context.Context) (int, error) {
				return s.IncrementDailySlashCount(ctx)
			},
		},
		{
			name: "Connections",
			increment: func(s *CachedGuessCountsStore, ctx context.Context) (int, error) {
				return s.IncrementConnectionsCount(ctx)
			},
		},
		{
			name: "GuessTheNpc",
			increment: func(s *CachedGuessCountsStore, ctx context.Context) (int, error) {
				return s.IncrementGuessTheNpcCount(ctx)
			},
		},
		{
			name: "Hangman",
			increment: func(s *CachedGuessCountsStore, ctx context.Context) (int, error) {
				return s.IncrementHangmanCount(ctx)
			},
		},
		{
			name: "TerraTrivia",
			increment: func(s *CachedGuessCountsStore, ctx context.Context) (int, error) {
				return s.IncrementTerraTriviaCount(ctx)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeRepo := &testutils.FakeAnswerRepo{
				GuessCounts: domain.PlayerGuessCounts{
					DailySlashCount:  1,
					ConnectionsCount: 1,
					GuessTheNpcCount: 1,
					HangmanCount:     1,
					TerraTriviaCount: 1,
				},
			}

			store, err := NewGuessCountStore(ctx, fakeRepo, &domain.Broker{})
			if err != nil {
				t.Fatalf("newanswerstore failed: %v", err)
			}

			got, err := tt.increment(store, ctx)
			if err != nil {
				t.Fatalf("increment failed: %v", err)
			}
			if got != 2 {
				t.Errorf("Expected 2, got: %d", got)
			}
		})
	}
}
