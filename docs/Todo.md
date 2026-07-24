## Next Steps for final product

**Next items needed**

- [x] Move data to Atlas
- [x] Change backend to pull from Atlas rather than JSON
- [x] Create page for new game
    - [x] Plan and Implement TerraTrivia backend for minigame
    - [x] Create front-end components for game
    - [x] Integrate backend and front-end
- [x] Update go and npm packages
- [x] Updating flying enemies in categories and fix 'Fultures'
- [x] Seperate Staff/Wand in db

**Frontend changes**

- [x] Go through all the front-end and update it with new knowledge
    - [x] Home page
    - [x] Daily Slash
    - [x] Connections
    - [x] Guess the NPC
    - [x] Hangman
    - [x] TerraTrivia
    - [x] About
- [x] Create tabs at the top to navigate through games
- [x] Change About section link to sign at the bottom

**Backend changes**

- [x] Update database package to use dependency injection
- [x] Split up `server` package for `server` + `service`
- [x] Create `domain` package for shared types
- [x] Create `repo` package to replace `models` functionality
- [x] Implement in-memory guess store with periodic flush
- [x] Fix jobs after implementing domain and store
- [x] Change backend guess counts to use SSE - server sent events (used to be web-sockets)
- [x] Go through backend and finalize structure

**Hosting steps**

- [x] Add static adapter to Sveltekit
- [x] Reformat the site file structure:

```
Terrariadle/
├── main.go
├── internal/
├── frontend/
├── go.mod
├── go.sum
├── Makefile
└── docs/
```

- [x] Embed frontend build in go binary
- [x] Figure out versioning
- [x] Update Github to ignore certain files
- [x] Create oracle instance and set up networking
- [x] Update Cloudflare DNS to point to server
- [x] Add a new user to Oracle instance
- [x] Add .env files to Oracle instance
- [x] Create a systemd service for the app
- [x] Implement Caddy in place of Nginx
- [ ] Create new script that automatically builds and deploys
- [ ] Reset MongoDB Atlas password for db
- [ ] Run through site one final time
- [ ] Upload build binary and run service

**Finishing steps**

- [x] Update documentation ~~and wiki~~
- [x] Update Cloudflare domain to use Proxied (DNS -> records)
- [ ] Market the site:
    - [ ] Redit posts
    - [ ] Twitter
    - [ ] Reach out to Youtubers
    - [ ] Svelte promotions
- Add it to google search index, or anything equivalent

**Future nice to haves:**

- [ ] Consider GitHub Actions for automatic deployment
- [ ] Add a stats screen at the landing page for how many users have guessed
- [ ] Change cursor to the Terraria cursor (in progress)
- [ ] Update categories data
    - [ ] Add more categories
    - [ ] Update the options to contain only 4 options
- [ ] Add page for contacting support
- [ ] Add monitoring for better analytics
- [ ] Add some sort of logging that doesn't bog down my server
- [ ] Automate linting script as a check before a PR can be merged into main
- [ ] Automatically run the prettier format script when pushing to main
- [ ] Support mobile view (IMPORTANT)
- [ ] Implement Sitelinks for Google indexing
- [ ] Add error handling on try-catch loops on the frontend / api
- [ ] Add malformed input error protection in `/frontend/src/lib/api`, like so:

```ts
export async function parseJsonSafe<T = unknown>(
	res: Response,
): Promise<T | null> {
	try {
		return (await res.json()) as T;
	} catch {
		return null;
	}
}

export async function initializeDailySlashGame(
	fetchFn: typeof fetch,
	userId: string,
): Promise<DailySlashSession> {
	const res = await fetchFn(
		`/api/daily-slash/initialize-game?user_id=${userId}`,
	);

	if (!res.ok) {
		const err = await parseJsonSafe<ApiErrorBody>(res);
		throw new ApiError(
			res.status,
			err?.error ?? "Unable to initialize game",
		);
	}

	const body = await parseJsonSafe<DailySlashSession>(res);
	if (!body) {
		throw new ApiError(res.status, "Received empty or malformed response");
	}

	return body;
}
```
