package services

var specialChars = []string{"'", ".", "-", "1", " "}

type HangmanInitData struct {
	Attempts int
	Finished bool
	Phrase   []string
	Guesses  []string
}

type HangmanCheckData struct {
	Phrase    []string
	Finished  bool
	IsCorrect bool
	Attempts  int
}

type HangmanWinningData struct {
	PlayerCount int
	Position    int
	EnemyName   string
	EnemyPath   string
}
