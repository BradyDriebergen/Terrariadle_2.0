package utils

import (
	"sync"
	"terrariadle-backend/internal/types"
)

type MemoryStore struct {
	mu   sync.RWMutex
	data map[string]types.GameData
}

var MemStore *MemoryStore

func NewMemoryStore() {
	MemStore = &MemoryStore{data: make(map[string]types.GameData)}
}

func SetMemData(value types.GameData) {
	MemStore.mu.Lock()
	defer MemStore.mu.Unlock()
	MemStore.data["gameData"] = value
}

func GetMemData() (types.GameData, bool) {
	MemStore.mu.RLock()
	defer MemStore.mu.RUnlock()
	val, ok := MemStore.data["gameData"]
	return val, ok
}
