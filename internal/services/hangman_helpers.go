package services

import (
	"slices"
	"strings"
	"terrariadle-backend/internal/domain"
	"unicode"
)

func buildHangmanPhrase(name string, guesses []int) (phrase []string, guessedLetters []HangmanGuess) {
	guessedLetters = []HangmanGuess{}
	rawLetters := []string{}
	phrase = strings.Split(name, "")

	// Hydrate int guesses into characters
	for _, c := range guesses {
		letter := string(rune(c))

		rawLetters = append(rawLetters, letter)
		guessedLetters = append(guessedLetters, HangmanGuess{
			Letter:  letter,
			Correct: slices.Contains(phrase, letter),
		})
	}

	for i, c := range phrase {
		// Removes any unguessed letters
		if !slices.Contains(rawLetters, c) && !slices.Contains(specialChars, c) {
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
