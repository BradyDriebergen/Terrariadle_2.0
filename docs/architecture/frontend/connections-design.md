# Connections

Connections is a grouping puzzle based on the NYT game. The user is shown 16 options. To win, they need to find the 4 groups of 4 that share a common Terraria-themed category. The user only has 4 attempts until they lose the game.

## Page Load

When the page loads, the `/initialize-game` api called. The game session returned contains the following:

- `options` — The shuffled remaining unsolved options.
- `solved_categories` — Categories the user has guessed correctly.
- `attempts` — The user's remaining attempts.
- `finished` — Whether the game is already complete.

Already solved categories get sent separately so they can be rendered as banners at the top of the grid. The remaining options populate the interactive grid.

## Submitting a Guess

The player selects exactly 4 options from the grid and clicks "Check Connection". The submit button is disabled until only 4 items are selected. On submit, the frontend calls `/check-guess`. The response includes:

- `is_correct`: Whether the guess is correct.
- `one_away` — Whether 3 of the 4 guessed options were in the same category.
- `correct_guess` — The solved category object if`is_correct` is true. Else, returns null.
- `attempts` — The user's attempt count.
- `finished` — Whether the game is over.

## How the Check Works

The backend builds a map of each option to its category ID. For each guessed option it looks up the category and counts how many guessed options share the same one. When the count for any single category hits 4, the guess is correct. When it hits 3, `one_away` is set to true.

A wrong guess costs one attempt. Once all 4 categories are solved, the game is marked as finished and the user's finishing position is recorded.

## Running Out of Attempts

When the last attempt is exhausted without solving all categories, the frontend calls `/reveal-answers`. This is only available when attempts have reached 0. The backend solves the user's remaining unsolved categories and sends them back to the client.

## Winning

The game is won by correctly guessing all 4 categories before running out of attempts. When finished, the winning card calls `/winning-data` This returns the user's position. The winning card only displays if the player solved all 4 without running out of attempts.

A live player count is also streamed to the page via SSE. The count updates in real time as other users finish the game.
