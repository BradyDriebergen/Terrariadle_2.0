
`$state()` introduces reactivity. It declares that a variable can be changed and is kept in sync with the DOM
- `let numbers = $state([1, 2, 3, 4]);`

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