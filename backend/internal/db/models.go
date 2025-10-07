package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GuessDocument represents the MongoDB document structure
type GuessDocument struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID string             `bson:"userId,omitempty"`
	Games  []Game             `bson:"games"`
}

// Game represents one game entry inside GuessDocument
type Game struct {
	GameType string         `bson:"gameType"`           // required in Mongoose
	Guesses  []string       `bson:"guesses,omitempty"`  // default: []
	HasWon   bool           `bson:"hasWon,omitempty"`   // default: false
	Position int            `bson:"position,omitempty"` // default: -1
	Extra    map[string]any `bson:"extra,omitempty"`    // flexible field for game-specific data
}
