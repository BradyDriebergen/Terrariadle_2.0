package services

var specialChars = []string{"'", ".", "-", "1", " "}

type HangmanInitData struct {
	Attempts int      `json:"attempts"`
	Finished bool     `json:"finished"`
	Phrase   []string `json:"phrase"`
	Guesses  []string `json:"guesses"`
}

type HangmanCheckData struct {
	Phrase    []string `json:"phrase"`
	Finished  bool     `json:"finished"`
	IsCorrect bool     `json:"is_correct"`
	Attempts  int      `json:"attempts"`
}

type HangmanWinningData struct {
	PlayerCount int    `json:"player_count"`
	Position    int    `json:"position"`
	EnemyName   string `json:"enemy_name"`
	EnemyPath   string `json:"enemy_path"`
}
