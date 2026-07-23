# Daily Slash

Daily Slash is the Wordle-style game in this application. The goal of the game is to identify a random Terraria weapon by submitting guesses. When a guess is made, it compares the stats of the guessed weapon with the answer. There's no limit on the number of guesses you can make. The game is won when the daily weapon is guessed.

## Page Load

When the page loads, two APIs are called before the page renders:

1. `/initialize-game` returns the user's current game session.
2. `/search-items` returns the full list of searchable weapons for the dropdown.

The session returned from `initialize-game` includes:

- `previous_weapon`: A preview of yesterday's weapon.
- `guessed_ids`: The IDs of weapons the user has already guessed (used for filtering searchable weapons).
- `guesses`: The list of the user's guesses.
- `finished`: Whether the game is already complete.

## Submitting a Guess

The user picks a weapon from the searchable dropdown. On selection, the frontend calls `/check-guess` with the weapon's ID. This guess is checked against the answer in the backend.

The backend then responses containing a `finished` flag and a `guess_result` object with the guessed weapon data and its comparison checks. The new guess is prepended to the list of guessed weapons.

## How Comparison Checks Work

The backend compares the guessed weapon against the answer across 8 attributes. The result for each attribute is attached to the guess and displayed in the guess list:

| Attribute     | Type           | Logic                                                |
| ------------- | -------------- | ---------------------------------------------------- |
| `damage_type` | Boolean        | Exact match                                          |
| `damage`      | Numeric        | `Higher` / `Lower` / `Match`                         |
| `use_time`    | Ordered string | Ranked scale from `Snail` (0) to `Insanely Fast` (7) |
| `rarity`      | Ordered string | Ranked from `White` (0) to `Red` (10)                |
| `operation`   | Boolean        | Exact match (pre-Hardmode vs Hardmode)               |
| `material`    | Boolean        | Exact match                                          |
| `obtained`    | String array   | `Match` / `PartialMatch` / `NoMatch`                 |

`Higher` and `Lower` always refer to the direction of the answer relative to the guess. A `Higher` result means the answer's value is higher than what was guessed.

## Hints

Three hints are available and unlock progressively as the guess count grows:

| Hint | Unlocks After | Reveals                       |
| ---- | ------------- | ----------------------------- |
| 1    | 4 guesses     | Mode Obtained (Pre/Hardmode)  |
| 2    | 7 guesses     | Weapon Type                   |
| 3    | 12 guesses    | A blurred image of the weapon |

Once a hint has been fetched, the result is cached client-side. Clicking the button again just toggles the display without making another API call.

## Winning

The game ends when the user guesses the correct weapon. When finished, the winning card fetches: `/winning-data`. This returns the user's position.

A live player count is also streamed to the page via SSE. The count updates in real time as other users finish the game.
