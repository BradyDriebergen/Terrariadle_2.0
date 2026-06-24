<script lang="ts">
	let {
		guessCount = 0,
		finished = false
	}: {
		guessCount: number;
		finished: boolean;
	} = $props();

	const progressBar2ByGuess = [100, 75, 50, 25, 100, 60, 30, 100, 80, 60, 40, 20, 0];

	// Use 12 guesses when won is true
	let effectiveGuesses: number = $derived(finished ? 12 : guessCount);

	let progressBar1: number = $derived.by(() => {
		if (effectiveGuesses < 4) {
			return 100;
		} else if (effectiveGuesses < 7) {
			return 66;
		} else if (effectiveGuesses < 12) {
			return 33;
		}
		return 0;
	});

	let progressBar2: number = $derived.by(() => {
		if (effectiveGuesses >= 12) return progressBar2ByGuess[12];
		return progressBar2ByGuess[effectiveGuesses];
	});
</script>

<div class="wrapper">
	<div class="text">Hint Progress Bar</div>
	<img
		class="background"
		src="/loading-bar-assets/LoadingBackground.png"
		alt="Loading background"
	/>
	<img
		class="bar"
		src={'/loading-bar-assets/LoadingBar1.png'}
		alt="Loading bar 1"
		style="clip-path: inset(0 {progressBar1}% 0 0);"
	/>
	<img
		class="bar"
		src={'/loading-bar-assets/LoadingBar2.png'}
		alt="Loading bar 2"
		style="clip-path: inset(0 {progressBar2}% 0 0);"
	/>
	<img class="border" src="/loading-bar-assets/LoadingBorder.png" alt="Loading border" />
</div>

<style>
	.wrapper {
		position: relative;
		width: 300px;
		background-color: var(--color-button);
		padding: 45px;

		border-radius: 5px;
		border: thin solid black;
		transition: background-color 0.2s ease;
	}

	.text {
		margin: auto;
		margin-top: -60px;
		padding: 5px 8px;
		background-color: var(--color-lightblue);
		width: fit-content;
		text-align: center;

		border-radius: 5px;
		border: thin solid black;
	}

	img {
		position: absolute;
		top: 22px;
		left: 5px;
		width: 97%;
	}

	.border {
		z-index: 3;
	}

	.bar {
		z-index: 2;
		transition: clip-path 400ms ease-out;
	}

	.background {
		z-index: 1;
	}
</style>
