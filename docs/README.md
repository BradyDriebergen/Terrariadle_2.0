# Terrariadle Documentation

Welcome to the Terrariadle project's documentation. This table of contents goes over all the important developer documents and where you can find them. I appreciate you taking the time to read through my struggles, design choices, and how I've grown as a developer.

## Growth points as a developer

This section of documents goes over the struggles I had with developing the site and what I've learned.

- **[Choosing My Stack](./choosing-my-stack.md)** - What drew me to choosing this stack, and the benefits and shortfalls of each layer.

- **[Managing Go Packages](./managing-go-packages.md)** - My struggles with designing systems and process flows, and how this became my cleanest project yet. Specifically, my Go backend.

- **[Embedding Binaries](./embedding-binaries.md)** - How I was able to make running two separate services into a single service utilizing Go's `embed` package.

- **[Static Adapter](./static-adapter.md)** - My use of SvelteKit's `adapter-static` tool to generate static HTML/JS files, rather than a site with a node runtime. Plays into _Embedding Binaries_.

- **[Versioning](./versioning.md)** - How I learned to version my site without having to manually update my codebase with every release.

- **[Modifying Git History](./modifying-git-history.md)** - How I kept my Git history from showing specific files and why.

- **[Infrastructure](./infrastructure.md)** - Hosting my project, and the configurations/optimizations I used in my instance.

- **[Running the Service](./running-the-service.md)** - Setting up the `systemd` service to run on my Ubuntu shape.

## Architecture

This section goes over the architecture of the project, the backend blueprint, and how the games were designed to play.

#### Backend

- **[Endpoints](./architecture/backend/endpoints.md)** - List of my backend endpoints I use in my games.

- **[Packages](./architecture/backend/packages.md)** - Summarizes my Go packages and their responsibilities.

- **[Process Flow](./architecture/backend/process-flow.md)** - Contains the dependency map, and how data flows from the database to the client

#### Games (Frontend)

- **[Daily Slash](./architecture/frontend/daily-slash-design.md)** - Daily Slash's design and game sequence.

- **[Connections](./architecture/frontend/connections-design.md)** - Connections's design and game sequence.

- **[Guess The NPC](./architecture/frontend/guess-the-npc-design.md)** - Guess the NPC's design and game sequence.

- **[Hangman](./architecture/frontend/hangman-design.md)** - Hangman's design and game sequence.

- **[TerraTrivia](./architecture/frontend/terratrivia-design.md)** - TerraTrivia's design and game sequence.

## Maintenance

These notes are used by me to maintain the project and to refer back to if anything breaks.

- **[Building for Production](./maintenance/building-for-prod.md)** - Describes how to build for production.

- **[Curl API Calls](./maintenance/curl-api-calls.md)** - List of all `curl` calls I use to test the backend during development.'

- **[If the Backend is Overloaded](./maintenance/if-backend-is-overloaded.md)** - Gives tips and tricks to mediate issues if too much traffic hits the backend.

- **[If Versioning Breaks](./maintenance/if-versioning-breaks.md)** - Provides steps to fix versioning if it gets out of sync.

- **[Instance Setup](./maintenance/instance-setup.md)** - Overview of all commands and configs up to set up cloud instance.

- **[Logging](./maintenance/logging.md)** - Contains commands for viewing logs.

- **[`systemd` Text File](./maintenance/systemd-file.txt)** - Service file used to run Terrariadle

- **[Updating Dependencies](./maintenance/updating-dependencies.md)** - Instructions for updating Go and NPM dependencies.

---
