package store

import (
	"context"
	"terrariadle/internal/domain"
)

// struct and methods for mocking answer_repo
type fakeAnswerRepo struct {
	answerData      domain.AnswerRefs
	guessCounts     domain.PlayerGuessCounts
	getAnswerErr    error
	upsertAnswerErr error
	getGuessErr     error
	upsertGuessErr  error
}

func (f *fakeAnswerRepo) GetAnswerData(ctx context.Context) (domain.AnswerRefs, error) {
	if f.getAnswerErr != nil {
		return domain.AnswerRefs{}, f.getAnswerErr
	}
	return f.answerData, nil
}

func (f *fakeAnswerRepo) UpsertAnswerData(ctx context.Context, answerData *domain.AnswerRefs) error {
	if f.upsertAnswerErr != nil {
		return f.upsertAnswerErr
	}
	f.answerData = *answerData
	return nil
}

func (f *fakeAnswerRepo) GetGuessCounts(ctx context.Context) (domain.PlayerGuessCounts, error) {
	if f.getGuessErr != nil {
		return domain.PlayerGuessCounts{}, f.getGuessErr
	}
	return f.guessCounts, nil
}

func (f *fakeAnswerRepo) UpsertGuessCounts(ctx context.Context, guessCounts *domain.PlayerGuessCounts) error {
	if f.upsertGuessErr != nil {
		return f.upsertGuessErr
	}
	f.guessCounts = *guessCounts
	return nil
}

// struct and methods for mocking catalog_store
type fakeCatalogStore struct {
	weapons         []domain.Weapon
	weapon          domain.Weapon
	weaponFound     bool
	searchWeapons   []domain.SearchWeaponResult
	categories      []domain.Category
	category        domain.Category
	categoryFound   bool
	npcs            []domain.Npc
	npc             domain.Npc
	npcFound        bool
	searchNpcs      []domain.SearchNpcResult
	enemies         []domain.Enemy
	enemy           domain.Enemy
	enemyFound      bool
	triviaQuestions []domain.TriviaQuestion
	triviaQuestion  domain.TriviaQuestion
	triviaFound     bool
}

func (f *fakeCatalogStore) GetWeapons() []domain.Weapon {
	return f.weapons
}

func (f *fakeCatalogStore) GetWeapon(id int) (domain.Weapon, bool) {
	return f.weapon, f.weaponFound
}

func (f *fakeCatalogStore) GetSearchableWeapons() []domain.SearchWeaponResult {
	return f.searchWeapons
}

func (f *fakeCatalogStore) GetCategories() []domain.Category {
	return f.categories
}

func (f *fakeCatalogStore) GetCategory(id int) (domain.Category, bool) {
	return f.category, f.categoryFound
}

func (f *fakeCatalogStore) GetNpcs() []domain.Npc {
	return f.npcs
}

func (f *fakeCatalogStore) GetNpc(id int) (domain.Npc, bool) {
	return f.npc, f.npcFound
}

func (f *fakeCatalogStore) GetSearchableNpcs() []domain.SearchNpcResult {
	return f.searchNpcs
}

func (f *fakeCatalogStore) GetEnemies() []domain.Enemy {
	return f.enemies
}

func (f *fakeCatalogStore) GetEnemy(id int) (domain.Enemy, bool) {
	return f.enemy, f.enemyFound
}

func (f *fakeCatalogStore) GetTriviaQuestions() []domain.TriviaQuestion {
	return f.triviaQuestions
}

func (f *fakeCatalogStore) GetTriviaQuestion(id int) (domain.TriviaQuestion, bool) {
	return f.triviaQuestion, f.triviaFound
}
