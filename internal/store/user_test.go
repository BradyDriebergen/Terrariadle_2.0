package store

import (
	"context"
	"errors"
	"terrariadle/internal/domain"
	"terrariadle/internal/testutils"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// Tests if function gets existing users and creates new users
func TestGetOrCreateUser(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeUserRepo()

	store := NewUserStore(fakeRepo)

	// testing data setup
	newUserId := "1"
	newUser := testutils.GenerateUser(newUserId)

	existingUserId := "2"
	existingUser := testutils.GenerateUser(existingUserId)
	store.userCache[existingUserId] = existingUser

	tests := []struct {
		name   string
		userId string
		want   domain.User
	}{
		{name: "create user", userId: newUserId, want: newUser},
		{name: "get user", userId: existingUserId, want: existingUser},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := store.GetOrCreateUser(ctx, tt.userId)
			if err != nil {
				t.Fatalf("getorcreateuser failed: %v", err)
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.EquateApproxTime(time.Second)); diff != "" {
				t.Errorf("user store mismatch (-want +got):\n%s", diff)
			}

			if diff := cmp.Diff(tt.want, store.userCache[tt.userId], cmpopts.EquateApproxTime(time.Second)); diff != "" {
				t.Errorf("user store mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// Tests if getting user returns the user. Throws error otherwise.
func TestGetUser(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeUserRepo()

	store := NewUserStore(fakeRepo)

	tests := []struct {
		name   string
		userId string
		want   domain.User
		err    error
	}{
		{name: "error no user", userId: "1", want: domain.User{}, err: domain.MongoErrNotFound},
		{name: "get user", userId: "2", want: testutils.GenerateUser("2"), err: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				store.userCache[tt.userId] = tt.want
			}

			got, err := store.GetUser(ctx, tt.userId)
			if tt.err != nil {
				if !errors.Is(err, tt.err) {
					t.Fatalf("got err %v, want %v", err, tt.err)
				}
				return
			}
			if err != nil {
				t.Fatalf("getuser failed: %v", err)
			}

			if diff := cmp.Diff(tt.want, got, cmpopts.EquateApproxTime(time.Second)); diff != "" {
				t.Errorf("user store mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

// Checks if function successfully creates and updates the user
func TestUpsertUser(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeUserRepo()
	testUserId := "123"
	want := testutils.GenerateUser(testUserId)

	store := NewUserStore(fakeRepo)

	err := store.UpsertUser(ctx, want)
	if err != nil {
		t.Fatalf("upsertuser failed: %v", err)
	}

	if diff := cmp.Diff(want, fakeRepo.User, cmpopts.EquateApproxTime(time.Second)); diff != "" {
		t.Errorf("user store mismatch (-want +got):\n%s", diff)
	}

	if diff := cmp.Diff(want, store.userCache[testUserId], cmpopts.EquateApproxTime(time.Second)); diff != "" {
		t.Errorf("user store mismatch (-want +got):\n%s", diff)
	}

	want.DailySlash.Game.Finished = true

	err = store.UpsertUser(ctx, want)
	if err != nil {
		t.Fatalf("upsertuser failed: %v", err)
	}

	if diff := cmp.Diff(want, store.userCache[testUserId], cmpopts.EquateApproxTime(time.Second)); diff != "" {
		t.Errorf("user store mismatch (-want +got):\n%s", diff)
	}
}

// Checks if dropping all users removes them from the cache and database
func TestDropAllUsers(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeUserRepo()
	userId := "1"

	store := NewUserStore(fakeRepo)

	_, err := store.GetOrCreateUser(ctx, userId)
	if err != nil {
		t.Fatalf("getorcreateuser failed: %v", err)
	}

	err = store.DropAllUsers(ctx)
	if err != nil {
		t.Fatalf("dropallusers failed: %v", err)
	}

	if fakeRepo.User.UserID == userId {
		t.Errorf("want \"\", got %v", fakeRepo.User.UserID)
	}

	if len(store.userCache) != 0 {
		t.Errorf("want 0, got %v", len(store.userCache))
	}
}

// Tests if flushing dirty users updates them in the database
// and updates their status
func TestFlushDirty(t *testing.T) {
	ctx := context.Background()

	fakeRepo := testutils.GenerateFakeUserRepo()

	store := NewUserStore(fakeRepo)

	existingUserId := "1"
	existingUser := testutils.GenerateUser(existingUserId)
	existingUser.Dirty = true // Sets user to be flushed
	store.userCache[existingUserId] = existingUser

	err := store.FlushDirty(ctx)
	if err != nil {
		t.Fatalf("flushdirty failed: %v", err)
	}

	if diff := cmp.Diff(existingUser, fakeRepo.User, cmpopts.EquateApproxTime(time.Second)); diff != "" {
		t.Errorf("user store mismatch (-want +got):\n%s", diff)
	}

	if store.userCache[existingUserId].Dirty == true {
		t.Errorf("Dirty: want false, got %v", store.userCache[existingUserId].Dirty)
	}
}

// Tests if evicting stale removes users from the cache that haven't guessed
// in over an hour
func TestEvictStale(t *testing.T) {
	fakeRepo := testutils.GenerateFakeUserRepo()

	store := NewUserStore(fakeRepo)

	userToBeRemovedId := "1"
	userToBeRemoved := testutils.GenerateUser(userToBeRemovedId)
	userToBeRemoved.LastSeen = time.Now().Add(-1*time.Hour - time.Minute) // Sets user to be evicted
	store.userCache[userToBeRemovedId] = userToBeRemoved

	userToKeepId := "2"
	userToKeep := testutils.GenerateUser(userToKeepId)
	userToKeep.LastSeen = time.Now() // Sets user to be ignored from being evicted
	store.userCache[userToKeepId] = userToKeep

	store.EvictStale()

	if len(store.userCache) != 1 {
		t.Errorf("user cache length wanted 1, got %v", len(store.userCache))
	}
}
