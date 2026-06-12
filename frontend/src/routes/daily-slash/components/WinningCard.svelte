<script lang="ts">
	import { subscribeToPlayerCount } from '$lib/api/common';
	import { getDailySlashWinningData } from '$lib/api/daily-slash';
	import RemainingTime from '$lib/components/RemainingTime.svelte';
	import { backgrounds, borders, colors, type Rarity, type Weapon } from '$lib/types/daily-slash';
	import { ConvertPositionToString } from '$lib/utils/posToString';
	import { typewriter } from '$lib/utils/transitions';
	import { onMount } from 'svelte';
	import { scale } from 'svelte/transition';

	let { 
		weaponAnswer
	} : {
		weaponAnswer: Weapon;
	}= $props();

	let position = $state(0);
	let playerCount = $state(0);

	onMount(async () => {
		const winningData = await getDailySlashWinningData();

		position = winningData.position;
	});

	// Streams live player count to the user
	onMount(() => {
		return subscribeToPlayerCount('daily_slash', (count) => {
			playerCount = count;
		});
	});
</script>

<div
	class="wrapper"
	style="
        border-image-source: {borders[weaponAnswer.rarity as Rarity]}; 
        background: {backgrounds[weaponAnswer.rarity as Rarity]}"
	in:scale
>
	<h1>You Got It!</h1>
	
	{#if position !== 0}
		<p transition:typewriter={{ speed: 1 }}>
			You were the {ConvertPositionToString(position)} person to guess today's weapon!
		</p>
	{:else}
		<br />
	{/if}

	<img
		style="border-color: {colors[weaponAnswer.rarity as Rarity]}"
		src={`/weapons/${weaponAnswer.image_path}`}
		alt="Previous weapon"
		in:scale
	/>
	
	<h3 class="weapon-name" style="color: {colors[weaponAnswer.rarity as Rarity]}">{weaponAnswer.name}</h3>
	<p>{playerCount} people guessed todays weapon</p>

	<RemainingTime />
</div>

<style>
	.wrapper {
		border: 20px solid transparent;
		border-radius: 5px;
		border-image-slice: 17;
		border-image-repeat: round;

		background-repeat: repeat;
		background-size: 20%;

		display: flex;
		justify-content: center;
		flex-direction: column;
		align-items: center;
		text-align: center;
		padding: 0 20px;
		width: 260px;
		margin: 15px auto;
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
	}

	img {
		background-color: var(--color-button);
		padding: 18px;
		width: 45px;
		height: 45px;
		object-fit: contain;

		border-radius: 15px;
		border: 2px solid;
	}

	p {
		margin-top: 0;
	}
</style>
