package store

import (
	"context"
	"fmt"
	"sync"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
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

	return nil
}

func (s *CachedGuessCountsStore) IncrementDailySlashCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.DailySlashCount++

	gc := fromGuessCountDomain(s.guessCountsCache)
	if err := s.answerRepo.UpsertGuessCounts(ctx, &gc); err != nil {
		s.guessCountsCache.DailySlashCount-- // roll back on failure
		return 0, fmt.Errorf("increment-daily-slash-count: %w", err)
	}

	count := s.guessCountsCache.DailySlashCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeDailySlash,
		Count:    count,
	})

	return s.guessCountsCache.DailySlashCount, nil
}

func (s *CachedGuessCountsStore) IncrementConnectionsCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.ConnectionsCount++

	gc := fromGuessCountDomain(s.guessCountsCache)
	if err := s.answerRepo.UpsertGuessCounts(ctx, &gc); err != nil {
		s.guessCountsCache.ConnectionsCount-- // roll back on failure
		return 0, fmt.Errorf("increment-connections-count: %w", err)
	}

	count := s.guessCountsCache.ConnectionsCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeConnections,
		Count:    count,
	})

	return s.guessCountsCache.ConnectionsCount, nil
}

func (s *CachedGuessCountsStore) IncrementGuessTheNpcCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.GuessTheNpcCount++

	gc := fromGuessCountDomain(s.guessCountsCache)
	if err := s.answerRepo.UpsertGuessCounts(ctx, &gc); err != nil {
		s.guessCountsCache.GuessTheNpcCount-- // roll back on failure
		return 0, fmt.Errorf("increment-guess-the-npc-count: %w", err)
	}

	count := s.guessCountsCache.GuessTheNpcCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeGuessTheNpc,
		Count:    count,
	})

	return s.guessCountsCache.GuessTheNpcCount, nil
}

func (s *CachedGuessCountsStore) IncrementHangmanCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.HangmanCount++

	gc := fromGuessCountDomain(s.guessCountsCache)
	if err := s.answerRepo.UpsertGuessCounts(ctx, &gc); err != nil {
		s.guessCountsCache.HangmanCount-- // roll back on failure
		return 0, fmt.Errorf("increment-hangman-count: %w", err)
	}

	count := s.guessCountsCache.HangmanCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeHangman,
		Count:    count,
	})

	return s.guessCountsCache.HangmanCount, nil
}

func (s *CachedGuessCountsStore) IncrementTerraTriviaCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.TerraTriviaCount++

	gc := fromGuessCountDomain(s.guessCountsCache)
	if err := s.answerRepo.UpsertGuessCounts(ctx, &gc); err != nil {
		s.guessCountsCache.TerraTriviaCount-- // roll back on failure
		return 0, fmt.Errorf("increment-terratrivia-count: %w", err)
	}

	count := s.guessCountsCache.TerraTriviaCount

	s.mu.Unlock()

	s.broker.Publish(domain.GuessCountEvent{
		GameMode: domain.GameModeTerraTrivia,
		Count:    count,
	})

	return s.guessCountsCache.TerraTriviaCount, nil
}
