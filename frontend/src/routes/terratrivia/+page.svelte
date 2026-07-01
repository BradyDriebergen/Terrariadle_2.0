<script lang="ts">
	import { send } from "$lib/utils/transitions";
	import { flip } from "svelte/animate";
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

	const chunk = [
		"DIG", "GIN", "GC", "LAW", "ANGL", 
		"ERFI", "SH", "RAPI", "DHEA", "LING", 
		"MO", "LT", "EN", "PIC", "KA", "XE", 
		"VOR", "TEX", "RA", "IN"
	];

	let finished: boolean = $state(false);
    let chunks= $state(chunk.map((value, i) => ({id: i, value})));
	let selectedChunks: string[] = $state([]);
	let input: string | null = $derived(selectedChunks.length > 0 ? selectedChunks.join('') : null);

    function shuffle() {
		const result = [...chunks];
		for (let i = result.length - 1; i > 0; i--) {
			const j = Math.floor(Math.random() * (i + 1));
			[result[i], result[j]] = [result[j], result[i]];
		}
		chunks = result;
	}
</script>

{#if finished}
    <span class="color-cycle">TerraTrivia Results</span>
{/if}

<div class="game-window">
    {#if !finished}
        <selection out:slide={{ duration: 700, easing: cubicInOut }}>
            <h2>TerraTrivia</h2>
            <p>Use clues to decipher all 7 words!</p>
        </selection>
    {/if}

	<div class="clue-box">
		{#each clues as clue}
			<div class="clue-row">
				<span class="clue">{clue}</span>
				<span class="letter-count">6 letters</span>
			</div>
		{/each}
	</div>

	<div class="selection-menu">
		<button onclick={shuffle}>Shuffle</button>
		<button>Hint</button>
		<button onclick={() => selectedChunks = []}>Clear</button>
	</div>

	<div class="input-box">
		<span class:empty={selectedChunks.length > 0}>
            {input ?? "Start building words..."}
        </span>
		<button 
            onclick={() => selectedChunks.pop()}
            disabled={selectedChunks.length === 0}
        >
            ←
        </button>
	</div>
</div>

<div class="chunk-grid">
	{#each chunks as chunk (chunk.id)}
			<button 
				class="chunk" 
                class:chunk-placeholder={
                    selectedChunks.includes(chunk.value) || 
                    chunk.value === ""
                }
				disabled={
                    selectedChunks.includes(chunk.value) || 
                    selectedChunks.length >= 4
                }
                onclick={() => selectedChunks.push(chunk.value)} 
                animate:flip={{ duration: 220, easing: (t) => t }}
			>
				{!selectedChunks.includes(chunk.value) ? chunk.value : ""}
			</button>
	{/each}
</div>

<style>
	.game-window {
		background-color: var(--color-backgroundblue);
		width: fit-content;
		margin: 15px auto 15px auto;
		padding: 10px;
		border-radius: 15px;
		border: 2px solid black;
	}

    .game-window h2 {
		background-color: var(--color-lightblue);
		width: fit-content;
		margin: auto;
		margin-top: -38px;
		margin-bottom: 10px;
		padding: 10px 20px;

		border-radius: 15px;
		border: 2px solid black;
	}

	.clue-box {
		background-color: var(--color-button);
		width: 350px;
		padding: 4px 0;

		border-radius: 15px;
		border: 2px solid black;
	}

	.clue-row {
		display: flex; 
		justify-content: space-between;
        align-items: center;
		padding: 4px 8px;
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
        transition: background-color 0.1s ease;
	}

    .selection-menu button:hover {
        background-color: var(--color-lightblue);
        cursor: pointer;
    }

	.input-box {
		background-color: var(--color-button);
		border-radius: 8px;
		border: 2px solid black;
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 8px 12px;
		font-size: 20px;
		margin: 0px auto;
		width: 250px;
		min-height: 20px;
		cursor: default;
	}

    .input-box span {
        color: gray;
    }

    .input-box .empty {
        color: white;
        letter-spacing: 2px;
    }

	.input-box button {
		background-color: var(--color-lightblue);
		border-radius: 8px;
		border: 2px solid black;
		padding: 4px 8px 2px;
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

    .chunk-placeholder {
        border: none;
        background: none;
    }

	.chunk:hover {
		background-color: var(--color-lightblue);
		cursor: pointer;
	}

    .chunk-placeholder:hover {
        border: none;
        background: none;
        cursor: default;
    }

	.chunk:disabled {
		cursor: not-allowed;
		background-color: var(--color-button);
	}
    
    .chunk-placeholder:disabled {
        border: none;
        background: none;
        cursor: default;
    }
</style>