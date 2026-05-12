package domain

type User struct {
	UserID      string
	DailySlash  dailySlashGame
	Connections connectionGame
	GuessTheNPC guessTheNpcGame
	Hangman     hangmanGame
}

type game struct {
	Guesses  []int // default: []
	HasWon   bool  // default: false
	Position int   // default: -1
}

type dailySlashGame struct {
	Game   game
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

type connectionGame struct {
	Game     game
	Attempts int
}

type guessTheNpcGame struct {
	Game        game
	GuessedName string
}

type hangmanGame struct {
	Game     game
	Attempts int
}
