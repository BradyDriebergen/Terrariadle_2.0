<script lang="ts">
	import { cubicInOut } from "svelte/easing";
	import { slide } from "svelte/transition";

	const clues = [
		"Shroomite pickaxe", 
		"Hardmode swimming enemy", 
		"Buff gained from Palladium Armor", 
		"Hellstone tools prefix", 
		"Tool used to break blocks", 
		"Galatic material for rangers", 
		"Wet weather event"
	];

	const chunks = [
		"DIG", "GIN", "GC", "LAW", "ANGL", 
		"ERFI", "SH", "RAPI", "DHEA", "LING", 
		"MO", "LT", "EN", "PIC", "KA", "XE", 
		"VOR", "TEX", "RA", "IN"
	];

	let finished: boolean = $state(false);
	let selectedChunks: string[] = $state([]);
	let input: string | null = $derived(selectedChunks.length > 0 ? selectedChunks.join('') : null);

	function selectChunk(chunk: string) {
		if (selectedChunks.includes(chunk)) {
			selectedChunks = selectedChunks.filter(c => c !== chunk);
		} else {
			selectedChunks.push(chunk);
		}
	}
</script>

{#if !finished}
	<div class="title-box" out:slide={{ duration: 700, easing: cubicInOut }}>
		<h2>Wordaria</h2>
		<p>Use clues to decipher all 7 words!</p>
	</div>
{/if}

<div class="game-window">
	<div class="clue-box">
		{#each clues as clue}
			<div class="clue-row">
				<span class="clue">{clue}</span>
				<span class="letter-count">6 letters</span>
			</div>
		{/each}
	</div>

	<div class="selection-menu">
		<button>Shuffle</button>
		<button>Hint</button>
		<button onclick={() => selectedChunks = []}>Clear</button>
	</div>

	<div class="input-box">
		<span style:color={selectedChunks.length > 0 ? "" : "Grey"}>{input ?? "Start by forming some words"}</span>
		<button onclick={() => selectedChunks.pop()}>X</button>
	</div>
</div>

<div class="chunk-grid">
	{#each chunks as chunk}
		{#if selectedChunks.includes(chunk) || chunk === ""}
			<div></div>
		{:else}
			<button 
				class="chunk" 
				onclick={() => selectChunk(chunk)} 
				disabled={selectedChunks.length >= 4}
			>
				{chunk}
			</button>
		{/if}
	{/each}
</div>

<style>
	.title-box {
		background-color: var(--color-backgroundblue);
		width: fit-content;
		text-align: center;
		margin: 30px auto 15px auto;
		padding: 0 15px;

		border-radius: 15px;
		border: thin solid black;
	}

	.title-box h2 {
		background-color: var(--color-lightblue);
		width: fit-content;
		margin: auto;
		margin-top: -30px;
		margin-bottom: 10px;
		padding: 10px 20px;

		border-radius: 15px;
		border: 2px solid black;
	}
	
	.game-window {
		background-color: var(--color-backgroundblue);
		width: fit-content;
		margin: 0 auto 15px auto;
		padding: 10px;
		border-radius: 15px;
		border: 2px solid black;
	}

	.clue-box {
		background-color: var(--color-button);
		width: 350px;
		margin: 0 auto 15px auto;
		padding: 4px 0;

		border-radius: 15px;
		border: 2px solid black;
	}

	.clue-row {
		display: flex; 
		justify-content: space-between;
		padding: 4px;
	}

	.selection-menu {
		margin: 15px auto;
	}

	.selection-menu button {
		background-color: var(--color-button);
		border-radius: 8px;
		border: 2px solid black;
		padding: 5px 10px;
		font-size: 16px;
	}

	.input-box {
		background-color: var(--color-button);
		border-radius: 8px;
		border: 2px solid black;
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 5px 10px;
		font-size: 16px;
		margin: 5px auto;
		width: 250px;
		min-height: 20px;
		cursor: default;
	}

	.input-box button {
		background-color: var(--color-lightblue);
		border-radius: 8px;
		border: 2px solid black;
		padding: 2px 6px;
		transition: background-color 0.1s ease;
	}

	.input-box button:hover {
		background-color: var(--color-button);
		cursor: pointer;
	}

	.clue-row:nth-child(even) {
		background-color: var(--color-lightblue);
	}

	.chunk-grid {
		display: grid;
		grid-template-columns: repeat(4, 1fr);
		grid-auto-rows: 1fr;

		width: fit-content;
		height: fit-content;
		margin: 0 auto;
		gap: 5px;
	}

	.chunk {
		background-color: var(--color-button);
		width: 80px;
		height: 50px;
		display: grid;
		place-items: center;

		border-radius: 8px;
		border: 2px solid black;

		padding: 10px;
		font-size: 18px;
		transition: background-color 0.1s ease;
	}

	.chunk:hover {
		background-color: var(--color-lightblue);
		cursor: pointer;
	}

	.chunk:disabled {
		cursor: not-allowed;
		background-color: var(--color-button);
	}
</style>