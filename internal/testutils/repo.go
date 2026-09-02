package testutils

import (
	"context"
	"terrariadle/internal/domain"
)

// struct and methods for mocking answer_repo
type FakeAnswerRepo struct {
	AnswerData      domain.AnswerRefs
	GuessCounts     domain.PlayerGuessCounts
	GetAnswerErr    error
	UpsertAnswerErr error
	GetGuessErr     error
	UpsertGuessErr  error
}

func (f *FakeAnswerRepo) GetAnswerData(ctx context.Context) (domain.AnswerRefs, error) {
	if f.GetAnswerErr != nil {
		return domain.AnswerRefs{}, f.GetAnswerErr
	}
	return f.AnswerData, nil
}

func (f *FakeAnswerRepo) UpsertAnswerData(ctx context.Context, answerData *domain.AnswerRefs) error {
	if f.UpsertAnswerErr != nil {
		return f.UpsertAnswerErr
	}
	f.AnswerData = *answerData
	return nil
}

func (f *FakeAnswerRepo) GetGuessCounts(ctx context.Context) (domain.PlayerGuessCounts, error) {
	if f.GetGuessErr != nil {
		return domain.PlayerGuessCounts{}, f.GetGuessErr
	}
	return f.GuessCounts, nil
}

func (f *FakeAnswerRepo) UpsertGuessCounts(ctx context.Context, guessCounts *domain.PlayerGuessCounts) error {
	if f.UpsertGuessErr != nil {
		return f.UpsertGuessErr
	}
	f.GuessCounts = *guessCounts
	return nil
}

// struct and methods for mocking catalog_repo
type FakeCatalogRepo struct {
	Weapons         []domain.Weapon
	Categories      []domain.Category
	Npcs            []domain.Npc
	Enemies         []domain.Enemy
	TriviaQuestions []domain.TriviaQuestion
}

func (m *FakeCatalogRepo) GetWeapons(ctx context.Context) ([]domain.Weapon, error) {
	return m.Weapons, nil
}

func (m *FakeCatalogRepo) GetCategories(ctx context.Context) ([]domain.Category, error) {
	return m.Categories, nil
}

func (m *FakeCatalogRepo) GetNpcs(ctx context.Context) ([]domain.Npc, error) {
	return m.Npcs, nil
}

func (m *FakeCatalogRepo) GetEnemies(ctx context.Context) ([]domain.Enemy, error) {
	return m.Enemies, nil
}

func (m *FakeCatalogRepo) GetTriviaQuestions(ctx context.Context) ([]domain.TriviaQuestion, error) {
	return m.TriviaQuestions, nil
}
