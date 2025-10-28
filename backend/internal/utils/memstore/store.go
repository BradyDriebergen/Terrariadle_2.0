package memstore

import (
	"encoding/json"
	"fmt"
	"os"
	"terrariadle-backend/internal/domain"
)

// Global stores
var (
	GameData     = New[domain.GameData]()
	WeaponData   = New[[]domain.Weapon]()
	CategoryData = New[[]domain.Category]()
	NpcData      = New[[]domain.NPC]()
	HangmanData  = New[[]domain.Enemy]()
)

// Loads the files from json into memory for faster access
func InitializeDataFromJsonFiles() error {
	weapons, err := loadJSONData[domain.Weapon]("../data/weapons.json")
	if err != nil {
		return fmt.Errorf("error getting weapons from JSON file")
	}
	WeaponData.Set(weapons)

	categories, err := loadJSONData[domain.Category]("../data/categories.json")
	if err != nil {
		return fmt.Errorf("error getting categories from JSON file")
	}
	CategoryData.Set(categories)

	npcs, err := loadJSONData[domain.NPC]("../data/npcs.json")
	if err != nil {
		return fmt.Errorf("error getting npcs from JSON file")
	}
	NpcData.Set(npcs)

	enemies, err := loadJSONData[domain.Enemy]("../data/enemies.json")
	if err != nil {
		return fmt.Errorf("error getting enemies from JSON file")
	}
	HangmanData.Set(enemies)

	return nil
}

// Helper method for reading from JSON files
func loadJSONData[T any](path string) ([]T, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data []T
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("no data in %v", path)
	}

	return data, nil
}
