package store

import (
	"context"
	"fmt"
	"sync"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
	"time"
)

type GuessCountsStore interface {
	GetGuessCounts() domain.PlayerGuessCounts
	ResetGuessCounts(ctx context.Context) error
	IncrementDailySlashCount(ctx context.Context) (int, error)
	IncrementConnectionsCount(ctx context.Context) (int, error)
	IncrementGuessTheNpcCount(ctx context.Context) (int, error)
	IncrementHangmanCount(ctx context.Context) (int, error)
	IncrementTerraTriviaCount(ctx context.Context) (int, error)
}

type CachedGuessCountsStore struct {
	answerRepo       repo.AnswerRepo
	mu               sync.RWMutex
	guessCountsCache domain.PlayerGuessCounts
	generation       int // bumped on every reset, so in-flight increments from before the reset are ignored
	broker           domain.GuessCountBroker
}

func NewGuessCountStore(
	ctx context.Context,
	answerRepo repo.AnswerRepo,
	broker domain.GuessCountBroker,
) (*CachedGuessCountsStore, error) {

	guessCounts, err := answerRepo.GetGuessCounts(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-guess-count-store: failed to initialize: %w", err)
	}

	return &CachedGuessCountsStore{
		answerRepo:       answerRepo,
		broker:           broker,
		guessCountsCache: toGuessCountDomain(guessCounts),
	}, nil
}

func (s *CachedGuessCountsStore) GetGuessCounts() domain.PlayerGuessCounts {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.guessCountsCache
}

func (s *CachedGuessCountsStore) ResetGuessCounts(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	guessCounts := repo.PlayerGuessCounts{}

	if err := s.answerRepo.UpsertGuessCounts(ctx, &guessCounts); err != nil {
		return fmt.Errorf("set-guess-counts: upserting guess counts: %w", err)
	}

	s.guessCountsCache = toGuessCountDomain(guessCounts)
	s.generation++

	return nil
}

func (s *CachedGuessCountsStore) IncrementDailySlashCount(ctx context.Context) (int, error) {
	count, err := s.increment(ctx, repo.DailySlashCountField, func(c *domain.PlayerGuessCounts) *int { return &c.DailySlashCount })
	if err != nil {
		return 0, fmt.Errorf("increment-daily-slash-count: %w", err)
	}
	s.broker.Publish(domain.GuessCountEvent{GameMode: domain.GameModeDailySlash, Count: count})
	return count, nil
}

func (s *CachedGuessCountsStore) IncrementConnectionsCount(ctx context.Context) (int, error) {
	count, err := s.increment(ctx, repo.ConnectionsCountField, func(c *domain.PlayerGuessCounts) *int { return &c.ConnectionsCount })
	if err != nil {
		return 0, fmt.Errorf("increment-connections-count: %w", err)
	}
	s.broker.Publish(domain.GuessCountEvent{GameMode: domain.GameModeConnections, Count: count})
	return count, nil
}

func (s *CachedGuessCountsStore) IncrementGuessTheNpcCount(ctx context.Context) (int, error) {
	count, err := s.increment(ctx, repo.GuessTheNpcCountField, func(c *domain.PlayerGuessCounts) *int { return &c.GuessTheNpcCount })
	if err != nil {
		return 0, fmt.Errorf("increment-guess-the-npc-count: %w", err)
	}
	s.broker.Publish(domain.GuessCountEvent{GameMode: domain.GameModeGuessTheNpc, Count: count})
	return count, nil
}

func (s *CachedGuessCountsStore) IncrementHangmanCount(ctx context.Context) (int, error) {
	count, err := s.increment(ctx, repo.HangmanCountField, func(c *domain.PlayerGuessCounts) *int { return &c.HangmanCount })
	if err != nil {
		return 0, fmt.Errorf("increment-hangman-count: %w", err)
	}
	s.broker.Publish(domain.GuessCountEvent{GameMode: domain.GameModeHangman, Count: count})
	return count, nil
}

func (s *CachedGuessCountsStore) IncrementTerraTriviaCount(ctx context.Context) (int, error) {
	count, err := s.increment(ctx, repo.TerraTriviaCountField, func(c *domain.PlayerGuessCounts) *int { return &c.TerraTriviaCount })
	if err != nil {
		return 0, fmt.Errorf("increment-terratrivia-count: %w", err)
	}
	s.broker.Publish(domain.GuessCountEvent{GameMode: domain.GameModeTerraTrivia, Count: count})
	return count, nil
}

func (s *CachedGuessCountsStore) increment(ctx context.Context, repoField string, field func(*domain.PlayerGuessCounts) *int) (int, error) {
	s.mu.RLock()
	generation := s.generation
	s.mu.RUnlock()

	// Other users' positions depend on this write, so don't let a client disconnect abort it
	ctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
	defer cancel()

	guessCounts, err := s.answerRepo.IncrementGuessCount(ctx, repoField)
	if err != nil {
		return 0, err
	}
	updated := toGuessCountDomain(guessCounts)
	count := *field(&updated)

	s.mu.Lock()
	defer s.mu.Unlock()

	// Skip stale results: a reset happened mid-write, or a later increment already landed
	if p := field(&s.guessCountsCache); s.generation == generation && count > *p {
		*p = count
	}

	return count, nil
}
