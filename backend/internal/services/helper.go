package services

import (
	"errors"
	"fmt"
	"terrariadle-backend/internal/db"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Helper method to get users
func getUser(col *mongo.Collection, userId string) (*db.GuessDocument, error) {
	user, err := db.GetGuessRecord(col, bson.M{"userId": userId})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			newUser := db.GuessDocument{
				UserID: userId,
				Games: []db.Game{
					{
						GameType: "DailySlash",
						Guesses:  []int{},
						HasWon:   false,
						Position: 0,
						Extra:    map[string]any{},
					},
					{
						GameType: "Connections",
						Guesses:  []int{},
						HasWon:   false,
						Position: 0,
						Extra: map[string]any{
							"NameGuess": "",
						},
					},
					{
						GameType: "GuessTheNpc",
						Guesses:  []int{},
						HasWon:   false,
						Position: 0,
						Extra: map[string]any{
							"Attempts": 4,
						},
					},
					{
						GameType: "Hangman",
						Guesses:  []int{},
						HasWon:   false,
						Position: 0,
						Extra:    map[string]any{},
					}},
			}
			return &newUser, nil
		}
		return &db.GuessDocument{}, fmt.Errorf("failed to get user in check API: %s", err)
	}

	return user, nil
}

// Helper method to update user
func updateUser(col *mongo.Collection, user *db.GuessDocument) error {
	err := db.UpsertRecord(col, bson.M{"userId": user.UserID}, bson.M{"$set": user})
	return err
}

// Checks if a UUID is valid
func isValidUUID(id string) bool {
	if id == "" {
		return false
	}
	_, err := uuid.Parse(id)
	return err == nil
}
