# Guess the NPC

Guess the NPC presents the user with a quote pulled from a Terraria NPC's dialogue. The user is then able to guess until they've guessed the NPC that says the quote. Unlike Daily Slash, there's no attribute comparison. After identifying the NPC, a short bonus round becomes available.

## Page Load

When the page loads, two APIs are called before the page renders:

1. `/initialize-game` returns the user's current game session.
2. `/search-items` returns the full list of searchable weapons for the dropdown.

The session returned from `initialize-game` includes:

- `quote` — The NPC's quote.
- `guessed_ids` — IDs of NPCs the player has already guessed (used for filtering searchable NPCs).
- `guesses` — The list of the user's guesses.
- `finished` — Whether the game is already complete.

## Submitting a Guess

The player picks an NPC from the dropdown. On selection, the frontend calls `/check-guess` with the NPC ID. This guess is checked against the answer in the backend.

The response contains a `finished` flag and a `guess` object. The new guess is prepended to the guess list and the NPC is removed from the dropdown.

There's no attribute comparison here. The guess list is just a record of the NPCs the player has already tried. The game ends as soon as the correct NPC is guessed.

## Winning

When the game is finished, the winning card mounts and calls `/winning-data`. This returns:

- `position` — User's position when they guessed correctly.
- `names` — A list of name options for the bonus round.
- `guessed_name` — The name the user already picked in the bonus round (null if not yet played).
- `correct_name` — The correct answer (only populated after the bonus round has been attempted).

A live player count is also streamed to the page via SSE. The count updates in real time as other users finish the game.

## Bonus Round

After correctly guessing the NPC, a bonus game appears on the winning card. Terraria NPCs can randomly generate with different names, and this round picks a random name from the answer NPC for the user to guess.

The user is shown 4 name options. The user can then guess which name the answer NPC could have. This can only be played once per day. On selection, `/check-name-guess` is called.

The response returns the player's `guessed_name` and the `correct_name`.
