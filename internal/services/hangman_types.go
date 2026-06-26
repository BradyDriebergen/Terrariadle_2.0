package services

var specialChars = []string{"'", ".", "-", "1", " "}

type HangmanInitData struct {
	Attempts int            `json:"attempts"`
	Finished bool           `json:"finished"`
	Phrase   []string       `json:"phrase"`
	Guesses  []HangmanGuess `json:"guesses"`
}

type HangmanCheckData struct {
	Phrase   []string     `json:"phrase"`
	Guess    HangmanGuess `json:"guess"`
	Finished bool         `json:"finished"`
	Attempts int          `json:"attempts"`
}

type HangmanGuess struct {
	Letter  string `json:"letter"`
	Correct bool   `json:"correct"`
}

type HangmanWinningData struct {
	Position  int    `json:"position"`
	EnemyName string `json:"enemy_name"`
	EnemyPath string `json:"enemy_path"`
}
