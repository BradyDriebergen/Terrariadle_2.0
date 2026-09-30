package testutils

import (
	"context"
	"terrariadle/internal/domain"
)

type FakeUserStore struct {
	UserCache map[string]domain.User
}

func (f *FakeUserStore) GetOrCreateUser(ctx context.Context, userID string) (domain.User, error) {
	if user, ok := f.UserCache[userID]; ok {
		return user, nil
	}

	user := GenerateUser(userID)
	f.UserCache[userID] = user
	return user, nil
}

func (f *FakeUserStore) GetUser(ctx context.Context, userID string) (domain.User, error) {
	if user, ok := f.UserCache[userID]; ok {
		return user, nil
	}
	return domain.User{}, domain.MongoErrNotFound
}

func (f *FakeUserStore) UpsertUser(ctx context.Context, user domain.User) error {
	f.UserCache[user.UserID] = user
	return nil
}

func (f *FakeUserStore) DropAllUsers(ctx context.Context) error {
	f.UserCache = make(map[string]domain.User)
	return nil
}

func (f *FakeUserStore) FlushDirty(ctx context.Context) error {
	for userID, entry := range f.UserCache {
		if entry.Dirty {
			entry.Dirty = false
			f.UserCache[userID] = entry
		}
	}

	return nil
}

func (f *FakeUserStore) EvictStale() {
	f.UserCache = make(map[string]domain.User)
}

func GenerateFakeUserStore() *FakeUserStore {
	return &FakeUserStore{
		UserCache: make(map[string]domain.User),
	}
}
