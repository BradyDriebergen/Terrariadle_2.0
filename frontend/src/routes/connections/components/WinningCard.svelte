<script lang="ts">
	import { page } from '$app/state';
	import { subscribeToPlayerCount } from '$lib/api/common';
	import { getConnectionsWinningData } from '$lib/api/connections';
	import RemainingTime from '$lib/components/RemainingTime.svelte';
	import { ConvertPositionToString } from '$lib/utils/posToString';
	import { onMount } from 'svelte';
	import { slide } from 'svelte/transition';

	let {
		attempts = 0
	}: {
		attempts: number;
	} = $props();

	let position: number = $state(0);
	let playerCount: number = $state(0);

	onMount(async () => {
		const winningData = await getConnectionsWinningData(page.data.userId);

		position = winningData.position;
	});

	// Streams live player count to the user
	onMount(() => {
		return subscribeToPlayerCount('connections', (count) => {
			playerCount = count;
		});
	});
</script>

<span class="color-cycle">Connections Results</span>

<div class="container" in:slide>
	{#if attempts > 0}
		<h1>Perfect!</h1>
		<p>
			You were the {ConvertPositionToString(position)} person to solved the puzzle. <br />
			You had {attempts}
			{attempts === 1 ? 'attempt' : 'attempts'} remaining!
		</p>
	{:else}
		<h1>Better Luck Next Time!</h1>
	{/if}

	<p>{playerCount} players solved this puzzle successfully</p>
	<RemainingTime />
</div>

<style>
	.container {
		background-color: var(--color-button);
		width: fit-content;
		margin: 20px auto;
		padding: 0 20px;
		border-radius: 10px;
		border: 2px solid black;
	}
</style>
