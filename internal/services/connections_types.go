package services

type ConnectionsInitData struct {
	Attepts          int
	Finished         bool
	Options          []string
	SolvedCategories []SolvedCategory
}

type SolvedCategory struct {
	Name    string
	Options []string
}

type ConnectionsCheckData struct {
	Attempts     int
	IsCorrect    bool
	OneAway      bool
	CorrectGuess SolvedCategory
}

type ConnectionsWinningData struct {
	Position    int
	PlayerCount int
}
