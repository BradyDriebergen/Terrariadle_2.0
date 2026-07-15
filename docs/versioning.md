# Versioning Terrariadle: How I Set It Up and Why

When I was ready for the initial release of this project, I wanted to include versioning to show users what stage of development this website was on. The goal of this versioning system is to make it as easy as possible, and not require manual intervention.

In my past iterations of versioning, I would just update a string value on my frontend. This caused me to forget to manually bump the version number. Rather than this manual process, I wanted to automate versioning, so I decided to utilize GitHub for versioning.

## The Approach I Landed On

I combined two ideas:

- **Using Git tags for the version number:** Rather than storing the version number in the project, `git describe
--tags --always --dirty` derives the version from the repo's tag history at build time.
- **GitHub PR labels trigger automated tag creation:** Instead of running `git tag` by hand, I label a PR (`version:patch`, `version:minor`, `version:major`), and a GitHub Action reads that label on merge and pushes the appropriate tag for me.

This allows for me to only need to make PRs and label them.Everything else like computing the next version number, creating the tag, pushing it, and using it in both my Go backend and SvelteKit frontend happens automatically.

## How It Works, End to End

1. I fix a bug or add a feature in a new branch.
2. Once complete, I create a pull request and add one of three labels: `version:patch`, `version:minor`, or `version:major`, depending on the size of the change.
3. When squashing and merging to `main` (the default branch), a GitHub Action (`.github/workflows/release.yml`) reads the label off the merged PR, maps it to a bump level (which looks at the latest existing tag), computes the next semver version, and pushes the new tag.
4. I then build the project using the Makefile. This make command runs `git describe --tags --always --dirty` and injects the version number into both the frontend and backend of the app (shown below).
5. The built binary now contains the current version number of what's on main without me having to manually update it.

### Other things to note:

- If I don't add a label to the PR, the version number doesn't increment. This allows me to make a change that might not warrant a version change.
- Semantic versioning reset behavior is built into the GitHub Action: a major bump resets minor and patch to `0` (v1.3.8 -> v2.0.0), and a minor bump resets patch to `0` (v 1.3.8 -> 1.4.0).

## Backend injection

There is a package-level variable that determines what version is. It defaults to `dev` when I run it locally, but when built, the Makefile injects the version at compile time:

`/main.go`:

```go
var version = "dev"

func main() {
    log.Printf("terrariadle-server version=%s starting", version)
    ...
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

## Frontend injection

The frontend works similar to the backend, where the version number is injected at build time. `Vite`, which is used to build the frontend, defines the version by running the git command. This is then saved as `__APP_VERSION__`, which is used as a global value that can be used throughout the site.

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
		/* ... */
	}
	const __APP_VERSION__: string;
}
export {};
```

## Downsides to This Approach

- **Branch protection needed** This isn't so much a downside as it is a requirement. Because I'm basing the versioning off of a Git tag in `main`, I can no longer directly push to main without it breaking. This is easily fixed however with branch protections and not allowing `git push --force`.
- **Not completely automated.** I still have to go in and remember to label the PRs in order to upgrade the version. This is a lot easier to remember compared to updating a line in a file after implementing changes, but still something I need to remember to do.
- **Hard to revert/update version number** If I ever need to revert a change, or I forget to update a version number, there is no concrete process in place to revert/update a version number. Currently, the only way to reset a version number is deleting the version Git tag and creating a new tag with the new version number. While this works, this is a clunky and unpredictable.
- **Not completely battle tested.** I'm using my own custom logic rather than using something like `anothrNick/github-tag-action`. I wanted to do this to get the full experience of building out versioning, but It makes more code for me to maintain. However, I don't need to my project to trust someone else's action with write access to my repo. This isn't a major issue, but something I need to keep my eye on.

---

This was a fun system to learn, and I feel better about versioning. My versioning isn't perfect, but it at least gets me a stepping off point where I can grow from.
