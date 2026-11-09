# Post-Incident Review: Check-Guess Endpoint Hang on Winning Guesses

| Field                  | Value                                                                   |
| ---------------------- | ----------------------------------------------------------------------- |
| **Area of impact**     | Check-guess endpoints, guess count store, MongoDB Atlas                 |
| **Audience effected**  | All users                                                               |
| **Incident date**      | 11/2/26                                                                 |
| **Detected**           | 11/3/26                                                                 |
| **Service Restored**   | 11/3/26                                                                 |
| **Fix implemented**    | 11/5/26                                                                 |
| **Duration of impact** | 36 hours                                                                |
| **Severity**           | Major                                                                   |
| **Status**             | Service restored by restart. Root cause **presumed**, fixed, and pushe. |

---

## 1. Summary

After routine use of the application on 11/3/26, users were unable to complete any of the Terrariadle game-modes. Non-winning guesses work normally, but when a user submitted a **winning** guess, nothing happened. The request never returned a response or an error. This prevented everyone from completing any of the games.

The service had been running for almost a month between deployments without a restart before the issue appeared.

As a quick response, a deployment and reset was made to the systemd service. This action restored normal behavior, and testing verified that all games were working correctly.

This document will go over this issue and the presumed causes, giving context if future versions of this issue arise.

## 2. How the Issue Was Encountered

I found the issue when I decided to play Terrariadle later in the day (8:00 PM). I first loaded the game _Daily Slash_, and noticed that 0 players guessed that day. This was odd because typically, I would have tens if not hundreds of players that guessed correctly.

Concerned, I played the game like I usually would, and found that I couldn't guess the option **Cursed Flames**. I was able to guess any other option, but this option seemed to leave the site hanging, like nothing happened.

Looking through the other games, it was clear that correct guesses weren't loading correctly. I tried the _Guess the NPC_ and received the same result. By this point, I knew there was a major problem with the application and took action.

## 3. Who Was Affected

All players were affected with this issue, making this a major outage

## 4. Working Through the Cause

The investigation narrowed the cause through a series of hypotheses.

### 4.1 Initial suspicion: hanging context or something in the service

The first suspicion was a hanging context. A restart fixing the problem pointed to some process-level state getting stuck, as opposed to bad data.

### 4.2 Hypothesis: SSE broker blocking on broadcast (ruled out)

**Reasoning:** A winning guess triggers a broadcast of the updated guess count to SSE subscribers. A blocking send to a dead subscriber, or a lock held during a send, could stall the winning-guess request.

**Findings after reviewing the broker code:**

- `Publish` uses a non-blocking `select` with a `default` case. When a subscriber's buffer (8 events) is full, the event is dropped instead of blocking.
- `Unsubscribe` takes the write lock, so it can't run during a `Publish`. It also deletes the channel from the map before closing it, so there's no send-on-closed-channel race. That race would panic, not hang.
- A stuck client only blocks its own handler goroutine. Its buffer fills and further events are dropped.

**Conclusion:** The broker can't cause "every winning guess hangs while everything else works." Ruled out.

### 4.3 Narrowing to the winning-guess path

Comparing the winning and non-winning paths in `CheckGuess`:

- Both call `userCache.GetUser` and `userCache.UpsertUser`. The user cache is write-behind, so neither waits on MongoDB.
- **Only a winning guess** calls `guessCountCache.IncrementDailySlashCount(ctx)`, which also calls `Publish`.

So the hang had to be inside the increment.

The fact that the service ran for about a month first also mattered. A deterministic deadlock would have appeared on the first win. This pointed to a rare event that left the process permanently stuck.

### 4.4 Hypothesis: mutex not released on an error path (ruled out)

**Reasoning:** A classic cause of "works for weeks, then hangs" is a lock left held on an early-return error path. One transient error, and every later caller blocks on `Lock()` forever.

**Findings after reviewing `IncrementDailySlashCount`:** Both the error path and the success path call `s.mu.Unlock()`, and `Publish` runs after the unlock. Lock handling is correct. Ruled out.

**Real problem found:** The function holds `s.mu` **across a network call to MongoDB Atlas** (`answerRepo.UpsertGuessCounts`), using the request's context, which has **no deadline**:

```go
s.mu.Lock()
s.guessCountsCache.DailySlashCount++
if err := s.answerRepo.UpsertGuessCounts(ctx, &s.guessCountsCache); err != nil { ... }
// ...
s.mu.Unlock()
```

