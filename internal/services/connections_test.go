package services

import (
	"context"
	"slices"
	"terrariadle/internal/domain"
	"terrariadle/internal/testutils"
	"testing"
)

func TestInitializeGame(t *testing.T) {
	ctx := context.Background()

	type wantData struct {
		attempts      int
		finished      bool
		optionsLength int
		solvedLength  int
	}

	tests := []struct {
		name string
		want wantData
		user *domain.User
	}{
		{
			name: "With no existing user",
			want: wantData{
				attempts:      4,
				finished:      false,
				optionsLength: 16,
				solvedLength:  0,
			},
			user: nil,
		},
		{
			name: "With an in-progress user",
			want: wantData{
				attempts:      3,
				finished:      false,
				optionsLength: 12,
				solvedLength:  1,
			},
			user: &domain.User{
				UserID: "testID",
				Connections: domain.ConnectionGame{
					Game: domain.Game{
						Guesses:  []int{0},
						Finished: false,
					},
					Attempts: 3,
				},
			},
		},
		{
			name: "With a winning user",
			want: wantData{
				attempts:      2,
				finished:      true,
				optionsLength: 0,
				solvedLength:  4,
			},
			user: &domain.User{
				UserID: "testID",
				Connections: domain.ConnectionGame{
					Game: domain.Game{
						Guesses:  []int{0, 1, 2, 3},
						Finished: true,
					},
					Attempts: 2,
				},
			},
		},
		{
			name: "With a loosing user",
			want: wantData{
				attempts:      0,
				finished:      true,
				optionsLength: 8,
				solvedLength:  2,
			},
			user: &domain.User{
				UserID: "testID",
				Connections: domain.ConnectionGame{
					Game: domain.Game{
						Guesses:  []int{0, 1},
						Finished: true,
					},
					Attempts: 0,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answerCache := testutils.GenerateFakeAnswerStore()
			guessCountCache := testutils.GenerateFakeGuessCountStore()
			catalogCache := testutils.GenerateFakeCatalogStore()
			userCache := testutils.GenerateFakeUserStore()

			var userID string
			if tt.user != nil {
				userID = tt.user.UserID
				userCache.UserCache[userID] = *tt.user
			}

			game := NewConnectionsGame(answerCache, guessCountCache, catalogCache, userCache)

			got, err := game.InitializeGame(ctx, userID)
			if err != nil {
				t.Fatalf("initializegame failed: %v", err)
			}

			if got.Attempts != tt.want.attempts {
				t.Errorf("Attempts: want %v, got %v", tt.want.attempts, got.Attempts)
			}
			if got.Finished != tt.want.finished {
				t.Errorf("Finished: want %v, got %v", tt.want.finished, got.Finished)
			}
			if len(got.Options) != tt.want.optionsLength {
				t.Errorf("Option length: want %v, got %v", tt.want.optionsLength, len(got.Options))
			}
			if len(got.SolvedCategories) != tt.want.solvedLength {
				t.Errorf("Solved categories length: want %v, got %v", tt.want.solvedLength, len(got.SolvedCategories))
			}

			var guesses []int
			if tt.user != nil {
				guesses = tt.user.Connections.Game.Guesses
			}

			var want []string
			for _, o := range answerCache.Answers.Connections.Options {
				if !slices.Contains(guesses, o.CategoryID) {
					want = append(want, o.Option)
				}
			}

			gotOpts := slices.Clone(got.Options)
			slices.Sort(gotOpts)
			slices.Sort(want)

			if !slices.Equal(gotOpts, want) {
				t.Errorf("Options = %q, want %q", gotOpts, want)
			}
		})
	}
}

func TestCheckGuess(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name string
		want ConnectionsCheckData
		user *domain.User
	}{
		{
			name: "Correct Guess",
			want: ConnectionsCheckData{
				Attempts:  4,
				IsCorrect: true,
				OneAway:   false,
				CorrectGuess: SolvedCategory{
					Name:    "cat1",
					Options: []string{"opt1_1", "opt2_1", "opt3_1", "opt4_1"},
				},
				Finished: false,
			},
			user: &domain.User{
				UserID: "testID",
				Connections: domain.ConnectionGame{
					Attempts: 4,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answerCache := testutils.GenerateFakeAnswerStore()
			guessCountCache := testutils.GenerateFakeGuessCountStore()
			catalogCache := testutils.GenerateFakeCatalogStore()
			userCache := testutils.GenerateFakeUserStore()

			var userID string
			if tt.user != nil {
				userID = tt.user.UserID
				userCache.UserCache[userID] = *tt.user
			}

			game := NewConnectionsGame(answerCache, guessCountCache, catalogCache, userCache)

			got, err := game.CheckGuess(ctx, userID, []string{"opt1_1", "opt2_1", "opt3_1", "opt4_1"})
			if err != nil {
				t.Fatalf("initializegame failed: %v", err)
			}

			if got.Attempts != tt.want.Attempts {
				t.Errorf("Attempts: want %v, got %v", tt.want.Attempts, got.Attempts)
			}
			if got.IsCorrect != tt.want.IsCorrect {
				t.Errorf("IsCorrect: want %v, got %v", tt.want.IsCorrect, got.IsCorrect)
			}
			if got.OneAway != tt.want.OneAway {
				t.Errorf("OneAway: want %v, got %v", tt.want.OneAway, got.OneAway)
			}
			if got.Finished != tt.want.Finished {
				t.Errorf("Finished: want %v, got %v", tt.want.Finished, got.Finished)
			}

			if got.IsCorrect && !slices.Equal(got.CorrectGuess.Options, tt.want.CorrectGuess.Options) {
				t.Errorf("CorrectGuess options: want %v, got %v", tt.want.CorrectGuess, got.CorrectGuess)
			}
		})
	}
}

func TestRevealAnswers(t *testing.T) {

}

func TestGetWinningData(t *testing.T) {

}
