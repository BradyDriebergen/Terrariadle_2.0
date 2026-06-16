<script lang="ts">
	import { send, receive } from '$lib/utils/transitions';
	import { flip } from 'svelte/animate';
	import { scale, slide } from 'svelte/transition';
	import WinningPanel from './components/WinningPanel.svelte';
	import Confetti from '$lib/components/Confetti.svelte';
	import { Tween } from 'svelte/motion';
	import { cubicInOut, cubicOut } from 'svelte/easing';
	import type { CategoryOption, SolvedCategory } from '$lib/types/connections';

	let { data } = $props();

	let attempts = $state<number>(0);
	let finished = $state<boolean>(false);
	let options = $state<CategoryOption[]>([]);
	let solvedCategories = $state<SolvedCategory[]>([]);

	// let options = $state<Option[]>(
	// 	data.options.map((label: string) => ({
	// 		id: crypto.randomUUID(),
	// 		label
	// 	}))
	// );
	// let attempts: number = $state(data.attempts);
	// let finished: boolean = $derived(data.finished);

	let transitioning: boolean = $state(false);

	let tempGuesses: Option[] = $state([]);
	let selectedOptions: Option[] = $state([]);

	let answerCategories: string[] = $state(updateAnswerCategories(data.guesses));
	let answerOptions: Record<string, string[]> = $state(updateAnswerOptions(data.guesses));
	let tempAnswerCategory: string = $state('');
	let tempAnswerOptions: string[] = $state([]);

	let showOneAway: boolean = $state(false);
	let timeout: ReturnType<typeof setTimeout> | undefined = $state(undefined);

	// Update methods for complex assignments
	function updateAnswerCategories(list: any) {
		return list.map((item: { category: any }) => item.category);
	}

	function updateAnswerOptions(list: any) {
		return list.reduce(
			(acc: Record<string, string[]>, item: { category: string; options: string[] }) => {
				acc[item.category] = item.options;
				return acc;
			},
			{}
		);
	}

	// Adds selected category to string
	function selectCategory(opt: Option) {
		if (selectedOptions.includes(opt)) {
			selectedOptions = selectedOptions.filter((s) => s !== opt);
		} else {
			selectedOptions.push(opt);
		}
	}

	// Shuffles the remaining panels
	function shuffle() {
		const result = [...options];
		for (let i = result.length - 1; i > 0; i--) {
			const j = Math.floor(Math.random() * (i + 1));
			[result[i], result[j]] = [result[j], result[i]];
		}
		options = result;
	}

	// Makes the banner appear when correct guess is made
	const MAX = 4;
	let index = $state(0);
	async function updateAnswerPanes() {
		index--;
		if (index === 0) {
			tempGuesses = [];

			if (finished && attempts === 0) return;

			answerCategories.push(tempAnswerCategory);
			answerOptions[tempAnswerCategory] = tempAnswerOptions;
			tempAnswerCategory = '';
			tempAnswerOptions = [];
		}
	}

	function toggleOneAway() {
		showOneAway = true;
		clearTimeout(timeout);

		timeout = setTimeout(() => {
			showOneAway = false;
		}, 2000);
	}

	let x = new Tween(0, {
		duration: 80,
		easing: cubicOut
	});

	async function triggerShake() {
		await x.set(10);
		await x.set(-10);
		await x.set(7);
		await x.set(-7);
		await x.set(0);
	}

	async function submitGuess() {
		const guesses = options
			.filter((option) => selectedOptions.includes(option))
			.map((option) => option.label);

		fetch('http://localhost:3000/api/connections/check-guess', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ userId: data.userId, guess: guesses })
		})
			.then((r) => r.json())
			.then(async (res) => {
				if (res.guess.id !== 0) {
					tempAnswerCategory = res.guess.category;
					tempAnswerOptions = res.guess.options;

					options = options.filter((s: Option) => !selectedOptions.includes(s));
					tempGuesses.push(...selectedOptions);
				} else {
					if (res.oneAway) {
						toggleOneAway();
					}

					await triggerShake();
					attempts--;
				}

				finished = res.finished;
				if (finished && attempts === 0) {
					const rawData = await fetch(
						`http://localhost:3000/api/connections/initialize-game/${data.userId}`
					);
					const dataJson = await rawData.json();

					options = dataJson.options;
					answerCategories = updateAnswerCategories(dataJson.guesses);
					answerOptions = updateAnswerOptions(dataJson.guesses);
				}

				selectedOptions = [];
			});
	}
</script>

{#if !finished}
	<div class="title-box" out:slide={{ duration: 700, easing: cubicInOut }}>
		<h2>Connections</h2>
		<p>Find groups of 4 with something in common!</p>
	</div>
{/if}

{#if showOneAway && !finished}
	<span class="one-away-msg" transition:scale>One Away!</span>
{/if}

{#if finished && !transitioning}
	<WinningPanel won={attempts > 0} {attempts} userId={data.userId} />
	<Confetti won={attempts > 0} />
{/if}

<div class="grid">
	{#each answerCategories as banner, index}
		<div class="answer-pane pane-{index}" id={String(index)} style="grid-column: span 4;" in:scale>
			<h4>{banner}</h4>
			<span
				>{answerOptions[banner][0]}, {answerOptions[banner][1]}, {answerOptions[banner][2]}, {answerOptions[
					banner
				][3]}</span
			>
		</div>
	{/each}

	{#each tempGuesses as category (category.id)}
		<button class="category" in:receive={{ key: category.label }}>
			<span>{category.label}</span>
		</button>
	{/each}

	{#each options as category (category.id)}
		<button
			type="button"
			class="category"
			class:Selected={selectedOptions.includes(category)}
			style:transform={selectedOptions.includes(category)
				? `translateX(${x.current}px)`
				: undefined}
			onclick={() => selectCategory(category)}
			disabled={selectedOptions.length === 4 && !selectedOptions.includes(category)}
			out:send={{ key: category.label }}
			animate:flip={{ duration: 220, easing: (t) => t }}
			onoutroend={() => {
				updateAnswerPanes();
				transitioning = false;
			}}
			onoutrostart={() => {
				index = MAX;
				transitioning = true;
			}}
		>
			<span>{category.label}</span>
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
			<button onclick={() => (selectedOptions = [])}>Deselect All</button>
			<button onclick={submitGuess} disabled={selectedOptions.length !== 4}>Check Connection</button
			>
		</div>
	{/if}
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

	.category {
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

	.category:hover,
	.category.Selected:hover {
		background-color: var(--color-lightblue);
		cursor: pointer;
	}

	.category:disabled:hover {
		background-color: var(--color-button);
		cursor: default;
	}

	.category.Selected,
	.category.Selected:hover {
		background-color: rgba(139, 31, 31, 0.862);
	}

	.category span {
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
