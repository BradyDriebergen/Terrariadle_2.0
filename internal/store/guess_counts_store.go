package store

import (
	"context"
	"fmt"
	"sync"
	"terrariadle/internal/domain"
	"terrariadle/internal/repo"
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
		guessCountsCache: guessCounts,
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

	// builds a new guess count struct and initializes all fields to 0
	guessCounts := domain.PlayerGuessCounts{}

	if err := s.answerRepo.UpsertGuessCounts(ctx, &guessCounts); err != nil {
		return fmt.Errorf("set-guess-counts: upserting guess counts: %w", err)
	}

	s.guessCountsCache = guessCounts

	return nil
}

func (s *CachedGuessCountsStore) IncrementDailySlashCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.DailySlashCount++

	if err := s.answerRepo.UpsertGuessCounts(ctx, &s.guessCountsCache); err != nil {
		s.guessCountsCache.DailySlashCount-- // roll back on failure
		s.mu.Unlock()
		return 0, fmt.Errorf("increment-daily-slash-count: %w", err)
	}

	count := s.guessCountsCache.DailySlashCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeDailySlash,
		Count:    count,
	})

	return count, nil
}

func (s *CachedGuessCountsStore) IncrementConnectionsCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.ConnectionsCount++

	if err := s.answerRepo.UpsertGuessCounts(ctx, &s.guessCountsCache); err != nil {
		s.guessCountsCache.ConnectionsCount-- // roll back on failure
		s.mu.Unlock()
		return 0, fmt.Errorf("increment-connections-count: %w", err)
	}

	count := s.guessCountsCache.ConnectionsCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeConnections,
		Count:    count,
	})

	return count, nil
}

func (s *CachedGuessCountsStore) IncrementGuessTheNpcCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.GuessTheNpcCount++

	if err := s.answerRepo.UpsertGuessCounts(ctx, &s.guessCountsCache); err != nil {
		s.guessCountsCache.GuessTheNpcCount-- // roll back on failure
		s.mu.Unlock()
		return 0, fmt.Errorf("increment-guess-the-npc-count: %w", err)
	}

	count := s.guessCountsCache.GuessTheNpcCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeGuessTheNpc,
		Count:    count,
	})

	return count, nil
}

func (s *CachedGuessCountsStore) IncrementHangmanCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.HangmanCount++

	if err := s.answerRepo.UpsertGuessCounts(ctx, &s.guessCountsCache); err != nil {
		s.guessCountsCache.HangmanCount-- // roll back on failure
		s.mu.Unlock()
		return 0, fmt.Errorf("increment-hangman-count: %w", err)
	}

	count := s.guessCountsCache.HangmanCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeHangman,
		Count:    count,
	})

	return count, nil
}

func (s *CachedGuessCountsStore) IncrementTerraTriviaCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.TerraTriviaCount++

	if err := s.answerRepo.UpsertGuessCounts(ctx, &s.guessCountsCache); err != nil {
		s.guessCountsCache.TerraTriviaCount-- // roll back on failure
		s.mu.Unlock()
		return 0, fmt.Errorf("increment-terratrivia-count: %w", err)
	}

	count := s.guessCountsCache.TerraTriviaCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeTerraTrivia,
		Count:    count,
	})

	return count, nil
}
