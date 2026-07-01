package services

type TerraTriviaInitData struct {
	Finished    bool         `json:"finished"`
	TriviaItems []TriviaItem `json:"trivia_items"`
	Chunks      []string     `json:"chunks"`
}

type TerraTriviaCheckData struct {
	Finished    bool       `json:"finished"`
	GuessResult TriviaItem `json:"guess_result"`
	IsCorrect   bool       `json:"is_correct"`
}

type TriviaItem struct {
	ID          int    `json:"id"`
	Clue        string `json:"clue"`
	LetterCount int    `json:"letter_count"`
	Answer      string `json:"answer"`
}

type TerraTriviaWinningData struct {
	Position int `json:"position"`
}
