# Terrariadle Go Packages

## `api`

The HTTP layer. This is where the `Server` lives, routes are registered, and requests are handled. It's the only package that touches Go's `HTTP` package.

`NewServer` wires all the service dependencies into a single struct and registers the api routes. Routes are organized into separate files per game (`daily_slash.go`, `connections.go`, etc.) and registered in `routes.go`. Each handler follows the same pattern: parse the request, call a service method, and write the result.

```go
func (s *Server) checkDailySlashGuess(w http.ResponseWriter, r *http.Request) {
    var req DailySlashGuessRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil { ... }

    result, err := s.dailySlash.CheckGuess(r.Context(), req.UserID, req.Guess)
    if err != nil {
        handleError(w, err)
        return
    }
    writeJSON(w, http.StatusOK, result)
}
```

Error handling is centralized in `helpers.go`. `handleError` reads `domain.AppError` values and maps their error codes to the assocaited HTTP status.

The `guessCountStream` handler in `common.go` manages the SSE connection for live player counts. It subscribes to the broker, writes an initial count, and then streams updates as they arrive until the request context is cancelled.

The `spaHandler` serves the embedded SvelteKit static files and falls back to `index.html` for any path that doesn't match a file. Svelte takes care of the routing from there.

## `db`

Generic MongoDB read/write functionality. This is the only package that imports the `MongoDB` driver. This package is responsible for directly accessing the database.

It exposes a small set of generic functions that cover everything the app needs:

```go
FindOne[T](ctx, db, collectionName, filter)  // fetches a single document
GetAll[T](ctx, db, collectionName)           // fetches all documents
Upsert(ctx, db, collectionName, filter, doc) // creates or update
DeleteAll(ctx, db, collectionName)           // wipes a collection
```

The `Filter` type is just a `map[string]any`. The generic type parameter on `FindOne` and `GetAll` allow function calls to return any shape of data to map to a struct from the database.

`Connect` establishes the MongoDB client and pings the server before returning, so a connection failure is caught immediately at startup.

## `domain`

Shared types and interfaces used across all packages. `domain` has no imports of its own, so it can be used anywhere without creating circular imports. Below are some examples of shared code:

- **Game types**: `User`, `Weapon`, `WeaponChecks`, `Category`, `Npc`, `Enemy`, `TriviaQuestion`
- **Answer types**: `DailyAnswers`, `PlayerGuessCounts`, and the per-game answer structs
- **`AppError`**: The error type that flows from `services` up through `api`
- **`Broker`**: A pub/sub implementation for the SSE player count stream
- **Time utilities**: `NextMidnight` and `TimeUntilNextMidnight`, used by the puzzle refresh job to calculate when to run and the remaining time API.

## `jobs`

Background goroutines handle reseting puzzles and updating users:

**`PuzzleRefreshJob`** fires once at midnight every day. It picks new random answers for all five games, resets the player guess counts, and drops all **user** data from the database. On startup, it checks whether the last reset time has already passed. If the server was down at midnight, it runs the refresh immediately to catch up rather than serving yesterday's puzzles.

**`StartFlushJob`** runs on a 30-second timer. Every second, it calls `FlushDirty` to write any modified users back to MongoDB, then `EvictStale` to remove users from the in-memory cache who haven't been active in over an hour. A final flush also runs on shutdown to make sure no writes are lost when the server stops.

Both jobs receive a context from `main.go`. When the context is cancelled (on shutdown), the jobs exit their loops.

## `repo`

App-specific MongoDB queries built on top of `db`. The `repo` layer knows what collections exist and what the MongoDB document shapes look like. It translates between those raw document types and the `domain` types that the rest of the app uses.

There are three repos:

- **`UserRepo`**: reads and writes user game data from the `user_data` collection.

- **`CatalogRepo`**: reads all game content from five separate collections. Each method calls `db.GetAll[T]`, then converts the result to domain types. This is used for storing all the game data in memory, rather than calling the database anytime we need it.

- **`AnswerRepo`**: reads and writes two singleton documents: the daily puzzle answers and the player guess counts.

Each repo is defined as an interface, which lets `store` depend on the interface rather than the concrete MongoDB implementation. This keeps the layers clean and testable.

## `services`

This is where each game's logic lives. There are five game services plus a `Common` service that handles the SSE stream and the user's finished-game status.

Each service receives `store` interfaces as its dependencies and operates entirely on domain types. None of the services know about HTTP or MongoDB.

The pattern is consistent across all five games:

- **`InitializeGame`**: loads the user's existing game state from the user store and rebuilds it into a response. For games like Daily Slash, this means hydrating stored weapon IDs back into full `Weapon` objects using the catalog store.
- **`CheckGuess`**: validates the guess, compares it against the answer in the answer store, updates the user, and returns the result.
- **`GetWinningData`**: fetches data for the post-game winning card after the game is finished.

Some services have their own logic. For example, Connections supports an API for revealing answers when attempts are 0. Guess the NPC supports a mini game alongside its main game.

All comparison logic lives in the `..._helpers.go` files alongside each service. For example, `generateWeaponChecks` in `daily_slash_helpers.go` handles the full attribute comparison for Daily Slash, and `buildHangmanPhrase` in `hangman_helpers.go` rebuilds the masked phrase from the raw guess history.

## `store`

In-memory cache layer. The store sits between `services` and `repo`, holding frequently-read data in memory so the app doesn't hit MongoDB on every request. There are four stores:

**`CachedCatalogStore`** loads all game content (weapons, categories, NPCs, enemies, trivia questions) from the catalog repo at startup and stores everything in `map[int]T` caches for O(1) lookups by ID. The catalog never changes, so there's no write path and no mutex needed.

```go
weapon, ok := catalogStore.GetWeapon(id)
```

**`CachedAnswerStore`** holds the current day's puzzle answers. The `DailyAnswers` struct contains what today's correct answer is. `GetAnswers()` returns from cache under a read lock. `UpsertAnswers` writes to MongoDB first, then updates the cache under a write lock.

**`CachedUserStore`** holds active users in a `map[string]domain.User`. On a read, it checks the cache first and only falls back to the repo if the user isn't there. On a write (`UpsertUser`), it writes to MongoDB and updates the cache atomically. `FlushDirty` and `EvictStale` are called by the flush job. Dirty users are persisted to MongoDB, and users who haven't been seen in over an hour are evicted from memory.

**`CachedGuessCountsStore`** tracks how many players have won each game today. Each `Increment*` method is fully atomic: it acquires the write lock, increments the counter, persists to MongoDB, and then publishes to the broker so connected SSE clients receive the updated count in real time. If the DB write fails, the in-memory count is rolled back.

## `web`

A single file that uses Go's `//go:embed` directive to bundle the compiled SvelteKit frontend into the binary at build time.

```go
//go:embed all:build
var embeddedFiles embed.FS

func Assets() fs.FS {
    sub, _ := fs.Sub(embeddedFiles, "build")
    return sub
}
```

`Assets()` returns an `fs.FS` pointing at the `build/` directory inside the embedded filesystem. This is passed directly into `api.NewServer`, which uses it to serve static files through the `spaHandler`. The result is that the entire frontend ships inside the Go binary, so only one process needs to run on the server.
