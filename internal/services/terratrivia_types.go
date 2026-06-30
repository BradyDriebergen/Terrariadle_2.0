package services

type TerraTriviaInitData struct {
	Finished    bool         `json:"finished"`
	TriviaItems []TriviaItem `json:"trivia_items"`
	Chunks      []string     `json:"chunks"`
}

type TriviaItem struct {
	Clue        string `json:"clue"`
	LetterCount int    `json:"letter_count"`
	Answer      string `json:"answer"`
}
