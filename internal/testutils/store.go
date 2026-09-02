package testutils

import "terrariadle/internal/domain"

// struct and methods for mocking catalog_store
type FakeCatalogStore struct {
	Weapons         []domain.Weapon
	SearchWeapons   []domain.SearchWeaponResult
	Categories      []domain.Category
	Npcs            []domain.Npc
	SearchNpcs      []domain.SearchNpcResult
	Enemies         []domain.Enemy
	TriviaQuestions []domain.TriviaQuestion
}

func (f *FakeCatalogStore) GetWeapons() []domain.Weapon {
	return f.Weapons
}

func (f *FakeCatalogStore) GetWeapon(id int) (domain.Weapon, bool) {
	for i := range f.Weapons {
		if f.Weapons[i].ID == id {
			return f.Weapons[i], true
		}
	}
	return domain.Weapon{}, false
}

func (f *FakeCatalogStore) GetSearchableWeapons() []domain.SearchWeaponResult {
	return f.SearchWeapons
}

func (f *FakeCatalogStore) GetCategories() []domain.Category {
	return f.Categories
}

func (f *FakeCatalogStore) GetCategory(id int) (domain.Category, bool) {
	for i := range f.Categories {
		if f.Categories[i].ID == id {
			return f.Categories[i], true
		}
	}
	return domain.Category{}, false
}

func (f *FakeCatalogStore) GetNpcs() []domain.Npc {
	return f.Npcs
}

func (f *FakeCatalogStore) GetNpc(id int) (domain.Npc, bool) {
	for i := range f.Npcs {
		if f.Npcs[i].ID == id {
			return f.Npcs[i], true
		}
	}
	return domain.Npc{}, false
}

func (f *FakeCatalogStore) GetSearchableNpcs() []domain.SearchNpcResult {
	return f.SearchNpcs
}

func (f *FakeCatalogStore) GetEnemies() []domain.Enemy {
	return f.Enemies
}

func (f *FakeCatalogStore) GetEnemy(id int) (domain.Enemy, bool) {
	for i := range f.Enemies {
		if f.Enemies[i].ID == id {
			return f.Enemies[i], true
		}
	}
	return domain.Enemy{}, false
}

func (f *FakeCatalogStore) GetTriviaQuestions() []domain.TriviaQuestion {
	return f.TriviaQuestions
}

func (f *FakeCatalogStore) GetTriviaQuestion(id int) (domain.TriviaQuestion, bool) {
	for i := range f.TriviaQuestions {
		if f.TriviaQuestions[i].ID == id {
			return f.TriviaQuestions[i], true
		}
	}
	return domain.TriviaQuestion{}, false
}
