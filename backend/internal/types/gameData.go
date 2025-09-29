package types

type GameData struct {
	DailySlash  Weapons    `json:"dailySlash"`
	Connections []Category `json:"connections"`
	GuessTheNpc NPC        `json:"guessTheNpc"`
	Hangman     Enemy      `json:"hangman"`
}
