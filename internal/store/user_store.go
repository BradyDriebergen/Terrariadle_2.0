package store

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
)

type UserStore struct {
	mu        sync.RWMutex
	userCache map[string]domain.User
	userRepo  *repo.UserRepo
}

func NewUserStore(userRepo *repo.UserRepo) *UserStore {
	return &UserStore{
		userCache: make(map[string]domain.User),
		userRepo:  userRepo,
	}
}

// GetUser returns from cache if present, otherwise fetches from repo and caches.
func (s *UserStore) GetUser(ctx context.Context, userID string) (domain.User, error) {
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
	s.userCache[userID] = user
	s.mu.Unlock()

	return user, nil
}

func (s *UserStore) UpsertUser(ctx context.Context, user *domain.User) error {
	if err := s.userRepo.UpsertUserData(ctx, user); err != nil {
		return fmt.Errorf("store: UpsertUser: %w", err)
	}

	s.mu.Lock()
	s.userCache[user.UserID] = *user
	s.mu.Unlock()

	return nil
}

func (s *UserStore) DropAllUsers(ctx context.Context) error {
	if err := s.userRepo.DropAllUserData(ctx); err != nil {
		return fmt.Errorf("store: DropAllUsers: %w", err)
	}

	s.mu.Lock()
	s.userCache = make(map[string]domain.User)
	s.mu.Unlock()

	return nil
}

func createNewUser(userID string) domain.User {
	emptyGame := domain.Game{
		Guesses:  []int{},
		HasWon:   false,
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
	}
}
