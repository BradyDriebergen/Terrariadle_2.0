## What is SvelteKit?

SvelteKit builds on Svelte by providing a complete application framework that handles routing, server-side rendering, data fetching, TypeScript integration, deployment, and more out of the box. It combines the speed and SEO benefits of server-rendered apps with the smooth navigation of single-page apps, adapting seamlessly to your needs. A typical SvelteKit project includes familiar configuration files (`package.json`, `svelte.config.js`, `vite.config.js`) and organized directories for source code (`src`) and static assets, making it simple to start small and scale up as your app grows.

---

## Routing

SvelteKit uses filesystem-based routing, meaning pages and their dependencies are contained within folders. Below is an example of how pages work in SvelteKit:

```
.
└── routes/
    ├── about/
    │   └── +page.svelte
    └── +page.svelte
```
As you can see, this site contains two pages. The landing page, contained within `routes/`, and the `about/` page.

*Unlike traditional multi-page apps, navigating to `/about` and back updates the contents of the current page, like a single-page app.*

---

## `+layout.svelte`

Different parts of your app will often share common UI (like headers and footers). Rather than making a separate component called on each page, we can make a `+layout.svelte` page. Below is an example of one:

```
.
└── Routes/
    ├── About/
    │   └── +page.svelte
    ├── +layout.svelte
    └── +page.svelte
```

*+layout.svelte*
```
<script>
	let { children } = $props();
</script>

<nav>
	<a href="/">home</a>
	<a href="/about">about</a>
</nav>

{@render children()}
```
What this does is applies all the elements to every route in the same directory. `children()` is the content within each of the `+page.svelte` in each of the directories. Not only is this useful for Headers and Footers, but if multiple routes share the same elements, it's a convenient way to repeat these elements without a component.

---

## `/[slug]`

Sometimes, we want to create multiple pages for things like blogs or articles. A problem we face is that we would need to create a new directories if we had something like `/blog/one`, `/blog/two`, `/blog/three`, etc. Rather than having a ton of folders, we can use the `[slug]` directory:

```
.
└── Routes/
    ├── blog/
    │   ├── [slug]/
    │   │   └── +page.svelte
    │   └── +page.svelte
    ├── +layout.svelte
    └── +page.svelte
```
This allows for pages like `/blog/one`, `/blog/two`, and `/blog/three` to be accessible. However, in its current state, it only displays what's in `+page.svelte` within `[slug]`. We'll go over how to change these in the future.

---

## Loading 

Every page of your app can declare a `load` function in a `+page.server.js`. This module ever only runs on the server, including for client-side navigations. This is useful for when you want to initialize data starting out or call to an API.

`+page.server.js`: for **secure, private, server-only preloading**
`+page.js`: for **public, preloaded data available to both server and client**
`+page.svelte`: for **runtime, user-triggered interactions and client-only fetches**