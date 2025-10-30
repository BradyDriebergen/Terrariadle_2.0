package store

import (
	"terrariadle-backend/internal/jsonreader"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/utils/atomicstore"
)

// Global stores
var (
	GameData        = atomicstore.New[models.GameData]()
	WeaponsCache    = atomicstore.New[[]jsonreader.Weapon]()
	CategoriesCache = atomicstore.New[[]jsonreader.Category]()
	NpcsCache       = atomicstore.New[[]jsonreader.Npc]()
	EnemiesCache    = atomicstore.New[[]jsonreader.Enemy]()
)

// Loads the files from json into memory for faster access
func InitializeStoreFromJson() error {
	weapons, err := jsonreader.GetWeaponsFromJSON()
	if err != nil {
		return err
	}
	WeaponsCache.Set(weapons)

	categories, err := jsonreader.GetCategoriesFromJson()
	if err != nil {
		return err
	}
	CategoriesCache.Set(categories)

	npcs, err := jsonreader.GetNpcsFromJson()
	if err != nil {
		return err
	}
	NpcsCache.Set(npcs)

	enemies, err := jsonreader.GetEnemiesFromJson()
	if err != nil {
		return err
	}
	EnemiesCache.Set(enemies)

	return nil
}
