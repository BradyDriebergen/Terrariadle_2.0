  
`$state()` introduces reactivity. It declares that a variable can be changed and is kept in sync with the DOM
- `let numbers = $state([1, 2, 3, 4]);`
When you use the state rune, it triggers the reactivity in the site that updates the value. It does this through a proxy. If you need to update a state value without triggering this reactivity, you can use the `raw` property
- `numbers.raw[4] = 31;`
Note that this should be used sparingly, it breaks the reactive guarantees of Svelte.

---

`$derived()` extends state and allows changes made to state to also changed the derived.
- `let total = $derived(numbers.reduce((t, n) => t + n, 0));`

---

`$inspect()` allows you to automatically log a snapshot whenever state changes
- `console.log($state.snapshot(numbers));` *Only prints it once*
- `$inspect(numbers)` *Prints when numbers is changed*

---

`$effect()` works like useEffect in React. It's a reactive effect that runs automatically whenever the state it depends on changes.
```
$effect(() => {
	const id = setInterval(() => {
		elapsed += 1
	}, interval)
});
```
Basically, effects run once when the component mounts and again when the states within it changes. 
If you return a function from inside it, it's considered a clean up function, which runs before the effect runs again or when the component is destroyed.
```
$effect(() => {
	const id = setInterval(() => {
		elapsed += 1
	}, interval)

	return () => {
		clearInterval(id);
	}
});
```
Here's where you should use them:
- Setting up or clearing intervals/timeouts
- Changing page layout when tracking window width
- Subscribing/unsubscribing to events or APIs
- Logging, network requests, or updating browser APIs (e.g. `localStorage`)
You can also have shared states between files. You would need a `<filename>.svelte.js` and declare them like the following:
```
export const counter = $state({
	count: 0
});
```
After this, you this state can be shared across all the files, you can import it by:
- `import { counter } from <filename>.svelte.js`

---

`$props` is used to pass in arguments from parents to children.
- *Child*
```
<script>
	let { answer } = $props();
</script>
  ``` 
-  *Parent*
```
<Nested answer={42} />
```
You can have default values as well:
- `let { answer = 'a mystery'} = $props();`
If you have a lot of props, you can simplify it
```
<script>
	import PackageInfo from './PackageInfo.svelte';

	const pkg = {
		name: 'svelte',
		version: 5,
		description: 'blazing fast',
		website: 'https://svelte.dev'
	};
</script>

<PackageInfo {...pkg} />
```
You can also pass handlers as props:
```
<Stepper 
	increment={() => value += 1}
	decrement={() => value -= 1}	
/>
```
You can use the simplify way to pass in these values too.

---

An alternative to props is `Context`. Context is a way for components to 'talk' to each other without passing around data as props. It uses a Key/Value system.
```
import { setContext } from 'svelte'

setContext('canvas', { addItem })

function addItem(fn) {
	$effect(() => {
		items.add(fn);
		return () => items.delete(fn);
	});
}

---

import { getContext } from 'svelte';

getContext('canvas').addItem(draw);
```
Contexts can include anything, including reactive `$state`. This becomes very useful when you don't want to have to pass elements through layers of components. Contexts can be called in any child of the tree.

---

`if` blocks can be used like the following:
```
{#if <condition>}
	do something...
{:else}
	do something else
{/if}
```
`{#...}` opens a block. `{/...}` closes a block. `{:...}` _continues_ a block.

---

`else` blocks can be used like the following:
```
{#each colors as color}
	<button
		style="background: {color}"
>	</button>
{/each}
```
You can also use the item index for it:
```
{#each colors as color, i}
	<button
		style="background: {color}"
	>{i + 1}</button>
{/each}
```

---

`await` blocks allow data to be initialized before it displays.
*App.svelte*
```
<script>
	import { roll } from './utils.js';

	let promise = $state(roll());
</script>

<button onclick={() => promise = roll()}>
	roll the dice
</button>

{#await promise}
	<p>...rolling</p>
{:then number}
	<p>you rolled a {number}</p>
{:catch error}
	<p style="color: red;">{error.message}</p>
{/await}
```
*utils.js*
```
export async function roll() {
	// Fetch a random number from 1 to 6
	// (with a delay, so that we can see it)
	return new Promise((fulfil, reject) => {
		setTimeout(() => {
			// simulate a flaky network
			if (Math.random() < 0.3) {
				reject(new Error('Request failed'));
				return;
			}

			fulfil(Math.floor(Math.random() * 6) + 1);
		}, 1000);
	});
}
```

