package store

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"terrariadle/internal/domain"
	"terrariadle/internal/repo"
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

func NewAnswerStore(ctx context.Context, ar repo.AnswerRepo, cs CatalogStore) (*CachedAnswerStore, error) {
	ad, err := ar.GetAnswerData(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-answer-store: failed to initialize: %w", err)
	}

	currentWeapon, ok := cs.GetWeapon(ad.DailySlash.CurrentWeaponID)
	if !ok {
		return nil, fmt.Errorf("failed to find current weapon with id: %v", ad.DailySlash.CurrentWeaponID)
	}

	prevWeapon, ok := cs.GetWeapon(ad.DailySlash.PrevWeaponID)
	if !ok {
		return nil, fmt.Errorf("failed to find previous weapon with id: %v", ad.DailySlash.PrevWeaponID)
	}

	npc, ok := cs.GetNpc(ad.GuessTheNpc.NpcID)
	if !ok {
		return nil, fmt.Errorf("failed to find npc with id: %v", ad.GuessTheNpc.NpcID)
	}

	enemy, ok := cs.GetEnemy(ad.Hangman.EnemyID)
	if !ok {
		return nil, fmt.Errorf("failed to find enemy with id: %v", ad.Hangman.EnemyID)
	}
	enemy.Name = strings.ToUpper(enemy.Name)

	triviaQuestions := make([]domain.TriviaQuestion, 7)
	for i, q := range ad.TerraTrivia.QuestionIDs {
		triviaQuestions[i], ok = cs.GetTriviaQuestion(q)
		if !ok {
			return nil, fmt.Errorf("failed to find enemy with id: %v", ad.Hangman.EnemyID)
		}
	}

	dailyAnswers := domain.DailyAnswers{
		DailySlash: domain.WeaponAnswer{
			CurrentWeapon: currentWeapon,
			PrevWeapon:    prevWeapon,
		},
		Connections: domain.ConnectionAnswer{
			CategoryIDs: ad.Connections.CategoryIDs,
			Options:     ad.Connections.Options,
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
		TerraTrivia: domain.TerraTriviaAnswer{
			Questions: triviaQuestions,
		},
		ResetTime:     ad.ResetTime,
		NextResetTime: ad.NextResetTime,
	}

	return &CachedAnswerStore{
		answerRepo:   ar,
		catalogStore: cs,
		answerCache:  dailyAnswers,
	}, nil
}

func (s *CachedAnswerStore) GetAnswers() domain.DailyAnswers {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.answerCache
}

func (s *CachedAnswerStore) UpsertAnswers(ctx context.Context, answer domain.DailyAnswers) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	triviaQuestionIDs := make([]int, 7)
	for i, q := range answer.TerraTrivia.Questions {
		triviaQuestionIDs[i] = q.ID
	}

	refs := domain.AnswerRefs{
		DailySlash: domain.WeaponRef{
			CurrentWeaponID: answer.DailySlash.CurrentWeapon.ID,
			PrevWeaponID:    answer.DailySlash.PrevWeapon.ID,
		},
		Connections: domain.ConnectionRef{
			CategoryIDs: answer.Connections.CategoryIDs,
			Options:     answer.Connections.Options,
		},
		GuessTheNpc: domain.NpcRef{
			NpcID:       answer.GuessTheNpc.NpcID,
			Quote:       answer.GuessTheNpc.Quote,
			Name:        answer.GuessTheNpc.Name,
			NameOptions: answer.GuessTheNpc.NameOptions,
		},
		Hangman: domain.HangmanRef{
			EnemyID: answer.Hangman.Enemy.ID,
		},
		TerraTrivia: domain.TerraTriviaRef{
			QuestionIDs: triviaQuestionIDs,
		},
		ResetTime:     answer.ResetTime,
		NextResetTime: answer.NextResetTime,
	}

	if err := s.answerRepo.UpsertAnswerData(ctx, &refs); err != nil {
		return fmt.Errorf("set-answer: upserting answers: %w", err)
	}

	s.answerCache = answer

	return nil
}
