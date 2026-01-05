package models

import (
	"context"
	"errors"
	"fmt"
	"terrariadle-backend/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserData represents the MongoDB document structure
type UserData struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"userId,omitempty"`
	DailySlash  dailySlashGame     `bson:"dailySlash,omitempty"`
	Connections connectionGame     `bson:"connections,omitempty"`
}

// Game represents one game entry inside GuessDocument
type game struct {
	Guesses  []int `bson:"guesses,omitempty"`  // default: []
	HasWon   bool  `bson:"hasWon,omitempty"`   // default: false
	Position int   `bson:"position,omitempty"` // default: -1
}

type dailySlashGame struct {
	Game   game           `bson:"game"`
	Checks []WeaponChecks `bson:"checks"`
}

// Little checklist for how this works:
/*
	Damage: 0 = lower, 1 = match, 2 = over
	UseTime: 0 = slower, 1 = match, 2 = faster
	Rarity: 0 = earlier, 1 = match, 2 = later
	Obtained: 0 = no match, 1 = partial, 2 = exact
*/
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

// Retrieves the user data
func GetUserData(userId string) (*UserData, error) {
	collection := db.GetCollection("terrariadle", "user_guesses")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user UserData
	res := collection.FindOne(ctx, bson.M{"userId": userId})
	err := res.Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, mongo.ErrNoDocuments
		}
		return nil, fmt.Errorf("failed to find record %v", err)
	}

	return &user, nil
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
			}
			return &newUser, nil
		}
		return &UserData{}, fmt.Errorf("failed to get user in check API: %s", err)
	}

	return user, nil
}

// Updates or creates the user data in the database
func UpdateUserData(user *UserData) error {
	collection := db.GetCollection("terrariadle", "user_guesses")
	err := db.UpsertRecord(collection, bson.M{"userId": user.UserID}, bson.M{"$set": user})
	return err
}

func DeleteUserData() error {
	collection := db.GetCollection("terrariadle", "user_guesses")
	err := db.DropCollection(collection)
	return err
}
