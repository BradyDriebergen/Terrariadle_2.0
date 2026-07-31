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
- [x] Create new script that automatically builds and deploys
- [x] Reset MongoDB Atlas password for db
- [x] Update Kofi
- [x] Run through site one final time
- [x] Upload build binary and run service
- [x] Make Github repo public

**Finishing steps**

- [x] Update documentation ~~and wiki~~
- [x] Update Cloudflare domain to use Proxied (DNS -> records)
- [ ] Market the site:
    - [ ] Redit posts
    - [ ] Twitter
    - [ ] Reach out to Youtubers
    - [x] Svelte promotions
- [x] Add it to google search index, or anything equivalent
- [x] Fix versioning to ignore pushes to main without tag
- [x] Fix mobile issue with daily slash and guess the NPC
- [ ] Route terrariadle.net to terrariadle.com
- [x] Fix the following errors from the browser console:

Cross-Origin Request Blocked: The Same Origin Policy disallows reading the remote resource at https://static.cloudflareinsights.com/beacon.min.js/v4513226cdae34746b4dedf0b4dfa099e1781791509496. (Reason: CORS request did not succeed). Status code: (null).

None of the “sha512” hashes in the integrity attribute match the content of the subresource at “https://static.cloudflareinsights.com/beacon.min.js/v4513226cdae34746b4dedf0b4dfa099e1781791509496”. The computed hash is “z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==”.

- [x] Implement Sitelinks for Google indexing
    - Verify domain ownership (DNS TXT record or HTML file upload).
    - Submit sitemap.xml under Sitemaps.
    - Use URL Inspection to request indexing on your key pages manually — speeds up first crawl significantly rather than waiting for organic discovery.

- [ ] Add error handling on try-catch loops on the frontend / api

**Future nice to haves:**

- [ ] Consider GitHub Actions for automatic deployment
- [ ] Add a stats screen at the landing page for how many users have guessed
- [ ] Change cursor to the Terraria cursor
- [ ] Update categories data
    - [ ] Add more categories
    - [ ] Update the options to contain only 4 options
- [ ] Add page for contacting support
- [x] Support mobile view (IMPORTANT)
- [x] Fix Linter errors
- [ ] Add dark mode
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

- [ ] Add monitoring for better analytics
- [ ] Add some sort of logging that doesn't bog down my server
