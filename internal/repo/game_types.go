package repo

import "time"

type gameData struct {
	DailySlash    weaponData        `bson:"dailySlash"`
	Connections   connectionsData   `bson:"connections"`
	GuessTheNpc   npcData           `bson:"guessTheNpc"`
	Hangman       hangmanData       `bson:"hangman"`
	GuessCounts   playerGuessCounts `bson:"guessCounts"`
	ResetTime     time.Time         `bson:"resetTime"`
	NextResetTime time.Time         `bson:"nextResetTime"`
}

type playerGuessCounts struct {
	DailySlashCount  int `bson:"dailySlashCount"`
	ConnectionsCount int `bson:"connectionsCount"`
	GuessTheNpcCount int `bson:"guessTheNpcCount"`
	HangmanCount     int `bson:"hangmanCount"`
}

type weaponData struct {
	CurrentWeaponID int `bson:"currentWeaponId"`
	PrevWeaponID    int `bson:"prevWeaponId"`
}

type connectionsData struct {
	CategoryIDs []int `bson:"categoryIds"`
}

type npcData struct {
	ID    int      `bson:"id"`
	Quote string   `bson:"quote"`
	Names []string `bson:"names"`
}

type hangmanData struct {
	ID int `bson:"id"`
}
