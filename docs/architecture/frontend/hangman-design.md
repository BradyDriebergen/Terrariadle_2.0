# Hangman

This game is pretty self explanitory. The user tries to guess the Terraria enemy letter-by-letter before running out of attempts.

## Page Load

When the page loads, the `/initialize-game` api called. The game session returned contains the following:

- `phrase`: The enemy name as a character array, with unguessed letters replaced by `_`. Spaces and special characters (like hyphens) are always visible.
- `guesses`: The user's guessed letters.
- `attempts`: The user's remaining attempts.
- `finished`: Whether the game is already complete.

The phrase is split into words on the frontend and rendered as a grid of individual letter tiles so the word boundaries are clear.

## Submitting a Guess

The player clicks a letter on the on-screen keyboard. Each button corresponds to one uppercase letter. On click, `/check-guess` is called with the letter as the guess.

The backend validates the guess and returns:

- `phrase`: The updated phrase.
- `guess`: The guessed letter and whether it was correct.
- `attempts`: The updated attempt count.
- `finished`: Whether the game is over.

The keyboard visually marks each letter as correct or incorrect after it's been guessed.

## How the Phrase Updates

The phrase is managed on the backend. After each guess, the backend rebuilds the phrase from scratch and replaces any letter that hasn't been guessed yet with `_`. Special characters are ignored.

The backend checks for a win by scanning the phrase for any remaining `_` characters. If none are found, the game is finished.

## Running Out of Attempts

Each incorrect guess decrements the attempt count. When `attempts` reaches 0, the game is lost.

The Guide character at the top of the page has several states. Each stage represents the remaining attempts. With each lost attempt, a limb is added to the Guide. When attempts hit 0, the page background fades to the Underworld and the Wall of Flesh appears in place of the guide.

## Winning

The game ends when either all letters are revealed or attempts reach 0. When finished, the winning card calls `/winning-data`

This returns the enemy's name and image path so they can be shown on the winning card, as well as the user's position. The winning card appears for both win and loss, but the content and position reflect the actual outcome.

A live player count is also streamed to the page via SSE. The count updates in real time as other users finish the game.
