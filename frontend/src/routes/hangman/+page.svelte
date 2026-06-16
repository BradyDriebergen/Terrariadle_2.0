<script lang="ts">
	import { slide } from 'svelte/transition';
	import Guide from './components/Guide.svelte';
	import Keyboard from './components/Keyboard.svelte';
	import WinningComponent from './components/WinningComponent.svelte';
	import { cubicInOut } from 'svelte/easing';

	let { data } = $props();

	let attempts: number = $state(data.attempts);
	let finished: boolean = $state(data.finished);
	let guessedLetters: string[] = $state(data.guessedLetters);
	let enemy: string[] = $state(data.phrase);

	// Audio to play after guess
	let audio = $state<HTMLAudioElement>();

	// Split into words for better wrapping
	let enemyWords = $derived.by(() => {
		const enemyString = enemy.join('');
		return enemyString.split(' ').map((word) => word.split(''));
	});

	async function onKeyPressed(letter: string) {
		fetch('http://localhost:3000/api/hangman/check-guess', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ userId: data.userId, guess: letter })
		})
			.then((r) => r.json())
			.then((data) => {
				enemy = data.newPhrase;
				if (!guessedLetters.includes(letter)) {
					guessedLetters.push(letter);
				}

				finished = data.finished;

				if (!data.correct) {
					attempts--;
					if (attempts <= 0) {
						audio?.play();
					}
				}
			});
	}

	// Effect for updating the background with fade
	let failed: boolean = $derived(attempts <= 0);
	$effect(() => {
		if (failed) {
			document.body.style.setProperty('--bg-image', "url('/backgrounds/Underworld.png')");
			document.body.style.setProperty('--bg-opacity', '1');
		}

		return () => {
			document.body.style.setProperty('--bg-opacity', '0');
		};
	});
</script>

{#if !finished}
	<!-- out:slide={{ duration: 700, easing: cubicInOut }} -->
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
	<WinningComponent {attempts} userId={data.userId} />
{/if}

<div class="phrase-container">
	{#each enemyWords as word}
		<div class="word">
			{#each word as letter}
				<span>{letter}</span>
			{/each}
		</div>
	{/each}
</div>
{#if !finished}
	<Keyboard {onKeyPressed} enemyLetters={enemy} {guessedLetters} />
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
