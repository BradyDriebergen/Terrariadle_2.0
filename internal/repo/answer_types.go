package repo

import "time"

type answerData struct {
	DailySlash    weaponData      `bson:"dailySlash"`
	Connections   connectionData  `bson:"connections"`
	GuessTheNpc   npcData         `bson:"guessTheNpc"`
	Hangman       hangmanData     `bson:"hangman"`
	TerraTrivia   terraTriviaData `bson:"terratrivia"`
	ResetTime     time.Time       `bson:"resetTime"`
	NextResetTime time.Time       `bson:"nextResetTime"`
}

type weaponData struct {
	CurrentWeaponID int `bson:"currentWeaponId"`
	PrevWeaponID    int `bson:"prevWeaponId"`
}

type connectionData struct {
	CategoryIDs []int              `bson:"categoryIds"`
	Options     []connectionOption `bson:"options"`
}

type connectionOption struct {
	Option     string `bson:"option"`
	CategoryID int    `bson:"categoryId"`
}

type npcData struct {
	NpcID       int      `bson:"npcId"`
	Quote       string   `bson:"quote"`
	Name        string   `bson:"name"`
	NameOptions []string `bson:"names"`
}

type hangmanData struct {
	EnemyID int `bson:"enemyId"`
}

type terraTriviaData struct {
	QuestionIDs []int `bson:"questionIds"`
}

type playerGuessCounts struct {
	DailySlashCount  int `bson:"dailySlashCount"`
	ConnectionsCount int `bson:"connectionsCount"`
	GuessTheNpcCount int `bson:"guessTheNpcCount"`
	HangmanCount     int `bson:"hangmanCount"`
	TerraTriviaCount int `bson:"terratriviaCount"`
}
