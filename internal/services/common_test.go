package services

import (
	"context"
	"errors"
	"terrariadle/internal/domain"
	"terrariadle/internal/testutils"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Checks if returned guess counts are correct, and it throws a not found error when
// misinput
func TestGetGuessCount(t *testing.T) {
	userStore := testutils.GenerateFakeUserStore()
	guessCountStore := testutils.GenerateFakeGuessCountStore()

	service := NewSseStream(guessCountStore, userStore)

	tests := []struct {
		name   string
		gameId domain.GameMode
		want   int
		err    domain.ErrCode
	}{
		{name: "daily slash", gameId: domain.GameModeDailySlash, want: guessCountStore.GuessCounts.DailySlashCount, err: ""},
		{name: "connections", gameId: domain.GameModeConnections, want: guessCountStore.GuessCounts.ConnectionsCount, err: ""},
		{name: "guess the npc", gameId: domain.GameModeGuessTheNpc, want: guessCountStore.GuessCounts.GuessTheNpcCount, err: ""},
		{name: "hangman", gameId: domain.GameModeHangman, want: guessCountStore.GuessCounts.HangmanCount, err: ""},
		{name: "terratrivia", gameId: domain.GameModeTerraTrivia, want: guessCountStore.GuessCounts.TerraTriviaCount, err: ""},
		{name: "gamemode not found error", gameId: domain.GameMode("wont_work"), want: 0, err: domain.ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.GetGuessCount(tt.gameId)

			if tt.err != "" {
				var appErr *domain.AppError

				if !errors.As(err, &appErr) {
					t.Fatalf("expected *AppError with code %q, got %v", tt.err, err)
				}
				if appErr.Code != tt.err {
					t.Fatalf("got code %q, want %q", appErr.Code, tt.err)
				}

				return
			}

			if err != nil {
				t.Fatalf("GetGuessCount failed: %v", err)
			}
			if got != tt.want {
				t.Errorf("want %v, got %v", tt.want, got)
			}
		})
	}
}

// Tests getting a user's finished game statuses or returning all
// false when user doesn't exist
func TestGetUserFinishedGames(t *testing.T) {
	ctx := context.Background()

	userStore := testutils.GenerateFakeUserStore()
	guessCountStore := testutils.GenerateFakeGuessCountStore()

	userId := "123"
	user := domain.User{
		UserID: userId,
		DailySlash: domain.DailySlashGame{
			Game: domain.Game{
				Finished: true,
			},
		},
	}
	userStore.UserCache[userId] = user

	service := NewSseStream(guessCountStore, userStore)

	got := service.GetUserFinishedGames(ctx, userId)

	if !got.DailySlash {
		t.Errorf("want true, got %v", got.DailySlash)
	}

	if user.DailySlash.Game.Finished != got.DailySlash {
		t.Errorf("want %v, got %v", user.DailySlash.Game.Finished, got.DailySlash)
	}

	// Checks that if a user doesn't exist, it defaults statuses to false
	gotNoUser := service.GetUserFinishedGames(ctx, "thisUserDoesn'tExist")

	if diff := cmp.Diff(UserGameStatuses{}, gotNoUser); diff != "" {
		t.Errorf("user store mismatch (-want +got):\n%s", diff)
	}
}
