package models

import (
	"errors"
	"fmt"
	"terrariadle-backend/internal/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserData struct {
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
	Checks []WeaponChecks `bson:"checks"`
}

type WeaponChecks struct {
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

type UserRepo struct {
	database *db.MongoDB
}

func NewUserRepo(db *db.MongoDB) *UserRepo {
	return &UserRepo{
		database: db,
	}
}

// Tries to get user from db. If user doesn't exist, create a new one.
func GetUser(userId string) (*UserData, error) {

}

// Retrieves the user data or creates a new user if not found
func GetOrCreateUser(userId string) (*UserData, error) {
	user, err := GetUserData(userId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			emptyGame := game{
				Guesses:  []int{},
				HasWon:   false,
				Position: 0,
			}

			newUser := UserData{
				UserID: userId,
				DailySlash: dailySlashGame{
					Game:   emptyGame,
					Checks: []WeaponChecks{},
				},
				Connections: connectionGame{
					Game:     emptyGame,
					Attempts: 4,
				},
				GuessTheNPC: guessTheNpcGame{
					Game:        emptyGame,
					GuessedName: "",
				},
				Hangman: hangmanGame{
					Game:     emptyGame,
					Attempts: 6,
				},
			}
			return &newUser, nil
		}
		return &UserData{}, fmt.Errorf("failed to get user in check API: %s", err)
	}

	return user, nil
}

func UpdateUserData(user *UserData) error {

}

func DeleteUserData() error {

}
