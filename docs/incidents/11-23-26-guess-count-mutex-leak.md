# Code Review: Winning-Guess Hang, Leaked Mutex in Guess Count Store

| Field           | Value                                                                       |
| --------------- | --------------------------------------------------------------------------- |
| **Review date** | 11/23/26                                                                    |
| **Area**        | `internal/store/guess_counts_store.go` (all `Increment*Count` methods)      |
| **Symptom**     | Wrong guesses work. Correct guesses never get a response in any game mode.  |
| **Verdict**     | **Root cause found in code.** A single DB error leaves `s.mu` locked forever. |
| **Related**     | [11-5-26 winning-guess hang](11-5-26-winning-guess-hang.md), [11-9-26 dropped Mongo connection](11-9-26-dropped-mongo-connection.md) |

---

## 1. Summary

Every `Increment*Count` method takes `s.mu.Lock()` and then **returns on the error path without calling `s.mu.Unlock()`**:

```go
// internal/store/guess_counts_store.go:67-88 (same shape at :90, :113, :136, :159)
func (s *CachedGuessCountsStore) IncrementDailySlashCount(ctx context.Context) (int, error) {
	s.mu.Lock()

	s.guessCountsCache.DailySlashCount++

	gc := fromGuessCountDomain(s.guessCountsCache)
	if err := s.answerRepo.UpsertGuessCounts(ctx, &gc); err != nil {
		s.guessCountsCache.DailySlashCount-- // roll back on failure
		return 0, fmt.Errorf("increment-daily-slash-count: %w", err) // <-- s.mu is never released
	}

	count := s.guessCountsCache.DailySlashCount
	s.mu.Unlock()
	...
}
```

After **one** failed `UpsertGuessCounts` call, the lock is held for the life of the process. Every later winning guess, in every game mode, blocks forever on `s.mu.Lock()`. `sync.Mutex` ignores context cancellation, so these goroutines never return a response or an error. That matches the browser behavior: the request is sent and the response never comes.

This bug has been in the code since the file was created (`6a09d4b`, 6/4/26). It stays hidden until the first MongoDB write error, so it can go weeks without showing up.

> **Correction to the 11/5 incident doc, §4.4:** That section says both the error path and the success path call `s.mu.Unlock()`. They don't. Only the success path unlocks. That review ruled out the actual cause.

## 2. Why the Earlier Fixes Did Not Stop It

The 11/5 and 11/9 fixes (`SetTimeout(5s)`, idle-connection pruning, pool sizing, TCP keepalive) make MongoDB calls **fail faster**. They don't make them **never fail**. With this bug, a fast failure does as much damage as a slow one:

| Before 11/5 fix | After 11/5 fix |
| --- | --- |
| A stalled write blocks until the client disconnects. `r.Context()` is then canceled, the upsert returns `context canceled`, and **the lock leaks**. | Any write slower than 5s (Atlas failover or maintenance, a cold TLS handshake, OCI network blip) returns `context deadline exceeded`, and **the lock leaks**. |

The timeout made this failure mode **more likely**, not less. Any transient Atlas or network error is now enough to cause an outage that lasts until a restart. That also fits the local instance working fine: a fresh process has no leaked lock.

## 3. Why Only Winning Guesses Hang

- Wrong guesses call `userCache.GetUser` and `userCache.UpsertUser`. Neither touches the guess-count mutex.
- Only a correct guess calls `guessCountCache.Increment*Count`.
- **Diagnostic note:** `UpsertUser` (`internal/store/user_store.go:84`) writes to MongoDB **synchronously**. It is not write-behind, as the 11/5 doc assumed. If wrong guesses still succeed, MongoDB is reachable and healthy **right now**. The hang is process-local lock state, not a live connectivity problem.

## 4. Blast Radius Once the Lock Leaks

| Caller | Effect |
| --- | --- |
| `Increment*Count` (all 5 game modes) | Blocks forever, so winning guesses hang. |
| `GetGuessCounts` (`RLock`) used by the SSE stream (`internal/services/common.go:32-40`) | Blocks forever, so live player counts stop loading. |
| `ResetGuessCounts` at midnight (`internal/jobs/puzzle_refresh.go:72`) | Blocks forever. The refresh goroutine is stuck, so **`DropAllUsers` never runs**, and **no future midnight refresh ever runs**. The next day's puzzle may already be written, but users and counts are never reset. |

