package store

import (
	"context"
	"fmt"
	"sync"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
)

type AnswerStore struct {
	answerRepo   *repo.AnswerRepo
	catalogStore *CatalogStore
	mu           sync.RWMutex
	answerCache  domain.DailyAnswers
}

func NewAnswerStore(ctx context.Context, answerRepo *repo.AnswerRepo, catalogStore *CatalogStore) (*AnswerStore, error) {
	answerData, err := answerRepo.GetAnswerData(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-answer-store: failed to initialize: %w", err)
	}

	answers, err := toDomain(answerData, catalogStore)
	if err != nil {
		return nil, fmt.Errorf("new-answer-store: failed to initialize: %w", err)
	}

	return &AnswerStore{
		answerRepo:   answerRepo,
		catalogStore: catalogStore,
		answerCache:  answers,
	}, nil
}

func (s *AnswerStore) GetAnswers() domain.DailyAnswers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.answerCache
}

func (s *AnswerStore) UpsertAnswers(ctx context.Context, answer domain.DailyAnswers) error {
	ad := fromDomain(answer)

	// Validate that all IDs resolve before committing
	if _, err := toDomain(ad, s.catalogStore); err != nil {
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

func toDomain(ad repo.AnswerData, cs *CatalogStore) (domain.DailyAnswers, error) {
	currentWeapon, ok := cs.GetWeapon(ad.DailySlash.CurrentWeaponID)
	if !ok {
		return domain.DailyAnswers{}, fmt.Errorf("failed to find current weapon with id: %v", ad.DailySlash.CurrentWeaponID)
	}

	prevWeapon, ok := cs.GetWeapon(ad.DailySlash.PrevWeaponID)
	if !ok {
		return domain.DailyAnswers{}, fmt.Errorf("failed to find previous weapon with id: %v", ad.DailySlash.PrevWeaponID)
	}

	npc, ok := cs.GetNpc(ad.GuessTheNpc.NpcID)
	if !ok {
		return domain.DailyAnswers{}, fmt.Errorf("failed to find npc with id: %v", ad.GuessTheNpc.NpcID)
	}

	enemy, ok := cs.GetEnemy(ad.Hangman.EnemyID)
	if !ok {
		return domain.DailyAnswers{}, fmt.Errorf("failed to find enemy with id: %v", ad.Hangman.EnemyID)
	}

	return domain.DailyAnswers{
		DailySlash: domain.WeaponAnswer{
			CurrentWeapon: currentWeapon,
			PrevWeapon:    prevWeapon,
		},
		Connections: domain.ConnectionAnswer{
			CategoryIDs: ad.Connections.CategoryIDs,
			Options:     optionsToAnswers(ad.Connections.Options),
		},
		GuessTheNpc: domain.NpcAnswer{
			NpcID:       npc.ID,
			Npc:         npc.NPC,
			Quote:       ad.GuessTheNpc.Quote,
			Name:        ad.GuessTheNpc.Name,
			NameOptions: ad.GuessTheNpc.NameOptions,
		},
		Hangman: domain.HangmanAnswer{
			Enemy: enemy,
		},
		GuessCounts: domain.PlayerGuessCounts{
			DailySlashCount:  ad.GuessCounts.DailySlashCount,
			ConnectionsCount: ad.GuessCounts.ConnectionsCount,
			GuessTheNpcCount: ad.GuessCounts.GuessTheNpcCount,
			HangmanCount:     ad.GuessCounts.HangmanCount,
		},
		ResetTime:     ad.ResetTime,
		NextResetTime: ad.NextResetTime,
	}, nil
}

func fromDomain(da domain.DailyAnswers) repo.AnswerData {
	return repo.AnswerData{
		DailySlash: repo.WeaponData{
			CurrentWeaponID: da.DailySlash.CurrentWeapon.ID,
			PrevWeaponID:    da.DailySlash.PrevWeapon.ID,
		},
		Connections: repo.ConnectionData{
			CategoryIDs: da.Connections.CategoryIDs,
			Options:     answersToOptions(da.Connections.Options),
		},
		GuessTheNpc: repo.NpcData{
			NpcID:       da.GuessTheNpc.NpcID,
			Quote:       da.GuessTheNpc.Quote,
			Name:        da.GuessTheNpc.Name,
			NameOptions: da.GuessTheNpc.NameOptions,
		},
		Hangman: repo.HangmanData{
			EnemyID: da.Hangman.Enemy.ID,
		},
		GuessCounts: repo.PlayerGuessCounts{
			DailySlashCount:  da.GuessCounts.DailySlashCount,
			ConnectionsCount: da.GuessCounts.ConnectionsCount,
			GuessTheNpcCount: da.GuessCounts.GuessTheNpcCount,
			HangmanCount:     da.GuessCounts.HangmanCount,
		},
		ResetTime:     da.ResetTime,
		NextResetTime: da.NextResetTime,
	}
}

func answersToOptions(answers []domain.ConnectionOption) []repo.ConnectionOption {
	options := make([]repo.ConnectionOption, len(answers))
	for i, a := range answers {
		options[i] = repo.ConnectionOption{
			Option:     a.Option,
			CategoryID: a.CategoryID,
		}
	}
	return options
}

func optionsToAnswers(options []repo.ConnectionOption) []domain.ConnectionOption {
	answers := make([]domain.ConnectionOption, len(options))
	for i, o := range options {
		answers[i] = domain.ConnectionOption{
			Option:     o.Option,
			CategoryID: o.CategoryID,
		}
	}
	return answers
}
