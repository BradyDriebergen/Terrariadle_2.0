package store

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"terrariadle/internal/domain"
	"terrariadle/internal/repo"
	"time"
)

type UserStore interface {
	GetOrCreateUser(ctx context.Context, userID string) (domain.User, error)
	GetUser(ctx context.Context, userID string) (domain.User, error)
	UpsertUser(ctx context.Context, user domain.User) error
	DropAllUsers(ctx context.Context) error
	FlushDirty(ctx context.Context) error
	EvictStale()
}

type CachedUserStore struct {
	mu        sync.RWMutex
	userCache map[string]domain.User
	userRepo  repo.UserRepo
}

func NewUserStore(userRepo repo.UserRepo) *CachedUserStore {
	return &CachedUserStore{
		userCache: make(map[string]domain.User),
		userRepo:  userRepo,
	}
}

// GetOrCreateUser returns from cache if present, otherwise fetches from repo and caches.
func (s *CachedUserStore) GetOrCreateUser(ctx context.Context, userID string) (domain.User, error) {
	s.mu.RLock()
	if user, ok := s.userCache[userID]; ok {
		s.mu.RUnlock()
		return user, nil
	}
	s.mu.RUnlock()

	user, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		if errors.Is(err, repo.ErrNotFound) {
			// If no user is found, make a new one and put it in the cache
			user = createNewUser(userID)
		} else {
			return domain.User{}, fmt.Errorf("store: GetUser: %w", err)
		}
	}

	s.mu.Lock()
	// Prevents creating two new users from race condition
	if _, ok := s.userCache[userID]; !ok {
		s.userCache[userID] = user
	}
	s.mu.Unlock()

	return user, nil
}

func (s *CachedUserStore) GetUser(ctx context.Context, userID string) (domain.User, error) {
	s.mu.RLock()
	if user, ok := s.userCache[userID]; ok {
		s.mu.RUnlock()
		return user, nil
	}
	s.mu.RUnlock()

	user, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("store: GetUser: %w", err)
	}

	s.mu.Lock()
	s.userCache[userID] = user
	s.mu.Unlock()

	return user, nil
}

func (s *CachedUserStore) UpsertUser(ctx context.Context, user domain.User) error {
	if err := s.userRepo.UpsertUserData(ctx, user); err != nil {
		return fmt.Errorf("store: UpsertUser: %w", err)
	}

	s.mu.Lock()
	s.userCache[user.UserID] = user
	s.mu.Unlock()

	return nil
}

func (s *CachedUserStore) DropAllUsers(ctx context.Context) error {
	if err := s.userRepo.DropAllUserData(ctx); err != nil {
		return fmt.Errorf("store: DropAllUsers: %w", err)
	}

	s.mu.Lock()
	s.userCache = make(map[string]domain.User)
	s.mu.Unlock()

	return nil
}

func (s *CachedUserStore) FlushDirty(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for userID, entry := range s.userCache {
		if !entry.Dirty {
			continue
		}

		if err := s.userRepo.UpsertUserData(ctx, entry); err != nil {
			return fmt.Errorf("store: flush-dirty: %w", err)
		}

		entry.Dirty = false
		s.userCache[userID] = entry
	}

	return nil
}

func (s *CachedUserStore) EvictStale() {
	s.mu.Lock()
	defer s.mu.Unlock()

	cutoff := time.Now().Add(-1 * time.Hour)
	for userID, entry := range s.userCache {
		if entry.LastSeen.Before(cutoff) {
			delete(s.userCache, userID)
		}
	}
}

func createNewUser(userID string) domain.User {
	emptyGame := domain.Game{
		Guesses:  []int{},
		Finished: false,
		Position: 0,
	}

	return domain.User{
		UserID: userID,
		DailySlash: domain.DailySlashGame{
			Game:   emptyGame,
			Checks: []domain.WeaponChecks{},
		},
		Connections: domain.ConnectionGame{
			Game:     emptyGame,
			Attempts: 4,
		},
		GuessTheNPC: domain.GuessTheNpcGame{
			Game:        emptyGame,
			GuessedName: "",
		},
		Hangman: domain.HangmanGame{
			Game:     emptyGame,
			Attempts: 6,
		},
		TerraTrivia: domain.TerraTriviaGame{
			Game: emptyGame,
		},
		LastSeen: time.Now(),
		Dirty:    true,
	}
}
