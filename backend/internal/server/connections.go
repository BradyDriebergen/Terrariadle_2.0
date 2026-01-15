package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"terrariadle-backend/internal/jsonreader"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
)

type connectionsInit struct {
	Attepts    int
	HasWon     bool
	Categories []jsonreader.Category
	Options    []string
}

// Gets the Connections data nessasary for the game
func getConnectionInitData(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	initData, err := initConnections(userId)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"attempts": initData.Attepts,
		"options":  initData.Options,
		"guesses":  initData.Categories,
		"finished": initData.HasWon,
	})
}

func initConnections(userId string) (connectionsInit, error) {
	returnData := connectionsInit{
		Attepts:    0,
		Categories: []jsonreader.Category{},
		Options:    []string{},
		HasWon:     false,
	}

	// Initial pulls from database and cache
	gameData, ok := store.GameData.Get()
	if !ok {
		return returnData, fmt.Errorf("failed to get game data")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return returnData, fmt.Errorf("failed to get user information")
	}

	// If user won, send all the right answers, if not, send correctly guessed categories
	if user.Connections.Game.HasWon {
		returnData.Categories = gameData.Connections
	} else {
		for i := range 4 {
			if slices.Contains(user.Connections.Game.Guesses, gameData.Connections[i].ID) {
				returnData.Categories = append(returnData.Categories, gameData.Connections[i])
			} else {
				returnData.Options = append(returnData.Options, gameData.Connections[i].Options...)
			}
		}
	}

	shuffle(returnData.Options)

	returnData.Attepts = user.Connections.Attempts
	returnData.HasWon = user.Connections.Game.HasWon

	return returnData, nil
}

type connectionsGuessRequestBody struct {
	UserID string   `json:"userId"`
	Guess  []string `json:"guess"`
}

type connectionsCheck struct {
	Guess    jsonreader.Category
	OneAway  bool
	GameOver bool
}

// Checks Connections guess and updates the database
func postConnectionsGuess(w http.ResponseWriter, r *http.Request) {
	var req connectionsGuessRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Initial checks
	if !isValidUUID(req.UserID) {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid user ID"})
		return
	}
	if len(req.Guess) != 4 {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid length of guess, 4 is required"})
		return
	}

	guessData, err := checkConnectionsGuess(req.UserID, req.Guess)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"guess":    guessData.Guess,
		"oneAway":  guessData.OneAway,
		"finished": guessData.GameOver,
	})
}

func checkConnectionsGuess(userId string, guess []string) (connectionsCheck, error) {
	returnData := connectionsCheck{
		Guess:    jsonreader.Category{},
		OneAway:  false,
		GameOver: false,
	}

	// Initial pull from database and cache
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return returnData, fmt.Errorf("failed to get user information")
	}
	gameData, ok := store.GameData.Get()
	if !ok {
		return returnData, fmt.Errorf("failed to get game data")
	}

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
			returnData.OneAway = true
		} else if cgs == 4 {
			returnData.Guess = category
			user.Connections.Game.Guesses = append(user.Connections.Game.Guesses, returnData.Guess.ID)

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
	if returnData.Guess.ID == 0 {
		user.Connections.Attempts--
		if user.Connections.Attempts == 0 {
			user.Connections.Game.HasWon = true
		}
	}

	err = models.UpdateUserData(user)
	if err != nil {
		return returnData, fmt.Errorf("failed to update user")
	}

	returnData.GameOver = user.Connections.Game.HasWon

	return returnData, nil
}

// Gets the position and players guessed numbers
func getConnectionsWinningData(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	pos, count, err := connectionsWinningData(userId)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"pos":   pos,
		"count": count,
	})
}

func connectionsWinningData(userId string) (int, int, error) {
	// Initial pull from database and cache
	gameData, ok := store.GameData.Get()
	if !ok {
		return 0, 0, fmt.Errorf("failed to get game data")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get user information")
	}

	// Gets winning position and player count
	if user.Connections.Game.HasWon {
		return user.Connections.Game.Position, gameData.GuessCounts.ConnectionsCount, nil
	}

	return 0, 0, fmt.Errorf("player doesn't exist")
}
