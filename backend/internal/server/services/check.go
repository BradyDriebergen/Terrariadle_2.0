package services

import (
	"errors"
	"fmt"
	"terrariadle-backend/internal/db"
	"terrariadle-backend/internal/types"
	"terrariadle-backend/internal/utils/cache"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckDailySlashGuess(userId string, weaponId int) (bool, types.Weapon, error) {
	// Check if weaponId is valid
	// Check if userId is valid

	col := db.GetCollection("terrariadle", "user_guesses")
	user, err := getUser(col, userId)
	if err != nil {
		return false, types.Weapon{}, fmt.Errorf("failed to get user in check API")
	}

	gameData := cache.GetGameData()

	won := false
	if weaponId == gameData.DailySlash.CurrentWeapon.ID {
		won = true
	}

	for i := range user.Games {
		if user.Games[i].GameType == "DailySlash" {
			user.Games[i].Guesses = append(user.Games[i].Guesses, weaponId)
			if won {
				user.Games[i].HasWon = true
				gameData.GuessCounts.DailySlashCount += 1
				user.Games[i].Position = gameData.GuessCounts.DailySlashCount
				cache.SetGameData(gameData)
			}
			break
		}
	}
	err = updateUser(col, user)
	if err != nil {
		return false, types.Weapon{}, err
	}

	fmt.Printf("Guess count: %d\n", gameData.GuessCounts.DailySlashCount)

	// Get guessed weapon from database and return it here
	return won, types.Weapon{}, nil
}

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

func updateUser(col *mongo.Collection, user *db.GuessDocument) error {
	err := db.UpsertRecord(col, bson.M{"userId": user.UserID}, bson.M{"$set": user})
	return err
}
