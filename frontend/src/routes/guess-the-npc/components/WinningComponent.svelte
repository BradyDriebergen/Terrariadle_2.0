<script lang="ts">
	import { subscribeToPlayerCount } from '$lib/api/common';
	import { checkNpcName, getNpcWinningData } from '$lib/api/guess-the-npc';
	import RemainingTime from '$lib/components/RemainingTime.svelte';
	import type { NpcGuess } from '$lib/types/guess-the-npc';
	import { ConvertPositionToString } from '$lib/utils/posToString';
	import { typewriter } from '$lib/utils/transitions';
	import { onMount } from 'svelte';
	import { scale } from 'svelte/transition';

	let { npc }: { npc: NpcGuess } = $props();

	let position: number = $state(0);
	let playerCount: number = $state(0);

	let names: string[] = $state([]);
	let guessedName = $state('');
	let correctName = $state('');

	onMount(async () => {
		const winningData = await getNpcWinningData();

		position = winningData.position;
		names = winningData.names;
		guessedName = winningData.guessed_name;
		correctName = winningData.correct_name;
	});

	// Streams live player count to the user
	onMount(() => {
		return subscribeToPlayerCount('guess_the_npc', (count) => {
			playerCount = count;
		});
	});

	async function guessName(name: string) {
		try {
			const res = await checkNpcName(name);

			guessedName = res.guessed_name;
			correctName = res.correct_name;
		} catch (e) {
			// handle error here
			console.log(e);
		}
	}
</script>

<div class="winning-container" in:scale>
	<h1>Spot On!</h1>

	<div class="npc-housing">
		<div
			class="candle-light"
			style="bottom: {npc.name === 'Traveling Merchant' ? '18' : '35'}px"
		></div>
		<img
			class="candle"
			style="bottom: {npc.name === 'Traveling Merchant' ? '18' : '35'}px"
			src="/furniture/Candle.png"
			alt=""
		/>
		<img src={'/furniture/' + npc.name.replace(' ', '') + 'Table.png'} alt="" />
		<img src={'/furniture/' + npc.name.replace(' ', '') + 'Chair.png'} alt="" />
		<img class="npc-image" src={'/npcs/' + npc.path} alt="" />
	</div>

	{#if position !== 0}
		<p transition:typewriter={{ speed: 1 }}>
			You were the {ConvertPositionToString(position)} person to guess today's NPC!
		</p>
	{:else}
		<br />
	{/if}
	<p>{playerCount} people guessed todays weapon</p>

	<div class="bonus-container">
		<h2 class="bonus-title">Bonus Round!</h2>
		<p>Out of the following names, what name can the {npc.name} have?</p>
		{#if guessedName && correctName === guessedName}
			<h2 in:scale={{ duration: 600 }} class="winning-card">Nailed It!</h2>
		{:else if guessedName && correctName !== guessedName}
			<h2 in:scale={{ duration: 600 }}>Better Luck Next Time</h2>
		{/if}
		<div class="bonus-options">
			{#each names as name}
				<button
					class="bonus-button"
					onclick={() => guessName(name)}
					disabled={!!guessedName}
					class:Success={correctName === name}
					class:Fail={guessedName === name && correctName !== guessedName}
				>
					{name}
				</button>
			{/each}
		</div>
	</div>

	<RemainingTime />
</div>

<style>
	.winning-container {
		border: 20px solid transparent;
		border-image: url('/daily-slash/borders/Wood.png');
		border-image-slice: 17;
		border-image-repeat: round;
		border-radius: 5px;

		background-image: url('/daily-slash/backgrounds/WoodWall.png');
		background-repeat: repeat;
		background-size: 20%;

		width: fit-content;
		margin: auto;
		margin-top: 10px;
		padding: 20px;
		padding-bottom: 0;
	}

	h1 {
		margin-top: 5px;
	}

	.npc-housing {
		position: relative;
	}

	.candle-light {
		position: absolute;
		width: 40px;
		height: 40px;
		background-color: rgb(255, 204, 62);
		border-radius: 50%;
		filter: blur(18px);
		left: 50%;
		margin-left: -59px;
		margin-bottom: 5px;
	}

	.candle {
		position: absolute;
		margin-left: 19px;
		margin-bottom: 5px;
	}

	.npc-image {
		padding-left: 20px;
	}

	.bonus-container {
		width: 350px;
		margin: 50px auto 20px auto;
		padding: 10px;
		background: rgb(31, 47, 82);

		border-radius: 15px;
		border: 2px solid black;
	}

	.bonus-title {
		background-color: var(--color-lightblue);
		width: fit-content;
		margin: auto;
		margin-top: -40px;
		margin-bottom: 10px;
		padding: 10px 20px;

		border-radius: 15px;
		border: 2px solid black;
	}

	.bonus-options {
		display: grid;
		grid-template-columns: repeat(2, 1fr);
		grid-auto-rows: 1fr;

		width: 100%;
		height: 100%;
		margin: 0 auto;
		gap: 5px;
	}

	.bonus-button {
		background-color: var(--color-button);
		position: relative;
		height: 50px;
		font-size: 16px;
		text-align: center;

		border-radius: 5px;
		border: thin solid black;
		transition: background-color 0.2s ease;
	}

	.bonus-button:not(:disabled):hover {
		cursor: pointer;
		background-color: var(--color-lightblue);
	}

	.bonus-button:disabled {
		cursor: not-allowed;
	}

	.bonus-button.Success {
		background-color: var(--color-green);
	}

	.bonus-button.Fail {
		background-color: var(--color-red);
	}

	.winning-card {
		color: yellow;
	}
</style>
