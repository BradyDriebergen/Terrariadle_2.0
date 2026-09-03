package repo

import (
	"context"
	"terrariadle/internal/db"
	"terrariadle/internal/domain"
	"testing"
)

// Checks if the correct user is returned, and that a ErrNotFound
// error is thrown.
func TestGetUser(t *testing.T) {
	ctx := context.Background()
	userRepo := mockUserRepo(t)

	user := newUserData("2")
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

			if want.UserID != got.UserID {
				t.Errorf("wanted %v, got %v", want.UserID, got.UserID)
			}
		})
	}
}

// Tests adding user and updating the same user
func TestUpsertUser(t *testing.T) {
	ctx := context.Background()
	userRepo := mockUserRepo(t)

	users := []domain.User{newUser("1"), newUser("2")}

	for i := range users {
		err := userRepo.UpsertUserData(ctx, users[i])
		if err != nil {
			t.Fatalf("upsertuserdata failed: %v", err)
		}

		got, err := db.FindOne[userData](ctx, testMongo, userRepo.userCollection, db.Filter{"userId": users[i].UserID})
		if err != nil {
			t.Fatalf("findone failed: %v", err)
		}

		if users[i].UserID != got.UserID {
			t.Errorf("wanted %v, got %v", users[i].UserID, got.UserID)
		}
	}
}

// Test dropping all users, but keeping the collection
func TestDropAllUserData(t *testing.T) {
	ctx := context.Background()
	userRepo := mockUserRepo(t)

	users := []domain.User{newUser("1"), newUser("2")}

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

func newUserData(id string) userData {
	return userData{UserID: id}
}

func newUser(id string) domain.User {
	return domain.User{UserID: id}
}
