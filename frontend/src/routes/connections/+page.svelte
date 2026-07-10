<script lang="ts">
	import { send, receive } from '$lib/utils/transitions';
	import { flip } from 'svelte/animate';
	import { scale, slide } from 'svelte/transition';
	import WinningCard from './components/WinningCard.svelte';
	import Confetti from '$lib/components/Confetti.svelte';
	import { Tween } from 'svelte/motion';
	import { cubicInOut, cubicOut } from 'svelte/easing';
	import type { CategoryOption, SolvedCategory } from '$lib/types/connections';
	import { checkCategoryGuess, revealConnectionsAnswers } from '$lib/api/connections';
	import type { PageData } from './$types';
	import { page } from '$app/state';

	let { data }: { data: PageData } = $props();

	let attempts: number = $state(0);
	let finished: boolean = $state(false);
	let options: CategoryOption[] = $state([]);
	let solvedCategories: SolvedCategory[] = $state([]);

	$effect(() => {
		// Initialize data once pre-fetch is finished
		if (data.gameContext) {
			attempts = data.gameContext.attempts;
			finished = data.gameContext.finished;
			options = data.gameContext.options.map((value, i) => ({
				id: i,
				value,
				selected: false
			})) as CategoryOption[];
			solvedCategories = data.gameContext.solved_categories as SolvedCategory[];
		}
	});

	let selectedOptionCount: number = $derived(options.filter((o) => o.selected).length);
	let animatingOptions: CategoryOption[] = $state([]);
	let inFlightCategories: SolvedCategory[] = $state([]);

	let showOneAway: boolean = $state(false);

	// Used to delay winning panel from showing on last successful guess
	let transitioning: boolean = $state(false);

	// Used to prevent users from guessing during the transition
	let loadingGuess: boolean = $state(false);

	// Used for shake functionality
	let shakeTween = new Tween(0, {
		duration: 90,
		easing: cubicOut
	});

	async function triggerShake() {
		await shakeTween.set(8);
		await shakeTween.set(-8);
		await shakeTween.set(5);
		await shakeTween.set(-5);
		await shakeTween.set(0);
	}

	let timeout: ReturnType<typeof setTimeout> | undefined = undefined;
	function toggleOneAway() {
		showOneAway = true;
		clearTimeout(timeout);

		timeout = setTimeout(() => {
			showOneAway = false;
		}, 2000);
	}

	async function submitGuess() {
		loadingGuess = true;

		const guess = options.filter((option) => option.selected).map((option) => option.value);
		const userId = page.data.userId;

		let guessResult;
		try {
			guessResult = await checkCategoryGuess(guess, userId);
		} catch (e) {
			loadingGuess = false;
			// TODO: handle error here
			console.error(e);
			return;
		}

		attempts = guessResult.attempts;

		if (guessResult.finished) {
			transitioning = true;
			finished = guessResult.finished;
		}

		if (guessResult.is_correct) {
			animatingOptions = [...options.filter((o) => o.selected)];
			inFlightCategories = [guessResult.correct_guess];

			options = options.filter((o) => !o.selected);

			loadingGuess = false;
			return;
		}

		if (guessResult.one_away) toggleOneAway();

		await triggerShake();
		deselectOptions();
		loadingGuess = false;

		if (attempts === 0) {
			let answers;
			try {
				answers = await revealConnectionsAnswers(userId);
			} catch (e) {
				// TODO: handle error here
				console.error(e);
				return;
			}

			options = [];
			solvedCategories = [];
			inFlightCategories = [...answers.revealed_categories];
		}
	}

	// Makes the banner appear when correct guess is made
	const MAX = 4;
	let index = $state(0);
	function updateAnswerPanes() {
		index--;
		if (index === 0) {
			animatingOptions = [];

			inFlightCategories.forEach((category: SolvedCategory) => {
				solvedCategories.push(category);
			});

			inFlightCategories = [];
		}
	}

	function deselectOptions() {
		options.forEach((o) => (o.selected = false));
	}

	function shuffle() {
		const result = [...options];
		for (let i = result.length - 1; i > 0; i--) {
			const j = Math.floor(Math.random() * (i + 1));
			[result[i], result[j]] = [result[j], result[i]];
		}
		options = result;
	}
</script>

