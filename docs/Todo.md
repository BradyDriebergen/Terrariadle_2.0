## Next Steps for final product

**Next items needed**
- [x] Move data to Atlas
- [x] Change backend to pull from Atlas rather than JSON
- [ ] Create page for new game
    - [ ] Plan and Implement Wordaria backend for minigame
    - [ ] Create front-end components for game
    - [ ] Integrate backend and front-end
- [ ] Update go and npm packages
- [ ] Updating flying enemies in categories and fix 'Fultures'

**Frontend changes**
- [ ] Go through all the front-end and update it with new knowledge
    - [ ] Home page
    - [ ] Daily Slash
    - [ ] Connections
    - [ ] Guess the NPC
    - [ ] Hangman
    - [ ] About
- [ ] Create tabs at the top to navigate through games
- [ ] Change About section link to sign at the bottom

**Backend changes**
- [x] Update database package to use dependency injection
- [ ] Split up `server` package for `handler` + `service`
- [x] Create `domain` package for shared types
- [x] Create `repo` package to replace `models` functionality
- [x] Implement in-memory guess store with periodic flush
- [x] Fix jobs after implementing domain and store
- [ ] Change backend guess counts to use web-sockets
- [ ] Go through backend and finalize structure

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
- [ ] Add monitering for better analytics
- [ ] Update categories data
    - [ ] Add more categories
    - [ ] Update the options to contain only 4 options

**Future nice to haves:**

**CD/CI changes**
- [ ] Github cleanup
    - [ ] Reset commits so no previous history was there
    - [ ] Implement Github actions
    - [ ] Configure Github secrets
    - [ ] Setup container registry
    - [ ] Verify that flow works from end to end

