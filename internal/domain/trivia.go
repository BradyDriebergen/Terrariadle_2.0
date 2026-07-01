package domain

type TriviaQuestion struct {
	ID         int
	Answer     string
	Clue       string
	Chunks     []string
	ChunkCount int
}
