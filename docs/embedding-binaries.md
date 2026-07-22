# Embedding Static Files in Go

### Table of Contents

- [What is it?](#what-is-embedding-files-into-go-binaries)
- [Layer 1: Embedding](#1-embedding-getting-the-files-into-the-binary)
- [Layer 2: Serving](#2-serving-turning-the-embedded-fs-into-http-responses)
- [Layer 3: SPA Fallback](#3-spa-fallback-handling-client-side-routes)
- [Wiring it in](#wiring-it-into-the-server)
- [Other Issues](#other-issues-i-ran-into)

---

In my old implementation of this project, I was running the frontend and backend as two separate services on my server. It was never really a major problem, but it resulted in two things to deploy, two ports to manage, and a CORS configuration just to let the frontend talk to the API.

I expected to do the same for this project, until I did more research on embedding files into a Go binary. In short, I was able to bundle my frontend and backend into one executable. I was blown away at this capability because it allowed me to only need to run one service on my server. This reduced complexity also meant I only had to manage one port. I decided it would be perfect for this project.

_Also, in my old implementation, I never created systemd service. I actually ran the old site on two tmux instances using nginx as a proxy (that didn't even work properly). It was super buggy and probably what caused my site to crash so often._

## What is Embedding Files Into Go Binaries?

Go's `embed` package bundles external files directly into a compiled binary. When compiling, the Go compiler reads every file under the specified path (declared by `//go:embed <folder name>`) and copies the raw bytes into the binary as static data. The result is a single, self-contained executable with no separate folder of static assets to ship alongside it.

## The Three Layers

I struggled to get this working at first. It required understanding three separate concerns. I kept mixing these up at first, so I'm documenting them separately here for anyone who's curious.

### 1. Embedding: getting the files into the binary

I compile the SvelteKit project with the static adapter, which outputs a `build/` directory containing `index.html`, static assets, and a JS bundle. To learn about this more in detail, see the [static-adapter](static-adapter.md) document. I copy that `build/` directory into my Go module (it has to live inside the module, `//go:embed` can't reach outside it), then embed it:

```go
package web

import (
	"embed"
	"io/fs"
)

//go:embed all:build
var embeddedFiles embed.FS

func Assets() fs.FS {
	sub, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		panic(err)
	}
	return sub
}
```

At compile time, the Go compiler reads every file under `build/` and copies the bytes into the binary as static data. After it's done compiling, the `build/` folder on disk is no longer needed (until the next compile).

The `Assets` method strips the `build/` prefix, so paths are rooted at `index.html` rather than `build/index.html`.

**Important:** the `all:` prefix on the embed directive is required to compile files that start with symbols like `" * : < > ? | \`. Go's `embed` package silently excludes files and directories starting with those symbols by default (meant to skip things like `.git` presumably). SvelteKit puts nearly all of its build output under `_app/`, so without `all:`, the entire JS/CSS bundle gets excluded from the binary with no error at compile time.

### 2. Serving: turning the embedded FS into HTTP responses

```go
fileServer := http.FileServer(http.FS(s.frontend))
```

This line looks a little funky, but all it's doing is converting a `fs.FS` into a `http.FS()`, which converts it again into a `http.FileSystem` interface. `http.FileServer` uses this file system to handle exact file-to-URL matching. It looks up the path in the embedded FS and serves the bytes with the correct headers. This is perfect for serving the SvelteKit frontend when it compiles down to static assets (JS, CSS, images) since those have exact filenames.

### 3. SPA fallback: handling client-side routes

This is the layer that gave me the most trouble. My SvelteKit app has client-side routes (e.g. `/connections` or `/daily-slash`) that don't correspond to real files in `build/`. The routing for those happen in the browser via the JS bundle _after_ `index.html` has loaded. If a browser makes a fresh request directly to one of those paths, whether through a page refresh, a shared link, etc, the plain `http.FileServer` can't find a matching file and returns a bare 404.

The fix is a small wrapper handler that checks whether a requested path exists as a real file. If it does, serve it normally. If it doesn't, rewrite the request to `/` so `index.html` gets served instead, and let the SvelteKit client-side router take over from there:

```go
func (s *Server) spaHandler() http.Handler {
	fileServer := http.FileServer(http.FS(s.frontend))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}
		if _, err := fs.Stat(s.frontend, path); err != nil {
			r.URL.Path = "/" // fall back to index.html for client-side routes
		}
		fileServer.ServeHTTP(w, r)
	})
}
```

An important nuance I initially got wrong: this handler has no knowledge of what routes actually exist in the app. It can't distinguish a legit SPA route from a mistyped URL. It only knows "no file matches this path." The actual 404 decision gets pushed to SvelteKit's client-side router, which renders its own `+error.svelte` 404 page once it loads and checks the URL against its known routes. So a bad URL still visually 404s for the user, it just happens client-side rather than as a raw HTTP 404 from the server.

## Wiring It Into the Server

Since my API routes were already prefixed with `/api/...`, adding the frontend was pretty easy. I only had to register `/` as a catch-all. It also doesn't interfere with the API routes as long as it's registered last.

```go
func (s *Server) newMux() http.Handler {
	mux := http.NewServeMux()

	s.registerCommonRoutes(mux)
	s.registerDailySlashRoutes(mux)
	s.registerConnectionsRoutes(mux)
	s.registerGuessTheNpcRoutes(mux)
	s.registerHangmanRoutes(mux)
	s.registerTerraTriviaRoutes(mux)

	mux.Handle("/", s.spaHandler())

	return withCORS(mux)
}
```

## Other Issues I Ran Into

**The white-screen MIME type error:** After wiring everything up for the first time, loading the site gave a white screen with console errors like:

```
Loading module from ".../start.Brgy2uXm.js" was blocked because of a disallowed MIME type ("text/html").
```

This happens because of the missing `all:` prefix described above. `_app/` never made it into the embedded binary in the first place. This caused the SPA fallback to treat every asset request as an unmatched route.

## Takeaway

While this `embed` package was hard to understand at first, I quickly learned more about how Go can handle embedded static files. I think this feature is super cool. It allows me to only have to run one HTTP server through Go to serve both my backend endpoints and my frontend pages. Super cool learning experience and I recommend trying it out for yourself.
