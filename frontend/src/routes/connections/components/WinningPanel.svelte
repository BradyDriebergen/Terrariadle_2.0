<script lang="ts">
	import RemainingTime from '$lib/components/RemainingTime.svelte';
	import { ConvertPositionToString } from '$lib/utils/posToString';
	import { onMount } from 'svelte';
	import { slide } from 'svelte/transition';

	let { won, attempts, userId } = $props();

	let pos = $state(0);
	let count = $state(0);

	onMount(async () => {
		const winningData = await fetch(`http://localhost:3000/api/connections/winning-data/${userId}`);
		const winningDataJson = await winningData.json();

		pos = winningDataJson.pos;
		count = winningDataJson.count;
	});
</script>

<span class="color-cycle">Connections Results</span>

<div class="container" in:slide>
	{#if won}
		<h1>Perfect!</h1>
		<p>
			You were the {ConvertPositionToString(pos)} person to solved the puzzle. <br />
			You had {attempts}
			{attempts === 1 ? 'attempt' : 'attempts'} remaining!
		</p>
	{:else}
		<h1>Better Luck Next Time!</h1>
	{/if}

	<p>{count} players solved this puzzle successfully</p>
	<RemainingTime />
</div>

<style>
	.container {
		background-color: var(--color-button);
		width: fit-content;
		margin: auto;
		margin-bottom: 20px;
		padding: 0 20px;
		border-radius: 10px;
		border: 2px solid black;
	}
</style>
