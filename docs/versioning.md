# Versioning Terrariadle: How I Set It Up and Why

When I was ready for the initial release of this project, I wanted to include a versioning system that represented the current stage of development. In my past iterations, I would just update a string value on my frontend after any change. This process came with a slew of problems, including forgetting to bump up the version number, ignoring any backend changes, and having no clear boundaries between what type of update it was (feature/bug/release). In this new versioning system, I wanted it to be as easy as possible to increment version numbers and not require manual intervention.

## The Approach I Landed On

I combined two ideas:

- **Using Git tags for the version number:** Rather than storing the version number in the project, I store the version number as a Git tag (`v0.0.0`). The version number could then get derived by the command `git describe --tags --always --dirty` at build time.
- **GitHub PR labels trigger automated tag creation:** Instead of running `git tag` by hand, I would label a PR (`version:patch`, `version:minor`, `version:major`), and a GitHub Action reads that label on merge and pushes the appropriate tag for me.

This requires me to only make PRs and label them. Everything else like computing the next version number, creating the tag, pushing it, and using it in both my Go backend and SvelteKit frontend happens automatically.

## How It Works

1. I fix a bug or add a feature in a new branch.
2. Once complete, I create a pull request and add one of three labels: `version:patch`, `version:minor`, or `version:major`, depending on the size of the change.
3. When squashing and merging to `main` (the default branch), a GitHub Action (`.github/workflows/release.yml`) reads the label off the merged PR, looks at the latest existing tag, computes the next semver version, and pushes the new tag.
4. I then build the project using the Makefile. This make command runs `git describe --tags --always --dirty` and injects the version number into both the frontend and backend of the app (shown below).
5. The built binary now contains the current version number of what's on main without me having to manually update it.

### Other things to note:

- If I don't add a label to the PR, the version number doesn't increment. This allows me to make a change that might not warrant a version change.
- Semantic versioning reset behavior is built into the GitHub Action: a major bump resets minor and patch to 0 (`v1.3.8` -> `v2.0.0`), and a minor bump resets patch to 0 (`v1.3.8` -> `v1.4.0`).

## Backend Injection

There is a package-level variable that determines what version is. It defaults to `dev` when I run it locally, but when built (via `go build ...`), the Makefile injects the version at compile time:

`/main.go`:

```go
var version = "dev"

func main() {
    ...
    log.Printf("Starting Terrariadle version %s, listening on port :8080\n", version)
}
```

`/Makefile`:

```makefile
build:
	cd frontend && npm run build
	rm -rf internal/web/build
	cp -r frontend/build internal/web/build
	go build -ldflags "-X main.version=$$(git describe --tags --always --dirty)" -o bin/terrariadle .
```

## Frontend Injection

The frontend works similar to the backend, where the version number is injected at build time. `Vite`, which is used to build the frontend, defines the version by running the Git tag command. This is then saved as `__APP_VERSION__`, which is used as a global value that can be used throughout the site.

`frontend/vite.config.ts`:

```ts
const appVersion = execSync("git describe --tags --always --dirty")
	.toString()
	.trim();

export default defineConfig({
	plugins: [sveltekit()],
	server: {
		// Api proxy used for dev
	},
	define: {
		__APP_VERSION__: JSON.stringify(appVersion),
	},
});
```

`/frontend/src/app.d.ts`:

```ts
declare global {
	namespace App {
		// global values
	}
	const __APP_VERSION__: string;
}
export {};
```

## Downsides to This Approach

- **Branch protection needed** This isn't so much a downside as it is a requirement. Because I'm basing the versioning off of a Git tag in `main`, I can no longer directly push to main without it breaking. This is easily fixed however with branch protections and not allowing `git push --force`.
- **Not completely automated.** I still have to go in and remember to label the PRs in order to upgrade the version. This is a lot easier to remember compared to updating a line in a file after implementing a change, but still something I need to remember to do.
- **Hard to revert/update version number** If I ever need to revert a change, or I forget to update a version number, there is no concrete process in place to revert/update a version number. Currently, the only way to reset a version number is deleting the version Git tag and creating a new tag with the new version number. While this works, this is a clunky and unpredictable.
- **Not completely battle tested.** I'm using my own custom logic rather than using a pre-built dependency like `anothrNick/github-tag-action`. I wanted to make my own logic to get the full experience of building out versioning. It also allows me to not trust someone else's Github Action for updating my tags. However, it makes more code for me to maintain. This isn't a major issue, but something I need to keep my eye on.

---

This was a fun system to learn, and I feel a lot better about versioning. My implementation isn't perfect, but it at least gets me a stepping off point where I can grow from.