---

Event handlers can handle any event in the DOM, this can vary between onclick, onpointermove, and more. Here is an example of this:
```
<script>
	let m = $state({ x: 0, y: 0 });

	function onpointermove(event) {
		m.x = event.clientX;
		m.y = event.clientY;
	}
</script>

<div {onpointermove}>
	The pointer is at {Math.round(m.x)} x {Math.round(m.y)}
</div>
```
You can also do inline calls:
```
<div onpointermove={(event) => {
		m.x = event.clientX;
		m.y = event.clientY;
}}>
	The pointer is at {Math.round(m.x)} x {Math.round(m.y)}
</div>
```

---

`bind:` extends on state where whenever the value of state is changed, it will propagate to the rest of the component
```
<script>
	let name = $state('world');
</script>

<input bind:value={name} />

<h1>Hello {name}!</h1>
```
bind is a simpler way of making an effect or derived updating the rest of the component. It can be used in almost all inputs. You can also use it for updating elements:
```
<select
	bind:value={selected}
	onchange={() => (answer = '')}
>
```

Media elements are another useful feature for Svelte. It's a bit complicated to put in this note but the jist of it is you can assign bindings to media's elements
```
<audio
	{src}
	bind:currentTime={time}
	bind:duration
	bind:paused
	onended={() => {
		time = 0;
	}}
></audio>
```
There is a lot more you can do with this. I recommend reading https://svelte.dev/tutorial/svelte/media-elements for more context.

There are also many different bindings you can observe
```
<div bind:clientWidth={w} bind:clientHeight={h}>
	<span style="font-size: {size}px" contenteditable>
		edit this text
	</span>

	<span class="size">{w} x {h}px</span>
</div>
```
Do more research when looking at properties like this. These ones in particular are read only.

`bind.this` is another useful way to get a readonly binding to an element in your component.
```
<script>
	import { paint } from './gradient.js';

	let canvas;
	
	$effect(() => {
		const context = canvas.getContext('2d');

		let frame = requestAnimationFrame(function loop(t) {
			frame = requestAnimationFrame(loop);
			paint(context, t);
		});

		return () => {
			cancelAnimationFrame(frame);
		};
	});
</script>

<canvas bind:this={canvas} width={32} height={32}></canvas>

<style>
	canvas {
		position: fixed;
		left: 0;
		top: 0;
		width: 100%;
		height: 100%;
		background-color: #666;
		mask: url(./svelte-logo-mask.svg) 50% 50% no-repeat;
		mask-size: 60vmin;
		-webkit-mask: url(./svelte-logo-mask.svg) 50% 50% no-repeat;
		-webkit-mask-size: 60vmin;
	}
</style>
```
What exactly is happening here is the `$effect` is trying to create a canvas context. We bind the `canvas` value to it so the effect knows where to point. This one's a bit harder to understand but it makes sense when you realize that `$effect` isn't connected to the canvas element other than the bind. Another benefit of this is the value canvas is undefined until the component has mounted, so it will render only when it's shown on the DOM.

---

`multiple` is a way to generate a new type of checkbox. Useful for saving multiple values to an array and not having the bloat of a checkbox.
```
<select multiple bind:value={flavours}>
{#each ['cookies and cream', 'mint choc chip', 'raspberry ripple'] as flavour}
	<option>{flavour}</option>
{/each}
</select>
```
Not a lot of uses but good to know.

---

Classes can be easily changed by using ternary operators
```
<button
	class="card {flipped ? 'flipped' : ''}"
	onclick={() => flipped = !flipped}
>
<script>
.card.flipped {
	transform: rotateY(0);
}
</script
```
You can also do the same with this format
```
<button
	class={["card", { flipped }]}
	onclick={() => flipped = !flipped}
>
```
Check out this cool animation you can do: https://svelte.dev/tutorial/svelte/classes

---

