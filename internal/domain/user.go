package domain

import "time"

type User struct {
	UserID      string
	DailySlash  DailySlashGame
	Connections ConnectionGame
	GuessTheNPC GuessTheNpcGame
	Hangman     HangmanGame
	LastSeen    time.Time
	Dirty       bool
}

type Game struct {
	Guesses  []int // default: []
	Finished bool  // default: false
	Position int   // default: -1
}

type DailySlashGame struct {
	Game   Game
	Checks []WeaponChecks
}

type CompareResult int

const (
	Lower  CompareResult = -1
	Match  CompareResult = 0
	Higher CompareResult = 1
)

type WeaponChecks struct {
	WeaponID   int
	DamageType bool
	Damage     CompareResult
	UseTime    CompareResult
	Rarity     CompareResult
	Operation  bool
	Material   bool
	Obtained   CompareResult
}

type ConnectionGame struct {
	Game     Game
	Attempts int
}

type GuessTheNpcGame struct {
	Game        Game
	GuessedName string
}

type HangmanGame struct {
	Game     Game
	Attempts int
}
