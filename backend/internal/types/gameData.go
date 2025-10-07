package types

import "time"

type GameData struct {
	DailySlash    WeaponData `bson:"dailySlash" json:"dailySlash"`
	Connections   []Category `bson:"connections" json:"connections"`
	GuessTheNpc   NPCdata    `bson:"guessTheNpc" json:"guessTheNpc"`
	Hangman       Enemy      `bson:"hangman" json:"hangman"`
	ResetTime     time.Time  `bson:"resetTime" json:"resetTime"`
	NextResetTime time.Time  `bson:"nextResetTime" json:"nextResetTime"`
}
