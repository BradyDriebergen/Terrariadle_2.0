package services

import (
	"fmt"
	"slices"
	"terrariadle-backend/internal/jsonreader"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
)

func InitializeConnectionsGame(userId string) (int, []string, []jsonreader.Category, bool, error) {
	gameData, ok := store.GameData.Get()
	if !ok {
		return 0, []string{}, []jsonreader.Category{}, false, fmt.Errorf("failed to get data from memstore")
	}
	// Gets user from database
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return 0, []string{}, []jsonreader.Category{}, false, fmt.Errorf("failed to get user in initialization API")
	}

	guessedCategories := []jsonreader.Category{}
	guessOptions := []string{}
	if user.Connections.Game.HasWon {
		guessedCategories = gameData.Connections
	} else {
		for i := range 4 {
			if slices.Contains(user.Connections.Game.Guesses, gameData.Connections[i].ID) {
				guessedCategories = append(guessedCategories, gameData.Connections[i])
			} else {
				guessOptions = append(guessOptions, gameData.Connections[i].Options...)
			}
		}
	}

	shuffle(guessOptions)

	return user.Connections.Attempts, guessOptions, guessedCategories, user.Connections.Game.HasWon, nil
}

// return oneaway, guessedCategory, won,
func CheckConnectionsGuess(userId string, guess []string) (jsonreader.Category, bool, bool, error) {
	if !isValidUUID(userId) {
		return jsonreader.Category{}, false, false, fmt.Errorf("invalid user ID")
	}
	if len(guess) != 4 {
		return jsonreader.Category{}, false, false, fmt.Errorf("invalid length of guesses, 4 is required")
	}

	// Gets user from database
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return jsonreader.Category{}, false, false, fmt.Errorf("failed to get user in check API")
	}
	// Gets memstore data
	gameData, ok := store.GameData.Get()
	if !ok {
		return jsonreader.Category{}, false, false, fmt.Errorf("failed to get data from memstore")
	}

	guessedCategory := jsonreader.Category{}
	oneAway := false
	for _, category := range gameData.Connections {
		counts := make(map[string]int)

		// Build frequency map from answers
		for _, a := range category.Options {
			counts[a]++
		}

		cgs := 0
		// Check guesses against the map
		for _, g := range guess {
			if counts[g] > 0 {
				cgs++
				counts[g]--
			}
		}

		if cgs == 3 {
			oneAway = true
		} else if cgs == 4 {
			guessedCategory = category
			user.Connections.Game.Guesses = append(user.Connections.Game.Guesses, guessedCategory.ID)

			if len(user.Connections.Game.Guesses) == 4 {
				user.Connections.Game.HasWon = true
				gameData.GuessCounts.ConnectionsCount++
				user.Connections.Game.Position = gameData.GuessCounts.ConnectionsCount
				store.GameData.Set(gameData)
			}

			break
		}
	}

	// Checks if a wrong guess was made
	if guessedCategory.ID == 0 {
		user.Connections.Attempts--
		if user.Connections.Attempts == 0 {
			user.Connections.Game.HasWon = true
		}
	}

	err = models.UpdateUserData(user)
	if err != nil {
		return jsonreader.Category{}, false, false, err
	}

	return guessedCategory, oneAway, user.Connections.Game.HasWon, nil
}

// Gets the winning position and total players for connections
func GetConnectionsWinningData(userId string) (int, int, error) {
	gameData, ok := store.GameData.Get()
	if !ok {
		return 0, 0, fmt.Errorf("failed to get data from memstore")
	}

	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get user in player position API")
	}

	// Gets winning position and player count
	if user.Connections.Game.HasWon {
		return user.Connections.Game.Position, gameData.GuessCounts.ConnectionsCount, nil
	}

	return 0, 0, fmt.Errorf("player doesn't exist")
}
