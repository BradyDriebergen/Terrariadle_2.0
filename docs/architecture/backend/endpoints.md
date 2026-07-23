# Terrariadle API Reference

## Common

| Method | Path                  | Params    | Description                |
| ------ | --------------------- | --------- | -------------------------- |
| GET    | `/api/health`         | none      | Health check               |
| GET    | `/api/remaining-time` | none      | Time until daily reset     |
| GET    | `/api/finished-games` | `user_id` | Games a user has completed |

## Daily Slash

| Method | Path                                | Params / Body              | Description                   |
| ------ | ----------------------------------- | -------------------------- | ----------------------------- |
| GET    | `/api/daily-slash/initialize-game`  | `user_id`                  | Start/resume game             |
| GET    | `/api/daily-slash/search-items`     | none                       | Gets searchable weapons       |
| GET    | `/api/daily-slash/hint`             | `hint (int)`               | Get a hint                    |
| POST   | `/api/daily-slash/check-guess`      | `{ user_id, guess (int) }` | Submit guess                  |
| GET    | `/api/daily-slash/winning-data`     | `user_id`                  | Get results                   |
| GET    | `/api/guess-count?mode=daily-slash` | none                       | Live player guess count (SSE) |

## Connections

| Method | Path                                | Params / Body                   | Description                   |
| ------ | ----------------------------------- | ------------------------------- | ----------------------------- |
| GET    | `/api/connections/initialize-game`  | `user_id`                       | Start/resume game             |
| POST   | `/api/connections/check-guess`      | `{ user_id, guess: string[4] }` | Submit guess                  |
| POST   | `/api/connections/reveal-answers`   | `{ user_id }`                   | Reveal correct categories     |
| GET    | `/api/connections/winning-data`     | `user_id`                       | Get results                   |
| GET    | `/api/guess-count?mode=connections` | none                            | Live player guess count (SSE) |

## Guess the NPC

| Method | Path                                  | Params / Body                 | Description                   |
| ------ | ------------------------------------- | ----------------------------- | ----------------------------- |
| GET    | `/api/guess-the-npc/initialize-game`  | `user_id`                     | Start/resume game             |
| GET    | `/api/guess-the-npc/search-items`     | none                          | Get searchable NPCs           |
| POST   | `/api/guess-the-npc/check-guess`      | `{ user_id, guess (int) }`    | Submit guess                  |
| POST   | `/api/guess-the-npc/check-name-guess` | `{ user_id, guess (string) }` | Submit mini-game guess        |
| GET    | `/api/guess-the-npc/winning-data`     | `user_id`                     | Get results                   |
| GET    | `/api/guess-count?mode=guess-the-npc` | none                          | Live player guess count (SSE) |

## Hangman

| Method | Path                            | Params / Body                 | Description                   |
| ------ | ------------------------------- | ----------------------------- | ----------------------------- |
| GET    | `/api/hangman/initialize-game`  | `user_id`                     | Start/resume game             |
| POST   | `/api/hangman/check-guess`      | `{ user_id, guess (letter) }` | Submit guess                  |
| GET    | `/api/hangman/winning-data`     | `user_id`                     | Get results                   |
| GET    | `/api/guess-count?mode=hangman` | none                          | Live player guess count (SSE) |

## TerraTrivia

| Method | Path                                | Params / Body                 | Description                   |
| ------ | ----------------------------------- | ----------------------------- | ----------------------------- |
| GET    | `/api/terratrivia/initialize-game`  | `user_id`                     | Start/resume game             |
| POST   | `/api/terratrivia/check-guess`      | `{ user_id, guess (string) }` | Submit guess                  |
| GET    | `/api/terratrivia/winning-data`     | `user_id`                     | Get results                   |
| GET    | `/api/guess-count?mode=terratrivia` | none                          | Live player guess count (SSE) |
