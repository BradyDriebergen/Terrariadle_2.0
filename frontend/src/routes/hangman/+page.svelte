<script lang="ts">
	import { slide } from 'svelte/transition';
	import Guide from './components/Guide.svelte';
	import Keyboard from './components/Keyboard.svelte';
	import WinningCard from './components/WinningCard.svelte';
	import { cubicInOut } from 'svelte/easing';
	import type { PageData } from './$types';
	import { get } from 'svelte/store';
	import { userIdStore } from '$lib/store/session';
	import { checkEnemyGuess } from '$lib/api/hangman';
	import type { HangmanGuess } from '$lib/types/hangman';

	let { data }: { data: PageData } = $props();

	let attempts: number = $state(0);
	let finished: boolean = $state(false);
	let guesses: HangmanGuess[] = $state([]);
	let phrase: string[] = $state([]);

	$effect(() => {
		// Initialize data once pre-fetch is finished
		if (data.gameContext) {
			attempts = data.gameContext.attempts;
			finished = data.gameContext.finished;
			guesses = data.gameContext.guesses;
			phrase = data.gameContext.phrase
		}
	});

	// Audio to play after guess
	let audio = $state<HTMLAudioElement>();

	// Split into words for better wrapping
	let phraseWords = $derived.by(() => {
		const enemyString = phrase.join('');
		return enemyString.split(' ').map((word) => word.split(''));
	});

	async function onKeyPressed(letter: string) {
		let res;
		try {
			const userId = get(userIdStore);
			res = await checkEnemyGuess(userId, letter);
		} catch (e) {
			// handle error here
			console.error(e);
			return
		}

		phrase = res.phrase;
		guesses = [res.guess, ...guesses];
		attempts = res.attempts;
		finished = res.finished;

		if (attempts <= 0) {
			audio?.play();
		}
	}

	// Effect for updating the background with fade
	let failed: boolean = $derived(attempts <= 0);
	$effect(() => {
		if (failed) {
			document.body.style.setProperty('--bg-image', "url('/page-backgrounds/Underworld.png')");
			document.body.style.setProperty('--bg-opacity', '1');
		}

		return () => {
			document.body.style.setProperty('--bg-opacity', '0');
		};
	});
</script>

{#if data.gameContext}
	{#if !finished}
		<div class="title-box" out:slide={{ duration: 700, easing: cubicInOut }}>
			<h2>Hangman</h2>
			<p>Guess letters one by one to figure out the enemy before hanging the Guide!</p>
		</div>
	{:else}
		<div style="margin-top: -20px; margin-bottom: {attempts === 0 ? '15px' : '-20px'}">
			<span class="color-cycle">Hangman Results</span>
		</div>
	{/if}

	<Guide {attempts} />

	{#if finished}
		<WinningCard {attempts} />
	{/if}

	<div class="phrase-container">
		{#each phraseWords as word}
			<div class="word">
				{#each word as letter}
					<span>{letter}</span>
				{/each}
			</div>
		{/each}
	</div>
	{#if !finished}
		<Keyboard {onKeyPressed} {guesses} />
	{/if}
{:else}
	<p>Loading...</p>
{/if}

<!-- Audio that only plays after final guess is made -->
<audio bind:this={audio} src={'/hangman/GuideDeath.mp3'}></audio>

<style>
	/* global CSS used for fading the background */
	:global(body) {
		margin: 0;
		min-height: 100vh;
	}

	:global(body::before) {
		content: '';
		position: fixed;
		inset: 0;
		background-image: var(--bg-image, none);
		background-size: cover;
		background-position: center;
		opacity: var(--bg-opacity, 0);
		transition: opacity 0.5s ease;
		z-index: -1;
		pointer-events: none;
	}

	.title-box {
		background-color: var(--color-backgroundblue);
		width: fit-content;
		text-align: center;
		margin: auto;
		margin-top: 30px;
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

	.phrase-container {
		background-color: var(--color-backgroundblue);
		display: flex;
		flex-wrap: wrap;
		justify-content: center;
		gap: 20px;
		padding: 20px 25px;
		margin: 20px auto;
		width: fit-content;

		border-radius: 15px;
		border: 2px solid black;
		font-size: 25px;
	}

	.word {
		display: flex;
		gap: 5px;
		white-space: nowrap;
	}
</style>
