package services

type UserGameStatuses struct {
	DailySlash  bool `json:"daily_slash"`
	Connections bool `json:"connections"`
	GuessTheNpc bool `json:"guess_the_npc"`
	Hangman     bool `json:"hangman"`
	TerraTrivia bool `json:"terratrivia"`
}
