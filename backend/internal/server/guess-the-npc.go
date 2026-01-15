package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"terrariadle-backend/internal/jsonreader"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
)

type npcInit struct {
	Quote   string
	Guesses []jsonreader.SearchNpcResult
	HasWon  bool
}

// Gets the Guess The NPC data nessasary for the game
func getNpcInitGame(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	initData, err := initNpcGame(userId)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"quote":   initData.Quote,
		"guesses": initData.Guesses,
		"won":     initData.HasWon,
	})
}

func initNpcGame(userId string) (npcInit, error) {
	returnValue := npcInit{
		Quote:   "",
		Guesses: []jsonreader.SearchNpcResult{},
		HasWon:  false,
	}

	// Initial pull from database and cache
	gameData, ok := store.GameData.Get()
	if !ok {
		return returnValue, fmt.Errorf("failed to get game data")
	}
	npcs, ok := store.NpcsCache.Get()
	if !ok {
		return returnValue, fmt.Errorf("failed to get npcs")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return returnValue, fmt.Errorf("failed to get user information")
	}

	// Maps guessed npcs and returns puzzle data
	npcById := make(map[int]jsonreader.SearchNpcResult)
	for _, n := range npcs {
		npcById[n.ID] = jsonreader.SearchNpcResult{
			NpcId: n.ID,
			Name:  n.NPC,
			Path:  n.NPCPath,
		}
	}

	for _, id := range user.GuessTheNPC.Game.Guesses {
		if n, ok := npcById[id]; ok {
			returnValue.Guesses = append(returnValue.Guesses, n)
		}
	}

	returnValue.Quote = gameData.GuessTheNpc.Quote
	returnValue.HasWon = user.GuessTheNPC.Game.HasWon

	return returnValue, nil
}

// Gets the simplified npc list from the database that the users search through
func getNpcSearchItems(w http.ResponseWriter, r *http.Request) {
	searchItems, err := searchableNpcs()
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
	}

	writeJSON(w, http.StatusOK, searchItems)
}

func searchableNpcs() ([]jsonreader.SearchNpcResult, error) {
	result, ok := store.SearchNpcCache.Get()
	if !ok {
		return []jsonreader.SearchNpcResult{}, fmt.Errorf("failed to get search data for npcs")
	}

	return result, nil
}

// Checks Guess The NPC guess and updates the database
func postNpcGuess(w http.ResponseWriter, r *http.Request) {
	var req NpcGuessRequestBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if !isValidUUID(req.UserID) {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid user ID"})
	}

	won, guess, err := checkNpcGuess(req.UserID, req.Guess)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
	}

	writeJSON(w, http.StatusOK, map[string]any{"guess": guess, "won": won})
}

func checkNpcGuess(userId string, npcId int) (bool, jsonreader.SearchNpcResult, error) {
	// Initial pulls from database and cache
	gameData, ok := store.GameData.Get()
	if !ok {
		return false, jsonreader.SearchNpcResult{}, fmt.Errorf("failed to get game data")
	}
	npcs, ok := store.NpcsCache.Get()
	if !ok {
		return false, jsonreader.SearchNpcResult{}, fmt.Errorf("failed to get npcs")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return false, jsonreader.SearchNpcResult{}, fmt.Errorf("failed to get user information")
	}

	// Finds the guessed npc
	guessedNpc := jsonreader.SearchNpcResult{}
	found := false
	for _, n := range npcs {
		if n.ID == npcId {
			guessedNpc = jsonreader.SearchNpcResult{
				NpcId: n.ID,
				Name:  n.NPC,
				Path:  n.NPCPath,
			}
			found = true
			break
		}
	}
	if !found {
		return false, jsonreader.SearchNpcResult{}, fmt.Errorf("npc with id %d not found", npcId)
	}

	// Checks guess
	won := false
	if npcId == gameData.GuessTheNpc.ID {
		won = true
	}

	user.GuessTheNPC.Game.Guesses = append([]int{npcId}, user.GuessTheNPC.Game.Guesses...)

	if won {
		user.GuessTheNPC.Game.HasWon = true
		gameData.GuessCounts.GuessTheNpcCount += 1
		user.GuessTheNPC.Game.Position = gameData.GuessCounts.GuessTheNpcCount
		store.GameData.Set(gameData)
	}

	err = models.UpdateUserData(user)
	if err != nil {
		return false, jsonreader.SearchNpcResult{}, err
	}

	return won, guessedNpc, nil
}

// Gets the position and players guessed numbers
func getNpcWinningData(w http.ResponseWriter, r *http.Request) {

}
