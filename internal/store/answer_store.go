package store

import (
	"context"
	"fmt"
	"sync"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
)

type AnswerStore interface {
	GetAnswers() domain.DailyAnswers
	UpsertAnswers(ctx context.Context, answer domain.DailyAnswers) error
}

type CachedAnswerStore struct {
	answerRepo   repo.AnswerRepo
	catalogStore CatalogStore
	mu           sync.RWMutex
	answerCache  domain.DailyAnswers
}

func NewAnswerStore(ctx context.Context, answerRepo repo.AnswerRepo, catalogStore CatalogStore) (*CachedAnswerStore, error) {
	answerData, err := answerRepo.GetAnswerData(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-answer-store: failed to initialize: %w", err)
	}

	answers, err := toAnswerDomain(answerData, catalogStore)
	if err != nil {
		return nil, fmt.Errorf("new-answer-store: failed to initialize: %w", err)
	}

	return &CachedAnswerStore{
		answerRepo:   answerRepo,
		catalogStore: catalogStore,
		answerCache:  answers,
	}, nil
}

func (s *CachedAnswerStore) GetAnswers() domain.DailyAnswers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.answerCache
}

func (s *CachedAnswerStore) UpsertAnswers(ctx context.Context, answer domain.DailyAnswers) error {
	ad := fromAnswerDomain(answer)

	// Validate that all IDs resolve before committing
	if _, err := toAnswerDomain(ad, s.catalogStore); err != nil {
		return fmt.Errorf("set-answer: validation failed: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.answerRepo.UpsertAnswerData(ctx, &ad); err != nil {
		return fmt.Errorf("set-answer: upserting answers: %w", err)
	}

	s.answerCache = answer

	return nil
}
