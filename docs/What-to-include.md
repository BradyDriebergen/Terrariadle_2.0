## Architecture

High-level diagram or bullet flow of how a request moves through the system, e.g.:

```
Browser → Caddy (TLS, reverse proxy) → Go API → MongoDB Atlas
                                       ↳ SSE broker (live guess counts)
```

Worth calling out specifically since they're the most interesting engineering pieces:

- **Service layer structure** — domain → db → repo → store → services → handler
- **SSE broker pattern** — how live guess counts are pushed to clients, and how shutdown is handled gracefully
- **Write-behind caching** — in-memory cache backed by `sync.RWMutex`, flushed to Mongo asynchronously
- **Daily puzzle rotation** — the background refresh job that resets puzzles at midnight and how chunk-count math drives it
- **Typed error handling** — `AppError` pattern for consistent API error responses

---

## Design Choices

This is where you explain the _why_ behind decisions a casual reader might question. Some prompts:

- **Why static frontend + separate Go API** instead of a full SvelteKit backend?
- **Why SSE instead of polling or WebSockets** for live guess counts?
- **Why Mongo** over a relational database for puzzle data?
- **Why self-host on Oracle Cloud** instead of Vercel/Railway/etc.?
- **State management** — how localStorage-based sessions work, and why no user accounts (if that's the case)

---

## Screenshots

<!-- 3-5 screenshots, one per game mode ideally -->

---

Extra things I should talk about:

- hosting
- proxy-ing