If that call stalls:

- The goroutine holding the lock waits indefinitely.
- Every later winning guess queues on `s.mu.Lock()`. `sync.Mutex` doesn't respect contexts, so waiters stay parked even after their clients disconnect.
- Anything else using that lock, possibly including the SSE initial count, also hangs.

This also explains why **only** winning guesses were affected. Because the user cache is write-behind, the guess count write is the **only synchronous MongoDB call on a normal request**. If MongoDB connectivity went bad, the winning-guess path is the only place it would show.

A smaller issue: the upsert uses `$set` with the entire cached struct, overwriting every mode's count with the in-memory values. An atomic `$inc` would be more correct.

### 4.5 Hypothesis: connection pool exhaustion from leaked cursors (ruled out)

**Reasoning:** Unclosed `Find` cursors can slowly leak connections until the pool (default 100) is full. New operations then wait for a free connection, and without a deadline they wait forever.

**Findings after reviewing the DB helpers:**

- `FindOne` returns a single result and doesn't leave a cursor open. `Decode` reads data that has already been fetched.
- `GetAll` uses `defer cursor.Close(ctx)`, and `cursor.All` closes the cursor itself.
- Outside of startup, these are only used by `GetUser` and `GetOrCreateUser`.

**Conclusion:** No cursor leak. Ruled out.

### 4.6 Why nothing was logged

- Errors are not logged. They are passed up to the `api` package and returned as responses.
- In this incident, though, **no error was ever produced**. The MongoDB call never returned, so nothing reached the API layer. A hanging request with no error is consistent with a blocked network read.
- From a user's view, the hang looked permanent. Each stuck request held the lock until its own client gave up. If the connection was still bad, the next winner got stuck the same way.

## 5. Final Assumption on the Cause

> **Presumed root cause:** A pooled MongoDB connection went stale while idle, for example dropped by a NAT or firewall between Oracle Cloud and Atlas, or by an Atlas maintenance event. The driver's health checks use their own connections, so the cluster still looked healthy. The next winning guess checked out the stale connection inside `IncrementDailySlashCount` and blocked indefinitely on the write. The call had no deadline and held `s.mu`, so every later winning guess hung behind it until the service restarted.

**Confidence:** Most likely explanation. **Not confirmed.** Nothing else in the code accounts for a failure limited to this path.

**Contributing factors:**

1. A network call made while holding a mutex.
2. No deadline on MongoDB operations, at either the client or the call level.
3. No cap on idle connection time in the driver pool.
4. Errors were not logged, and no hang detection or monitoring was in place.

## 6. Resolving the Issue

### Immediate mitigation

- Restarted the systemd service, which restored normal behavior.

### Implemented fix after event

**Cap idle connection time and add a client-wide timeout** in `Connect`:

    ```go
    opts := options.Client().
    	ApplyURI(uri).
    	SetMaxConnIdleTime(2 * time.Minute). // drop pooled connections idle longer than this
    	SetTimeout(5 * time.Second)          // backstop deadline for every operation
    ```

    With these in place, the worst case becomes a short error response instead of an indefinite hang.

- Date fixes deployed: 11/5/26

## 7. Steps to Better Handle Issues Like This

- **Observability** - Features are being planned out to increase the logging and monitoring of the application
- **Monitoring and alerting service** - A service that will run on another machine that checks the health of the backend. Throughout the day, it will constantly check every endpoint to ensure that it is (a) returning a response, (b) not returning an error, and (c) returning the correct data/shape. This service will run independently of Terrariadle.
- **Automatic reset on breaking behavior** - If the above-mentioned service has an issue with calling the server, it has built in functionality that can reset the systemd service, providing a quick resolution to a similar issue. This restart is only tried once before giving up.
- **Notifications on outage** - The service will also provide a way to give me immediate notification in the event that an application breaking outage occurs like this incident.

## 9. Lessons Learned

I learned that not all issues are as clear as they seem to be. My first assumptions was that this was a hanging context issue. I quickly dismissed the assumption, realizing that the site ran for a month without issue and it only affected the winning guess.

After much testing, it couldn't be determined with certainty that the stale MongoDB connection was the issue. There are too many outside factors and not enough observability to correctly identify this. Steps will be made to ensure something like this doesn't happen again.
