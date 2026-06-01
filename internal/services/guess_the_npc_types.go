package services

import "terrariadle-backend/internal/domain"

type GuessTheNpcInitData struct {
	Quote    string
	Finished bool
	Guesses  []domain.NpcInfo
}

type GuessTheNpcCheckData struct {
	Finished bool
	Guess    domain.NpcInfo
}

type GuessTheNpcWinningData struct {
	Position    int
	PlayerCount int
	Names       []string
	GuessedName string
	CorrectName string
}

type GuessTheNpcMiniGameData struct {
	GuessedName string
	CorrectName string
}
