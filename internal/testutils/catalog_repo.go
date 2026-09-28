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

func (f *FakeCatalogRepo) GetWeapons(ctx context.Context) ([]domain.Weapon, error) {
	return f.Weapons, nil
}

func (f *FakeCatalogRepo) GetCategories(ctx context.Context) ([]domain.Category, error) {
	return f.Categories, nil
}

func (f *FakeCatalogRepo) GetNpcs(ctx context.Context) ([]domain.Npc, error) {
	return f.Npcs, nil
}

func (f *FakeCatalogRepo) GetEnemies(ctx context.Context) ([]domain.Enemy, error) {
	return f.Enemies, nil
}

func (f *FakeCatalogRepo) GetTriviaQuestions(ctx context.Context) ([]domain.TriviaQuestion, error) {
	return f.TriviaQuestions, nil
}

func GenerateFakeCatalogRepo() *FakeCatalogRepo {
	weapons := GenerateWeapons()
	npcs := GenerateNpcs()
	enemies := GenerateEnemies()
	categories := GenerateCategories()
	triviaQuestions := GenerateTriviaQuestions()

	return &FakeCatalogRepo{
		Weapons:         weapons,
		Npcs:            npcs,
		Enemies:         enemies,
		Categories:      categories,
		TriviaQuestions: triviaQuestions,
	}
}
