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


---


---

## `<form>`

While this above method is useful for single actions, most pages come with multiple actions. You are able to assign certain requests with names, like so:

*+page.server.js*
```
export const actions = {
	create: async ({ cookies, request }) => {
		const data = await request.formData();
		db.createTodo(cookies.get('userid'), data.get('description'));
	},

	delete: async ({ cookies, request }) => {
		const data = await request.formData();
		db.deleteTodo(cookies.get('userid'), data.get('id'));
	}
};
```

*+page.svelte*
```
<div class="centered">
	<h1>todos</h1>

	<form method="POST" action="?/create">
		<label>
			add a todo:
			<input
				name="description"
				autocomplete="off"
			/>
		</label>
	</form>

	<ul class="todos">
		{#each data.todos as todo (todo.id)}
			<li>
				<form method="POST" action="?/delete">
					<input type="hidden" name="id" value={todo.id} />
					<span>{todo.description}</span>
					<button aria-label="Mark as complete"></button>
				</form>
			</li>
		{/each}
	</ul>
</div>
```
For context, you can put `default: async` if you need a default call for a form. Naming allows you to specify which action you want to run. **Default actions cannot coexist with named actions**.

With inputs, there needs to be input validation. To prevent users from sending nonsensical data, SvelteKit offers a way to validate this data using `@sveltejs/kit`

*+page.server.js*
```
import { fail } from '@sveltejs/kit'

---

export const actions = {
	create: async ({ cookies, request }) => {
		const data = await request.formData();
		try {
			db.createTodo(cookies.get('userid'), data.get('description'));
		} catch (error) {
			return fail(422, {
				description: data.get('description'),
				error: error.message
			})
		}
	},

	delete: async ({ cookies, request }) => {
		const data = await request.formData();
		db.deleteTodo(cookies.get('userid'), data.get('id'));
	}
};
```

*+page.svelte*
```
<script>
	let { data, form } = $props();
</script>

---

{#if form?.error}
	<p class="error">{form.error}</p>
{/if}
```

As you can see, this library offers an easy way of sending status messages to the client. `fail` isn't the only return type you can pass from an action. Showing a success message can be done via the `form` prop.

---

## `use:enhance`

The nice thing about the way `<form>` works is that it the user doesn't need to have JavaScript to run it. However, this causes the browser to cause full-page reloads. To allow JavaScript and update the page, use the `enhance` directive:

```
<script>
	import { enhance } from '$app/forms'
</script>

<form method="POST" action="?/create" use:enhance>
	<label>
		add a todo:
		<input
			name="description"
			value={form?.description ?? ''}
			autocomplete="off"
			required
		/>
	</label>
</form>

<ul class="todos">
	{#each data.todos as todo (todo.id)}
		<li>
			<form method="POST" action="?/delete" use:enhance>
				<input type="hidden" name="id" value={todo.id} />
				<span>{todo.description}</span>
				<button aria-label="Mark as complete"></button>
			</form>
		</li>
	{/each}
</ul>
```

This will:
- update the `form` prop
- invalidate all data on a successful response, causing `load` functions to re-run
- navigate to the new page on a redirect response
- render the nearest error page if an error occurs

Because this is also using JavaScript, we can add animations and transitions to this. See Svelte notes for transitions.

Alongside transitions, we can go further by providing pending states, callbacks, and optimistic UI. 

*+page.svelte*
```
<script>
	import { fly, slide } from 'svelte/transition';
	import { enhance } from '$app/forms';

	let { data, form } = $props();

	let creating = $state(false);
	let deleting = $state([]);
</script>

<div class="centered">
	<h1>todos</h1>

	{#if form?.error}
		<p class="error">{form.error}</p>
	{/if}

	<form 
		method="POST" 
		action="?/create" 
		use:enhance={() => {
			creating = true;

			return async ({ update }) => {
				await update();
				creating = false;
			};
		}}
	>
		<label>
			add a todo:
			<input
				disabled={creating}
				name="description"
				value={form?.description ?? ''}
				autocomplete="off"
				required
			/>
		</label>
	</form>

	<ul class="todos">
		{#each data.todos.filter((todo) => !deleting.includes(todo.id)) as todo (todo.id)}
			<li in:fly={{ y: 20 }} out:slide>
				<form 
					method="POST" 
					action="?/delete" 
					use:enhance={() => {
						deleting = [...deleting, todo.id];
						return async ({ update }) => {
							await update();
							deleting = deleting.filter((id) => id !== todo.id);
						}
					}}
				>
					<input type="hidden" name="id" value={todo.id} />
					<span>{todo.description}</span>
					<button aria-label="Mark as complete"></button>
				</form>
			</li>
		{/each}
	</ul>

	{#if creating}
		<span class="saving">saving...</span>
	{/if}
</div>
```
This seems like a lot so I'll explain what's happening. The first half of this is creating a simple loading icon where while the data is being saved, it will display a span that displays that it's saving to the database. It does this within the `use:enhance` directive, updating the `creating` value when loading and done executing. The second half of this deletes todos but we don't have to display a deleting message. Instead, we can show it deleting immediately while our backend takes care of the rest. 

This is easier to see with an example: https://svelte.dev/tutorial/kit/customizing-use-enhance

---

## `+server.js`

SvelteKit allows you to make more than just pages and actions. You can also create API routes within the `+server.js` file. These endpoints work just like your basic CRUD (`GET`, `PUT`, `POST`, `PATCH` and `DELETE`).

## `GET`

*roll/+server.js*
```
export function GET() {
	const number = Math.floor(Math.random() * 6) + 1;

	return new Response(number, {
		headers: {
			'Content-Type': 'application/json'
		}
	});
}
```

*+page.svelte*
```
<script>
	/** @type {number} */
	let number = $state();

	async function roll() {
		const response = await fetch('/roll');
		number = await response.json();
	}
</script>

<button onclick={roll}>Roll the dice</button>

{#if number !== undefined}
	<p>You rolled a {number}</p>
{/if}
```

Super easy to make APIs in SvelteKit. Continuing off of this example, request handlers must return a `Response` object. Since it's common to return JSON, SvelteKit makes it easy to do so:

```
import { json } from '@sveltejs/kit'

export function GET() {
	const number = Math.floor(Math.random() * 6) + 1;

	return json(number);
}
```

## `POST`

Below is an example of a `POST` handler:

*+page.svelte*
```
<label>
		add a todo:
	<input
		type="text"
		autocomplete="off"
		onkeydown={async (e) => {
			if (e.key !== 'Enter') return;

			const input = e.currentTarget;
			const description = input.value;
			
			const response = await fetch('/todo', {
				method: 'POST',
				body: JSON.stringify({ description }),
				headers: {
					'Content-Type': 'application/json'
				}
			})

			const { id } = await response.json();

			const todos = [...data.todos, {
				id,
				description
			}];

			data = { ...data, todos };

			input.value = '';
		}}
	/>
</label>
```

*todo/+server.js*
```
import { json } from '@sveltejs/kit';
import * as database from '$lib/server/database.js';

export async function POST({ request, cookies }) {
	const { description } = await request.json();

	const userid = cookies.get('userid');
	const { id } = await database.createTodo({ userid, description });

	return json({ id }, { status: 201 });
}
```

As you can see, it's very easy to set up API's for SvelteKit. The nice thing about it is that you don't have to worry about naming, because it's built into the routing. I'm sure we'll go into API's you can use throughout your code without having to copy paste.
