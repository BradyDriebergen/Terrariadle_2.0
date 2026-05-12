package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

type userData struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"userId,omitempty"`
	DailySlash  dailySlashGame     `bson:"dailySlash,omitempty"`
	Connections connectionGame     `bson:"connections,omitempty"`
	GuessTheNPC guessTheNpcGame    `bson:"guessTheNpc,omitempty"`
	Hangman     hangmanGame        `bson:"hangman,omitempty"`
}

type game struct {
	Guesses  []int `bson:"guesses,omitempty"`
	HasWon   bool  `bson:"hasWon,omitempty"`
	Position int   `bson:"position,omitempty"`
}

type dailySlashGame struct {
	Game   game           `bson:"game"`
	Checks []weaponChecks `bson:"checks"`
}

type weaponChecks struct {
	DamageType bool `bson:"damageType"`
	Damage     int  `bson:"damage"`
	UseTime    int  `bson:"useTime"`
	Rarity     int  `bson:"rarity"`
	Operation  bool `bson:"operation"`
	Material   bool `bson:"material"`
	Obtained   int  `bson:"obtained"`
}

type connectionGame struct {
	Game     game `bson:"game"`
	Attempts int  `bson:"attempts"`
}

type guessTheNpcGame struct {
	Game        game   `bson:"game"`
	GuessedName string `bson:"guessedName"`
}

type hangmanGame struct {
	Game     game `bson:"game"`
	Attempts int  `bson:"attempts"`
}
