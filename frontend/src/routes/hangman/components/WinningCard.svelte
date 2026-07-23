<script lang="ts">
	import Confetti from '$lib/components/Confetti.svelte';
	import RemainingTime from '$lib/components/RemainingTime.svelte';
	import { ConvertPositionToString } from '$lib/utils/posToString';
	import { typewriter } from '$lib/utils/transitions';
	import { onMount } from 'svelte';
	import { scale } from 'svelte/transition';
	import { getHangmanWinningData } from '$lib/api/hangman';
	import { subscribeToPlayerCount } from '$lib/api/common';
	import { page } from '$app/state';

	let { attempts } = $props();

	let name: string = $state('');
	let path: string = $state('');
	let playerCount: number = $state(0);
	let position: number = $state(0);

	onMount(async () => {
		const winningData = await getHangmanWinningData(page.data.userId);

		position = winningData.position;
		name = winningData.enemy_name;
		path = winningData.enemy_path;
	});

	// Streams live player count to the user
	onMount(() => {
		return subscribeToPlayerCount('hangman', (count) => {
			playerCount = count;
		});
	});
</script>

<div class="container" class:fail={attempts == 0} class:success={attempts > 0} in:scale>
	<Confetti finished={attempts > 0} />

	{#if attempts == 6}
		<h2 class="title">Outstanding!</h2>
		<span class="subtitle">You guessed the enemy perfectly!</span>
		<br />
		<p>The enemy of the day was:</p>
	{:else if attempts > 0}
		<h2 class="title">Bingo!</h2>
		<p>The enemy of the day was:</p>
	{:else}
		<h2 class="title">Game Over!</h2>
		<p>The correct enemy was:</p>
	{/if}

	<h1>{name}</h1>

	{#if path}
		<div class="enemy-image">
			<img src={'/enemies' + path} alt="" />
		</div>
	{/if}

	{#if position !== 0}
		<p transition:typewriter={{ speed: 1 }}>
			You were the {ConvertPositionToString(position)} person to guess today's enemy!
		</p>
	{:else}
		<br />
	{/if}
	<p>{playerCount} people guessed todays enemy</p>
	<RemainingTime />
</div>

<style>
	.container {
		width: fit-content;
		min-width: 300px;
		margin: auto;
		margin-bottom: 20px;
		padding: 0 20px;
	}

	.fail {
		border: 20px solid transparent;
		border-image: url('/daily-slash/borders/HellstoneBrick.png');
		border-image-slice: 17;
		border-image-repeat: round;
		border-radius: 5px;

		background-image: url('/daily-slash/backgrounds/HellstoneWall.png');
		background-repeat: repeat;
		background-size: 20%;
	}

	.success {
		background-color: var(--color-button);
		border-radius: 10px;
		border: 2px solid black;
	}

	.title {
		font-size: 40px;
		margin: 20px auto 0 auto;
	}

	.enemy-image {
		border: 20px solid transparent;
		border-image: url('/hangman/EnemyBorder.png');
		border-image-slice: 16;
		border-image-repeat: round;

		background-image: url('/hangman/EnemyBackground.png');
		background-repeat: repeat;
		background-size: 70%;
		background-clip: padding-box;

		padding: 10px;
		margin: auto;
		width: fit-content;
	}

	.enemy-image img {
		max-width: 100px;
		max-height: 100px;
		width: 100px;
		height: auto;
		object-fit: contain;
	}

	@media (max-width: 700px) {
		.container {
			min-width: 0;
			max-width: 300px;
			width: 75%;
		}
	}
</style>
