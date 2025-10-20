package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"terrariadle-backend/internal/types"
)

type gameDataStore struct {
	mu   sync.RWMutex
	data types.GameData
}
type puzzleDataStore struct {
	mu         sync.RWMutex
	Weapons    []types.Weapon
	Categories []types.Category
	Npcs       []types.NPC
	Enemies    []types.Enemy
}

var GameStore *gameDataStore
var PuzzleStore *puzzleDataStore

func NewGameStore() {
	GameStore = &gameDataStore{data: types.GameData{}}
}

func SetGameData(value types.GameData) {
	GameStore.mu.Lock()
	defer GameStore.mu.Unlock()
	GameStore.data = value
}

func GetGameData() types.GameData {
	GameStore.mu.RLock()
	defer GameStore.mu.RUnlock()
	val := GameStore.data
	return val
}

func NewPuzzleStore() error {
	weapons, err := loadJSONData[types.Weapon]("../data/weapons.json")
	if err != nil {
		return fmt.Errorf("error getting weapons from JSON file")
	}

	categories, err := loadJSONData[types.Category]("../data/categories.json")
	if err != nil {
		return fmt.Errorf("error getting categories from JSON file")
	}

	npcs, err := loadJSONData[types.NPC]("../data/npcs.json")
	if err != nil {
		return fmt.Errorf("error getting npcs from JSON file")
	}

	enemies, err := loadJSONData[types.Enemy]("../data/enemies.json")
	if err != nil {
		return fmt.Errorf("error getting enemies from JSON file")
	}

	PuzzleStore = &puzzleDataStore{
		Weapons:    weapons,
		Categories: categories,
		Npcs:       npcs,
		Enemies:    enemies,
	}

	return nil
}

type PuzzleItem interface {
	types.Weapon | types.Category | types.NPC | types.Enemy
}

func GetPuzzleData[T PuzzleItem](gameType string) ([]T, error) {
	PuzzleStore.mu.RLock()
	defer PuzzleStore.mu.RUnlock()

	switch gameType {
	case "DailySlash":
		if data, ok := any(PuzzleStore.Weapons).([]T); ok {
			return data, nil
		}
		return nil, fmt.Errorf("type mismatch: expected []types.Weapon for DailySlash")
	case "Connections":
		if data, ok := any(PuzzleStore.Categories).([]T); ok {
			return data, nil
		}
		return nil, fmt.Errorf("type mismatch: expected []types.Category for DailyCategory")
	case "GuessTheNpc":
		if data, ok := any(PuzzleStore.Npcs).([]T); ok {
			return data, nil
		}
		return nil, fmt.Errorf("type mismatch: expected []types.NPC for DailyNPC")
	case "Hangman":
		if data, ok := any(PuzzleStore.Enemies).([]T); ok {
			return data, nil
		}
		return nil, fmt.Errorf("type mismatch: expected []types.Enemy for DailyEnemy")
	default:
		return nil, fmt.Errorf("unknown game type: %s", gameType)
	}
}

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
