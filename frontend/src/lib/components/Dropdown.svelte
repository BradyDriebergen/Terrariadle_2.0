<script lang="ts">
	let { selectItem, itemList, itemName } = $props();

	let input = $state('');
	let filtered = $derived([...itemList]);
	let dropdownIndex = $state(-1);
	let itemElements = $state<(HTMLButtonElement | null)[]>([]); // used for smooth scrolling

	// Effect for when input changes
	$effect(() => {
		const query = input.toLowerCase();
		filtered = itemList
			.filter((item: { name: string }) => item.name.toLowerCase().includes(query))
			.slice(0, 20);
		dropdownIndex = -1;
	});

	// Keep the selected item in view when dropdownIndex changes
	$effect(() => {
		if (dropdownIndex < 0) return;

		const el = itemElements[dropdownIndex];
		if (el) {
			el.scrollIntoView({ behavior: 'smooth', block: 'nearest', inline: 'nearest' });
		}
	});

	function onKeyDown(event: KeyboardEvent) {
		if (!input) return;
		if (event.key === 'Tab') {
			dropdownIndex++;
			event.preventDefault();
		} else if (event.key === 'ArrowDown') {
			dropdownIndex = Math.min(dropdownIndex + 1, filtered.length - 1);
			event.preventDefault();
		} else if (event.key === 'ArrowUp') {
			dropdownIndex = Math.max(dropdownIndex - 1, 0);
			event.preventDefault();
		} else if (event.key === 'Enter' && dropdownIndex >= 0 && filtered.length !== 0) {
			selectItem(filtered[dropdownIndex].id);
			input = '';
			event.preventDefault();
		}
	}
</script>

<div class="wrapper">
	<input
		type="text"
		placeholder="Type any {itemName} to guess..."
		bind:value={input}
		onkeydown={onKeyDown}
	/>

	{#if input}
		<div class="dropdown">
			{#if filtered.length !== 0}
				{#each filtered as item, i (item.id)}
					<button
						class="item"
						class:selected={i === dropdownIndex}
						bind:this={itemElements[i]}
						onclick={() => {
							selectItem(item.id);
							input = '';
						}}
					>
						<img src={`/${itemName}s/${item.path}`} alt={item.path} />
						<span>{item.name}</span>
					</button>
				{/each}
			{:else}
				<span class="no-items">No results...</span>
			{/if}
		</div>
	{/if}
</div>

<style>
	.wrapper {
		position: relative;
		display: inline-block;
		margin: 0;
	}

	input {
		background-color: var(--color-button);
		padding: 10px;
		width: 220px;
		height: 20px;
		font-size: 15px;
		text-align: left;

		border-radius: 5px;
		border: thin solid black;
		text-shadow: none;
		text-shadow:
			0px 0px 3px #000,
			0px 0px 2px #000;
	}

	.dropdown {
		position: absolute;
		width: 240px;
		top: 100%;
		left: 0;
		background: var(--color-button);
		max-height: 240px;

		overflow-y: auto;
		scrollbar-width: none;
		-ms-overflow-style: none;

		border-radius: 8px;
		z-index: 50;
		border: thin solid black;
	}
	.dropdown::-webkit-scrollbar {
		display: none;
	}

	.item {
		display: flex;
		gap: 8px;
		padding: 6px;
		align-items: center;
		cursor: pointer;
		border-radius: 6px;

		background: none;
		width: 100%;
		border: none;
	}
	.item:hover,
	.item.selected {
		background: var(--color-lightblue);
	}
	.item span {
		font-size: 14px;
	}
	.item img {
		width: 35px;
		height: 35px;
		object-fit: contain;
	}

	.no-items {
		display: flex;
		padding: 10px;
		align-items: center;
		cursor: pointer;
		border-radius: 6px;
	}
</style>