`use:` creates a way to apply an "action" to a element.
```
<div use:myAction></div>
```
Here, `myAction` is a function (called an _action_) that runs when the element is created and can set up custom logic for it. Here is an example of what you can do with it
```
function myAction(node) {
  // do something with the node
  node.style.color = 'red';

  return {
    destroy() {
      // optional cleanup
      node.style.color = '';
    }
  };
}
```
You can also do the same with parameters:
```
<script>
	import tippy from 'tippy.js';

	let content = $state('Hello!');

	function tooltip(node, fn) {
		$effect(() => {
			const tooltip = tippy(node, fn());

			return tooltip.destroy;
		});
	}
</script>

<input bind:value={content} />

<button use:tooltip={() => ({ content })}>
	Hover me
</button>
```

You can also bind properties within a `each` block. This is useful for dealing with lists of items you might need to modify:
```
<script>
	let todos = $state([
		{ done: false, text: 'finish Svelte tutorial' },
		{ done: false, text: 'build an app' },
		{ done: false, text: 'world domination' }
	]);

	let remaining = $derived(todos.filter((t) => !t.done).length);
</script>

<div class="centered">
	<h1>todos</h1>

	<ul class="todos">
		{#each todos as todo}
			<li class={{ done: todo.done }}>
				<input
					type="checkbox"
					bind:checked={todo.done}
				/>

				<input
					type="text"
					placeholder="What needs to be done?"
					bind:value={todo.text}
				/>
			</li>
		{/each}
	</ul>
```

---

Transitions are very useful for smoothly introducing elements into the DOM. Here is an example of a `fade` transition:
```
<script>
	import { fade } from 'svelte/transition';
	
	let visible = $state(true);
</script>

<label>
	<input type="checkbox" bind:checked={visible} />
	visible
</label>

{#if visible}
	<p transition:fade>
		Fades in and out
	</p>
{/if}
```
Here is an example of fly:
```
<script>
	import { fly } from 'svelte/transition';

	let visible = $state(true);
</script>

<label>
	<input type="checkbox" bind:checked={visible} />
	visible
</label>

{#if visible}
	<p transition:fly={{ y: 200, duration: 2000 }}>
		Fades in and out
	</p>
{/if}
```
Here is more examples of transitions: https://svelte.dev/docs/svelte/svelte-transition

You can also use multiple transitions at once:
```
{#if visible}
	<p in:fly={{ y: 200, duration: 2000 }} out:fade>
		Flies in, fades out
	</p>
{/if}
```

This in and out functionality also works with *defering* transitions. This means that transitions can happen between multiple elements.
```
<ul class="todos">
	{#each todos as todo (todo.id)}
		<li
			class={{ done: todo.done }}
			in:receive={{ key: todo.id }}
			out:send={{ key: todo.id }}
		>
			<label>
				<input type="checkbox" bind:checked={todo.done}/>
				<span>{todo.description}</span>
				<button onclick={() => remove(todo)} aria-label="Remove"></button>
			</label>
		</li>
	{/each}
</ul>

---

import { crossfade } from 'svelte/transition';
import { quintOut } from 'svelte/easing';

export const [send, receive] = crossfade({
	duration: (d) => Math.sqrt(d * 200),

	fallback(node, params) {
		const style = getComputedStyle(node);
		const transform = style.transform === 'none' ? '' : style.transform;

		return {
			duration: 600,
			easing: quintOut,
			css: (t) => `
				transform: ${transform} scale(${t});
				opacity: ${t}
			`
		};
	}
});
```
See, this works by creating a pair of transitions called `send` and `recieve`. When an element is sent, it looks for a corresponding element being received. When found, it generates a transition that transforms the elements to each others' positions.
For a more in-depth explanation, go to this link: https://svelte.dev/tutorial/svelte/deferred-transitions
You can also add animations to these transitions
```
<li
	class={{ done: todo.done }}
	in:receive={{ key: todo.id }}
	out:send={{ key: todo.id }}
	animate:flip= {{ duration: 200 }}
>
```

