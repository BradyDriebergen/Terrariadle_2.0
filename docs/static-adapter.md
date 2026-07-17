# Migrating Terrariadle to `adapter-static`

One of the main features that drew me to Svelte was that it compiled down to pure JavaScript. I loved that Svelte didn't need a virtual DOM to create a responsive and efficient app. With this in mid and alongside other reasons, I decided to make my frontend in SvelteKit.

Later in my development, I learned about embedding static files in Go binaries (see [embedding binaries](embedding-binaries.md) for more info). I realized that SvelteKit could compile my frontend down to plain static HTML/JS/CSS and embed it directly into my Go binary. My frontend and backend could combine into one deployable service instead of running a separate Node process on top of my backend. I was immediately hooked, and quickly learned I needed to convert my app to use `@sveltejs/adapter-static` for compiling.

## Static Adapter vs. a Normal Adapter (with server files)

With a normal adapter (like `adapter-node` or `adapter-auto`), SvelteKit ships with a Node server runtime. This means you can have a built-in backend to your frontend application. `+page.server.ts` / `+server.ts` files run on only the server, not the client. These files can perform tasks like reading cookies, calling private APIs with secrets, pre-building pages, etc. This is because there's an actual server process alive to handle it and not expose secrets to the client.

`adapter-static` removes that server entirely. There is no runtime after the build finishes, it just turns into files on disk. Because of that, there is no `+server.ts` / `+page.server.ts` support. Anything that needs to run at request time is no longer easily available.

Routing also changes with a static adapter. The route's HTML is generated once and served as a static file until its next build. The route also serves a generic HTML shell, and the actual page content/data loading happens in the client's browser after JS loads.

This all culminates to an interesting situation where features like pre-rendering and route handling become a bit more tricky.

## Prerendering vs. SPA Mode

This was the part I initially struggled with, so I wanted to share what I've learned.

**Prerendering** (`export const prerender = true`) runs `load` functions once, during `npm run build`, in a temporary Node process. While building, it calls my backend endpoints (in the `load` methods) and the result gets written directly into the generated HTML file. After the build finishes, the server serves the site with preloaded data. This data is locked in, and won't change again until I rebuild and redeploy. This means for a page showing a daily puzzle, every visitor gets served whatever puzzle existed at build time.

**SPA mode** (`export const ssr = false`) skips running `load` functions at build time entirely. Instead, the adapter builds one generic HTML file (the `index.html`) plus the JS bundle. When someone visits any page, the server just returns that HTML file regardless of the URL. The browser then downloads the JS, Svelte figures out which route it's on, and then runs the `load` function. This allows the data to always be fresh in the pages.

## What I Changed

Rather than having the Svelte config that SvelteKit ships with:

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

**`src/routes/+layout.ts`**

```ts
export const ssr = false;

export const load: LayoutLoad = async ({ fetch }) => {
	...
};
```

### Notes:

- **fallback: 'index.html'** is what makes SPA (Single Page Application) mode deployable on a static host. Without prerendered routes from SvelteKit's Node server, there's no HTML file to serve per URL, so every path falls back to `index.html`, and client-side routing takes care of the rest.
- **strict: true** makes the build fail if any route isn't covered by prerendering or the fallback. That way we can still take advantage of `error.svelte` pages.
- **ssr = false** (Server Side Rendering) is the actual switch that makes `load` functions run only in the browser. It applies everywhere, like `npm run dev`, `vite preview`, and the deployed build. This allows dev mode to behave the same way production does.

## Takeaway

I struggled with understanding how Svelte handled rendering at first. I had to read a lot into the documentation to learn about SSR and SPA. In the end however, I was left with static website that preformed close to how SvelteKit with its Node server preformed. I am also able to host both my frontend and backend on one service. I'm happy with the result, but interested to see if this will cause problems in the future.