This explains the "0 players guessed today" observation in the 11/5 incident. The count never moved because every increment was stuck behind the lock.

## 5. Other Race Conditions Found

These are lower priority. None causes the hang on its own.

1. **Unlocked read on return.** Each `Increment*Count` returns `s.guessCountsCache.XCount` *after* `s.mu.Unlock()`. This is a data race (`go test -race` would flag it). Under concurrent wins, a user can be told the wrong finishing position. Return the local `count` instead.
2. **Lost update / read-modify-write in `CheckGuess`.** `GetUser` copies the user, the copy is changed, and then `UpsertUser` writes it back, with no per-user lock. Two concurrent requests from the same user (double-click, two tabs) can both pass the `Finished` / `Contains` checks. The result can be a double increment of the win count or a lost guess.
3. **Full-document `$set` for counts.** `UpsertGuessCounts` overwrites all five counts with in-memory values. If more than one instance ever runs, or a reset races with an increment, counts get clobbered. Use `$inc` on a single field.
4. **Mutex held across network I/O.** `Increment*Count`, `ResetGuessCounts`, `UpsertAnswers`, and `FlushDirty` all hold a lock while calling MongoDB. With the 5s timeout this is bounded, but it serializes all winners behind DB latency. `FlushDirty` also blocks all user-cache reads for up to (dirty users × 5s).
5. **Errors silently dropped in the refresh job.** `puzzle_refresh.go:62-74` ignores the errors from `UpsertAnswers`, `ResetGuessCounts`, and `DropAllUsers`.
6. **No timeouts on either end.** `http.Server` has no `WriteTimeout`, and the client `fetch` in `frontend/src/lib/api/daily-slash.ts:48` has no `AbortController`. A server-side hang therefore looks like an endless spinner instead of an error.

## 6. Recommended Fix

**Primary fix (required).** Release the lock on every path with `defer`, and do the publish after the critical section:

```go
func (s *CachedGuessCountsStore) IncrementDailySlashCount(ctx context.Context) (int, error) {
	count, err := s.increment(ctx, func(c *domain.PlayerGuessCounts) *int { return &c.DailySlashCount })
	if err != nil {
		return 0, fmt.Errorf("increment-daily-slash-count: %w", err)
	}
	s.broker.Publish(domain.GuessCountEvent{GameMode: domain.GameModeDailySlash, Count: count})
	return count, nil
}

// One helper replaces five copies of the same lock logic.
func (s *CachedGuessCountsStore) increment(ctx context.Context, field func(*domain.PlayerGuessCounts) *int) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p := field(&s.guessCountsCache)
	*p++
	gc := fromGuessCountDomain(s.guessCountsCache)
	if err := s.answerRepo.UpsertGuessCounts(ctx, &gc); err != nil {
		*p-- // roll back
		return 0, err
	}
	return *p, nil
}
```

**Recommended follow-ups:**

- Switch the count write to an atomic `$inc` with `ReturnDocument(After)` (`FindOneAndUpdate`). Then the in-memory count is only a cache, and the mutex no longer has to cover network I/O.
- Use `context.WithoutCancel(ctx)` plus its own timeout for the count write. A client that navigates away mid-request then can't abort a write that other users depend on.
- Add a per-user lock, or an atomic Mongo update with `Finished: false` in the filter, to close the double-submit race in `CheckGuess`.
- Log errors in `handleError` and in the refresh job. This failure produced exactly one error, and it was never logged.
- Add `ReadTimeout`/`WriteTimeout` to `http.Server` and a fetch timeout on the client.
- Add a regression test: a fake `AnswerRepo` that returns an error once, then assert a second `IncrementDailySlashCount` call returns within a deadline. Run the store tests with `-race`.

## 7. Confidence

**High.** The leak is deterministic and visible in the source. It needs only one MongoDB write error, from any cause, during a winning guess. It explains every observed symptom:

- Only winning guesses hang, across all modes.
- The browser shows no response and no error.
- The service is fine after a restart and fine locally.
- The service runs for weeks between failures.
- Player counts get stuck.
- The earlier connectivity fixes didn't prevent a recurrence.

No MongoDB or Oracle data was needed for this conclusion. To confirm in production before deploying the fix, capture a goroutine dump from the hung process (`kill -QUIT <pid>`, or `net/http/pprof` if it's enabled). Expect many goroutines parked in `sync.(*Mutex).Lock` under `IncrementXCount`, and none inside a MongoDB call.