If you really want to get creative, you can create your own transitions:
```
function spin(node, { duration }) {
	return {
		duration,
		css: (t, u) => {
			const eased = elasticOut(t);

			return `
				tranform: scale(${eased}) rotate(${eased * 1080}deg);
				color: hsl(
					${Math.trunc(t * 360)},
					${Math.min(100, 1000 * u)}%,
					${Math.min(50, 500 * u)}%
				);`
		}
	};
}
```
While you should generally use CSS for transitions, some effects can't be achieved without JavaScript.
```
function typewriter(node, { speed = 1 }) {
	const valid = node.childNodes.length === 1 && node.childNodes[0].nodeType === Node.TEXT_NODE;

	if (!valid) {
		throw new Error(`This transition only works on elements with a single text node child`);
	}

	const text = node.textContent;
	const duration = text.length / (speed * 0.01);

	return {
		duration,
		tick: (t) => {
			const i = Math.trunc(text.length * t);
			node.textContent = text.slice(0, i);
		}
	};
}
```

Another cool thing you can do is tell when the intro and outro begin/end.
```
<p
	transition:fly={{ y: 200, duration: 2000 }}
	onintrostart={() => status = 'intro started'}
	onoutrostart={() => status = 'outro started'}
	onintroend={() => status = 'intro ended'}
	onoutroend={() => status = 'outro ended'}
>
```

If you plan on having transitions between multiple elements in one block, you can apply the `global` variable to affect every element inside the block:
```
{#if showItems}
	{#each items.slice(0, i) as item}
		<div transition:slide|global>
			{item}
		</div>
	{/each}
{/if}
```

`Key` blocks destroy and recreate their contents when the value of an expression changes. This is useful for adding transitions between elements that are constantly changing:
```
{#key i}
	<p in:typewriter={{ speed: 10 }}>
		{messages[i] || ''}
	</p>
{/key}
```
Super cool loading example here: https://svelte.dev/tutorial/svelte/key-blocks

---

Classes are one of those things that will be somewhat useful when creating svelte projects. Thankfully, svelte created a way to make classes a little more bearable. Classes have the ability to use `$state()` and `$derived()` within them, making it easier to update the DOM and separate out functionality.
```
class Box {
	width = $state(0);
	height = $state(0);
	area = $derived(this.width * this.height);

	constructor(width, height) {
		this.width = width;
		this.height = height;
	}

	embiggen(amount) {
		this.width += amount;
		this.height += amount;
	}
}
```
You can also declare private values easier too. Using the `#` symbol makes values private
```
class Box {
	#width = $state(0);
	#height = $state(0);
	area = $derived(this.#width * this.#height);

	constructor(width, height) {
		this.width = width;
		this.height = height;
	}

	get width() {
		return this.#width;
	}

	get height() {
		return this.#height;
	}

	set width(value) {
		this.#width = Math.max(0, Math.min(MAX_SIZE, value));
	}

	set height(value) {
		this.#height = Math.max(0, Math.min(MAX_SIZE, value));
	}

	embiggen(amount) {
		this.width += amount;
		this.height += amount;
	}
}
```

---

Svelte ships with several reactive libraries that allow easier declarations of things like `Map`, `Set`, `Date`, `URL` and `URLSearchParams`. 
```
<script>
	import { SvelteDate } from 'svelte/reactivity'
	let date = new SvelteDate();

	const pad = (n) => n < 10 ? '0' + n : n;

	$effect(() => {
		const interval = setInterval(() => {
			date.setTime(Date.now());
		}, 1000);

		return () => {
			clearInterval(interval);
		};
	});
</script>

<p>The time is {date.getHours()}:{pad(date.getMinutes())}:{pad(date.getSeconds())}</p>
```
The same can be accomplished with `$state(new Data())` but this makes it a little more concise. 

---

Prior to runes, Svelte use to use `stores` to handle reactive states. While the primary way of doing this now is through `$state()`, there is still times you'll run into it, including SvelteKit for now. It's worth knowing it in case you see it.
```
<script>
	import { writable } from 'svelte/store';
	export const count = writable(0);
</script>
```
Apparently, you can make custom stores for when you need it. See https://svelte.dev/docs/svelte/stores for more info.

---

