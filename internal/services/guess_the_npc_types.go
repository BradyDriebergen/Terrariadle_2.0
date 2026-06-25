package services

type GuessTheNpcInitData struct {
	Quote      string                 `json:"quote"`
	Finished   bool                   `json:"finished"`
	Guesses    []GuessTheNpcGuessData `json:"guesses"`
	GuessedIDs []int                  `json:"guessed_ids"`
}

type SearchNpcData struct {
	NpcID int    `json:"id"`
	Name  string `json:"name"`
	Path  string `json:"path"`
}

type GuessTheNpcCheckData struct {
	Finished bool                 `json:"finished"`
	Guess    GuessTheNpcGuessData `json:"guess"`
}

type GuessTheNpcWinningData struct {
	Position    int      `json:"position"`
	Names       []string `json:"names"`
	GuessedName string   `json:"guessed_name"`
	CorrectName string   `json:"correct_name"`
}

type GuessTheNpcMiniGameData struct {
	GuessedName string `json:"guessed_name"`
	CorrectName string `json:"correct_name"`
}

type GuessTheNpcGuessData struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
