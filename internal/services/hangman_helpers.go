package services

import (
	"slices"
	"strings"
	"terrariadle-backend/internal/domain"
	"unicode"
)

func buildHangmanPhrase(name string, guesses []int) (phrase []string, guessedLetters []string) {
	guessedLetters = []string{}
	phrase = strings.Split(name, "")

	// Hydrate int guesses into characters
	for _, c := range guesses {
		guessedLetters = append(guessedLetters, string(rune(c)))
	}

	for i, c := range phrase {
		// Removes any unguessed letters
		if !slices.Contains(guessedLetters, c) && !slices.Contains(specialChars, c) {
			phrase[i] = "_"
		}
	}

	return phrase, guessedLetters
}

func validateGuessedLetter(guess string) error {
	if len(guess) != 1 {
		return domain.InvalidInput("Guess must be a single letter", nil)
	}
	if !unicode.IsLetter(rune(guess[0])) {
		return domain.InvalidInput("Guess must be a letter", nil)
	}
	if !unicode.IsUpper(rune(guess[0])) {
		return domain.InvalidInput("Guess must be capitalized", nil)
	}

	return nil
}