`#snippet` is a useful tool for creating reusable components without having to create a new file for it. It's helpful for small chunks of code that aren't worth creating a whole component for. 
```
<tbody>
	{#snippet monkey(emoji, description)}
		<tr>
			<td>{emoji}</td>
			<td>{description}</td>
			<td>\u{emoji.charCodeAt(0).toString(16)}\u{emoji.charCodeAt(1).toString(16)}</td>
			<td>&amp#{emoji.codePointAt(0)}</td>
		</tr>
	{/snippet}

	{@render monkey('🙈', 'see no evil')}
	{@render monkey('🙉', 'hear no evil')}
	{@render monkey('🙊', 'speak no evil')}
</tbody>
```
Pretty good choice for elements that you can't foreach over. Also a way to make more concise code. You can also declare snippets in a file's `props`.
```
<FilteredList
	data={colors}
	field="name"
	{header}
	{row}
></FilteredList>

{#snippet header()}
	<header>
		<span class="color"></span>
		<span class="name">name</span>
		<span class="hex">hex</span>
		<span class="rgb">rgb</span>
		<span class="hsl">hsl</span>
	</header>
{/snippet}

{#snippet row(d)}
	<div class="row">
		<span class="color" style="background-color: {d.hex}"></span>
		<span class="name">{d.name}</span>
		<span class="hex">{d.hex}</span>
		<span class="rgb">{d.rgb}</span>
		<span class="hsl">{d.hsl}</span>
	</div>
{/snippet}
```
You can also do this using the `children` property. If a snippet doesn't have any parameters, you don't need to include the `snippet` tag and instead just the html element itself.
```
<FilteredList
	data={colors}
	field="name"
>
	<header>
		<span class="color"></span>
		<span class="name">name</span>
		<span class="hex">hex</span>
		<span class="rgb">rgb</span>
		<span class="hsl">hsl</span>
	</header>

	{#snippet row(d)}
		<div class="row">
			<span class="color" style="background-color: {d.hex}"></span>
			<span class="name">{d.name}</span>
			<span class="hex">{d.hex}</span>
			<span class="rgb">{d.rgb}</span>
			<span class="hsl">{d.hsl}</span>
		</div>
	{/snippet}
</FilteredList>

---

<script>
	let { data, field, children, row } = $props();
</script>

<div class="header">
	{@render children()}
</div>
```

---

Motion is a useful way to show that a value is changing. There are built in classes like `Tween` that help add motion to your user interfaces.
```
<script>
	import { Tween } from 'svelte/motion'
	import { cubicOut } from 'svelte/easing'
	
	let progress = new Tween(0, {
		duration: 400,
		easing: cubicOut
	});
</script>

<progress value={progress.current}></progress>

<button onclick={() => (progress.target = 0)}>
	0%
</button>
```
This makes a progress bar that smoothly moves when the progress value is changed. There are many other attributes to `Tween` to help with motions.

`Spring` is another handy one for creating smoother transitions for things like the cursor or movement.
```
<script>
import { Spring } from 'svelte/motion';
	
	let coords = new Spring({ x: 50, y: 50 }, {
		stiffness: 0.1,
		damping: 0.25
	});
	let size = new Spring(10);
</script>

<svg
	onmousemove={(e) => {
		coords.target = { x: e.clientX, y: e.clientY };
	}}
	onmousedown={() => (size.target = 30)}
	onmouseup={() => (size.target = 10)}
	role="presentation"
>
	<circle
		cx={coords.current.x}
		cy={coords.current.y}
		r={size.current}
	></circle>
</svg>
```

---

Bit of a niche change but the `contenteditable` attribute supports `textContext` and `innerHTML` bindings.
```
<div bind:innerHTML={html} contenteditable></div>

<pre>{html}</pre>
```

--- 

`$bindable()` is a way we can bind to properties in component props. This way, we can easily pass data from one another without a bunch of inputs and outputs. This is a great option for things like keyboards and keypads.
```
<Keypad bind:value={pin} {onsubmit} />

---

// Keypad.svelte
<script>
	let { value = $bindable(''), onsubmit } = $props();

	const select = (num) => () => (value += num);
	const clear = () => (value = '');
</script>
```

You can also bind component instances using the same format.
```
<Canvas bind:this={canvas} color={selected} size={size} />

---

<script>
	export function clear() {
		context.clearRect(0, 0, canvas.width, canvas.height);
	}
</script>
```
This is useful when you have buttons on a parent component that you want to call functions in the child component. Very useful. See https://svelte.dev/tutorial/svelte/component-this for more details.

---

Svelte has many different special elements that you can use for your projects.

