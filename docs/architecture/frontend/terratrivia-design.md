# TerraTrivia

TerraTrivia is inspired by the game Seven Little Words. The player is given 7 Terraria-themed clues and a pool of word chunks. There's no attempt limit. The user keeps building words until all 7 clues are solved.

## Page Load

When the page loads, the `/initialize-game` api called. The game session returned contains the following:

- `trivia_items`: An array of 7 items. Each item has a `clue`, a `letter_count`, and an `answer` (empty string until solved).
- `chunks` — A flat array of word fragments for all unsolved answers, shuffled server-side.
- `finished` — Whether all 7 answers have been found.

The backend always returns exactly 7 trivia items. On load, any previously solved items already have their `answer` populated so the users's progress is restored.

## Submitting a Guess

The player builds a word by clicking up to 4 chunk buttons. Selected chunks are joined together into a single string and displayed in an input box. After the user selects 2 or more chunks, a 500ms debounce timer triggers an automatic submission. This is so the user can't submit a guess with every tile change.

On submit, `/check-guess` is called. The backend performs an exact string match of the guess against all 7 answers. If it matches, the trivia item is marked as solved for the user. If it doesn't match any answer, the response comes back as `is_correct: false` with no penalty.

The response includes:

- `is_correct`: Whether the guess matched an answer.
- `guess_result`: The full `TriviaItem` (with the `answer` now populated), null when `is_correct` is false.
- `finished` — Whether all 7 answers are now solved.

## How the Chunk Pool Works

Each trivia question has a set of pre-defined chunks that spell out its answer.

When a correct guess is made, the frontend removes the used chunks from the local pool.

The chunk grid displays up to 20 buttons at a time. Any remaining slots in the grid are rendered as invisible placeholders to keep the layout stable. Might change this look in the future.

## Win State

The game is won when all 7 questions are solved. The winning card mounts and calls `/winning-data`. This returns the user's position.

A live player count is also streamed to the page via SSE. The count updates in real time as other users finish the game.
