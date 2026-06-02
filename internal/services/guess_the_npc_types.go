package services

import "terrariadle-backend/internal/domain"

type GuessTheNpcInitData struct {
	Quote    string           `json:"quote"`
	Finished bool             `json:"finished"`
	Guesses  []domain.NpcInfo `json:"guesses"`
}

type GuessTheNpcCheckData struct {
	Finished bool           `json:"finished"`
	Guess    domain.NpcInfo `json:"guess"`
}

type GuessTheNpcWinningData struct {
	Position    int      `json:"position"`
	PlayerCount int      `json:"player_count"`
	Names       []string `json:"names"`
	GuessedName string   `json:"guessed_name"`
	CorrectName string   `json:"correct_name"`
}

type GuessTheNpcMiniGameData struct {
	GuessedName string `json:"guessed_name"`
	CorrectName string `json:"correct_name"`
}
