package services

import (
	"fmt"
	"terrariadle-backend/internal/jsonreader"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
)

func GetNpcSearchItems() ([]jsonreader.SearchNpcResult, error) {
	result, ok := store.SearchNpcCache.Get()
	if !ok {
		return []jsonreader.SearchNpcResult{}, fmt.Errorf("failed to get search data for npcs")
	}

	return result, nil
}

// Initializes the daily slash game for a user, returning the puzzle data and guessed weapons
func InitializeNpcGame(userId string) (string, []jsonreader.SearchNpcResult, bool, error) {
	// Gets daily game data and weapons from memstore
	gameData, ok := store.GameData.Get()
	if !ok {
		return "", []jsonreader.SearchNpcResult{}, false, fmt.Errorf("failed to get data from memstore")
	}
	npcs, ok := store.NpcsCache.Get()
	if !ok {
		return "", []jsonreader.SearchNpcResult{}, false, fmt.Errorf("failed to get weapons from memstore")
	}

	// Gets the user from the collection
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return "", []jsonreader.SearchNpcResult{}, false, fmt.Errorf("failed to get user in user guesses API")
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

	var filtered []jsonreader.SearchNpcResult
	for _, id := range user.GuessTheNPC.Game.Guesses {
		if n, ok := npcById[id]; ok {
			filtered = append(filtered, n)
		}
	}

	return gameData.GuessTheNpc.Quote, filtered, user.GuessTheNPC.Game.HasWon, nil
}
