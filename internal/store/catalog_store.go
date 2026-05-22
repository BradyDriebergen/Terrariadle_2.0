package store

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
)

type CatalogStore interface {
	GetWeapons() []domain.Weapon
	GetWeapon(id int) (domain.Weapon, bool)
	GetCategories() []domain.Category
	GetCategory(id int) (domain.Category, bool)
	GetNpcs() []domain.Npc
	GetNpc(id int) (domain.Npc, bool)
	GetEnemies() []domain.Enemy
	GetEnemy(id int) (domain.Enemy, bool)
}

type CachedCatalogStore struct {
	catalogRepo   repo.CatalogRepo
	weaponCache   map[int]domain.Weapon
	categoryCache map[int]domain.Category
	npcCache      map[int]domain.Npc
	enemyCache    map[int]domain.Enemy
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

	enemyData, err := catalogRepo.GetEnemies(ctx)
	if err != nil {
		return nil, fmt.Errorf("new-catalog-store: failed to initialize: %w", err)
	}
	if len(enemyData) == 0 {
		return nil, fmt.Errorf("new-catalog-store: no enemies found")
	}
	enemies := indexByID(enemyData, func(e domain.Enemy) int { return e.ID })

	return &CachedCatalogStore{
		catalogRepo:   catalogRepo,
		weaponCache:   weapons,
		categoryCache: categories,
		npcCache:      npcs,
		enemyCache:    enemies,
	}, nil
}

func (s *CachedCatalogStore) GetWeapons() []domain.Weapon {
	return slices.Collect(maps.Values(s.weaponCache))
}

func (s *CachedCatalogStore) GetWeapon(id int) (domain.Weapon, bool) {
	w, ok := s.weaponCache[id]
	return w, ok
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

func (s *CachedCatalogStore) GetEnemies() []domain.Enemy {
	return slices.Collect(maps.Values(s.enemyCache))
}

func (s *CachedCatalogStore) GetEnemy(id int) (domain.Enemy, bool) {
	e, ok := s.enemyCache[id]
	return e, ok
}

func indexByID[T any](items []T, id func(T) int) map[int]T {
	m := make(map[int]T, len(items))
	for _, item := range items {
		m[id(item)] = item
	}
	return m
}
