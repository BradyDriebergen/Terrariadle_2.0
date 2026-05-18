package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

type userData struct {
	id          primitive.ObjectID `bson:"_id,omitempty"`
	userID      string             `bson:"userId,omitempty"`
	dailySlash  dailySlashGame     `bson:"dailySlash,omitempty"`
	connections connectionGame     `bson:"connections,omitempty"`
	guessTheNPC guessTheNpcGame    `bson:"guessTheNpc,omitempty"`
	hangman     hangmanGame        `bson:"hangman,omitempty"`
}

type game struct {
	guesses  []int `bson:"guesses,omitempty"`
	hasWon   bool  `bson:"hasWon,omitempty"`
	position int   `bson:"position,omitempty"`
}

type dailySlashGame struct {
	game   game           `bson:"game"`
	checks []weaponChecks `bson:"checks"`
}

type weaponChecks struct {
	damageType bool `bson:"damageType"`
	damage     int  `bson:"damage"`
	useTime    int  `bson:"useTime"`
	rarity     int  `bson:"rarity"`
	operation  bool `bson:"operation"`
	material   bool `bson:"material"`
	obtained   int  `bson:"obtained"`
}

type connectionGame struct {
	game     game `bson:"game"`
	attempts int  `bson:"attempts"`
}

type guessTheNpcGame struct {
	game        game   `bson:"game"`
	guessedName string `bson:"guessedName"`
}

type hangmanGame struct {
	game     game `bson:"game"`
	attempts int  `bson:"attempts"`
}