{#if !finished || transitioning}
	<div class="title-box" out:slide={{ duration: 700, easing: cubicInOut }}>
		<h2>Connections</h2>
		<p>Find groups of 4 with something in common!</p>
	</div>
{/if}

{#if showOneAway && !finished}
	<span class="one-away-msg" transition:scale>One Away!</span>
{/if}

{#if finished && !transitioning}
	<WinningCard {attempts} />
	<Confetti finished={attempts > 0} />
{/if}

{#if data.gameContext}
	<div class="grid">
		{#each solvedCategories as category, index}
			<div
				class="answer-pane pane-{index}"
				id={String(index)}
				style="grid-column: span 4;"
				in:scale
			>
				<h4>{category.name}</h4>
				<span>
					{category.options[0]},
					{category.options[1]},
					{category.options[2]},
					{category.options[3]}
				</span>
			</div>
		{/each}

		{#each animatingOptions as option (option.id)}
			<button class="option" in:receive={{ key: option.value }}>
				<span>{option.value}</span>
			</button>
		{/each}

		{#each options as option (option.id)}
			<button
				type="button"
				class="option"
				class:Selected={option.selected}
				style:transform={option.selected ? `translateX(${shakeTween.current}px)` : undefined}
				onclick={() => (option.selected = !option.selected)}
				disabled={(selectedOptionCount >= 4 && !option.selected) || transitioning}
				out:send={{ key: option.value }}
				animate:flip={{ duration: 220, easing: (t) => t }}
				onoutrostart={() => {
					index = MAX;
					transitioning = true;
				}}
				onoutroend={() => {
					updateAnswerPanes();
					transitioning = false;
				}}
			>
				<span>{option.value}</span>
			</button>
		{/each}
	</div>

	<div>
		<div class="attempts-bar">
			<span>Attempts Left:</span>
			{#each Array(attempts) as _, i}
				<img src="/emojis/LifeHeart.png" alt="Number of changes left" out:scale />
			{/each}
			{#if attempts === 0}
				<span>None</span>
			{/if}
		</div>

		{#if !finished}
			<div class="game-buttons">
				<button onclick={shuffle}>Shuffle</button>
				<button onclick={deselectOptions}>Deselect All</button>
				<button onclick={submitGuess} disabled={selectedOptionCount !== 4 || loadingGuess}>
					Check Connection
				</button>
			</div>
		{/if}
	</div>
{:else}
	<p>loading...</p>
{/if}

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

	.grid {
		display: grid;
		grid-template-columns: repeat(4, 1fr);
		grid-auto-rows: 1fr;

		width: 500px;
		height: 400px;
		margin: 0 auto;
		gap: 10px;
	}

	.one-away-msg {
		position: absolute;
		top: 150px;
		left: calc(50% - 35px);
		z-index: 1000;

		background-color: var(--color-green);
		padding: 10px;
		border-radius: 15px;
		border: 2px solid black;
	}

	.answer-pane {
		background-color: var(--color-button);
		border-radius: 15px;
		border: 2px solid black;
		padding: 3px 10px;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}

	.answer-pane.pane-0 {
		background: url('/category-backgrounds/GrassWall.png');
		background-repeat: repeat;
		background-size: 20%;
	}
	.answer-pane.pane-1 {
		background: url('/category-backgrounds/HardenedSandWall.png');
		background-repeat: repeat;
		background-size: 20%;
	}
	.answer-pane.pane-2 {
		background: url('/category-backgrounds/SnowWall.png');
		background-repeat: repeat;
		background-size: 20%;
	}
	.answer-pane.pane-3 {
		background: url('/category-backgrounds/JungleVineWall.png');
		background-repeat: repeat;
		background-size: 20%;
	}

	.answer-pane h4 {
		font-size: 25px;
		font-weight: 600;
		margin: 0 5px 5px 5px;
	}

	.answer-pane span {
		font-size: 14px;
	}

	.option {
		background-color: var(--color-button);
		width: 100%;
		height: 100%;
		display: grid;
		place-items: center;

		border-radius: 15px;
		border: 2px solid black;

		padding: 0;
		font-family: inherit;
		font-size: inherit;
		transition: background-color 0.1s ease;
	}

	.option:hover,
	.option.Selected:hover {
		background-color: var(--color-lightblue);
		cursor: pointer;
	}

	.option:disabled:hover {
		background-color: var(--color-button);
		cursor: default;
	}

	.option.Selected,
	.option.Selected:hover {
		background-color: rgba(139, 31, 31, 0.862);
	}

	.option span {
		padding: 5px;
		font-size: 20px;
	}

	.attempts-bar {
		display: flex;
		justify-content: center;
		align-items: center;
		margin: 20px;
	}

	.attempts-bar span {
		margin-right: 5px;
	}

	.attempts-bar img {
		margin: 2px;
	}

	.game-buttons button {
		background-color: var(--color-button);
		border-radius: 10px;
		border: 2px solid black;
		padding: 8px 10px;
		font-size: 18px;
		margin: 0 2px;
	}
	.game-buttons button:not([disabled]):hover {
		background-color: var(--color-lightblue);
		cursor: pointer;
	}
	.game-buttons button:disabled {
		background-color: rgba(0, 0, 0, 0.518);
	}
</style>
