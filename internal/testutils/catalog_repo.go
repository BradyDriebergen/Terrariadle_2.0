package testutils

import (
	"context"
	"terrariadle/internal/domain"
)

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

func GenerateFakeCatalogRepo() FakeCatalogRepo {
	weapons := GenerateWeapons()
	npcs := GenerateNpcs()
	enemies := GenerateEnemies()
	categories := GenerateCategories()
	triviaQuestions := GenerateTriviaQuestions()

	return FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}
}
