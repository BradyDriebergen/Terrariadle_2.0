<script lang="ts">
	import Guide from './components/Guide.svelte';
	import Keyboard from './components/Keyboard.svelte';

    let { data } = $props();

    let failedGuesses: number = $state(0);
    let finished: boolean = $state(false);

	// Audio to play after guess
	let audio = $state<HTMLAudioElement>();

	const correctEnemy = "ETHERIAN JAVELIN THROWER"
	let guessedLetters: string[] = $state([])
	let enemy = $derived.by(() => {
		return correctEnemy.split('').map(letter => {
			if (letter === ' ') return ' ';
			return guessedLetters.includes(letter) ? letter : '_';
		});
	});

    function onKeyPressed(letter: string) {
        guessedLetters.push(letter);

		if (!correctEnemy.split('').includes(letter)) {
			failedGuesses++;
		}

		if (failedGuesses >= 6) {
			audio?.play();
		}
    }

	// Effect for updating the background with fade
	let failed: boolean = $derived(failedGuesses >= 6);
	$effect(() => {
		if (failed) {
			document.body.style.setProperty(
				'--bg-image',
				"url('/backgrounds/Underworld.png')"
			);
			document.body.style.setProperty('--bg-opacity', '1');
		}

		return () => {
			document.body.style.setProperty('--bg-opacity', '0');
		};
	});
</script>

{#if !finished}
    <!-- out:slide={{ duration: 700, easing: cubicInOut }} -->
	<div class="title-box">
		<h2>Hangman</h2>
		<p>Guess letters one by one to figure out the enemy before hanging the Guide!</p>
	</div>
{/if}

<Guide {failedGuesses}/>

<div class="phrase-container">
	{#each enemy as letter}
		{#if letter === " "}
			<div style="width: 10px"></div>
		{:else}
			<span>{letter}</span>
		{/if}
	{/each}
</div>

<Keyboard {onKeyPressed} letters={enemy} />

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
		gap: 8px;
		padding: 20px 25px;
		margin: auto;
		margin-bottom: 20px;
		width: fit-content;
		border-radius: 15px;
		border: 2px solid black;
		font-size: 25px;
	}
</style>