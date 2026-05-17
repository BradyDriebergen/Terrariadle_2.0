package store

import (
	"maps"
	"slices"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
)

type CatalogStore struct {
	catalogRepo *repo.CatalogRepo

	weaponCache   map[int]domain.Weapon
	categoryCache map[int]domain.Category
	npcCache      map[int]domain.Npc
	enemyCache    map[int]domain.Enemy
}

func (s *CatalogStore) GetWeapons() []domain.Weapon {
	return slices.Collect(maps.Values(s.weaponCache))
}

func (s *CatalogStore) GetWeapon(id int) (domain.Weapon, bool) {
	w, ok := s.weaponCache[id]
	return w, ok
}

func (s *CatalogStore) GetCategories() []domain.Category {
	return slices.Collect(maps.Values(s.categoryCache))
}

func (s *CatalogStore) GetCategory(id int) (domain.Category, bool) {
	c, ok := s.categoryCache[id]
	return c, ok
}

func (s *CatalogStore) GetNpcs() []domain.Npc {
	return slices.Collect(maps.Values(s.npcCache))
}

func (s *CatalogStore) GetNpc(id int) (domain.Npc, bool) {
	n, ok := s.npcCache[id]
	return n, ok
}

func (s *CatalogStore) GetEnemies() []domain.Enemy {
	return slices.Collect(maps.Values(s.enemyCache))
}

func (s *CatalogStore) GetEnemy(id int) (domain.Enemy, bool) {
	e, ok := s.enemyCache[id]
	return e, ok
}
