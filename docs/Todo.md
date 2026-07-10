## Next Steps for final product

**Next items needed**
- [x] Move data to Atlas
- [x] Change backend to pull from Atlas rather than JSON
- [x] Create page for new game
    - [x] Plan and Implement TerraTrivia backend for minigame
    - [x] Create front-end components for game
    - [x] Integrate backend and front-end
- [ ] Update go and npm packages
- [x] Updating flying enemies in categories and fix 'Fultures'
- [x] Seperate Staff/Wand in db

**Frontend changes**
- [x] Go through all the front-end and update it with new knowledge
    - [x] Home page
    - [x] Daily Slash
    - [x] Connections
    - [x] Guess the NPC
    - [x] Hangman
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
- [ ] Add static adapter to Sveltekit
- [x] Reformat the site file structure:

Terrariadle/
├── main.go
├── internal/
├── frontend/
├── go.mod
├── go.sum
├── Makefile
└── docs/

- [ ] Figure out versioning
- [ ] Update Github to ignore certain files
- [ ] Add a new user to Oracle instance
- [ ] Add .env files to Oracle instance
- [ ] Reset MongoDB Atlas password for db
- [ ] Create a systemd service for the app
- [ ] Implement Caddy in place of Nginx
- [ ] Create new script that automatically builds and deploys
- [ ] Reset MongoDB Atlas password for db

**Finishing steps**
- [ ] Update documentation and wiki
- [ ] Market the site:
    - [ ] Redit posts
    - [ ] Twitter
    - [ ] Reach out to Youtubers
    - [ ] Svelte promotions
- [ ] Add page for contacting support
- [ ] Add monitoring for better analytics
- [ ] Add some sort of logging that doesn't bog down my server
- [ ] Update categories data
    - [ ] Add more categories
    - [ ] Update the options to contain only 4 options

**Future nice to haves:**

- [ ] Consider GitHub Actions for automatic deployment
- [ ] Add a stats screen at the beginning for how many users have guessed
- [ ] Change cursor to the Terraria cursor (in progress)