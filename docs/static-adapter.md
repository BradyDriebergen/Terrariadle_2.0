# Migrating Terrariadle to `adapter-static`

One of the main features that drew me to Svelte was that it compiled down to pure JavaScript. I loved that Svelte didn't need a virtual DOM to create a responsive and efficient app. Alongside other reasons, I decided to make my frontend in SvelteKit.

During my development, I learned about embedding static files in Go binaries (see [embedding binaries](embedding-binaries.md) for more info). I realized that SvelteKit could compile my frontend down to plain static HTML/JS/CSS and embed it directly into my Go binary. My frontend and backend could combine into one deployable service instead of running a separate Node process. I was immediately hooked, and quickly learned I needed to convert my app to use `@sveltejs/adapter-static` for compiling.

## Static Adapter vs. a Normal Adapter (with server files)

With a normal adapter (like `adapter-node` or `adapter-auto`), SvelteKit ships with a Node server runtime. This means you can have a built-in backend to your frontend application. `+page.server.ts` / `+server.ts` files run on only the server, not the client. These files can perform tasks like reading cookies, calling private APIs with secrets, pre-building pages, etc. This is because there's an actual server process alive to handle it and not expose secrets to the client.

`adapter-static` removes that server entirely. There is no runtime after the build finishes, it just turns into files on disk. Because of that, there is no `+server.ts` / `+page.server.ts` support. Anything that needs to run at request time is no longer easily available.

Routing also changes with a static adapter. The route's HTML is generated once and served as a static file until its next build. The route also serves a generic HTML shell, and the actual page content/data loading happens in the client's browser after JS loads.

This all culminates to an interesting situation where features like pre-rendering and route handling become a bit more tricky.

## Why SPA mode was the right call, not prerendering

Two things made this decision, and the second one turned out to be the harder constraint:

1. **Daily-changing data.** Puzzle content and guess counts change over time. Prerendering would freeze that data at whatever moment I last ran a build.
2. **Per-user data.** Several of my `load` functions need a user ID to fetch that user's game results. At build time, there is no visitor and no user ID — that value simply doesn't exist yet. Prerendering isn't just suboptimal here, it's structurally the wrong tool, since there's nothing valid to substitute in for a user who hasn't shown up yet.

Given both signals — time-sensitive data and per-user data — SPA mode across all routes was the clear choice.

## What I changed

Rather than having the config that SvelteKit ships with:

```js
...
kit: {
    adapter: adapter()
}
```

I instead updated my `svelte.config.js` file:

```js
import adapter from "@sveltejs/adapter-static";
import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

const config = {
	preprocess: vitePreprocess(),

	kit: {
		adapter: adapter({
			pages: "build",
			assets: "build",
			fallback: "index.html",
			precompress: false,
			strict: true,
		}),
	},
};

export default config;
```

- `fallback: 'index.html'` is what makes SPA mode deployable on a static host — without prerendered routes, there's no per-URL HTML file to serve, so every path falls back to this one shell, and client-side routing takes it from there.
- `strict: true` makes the build fail loudly if any route isn't covered by prerendering or the fallback, instead of silently shipping something broken.

**`src/routes/+layout.ts`**

```ts
export const ssr = false;

export const load: LayoutLoad = async ({ fetch }) => {
	...
};
```

`ssr = false` is the actual switch that makes `load` functions run only in the browser. It applies everywhere — `npm run dev`, `vite preview`, and the deployed static build — so dev mode now behaves the same way production does, no surprises after deploying.

## Prerendering vs. SPA mode — what actually happens

This was the part I initially struggled with, so I wanted to share what I've learned.

**Prerendering** (`export const prerender = true`) runs my `load` functions **once, during `npm run build`**, in a temporary Node process. Whatever my Go API returns _at that exact moment_ gets baked directly into the generated HTML file as static markup. After the build finishes, that file is frozen — nothing is "stored" in the traditional sense, but the API response is now permanently embedded in the output artifact itself. It won't change again until I rebuild and redeploy. For a route showing a daily-rotating puzzle, this means every visitor gets served whatever puzzle existed at build time, indefinitely.

**SPA mode** (`export const ssr = false`) skips running `load` functions at build time entirely. Instead, the adapter builds one generic shell HTML file (the `fallback`) plus the JS bundle. When someone visits any route, the server just returns that same shell regardless of URL; the browser downloads the JS, Svelte boots up, figures out which route it's on, and _then_ runs the `load` function — live, in the browser, hitting my Go API fresh every time.

## What I struggled with

- **Understanding _why_ prerendering causes staleness.** It wasn't obvious to me at first that "the frontend isn't storing anything" didn't mean the data stayed live — the API response gets embedded directly into the static HTML output at build time, so the output file itself becomes the stale artifact.
- **Forgetting to actually switch the adapter import.** I updated the adapter config object but left the import statement pointing at `@sveltejs/adapter-auto`. The build "succeeded" but silently fell back to `adapter-auto`, which logged `Could not detect a supported production environment` and produced both client _and_ server output — a sign the static adapter wasn't actually active. No `build/` folder ever appeared because it was never being generated. Fixing the import line resolved it immediately.

## What I learned

- `ssr = false` and `fallback` solve two different problems: `ssr = false` controls _when/where_ `load` functions execute (browser-only, always), while `fallback` controls _what gets served_ for routes that aren't prerendered on a static host.
- SvelteKit's client-side router re-runs `load` functions on every in-app navigation regardless of prerender/SPA settings — the prerender-vs-SPA distinction only really matters for the very first page load (or a hard refresh).
- Any route whose `load` function depends on cookies, auth state, or other per-user/per-request data is a strong signal that prerendering isn't viable for it, independent of whether the data is also time-sensitive.
- Good ways to verify a static adapter migration actually worked:
    - Build with the backend unreachable — if the build still succeeds, no API calls are happening at build time.
    - Inspect the generated HTML directly and confirm no API data is embedded in it.
    - Run `npm run preview` with the Network tab open and confirm data-fetch requests fire client-side, after the shell loads.
    - Manually navigate straight to a nested route URL (not via in-app navigation) to confirm the fallback shell is actually being served correctly by the host.
