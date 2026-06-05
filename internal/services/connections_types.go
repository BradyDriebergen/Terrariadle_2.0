package services

type ConnectionsInitData struct {
	Attempts         int              `json:"attempts"`
	Finished         bool             `json:"finished"`
	Options          []string         `json:"options"`
	SolvedCategories []SolvedCategory `json:"solved_categories"`
}

type SolvedCategory struct {
	Name    string   `json:"name"`
	Options []string `json:"options"`
}

type ConnectionsCheckData struct {
	Attempts     int            `json:"attempts"`
	IsCorrect    bool           `json:"is_correct"`
	OneAway      bool           `json:"one_away"`
	CorrectGuess SolvedCategory `json:"correct_guess"`
}

type ConnectionsWinningData struct {
	Position int `json:"position"`
}
