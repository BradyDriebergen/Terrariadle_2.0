package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"terrariadle-backend/internal/models"
	"terrariadle-backend/internal/store"
	"unicode"
)

var specialChars = []string{"'", ".", "-", "1", " "}

type hangmanInit struct {
	Phrase         []string
	GuessedLetters []string
	Attempts       int
	Finished       bool
}

func getHangmanInitGame(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	initData, err := initHangmanGame(userId)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"phrase":         initData.Phrase,
		"guessedLetters": initData.GuessedLetters,
		"attempts":       initData.Attempts,
		"finished":       initData.Finished,
	})
}

// phrase, guessedletters, attempts, error
func initHangmanGame(userId string) (hangmanInit, error) {
	returnValue := hangmanInit{
		Phrase:         []string{},
		GuessedLetters: []string{},
		Attempts:       0,
		Finished:       false,
	}

	// Initial pull from database and cache
	gameData, ok := store.GameData.Get()
	if !ok {
		return returnValue, fmt.Errorf("failed to get game data")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return returnValue, fmt.Errorf("failed to get user information")
	}

	// Initialized guessed letters from ids to actual letters
	for _, id := range user.Hangman.Game.Guesses {
		returnValue.GuessedLetters = append(returnValue.GuessedLetters, string(rune(id)))
	}

	// Create array of characters representing the enemy name
	returnValue.Phrase = strings.Split(gameData.Hangman.Name, "")

	// Used guessed letters to filter out the name and add _s
	for i, letter := range returnValue.Phrase {
		if !slices.Contains(returnValue.GuessedLetters, letter) && !slices.Contains(specialChars, letter) {
			returnValue.Phrase[i] = "_"
		}
	}

	// Return attempts and has won
	returnValue.Attempts = user.Hangman.Attempts
	returnValue.Finished = user.Hangman.Game.HasWon

	return returnValue, nil
}

type hangmanGuessReqBody struct {
	UserID string `json:"userId"`
	Guess  string `json:"guess"`
}

type hangmanCheck struct {
	Phrase   []string
	Correct  bool
	Finished bool
}

func postHangmanGuess(w http.ResponseWriter, r *http.Request) {
	var req hangmanGuessReqBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if !isValidUUID(req.UserID) {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid user ID"})
	}

	checkData, err := checkHangmanGuess(req.UserID, req.Guess)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"newPhrase": checkData.Phrase,
		"correct":   checkData.Correct,
		"finished":  checkData.Finished,
	})
}

func checkHangmanGuess(userId string, guess string) (hangmanCheck, error) {
	returnValue := hangmanCheck{
		Phrase:   []string{},
		Correct:  false,
		Finished: false,
	}

	// Initial checks of the guesses
	if len(guess) != 1 {
		return returnValue, fmt.Errorf("guess must be a single letter")
	}

	character := []rune(guess)

	if !unicode.IsLetter(character[0]) {
		return returnValue, fmt.Errorf("guess must be a letter")
	}
	if !unicode.IsUpper(character[0]) {
		return returnValue, fmt.Errorf("guess must capitalized")
	}

	// Initial pulls from database and cache
	gameData, ok := store.GameData.Get()
	if !ok {
		return returnValue, fmt.Errorf("failed to get game data")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return returnValue, fmt.Errorf("failed to get user information")
	}

	// Adds guess to user's model
	user.Hangman.Game.Guesses = append(user.Hangman.Game.Guesses, int(character[0]))

	// Splits the enemy's name
	newPhrase := strings.Split(gameData.Hangman.Name, "")

	correctGuess := false

	if !slices.Contains(newPhrase, guess) {
		user.Hangman.Attempts--
		if user.Hangman.Attempts <= 0 {
			user.Hangman.Game.HasWon = true
		}
	} else {
		correctGuess = true
	}

	guessedLetters := []string{}
	for _, id := range user.Hangman.Game.Guesses {
		guessedLetters = append(guessedLetters, string(rune(id)))
	}

	// Used guessed letters to filter out the name and add _s
	for i, letter := range newPhrase {
		if !slices.Contains(guessedLetters, letter) && !slices.Contains(specialChars, letter) {
			newPhrase[i] = "_"
		}
	}

	if !slices.Contains(newPhrase, "_") {
		user.Hangman.Game.HasWon = true
		gameData.GuessCounts.HangmanCount++
		user.Hangman.Game.Position = gameData.GuessCounts.HangmanCount
		store.GameData.Set(gameData)
	}

	returnValue.Phrase = newPhrase
	returnValue.Finished = user.Hangman.Game.HasWon
	returnValue.Correct = correctGuess

	err = models.UpdateUserData(user)
	if err != nil {
		return returnValue, err
	}

	return returnValue, nil
}

type hangmanWinData struct {
	Count    int
	Position int
	Name     string
	Path     string
}

func getHangmanWinningData(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")

	winData, err := hangmanWinningData(userId)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"pos":   winData.Position,
		"count": winData.Count,
		"name":  winData.Name,
		"path":  winData.Path,
	})
}

func hangmanWinningData(userId string) (hangmanWinData, error) {
	returnValue := hangmanWinData{
		Count:    0,
		Position: 0,
		Name:     "",
		Path:     "",
	}

	// Initial pull from database and cache
	gameData, ok := store.GameData.Get()
	if !ok {
		return returnValue, fmt.Errorf("failed to get game data")
	}
	user, err := models.GetOrCreateUser(userId)
	if err != nil {
		return returnValue, fmt.Errorf("failed to get user information")
	}

	// Gets winning position and player count
	if user.Hangman.Game.HasWon {
		returnValue.Count = gameData.GuessCounts.HangmanCount
		returnValue.Position = user.Hangman.Game.Position
		returnValue.Name = gameData.Hangman.Name
		returnValue.Path = gameData.Hangman.ImagePath

		return returnValue, nil
	}

	return returnValue, fmt.Errorf("player doesn't exist")
}
