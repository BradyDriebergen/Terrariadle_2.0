package store

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"terrariadle/internal/domain"
	"terrariadle/internal/repo"
)

type CatalogStore interface {
	GetWeapons() []domain.Weapon
	GetWeapon(id int) (domain.Weapon, bool)
	GetSearchableWeapons() []domain.SearchWeaponResult
	GetCategories() []domain.Category
	GetCategory(id int) (domain.Category, bool)
	GetNpcs() []domain.Npc
	GetNpc(id int) (domain.Npc, bool)
	GetSearchableNpcs() []domain.SearchNpcResult
	GetEnemies() []domain.Enemy
	GetEnemy(id int) (domain.Enemy, bool)
	GetTriviaQuestions() []domain.TriviaQuestion
	GetTriviaQuestion(id int) (domain.TriviaQuestion, bool)
}

type CachedCatalogStore struct {
	catalogRepo       repo.CatalogRepo
	weaponCache       map[int]domain.Weapon
	searchWeaponCache []domain.SearchWeaponResult
	categoryCache     map[int]domain.Category
	npcCache          map[int]domain.Npc
	searchNpcCache    []domain.SearchNpcResult
	enemyCache        map[int]domain.Enemy
	triviaCache       map[int]domain.TriviaQuestion
}

func NewCatalogStore(ctx context.Context, catalogRepo repo.CatalogRepo) (*CachedCatalogStore, error) {
	weaponData, err := catalogRepo.GetWeapons(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-catalog-store: failed to initialize: %w", err)
	}
	if len(weaponData) == 0 {
		return nil, fmt.Errorf("new-catalog-store: no weapons found")
	}
	weapons := indexByID(weaponData, func(w domain.Weapon) int { return w.ID })

	searchWeapons := make([]domain.SearchWeaponResult, 0, len(weaponData))
	for i := range weaponData {
		searchWeapons = append(searchWeapons, toSearchWeapon(weaponData[i]))
	}

	categoryData, err := catalogRepo.GetCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-catalog-store: failed to initialize: %w", err)
	}
	if len(categoryData) == 0 {
		return nil, fmt.Errorf("new-catalog-store: no categories found")
	}
	categories := indexByID(categoryData, func(c domain.Category) int { return c.ID })

	npcData, err := catalogRepo.GetNpcs(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-catalog-store: failed to initialize: %w", err)
	}
	if len(npcData) == 0 {
		return nil, fmt.Errorf("new-catalog-store: no npcs found")
	}
	npcs := indexByID(npcData, func(n domain.Npc) int { return n.ID })

	searchNpcs := make([]domain.SearchNpcResult, 0, len(npcData))
	for i := range npcData {
		searchNpcs = append(searchNpcs, toSearchNPC(npcData[i]))
	}

	enemyData, err := catalogRepo.GetEnemies(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-catalog-store: failed to initialize: %w", err)
	}
	if len(enemyData) == 0 {
		return nil, fmt.Errorf("new-catalog-store: no enemies found")
	}
	enemies := indexByID(enemyData, func(e domain.Enemy) int { return e.ID })

	triviaData, err := catalogRepo.GetTriviaQuestions(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-catalog-store: failed to initialize: %w", err)
	}
	if len(triviaData) == 0 {
		return nil, fmt.Errorf("new-catalog-store: no trivia questions found")
	}
	triviaQuestions := indexByID(triviaData, func(e domain.TriviaQuestion) int { return e.ID })

	return &CachedCatalogStore{
		catalogRepo:       catalogRepo,
		weaponCache:       weapons,
		searchWeaponCache: searchWeapons,
		categoryCache:     categories,
		npcCache:          npcs,
		searchNpcCache:    searchNpcs,
		enemyCache:        enemies,
		triviaCache:       triviaQuestions,
	}, nil
}

func (s *CachedCatalogStore) GetWeapons() []domain.Weapon {
	return slices.Collect(maps.Values(s.weaponCache))
}

func (s *CachedCatalogStore) GetWeapon(id int) (domain.Weapon, bool) {
	w, ok := s.weaponCache[id]
	return w, ok
}

func (s *CachedCatalogStore) GetSearchableWeapons() []domain.SearchWeaponResult {
	return s.searchWeaponCache
}

func (s *CachedCatalogStore) GetCategories() []domain.Category {
	return slices.Collect(maps.Values(s.categoryCache))
}

func (s *CachedCatalogStore) GetCategory(id int) (domain.Category, bool) {
	c, ok := s.categoryCache[id]
	return c, ok
}

func (s *CachedCatalogStore) GetNpcs() []domain.Npc {
	return slices.Collect(maps.Values(s.npcCache))
}

func (s *CachedCatalogStore) GetNpc(id int) (domain.Npc, bool) {
	n, ok := s.npcCache[id]
	return n, ok
}

func (s *CachedCatalogStore) GetSearchableNpcs() []domain.SearchNpcResult {
	return s.searchNpcCache
}

func (s *CachedCatalogStore) GetEnemies() []domain.Enemy {
	return slices.Collect(maps.Values(s.enemyCache))
}

func (s *CachedCatalogStore) GetEnemy(id int) (domain.Enemy, bool) {
	e, ok := s.enemyCache[id]
	return e, ok
}

func (s *CachedCatalogStore) GetTriviaQuestions() []domain.TriviaQuestion {
	return slices.Collect(maps.Values(s.triviaCache))
}

func (s *CachedCatalogStore) GetTriviaQuestion(id int) (domain.TriviaQuestion, bool) {
	e, ok := s.triviaCache[id]
	return e, ok
}

func indexByID[T any](items []T, id func(T) int) map[int]T {
	m := make(map[int]T, len(items))
	for i := range items {
		m[id(items[i])] = items[i]
	}
	return m
}

func toSearchWeapon(w domain.Weapon) domain.SearchWeaponResult {
	return domain.SearchWeaponResult{
		WeaponID: w.ID,
		Name:     w.Name,
		Path:     w.Info.ImagePath,
	}
}

func toSearchNPC(n domain.Npc) domain.SearchNpcResult {
	return domain.SearchNpcResult{
		NpcID: n.ID,
		Name:  n.NPC,
		Path:  n.NpcPath,
	}
}
