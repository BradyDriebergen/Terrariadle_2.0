package store

import (
	"context"
	"fmt"
	"sync"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
)

type UserStore struct {
	mu       sync.RWMutex
	cache    map[string]domain.User
	userRepo *repo.UserRepo
}

func NewUserStore(userRepo *repo.UserRepo) *UserStore {
	return &UserStore{
		cache:    make(map[string]domain.User),
		userRepo: userRepo,
	}
}

// GetUser returns from cache if present, otherwise fetches from repo and caches.
func (s *UserStore) GetUser(ctx context.Context, userID string) (domain.User, error) {
	s.mu.RLock()
	if user, ok := s.cache[userID]; ok {
		s.mu.RUnlock()
		return user, nil
	}
	s.mu.RUnlock()

	user, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("store: GetUser: %w", err)
	}

	s.mu.Lock()
	s.cache[userID] = user
	s.mu.Unlock()

	return user, nil
}

// UpsertUser writes through to the repo and updates the cache.
func (s *UserStore) UpsertUser(ctx context.Context, user *domain.User) error {
	if err := s.userRepo.UpsertUserData(ctx, user); err != nil {
		return fmt.Errorf("store: UpsertUser: %w", err)
	}

	s.mu.Lock()
	s.cache[user.UserID] = *user
	s.mu.Unlock()

	return nil
}

// DropAllUsers clears the repo and wipes the cache.
func (s *UserStore) DropAllUsers(ctx context.Context) error {
	if err := s.userRepo.DropAllUserData(ctx); err != nil {
		return fmt.Errorf("store: DropAllUsers: %w", err)
	}

	s.mu.Lock()
	s.cache = make(map[string]domain.User)
	s.mu.Unlock()

	return nil
}
