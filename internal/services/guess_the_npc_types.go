package services

import "terrariadle-backend/internal/domain"

type GuessTheNpcInitData struct {
	Quote    string
	Finished bool
	Guesses  []domain.NpcInfo
}
