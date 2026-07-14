**Sveltekit**

I chose this framework because of a few reasons:

- Tutorial was amazing, easy to learn from. Everything is in HTML/CSS/JS. Small learning curve
- Modern framework with awesome tools built in (Vite, Linter, prettier, etc.)
- Compiles down directly to HTML/JS. No virtual DOM, and allows for embedding in go binary
- Super simple reactivity with states and runes.
- Built in scoped styles
- File structure and page layout is very simple
- Simple framework.

Shortfalls

- Tiny ecosystem. Few projects, dependencies, stack-overflow posts, and very little documentation
- Server side rendering vs static site generating
- Updates change a lot of things (prop issues, svelte 4 -> svelte 5)
- No clear line between legacy features and current features, (svelte/store vs svelte/state).

**Go**

I chose this language because:

- Go is extremely lightweight, combining the speed of languages like C while also having a garbage collector.
- Go is a procedural programming language, opting for functions and packages compared to classes and abstraction.
- Concurrency is extremely easy. Mutexes and goroutines have a bit of a learning curve, but are super easy to use.
- Type assignment, using := without having to declare every type
- Standard library is amazing. Able to make a full backend with SSE events with only 2 outside dependencies.
- Go binaries allow me to run the backend and frontend from the same backend, only needing one process to run for the whole project.
- Error handling is very easy to read. You know where a potential error could happen and follow it down the rabbit hole really easily.

Challenges:

- Very high learning curve. Started out with OOP and learning Go's package structure was very different.
- Very easy to mess up a project without proper system design fundamentals. Started out with a web of dependency loops, and after learning about system design where I found out where I was going wrong. It's easy to mess up a project.
- Error handling can get a bit much at time. Busies up the code and makes certain parts hard to read.
- Easy to create race conditions with goroutines

---

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

- versioning
- embedding
- building
- hosting
- proxy-ing

Include a document on how to reset version number:

git tag v2.0.0
git push origin v2.0.0
