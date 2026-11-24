<script lang="ts">
	import { checkTriviaQuestionGuess } from '$lib/api/terratrivia';
	import type { TriviaItem } from '$lib/types/terratrivia.js';
	import { onDestroy } from 'svelte';
	import { cubicInOut } from 'svelte/easing';
	import { fly, slide } from 'svelte/transition';
	import WinningCard from './components/WinningCard.svelte';
	import { SvelteMap } from 'svelte/reactivity';
	import Confetti from '$lib/components/Confetti.svelte';
	import type { PageData } from './$types';
	import { page } from '$app/state';
	import { session } from '$lib/store/session.svelte';
	import { send } from '$lib/utils/transitions';
	import { flip } from 'svelte/animate';

	let { data }: { data: PageData } = $props();

	let finished: boolean = $state(false);
	let chunks: SvelteMap<number, string> = $state(new SvelteMap());
	let triviaItems: SvelteMap<number, TriviaItem> = $state(new SvelteMap());

	let pageLoading: boolean = $state(true);
	$effect(() => {
		// Initialize data once pre-fetch is finished
		if (data.gameContext) {
			finished = data.gameContext.finished;
			chunks = new SvelteMap(data.gameContext.chunks.entries());
			triviaItems = new SvelteMap<number, TriviaItem>(
				data.gameContext.trivia_items.map((item) => [item.id, item])
			);

			pageLoading = false;
		}
	});

	let selectedChunkIds: number[] = $state([]);
	let input: string = $derived(
		selectedChunkIds
			.map((id) => chunks.get(id))
			.filter((s): s is string => s !== undefined)
			.join('')
	);

	// Debouncer function for api calls
	let timeoutId: ReturnType<typeof setTimeout> | undefined;
	function handleClick() {
		if (selectedChunkIds.length < 2 || !input) return;

		clearTimeout(timeoutId);
		timeoutId = setTimeout(() => {
			submitGuess();
		}, 500);
	}
	onDestroy(() => clearTimeout(timeoutId));

	async function submitGuess() {
		const ids = [...selectedChunkIds];
		const guess = input;
		if (ids.length < 2 || !guess) return;

		try {
			const res = await checkTriviaQuestionGuess(page.data.userId, guess);
			if (!res.is_correct) return;

			triviaItems.set(res.guess_result.id, res.guess_result);
			finished = res.finished;
			if (finished) session.terratriviaStatus = true;

			for (const id of ids) chunks.delete(id);
			selectedChunkIds = selectedChunkIds.filter((id) => !ids.includes(id));
		} catch (e) {
			alert('Games have refreshed! Refresh the page to start guessing.');

			// TODO: handle error here
			console.error(e);
		}
	}

	function shuffle() {
		let shuffled = [...chunks];
		for (let i = shuffled.length - 1; i > 0; i--) {
			const j = Math.floor(Math.random() * (i + 1));
			[shuffled[i], shuffled[j]] = [shuffled[j], shuffled[i]];
		}

		chunks = new SvelteMap(shuffled);
	}
</script>

{#if !pageLoading}
	<div in:fly={{ y: 20, duration: 400 }}>
		{#if finished}
			<WinningCard />
			<Confetti finished />
		{/if}

		<div class="game-window">
			{#if !finished}
				<div out:slide={{ duration: 700, easing: cubicInOut }}>
					<h2>TerraTrivia</h2>
					<p>Use clues to decipher all 7 words!</p>
				</div>
			{/if}

			<div class="clue-box">
				{#each triviaItems as [id, item] (id)}
					<div class="clue-row">
						<span>{item.clue}</span>
						<span>{item.answer !== '' ? item.answer : item.letter_count + ' letters'}</span>
					</div>
				{/each}
			</div>

			{#if !finished}
				<div class="selection-menu">
					<button onclick={shuffle}>Shuffle</button>
					<button onclick={() => (selectedChunkIds = [])}>Clear</button>
				</div>

				<div class="input-box">
					<span class:empty={selectedChunkIds.length > 0}>
						{input || 'Start building words...'}
					</span>
					<button
						onclick={() => {
							selectedChunkIds.pop();
							handleClick();
						}}
						disabled={selectedChunkIds.length === 0}
					>
						<img src="/emojis/backspace.png" alt="backspace" />
					</button>
				</div>
			{/if}
		</div>

		{#if !finished}
			<div class="chunk-grid" out:slide={{ duration: 300, easing: cubicInOut }}>
				{#each [...chunks] as [id, chunk] (id)}
					<button
						class="chunk"
						class:selected-chunk={selectedChunkIds.includes(id)}
						disabled={selectedChunkIds.includes(id) || selectedChunkIds.length >= 4}
						out:send={{ key: id }}
						animate:flip={{ duration: 180, easing: (t) => t }}
						onclick={() => {
							selectedChunkIds.push(id);
							handleClick();
						}}
					>
						{!selectedChunkIds.includes(id) ? chunk : ''}
					</button>
				{/each}
			</div>
		{/if}
	</div>
{:else}
	<p>Loading...</p>
{/if}

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
		width: 380px;
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
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: var(--color-lightblue);
		border-radius: 8px;
		border: 2px solid black;
		padding: 6px 8px;
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

	.selected-chunk {
		border: none;
		background: none;
	}

	.chunk:hover {
		background-color: var(--color-lightblue);
		cursor: pointer;
	}

	.selected-chunk:hover {
		border: none;
		background: none;
		cursor: default;
	}

	.chunk:disabled {
		cursor: not-allowed;
		background-color: var(--color-button);
	}

	.selected-chunk:disabled {
		border: none;
		background: none;
		cursor: default;
	}

	@media (max-width: 650px) {
		.game-window {
			width: 90%;
		}

		.clue-box {
			width: 100%;
		}

		.clue-row {
			font-size: 12px;
		}
	}
</style>
