package types

import "time"

type GameData struct {
	DailySlash    WeaponData        `bson:"dailySlash" json:"dailySlash"`
	Connections   []Category        `bson:"connections" json:"connections"`
	GuessTheNpc   NPCdata           `bson:"guessTheNpc" json:"guessTheNpc"`
	Hangman       Enemy             `bson:"hangman" json:"hangman"`
	GuessCounts   PlayerGuessCounts `bson:"guessCounts" json:"guessCounts"`
	ResetTime     time.Time         `bson:"resetTime" json:"resetTime"`
	NextResetTime time.Time         `bson:"nextResetTime" json:"nextResetTime"`
}

type PlayerGuessCounts struct {
	DailySlashCount  int `bson:"dailySlashCount" json:"dailySlashCount"`
	ConnectionsCount int `bson:"connectionsCount" json:"connectionsCount"`
	GuessTheNpcCount int `bson:"guessTheNpcCount" json:"guessTheNpcCount"`
	HangmanCount     int `bson:"hangmanCount" json:"hangmanCount"`
}
