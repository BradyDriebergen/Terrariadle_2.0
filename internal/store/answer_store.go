package store

import (
	"sync"
	"terrariadle-backend/internal/domain"
	"terrariadle-backend/internal/repo"
)

type AnswerStore struct {
	mu       sync.RWMutex
	cache    map[string]domain.Game
	userRepo *repo.UserRepo
}