`svelte:window` adds an event listener to the window object, allowing you to do things like this
```
<script>
	let key = $state();
	let keyCode = $state();

	function onkeydown(event) {
		key = event.key;
		keyCode = event.keyCode;
	}
</script>

<svelte:window {onkeydown} />

<div style="text-align: center">
	{#if key}
		<kbd>{key === ' ' ? 'Space' : key}</kbd>
		<p>{keyCode}</p>
	{:else}
		<p>Focus this window and press any key</p>
	{/if}
</div>
```

You can also bind certain properties to it:
```
<script>
	let y = $state(0);
</script>

<svelte:window bind:scrollY={y}/>

<span>depth: {y}px</span>
```

Here is a list of all the properties you can bind to it:
- `innerWidth`
- `innerHeight`
- `outerWidth`
- `outerHeight`
- `scrollX`
- `scrollY`
- `online`

---

The `svelte:document` allows you to listen for events that fire on `document`. This is useful for things like `selectionchange`, which doesn't fire on `window`.

```
<script>
	let selection = $state('');

	const onselectionchange = (e) => {
		selection = document.getSelection().toString();
	};
</script>

<svelte:document {onselectionchange}/>

<h1>Select this text to fire events</h1>
<p>Selection: {selection}</p>
```
*The above example displays the selected text, pretty cool trick.*

---

`svelte:body` is the same as the two above but listen to events on the `document.body` element.
```
<script>
	import kitten from './kitten.png';

	let hereKitty = $state(false);
</script>

<svelte:body 
	onmouseenter={() => hereKitty = true}
	onmouseleave={() => hereKitty = false}
/>

<img
	class={{ curious: hereKitty }}
	alt="Kitten wants to know what's going on"
	src={kitten}
/>
```

---

`svelte:head` allows you to dynamically insert elements into the document's `<head>`. It's useful for things like `<title>` and `<meta>` tags.
```
<script>
	const themes = ['margaritaville', 'retrowave', 'spaaaaace', 'halloween'];
	let selected = $state(themes[0]);
</script>

<svelte:head>
	<link rel="stylesheet" href="/tutorial/stylesheets/{selected}.css" />
</svelte:head>

<h1>Welcome to my site!</h1>
```
*tutorial/stylesheets is a good place for examples of CSS*

https://svelte.dev/tutorial/svelte/svelte-head

---

`svelte:element` is useful in nicher situations. for example, rather than having a long `{#if...}` block:
```
{#if selected === 'h1'}
	<h1>I'm a <code>&lt;h1&gt;</code> element</h1>
{:else}
	<p>TODO others</p>
{/if}
```
You can instead have one `svelte:element`:
```
<svelte:element this={selected}>
	I'm a <code>&lt;{selected}&gt;</code> element
</svelte:element>
```

---

`svelte:boundry` is a useful way of preventing errors from leaving your app. I've never thought about error handling in my frontend but it seems like a useful tool for if things break in a certain area.

```
<svelte:boundary onerror={(e) => console.error(e)}>
	<FlakyComponent />

	// Can be failed(error) or failed(error, reset)
	{#snippet failed(error, reset)}
		<p>Oops! {error.message}</p>
		<button onclick={reset}>Reset</button>
	{/snippet}
</svelte:boundary>
```

This example is going to take some explaining. `<FlakyComponent />` contains a button that sets a value to null, breaking the component. `svelte:boundry` can catch this error and handle it accordingly. First, you can simply add a `#snippet` that catches this error and a `failed(error)` to have the error message. You can also add `reset` to the parameters to add the ability to reset the component. Lastly, you can add a `onerror` call, which is useful for sending information about the error to a reporting service, or adding UI outside of the error boundary itself.

---

While the `<script>` tag is used to run code when a component is initialized, sometimes, we need to run some code outside of the component. To do this, we need to declare a `<script module>` tag.
```
<script module>
	let current;
</script>
```
This acts kind of like a `protected` value in Java. Each component has access to this one variable. It's good for tracking multiple components that need to talk to each other.

You can also export functions in the `<script module>`
```
<script module>
	let current;

	export function pauseAll() {
		current?.pause();
	}
</script>

---

import AudioPlayer, { pauseAll } from './AudioPlayer.svelte';
```

---

Confetti example: https://svelte.dev/tutorial/svelte/congratulations