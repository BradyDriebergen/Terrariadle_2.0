package types

type GameData struct {
	DailySlash  WeaponData `json:"dailySlash"`
	Connections []Category `json:"connections"`
	GuessTheNpc NPCdata    `json:"guessTheNpc"`
	Hangman     Enemy      `json:"hangman"`
}
