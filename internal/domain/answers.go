package domain

import "time"

type AnswerRefs struct {
	DailySlash    WeaponRef
	Connections   ConnectionRef
	GuessTheNpc   NpcRef
	Hangman       HangmanRef
	TerraTrivia   TerraTriviaRef
	ResetTime     time.Time
	NextResetTime time.Time
}

type WeaponRef struct {
	CurrentWeaponID int
	PrevWeaponID    int
}

type ConnectionRef struct {
	CategoryIDs []int
	Options     []ConnectionOption
}

type ConnectionOptionRef struct {
	Option     string
	CategoryID int
}

type NpcRef struct {
	NpcID       int
	Quote       string
	Name        string
	NameOptions []string
}

type HangmanRef struct {
	EnemyID int
}

type TerraTriviaRef struct {
	QuestionIDs []int
}

type DailyAnswers struct {
	DailySlash    WeaponAnswer
	Connections   ConnectionAnswer
	GuessTheNpc   NpcAnswer
	Hangman       HangmanAnswer
	TerraTrivia   TerraTriviaAnswer
	ResetTime     time.Time
	NextResetTime time.Time
}

type WeaponAnswer struct {
	CurrentWeapon Weapon
	PrevWeapon    Weapon
}

type ConnectionAnswer struct {
	CategoryIDs []int
	Options     []ConnectionOption
}

type ConnectionOption struct {
	Option     string
	CategoryID int
}

type NpcAnswer struct {
	NpcID       int
	Npc         string
	Quote       string
	Name        string
	NameOptions []string
}

type HangmanAnswer struct {
	Enemy Enemy
}

type TerraTriviaAnswer struct {
	Questions []TriviaQuestion
}

type PlayerGuessCounts struct {
	DailySlashCount  int
	ConnectionsCount int
	GuessTheNpcCount int
	HangmanCount     int
	TerraTriviaCount int
}
