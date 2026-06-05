package repo

import "time"

type AnswerData struct {
	DailySlash    WeaponData     `bson:"dailySlash"`
	Connections   ConnectionData `bson:"connections"`
	GuessTheNpc   NpcData        `bson:"guessTheNpc"`
	Hangman       HangmanData    `bson:"hangman"`
	ResetTime     time.Time      `bson:"resetTime"`
	NextResetTime time.Time      `bson:"nextResetTime"`
}

type WeaponData struct {
	CurrentWeaponID int `bson:"currentWeaponId"`
	PrevWeaponID    int `bson:"prevWeaponId"`
}

type ConnectionData struct {
	CategoryIDs []int              `bson:"categoryIds"`
	Options     []ConnectionOption `bson:"options"`
}

type ConnectionOption struct {
	Option     string `bson:"option"`
	CategoryID int    `bson:"categoryId"`
}

type NpcData struct {
	NpcID       int      `bson:"npcId"`
	Quote       string   `bson:"quote"`
	Name        string   `bson:"name"`
	NameOptions []string `bson:"names"`
}

type HangmanData struct {
	EnemyID int `bson:"enemyId"`
}

type PlayerGuessCounts struct {
	DailySlashCount  int `bson:"dailySlashCount"`
	ConnectionsCount int `bson:"connectionsCount"`
	GuessTheNpcCount int `bson:"guessTheNpcCount"`
	HangmanCount     int `bson:"hangmanCount"`
}
