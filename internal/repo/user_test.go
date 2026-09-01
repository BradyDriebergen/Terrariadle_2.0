package repo

import (
	"context"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Checks if the correct user is returned, and that a ErrNotFound
// error is thrown.
func TestGetUser(t *testing.T) {
	ctx := context.Background()
	userRepo := mockUserRepo(t)

	user := generateUser1()
	want := toUser(user)

	tests := []struct {
		name    string
		seed    *userData
		id      string
		wantErr error
	}{
		{name: "ExistingUser", seed: &user, id: user.UserID},
		{name: "NonExistingUser", seed: nil, id: "does-not-exist", wantErr: ErrNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.seed != nil {
				err := db.Upsert(ctx, testMongo, userRepo.userCollection, db.Filter{"_id": tt.seed.ID}, *tt.seed)
				if err != nil {
					t.Fatalf("upsert failed: %v", err)
				}
			}

			got, err := userRepo.GetUser(ctx, tt.id)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Fatalf("got err %v, want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("getuser failed: %v", err)
			}

			// Bypass time issue, these don't run at the same time
			want.LastSeen = got.LastSeen

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("answer mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// Tests adding user and updating the same user
func TestUpsertUser(t *testing.T) {
	ctx := context.Background()
	userRepo := mockUserRepo(t)

	users := []domain.User{
		toUser(generateUser1()),
		toUser(generateUser2()),
	}

	for i := range users {
		err := userRepo.UpsertUserData(ctx, users[i])
		if err != nil {
			t.Fatalf("upsertuserdata failed: %v", err)
		}

		got, err := db.FindOne[userData](ctx, testMongo, userRepo.userCollection, db.Filter{"userId": users[i].UserID})
		if err != nil {
			t.Fatalf("findone failed: %v", err)
		}

		want := toUserData(users[i])
		got.ID = want.ID

		if diff := cmp.Diff(want, *got); diff != "" {
			t.Errorf("answer mismatch (-want +got):\n%s", diff)
		}
	}
}

// Test dropping all users, but keeping the collection
func TestDropAllUserData(t *testing.T) {
	ctx := context.Background()
	userRepo := mockUserRepo(t)

	users := []userData{
		generateUser1(),
		generateUser2(),
	}

	err := db.InsertMany(ctx, testMongo, userRepo.userCollection, users)
	if err != nil {
		t.Fatalf("insertmany failed: %v", err)
	}

	err = userRepo.DropAllUserData(ctx)
	if err != nil {
		t.Fatalf("dropalluserdata failed: %v", err)
	}

	got, err := db.GetAll[userData](ctx, testMongo, userRepo.userCollection)
	if err != nil {
		t.Fatalf("getall failed: %v", err)
	}

	if len(got) > 0 {
		t.Errorf("want 0 records, got: %d", len(got))
	}
}

// Helper method creating a collection for user
func mockUserRepo(t *testing.T) *MongoUserRepo {
	t.Helper()

	name := "repo_" + t.Name()

	t.Cleanup(func() {
		_ = db.DeleteAll(context.Background(), testMongo, name)
	})

	return NewUserRepo(testMongo, name)
}

// generates user data for testing
func generateUser1() userData {
	gameDoc := game{
		Guesses:  []int{2, 3, 4},
		HasWon:   true,
		Position: 2,
	}

	return userData{
		UserID: "2",
		DailySlash: dailySlashGame{
			Game:   gameDoc,
			Checks: []weaponChecks{},
		},
		Connections: connectionGame{
			Game:     gameDoc,
			Attempts: 2,
		},
		GuessTheNPC: guessTheNpcGame{
			Game:        gameDoc,
			GuessedName: "Alex",
		},
		Hangman: hangmanGame{
			Game:     gameDoc,
			Attempts: 4,
		},
		TerraTrivia: terraTriviaGame{
			Game: gameDoc,
		},
	}
}

// generates user data for testing
func generateUser2() userData {
	gameDoc := game{
		Guesses:  []int{1, 2, 3},
		HasWon:   false,
		Position: 1,
	}

	return userData{
		UserID: "1",
		DailySlash: dailySlashGame{
			Game:   gameDoc,
			Checks: []weaponChecks{},
		},
		Connections: connectionGame{
			Game:     gameDoc,
			Attempts: 3,
		},
		GuessTheNPC: guessTheNpcGame{
			Game:        gameDoc,
			GuessedName: "Steve",
		},
		Hangman: hangmanGame{
			Game:     gameDoc,
			Attempts: 6,
		},
		TerraTrivia: terraTriviaGame{
			Game: gameDoc,
		},
	}
}
