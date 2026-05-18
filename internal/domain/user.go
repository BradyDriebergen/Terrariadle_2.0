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
	HasWon   bool  // default: false
	Position int   // default: -1
}

type DailySlashGame struct {
	Game   Game
	Checks []WeaponChecks
}

// Little checklist for how this works:
/*
	Damage: 0 = lower, 1 = match, 2 = over
	UseTime: 0 = slower, 1 = match, 2 = faster
	Rarity: 0 = earlier, 1 = match, 2 = later
	Obtained: 0 = no match, 1 = partial, 2 = exact
*/
type WeaponChecks struct {
	DamageType bool
	Damage     int
	UseTime    int
	Rarity     int
	Operation  bool
	Material   bool
	Obtained   int
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
