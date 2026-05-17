package domain

import "time"

type DailyAnswers struct {
	DailySlash    WeaponAnswer
	Connections   []ConnectionAnswer
	GuessTheNpc   NpcAnswer
	Hangman       HangmanAnswer
	GuessCounts   PlayerGuessCounts
	ResetTime     time.Time
	NextResetTime time.Time
}

type WeaponAnswer struct {
	CurrentWeapon Weapon
	PrevWeapon    Weapon
}

type ConnectionAnswer struct {
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

type PlayerGuessCounts struct {
	DailySlashCount  int
	ConnectionsCount int
	GuessTheNpcCount int
	HangmanCount     int
}
