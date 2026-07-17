# Terrariadle Documentation

Welcome to the Terrariadle project documentation. This table of contents goes over all the important developer documents and where you can find them.

## Table of Contents

### Growth points as a developer

This section of documents goes over the struggles I had with developing the site.

- [Choosing My Stack](./choosing-my-stack.md)
    - This document goes over what drew me to choosing this stack, and the benefits and shortfalls of each layer.
- [Managing Go Packages](./managing-go-packages.md)
    - This document goes over my struggles with designing systems and process flows, and how this became my cleanest project yet. Specifically, my Go backend.
- [Embedding Binaries](./embedding-binaries.md)
    - This document goes over how I was able to make running two separate services into a single service utilizing Go's `embed` package.
- [Static Adapter](./static-adapter.md)
    - This document goes over my use of SvelteKit's `adapter-static` tool to generate static HTML/JS files, rather than a site with a node runtime. Plays into Embedding Binaries
- [Versioning](./versioning.md)
    - This document goes over how I learned to version my site without having to manually update my codebase with every release.
- [Infrastructure](./infrastructure.md)
    - This document goes over hosting my project, and the configurations/optimizations I used in my instance.

### Architecture

This section goes over the architecture of the project, and how the games were designed to play.

#### Backend

- [Endpoints](./architecture/backend/endpoints.md)
    - This document is a list of my backend endpoints I use in my games
- [Packages](./architecture/backend/packages.md)
    - This document summarizes my Go packages and their place in the flow.
- [Process Flow](./architecture/backend/process-flow.md)
    - This document contains the dependency map, and how data flows from the database to the client

#### Games (Frontend)

- [Daily Slash](./architecture/frontend/daily-slash-design.md)
    - This document contains Daily Slash's design and game sequence.
- [Connections](./architecture/frontend/connections-design.md)
    - This document contains Connections's design and game sequence.
- [Guess The NPC](./architecture/frontend/guess-the-npc-design.md)
    - This document contains Guess the NPC's design and game sequence.
- [Hangman](./architecture/frontend/hangman-design.md)
    - This document contains Hangman's design and game sequence.
- [TerraTrivia](./architecture/frontend/terratrivia-design.md)
    - This document contains TerraTrivia's design and game sequence.

### Dev-Notes

- [Deployment](./ops/deployment.md)
- [Versioning & Release Process](./ops/versioning.md)
- [Puzzle Refresh Job](./ops/puzzle-refresh.md)
- [robots.txt & Sitemap](./ops/seo-config.md)

### Development Journal

- [Dev Journal Index](./dev-journal/README.md)

### Contributing

- [Contributing Guide](./contributing.md)
- [Code Style & Conventions](./code-style.md)

---

_Keep this index up to date as new docs are added. Each linked file should exist under `docs/` with a matching path._
