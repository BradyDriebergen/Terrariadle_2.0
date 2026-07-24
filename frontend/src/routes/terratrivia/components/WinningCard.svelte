<script lang="ts">
	import { page } from '$app/state';
	import { subscribeToPlayerCount } from '$lib/api/common';
	import { getTerraTriviaWinningData } from '$lib/api/terratrivia';
	import RemainingTime from '$lib/components/RemainingTime.svelte';
	import { ConvertPositionToString } from '$lib/utils/posToString';
	import { onMount } from 'svelte';
	import { slide } from 'svelte/transition';

	let position: number = $state(0);
	let playerCount: number = $state(0);

	onMount(async () => {
		const winningData = await getTerraTriviaWinningData(page.data.userId);

		position = winningData.position;
	});

	// Streams live player count to the user
	onMount(() => {
		return subscribeToPlayerCount('terratrivia', (count) => {
			playerCount = count;
		});
	});
</script>

<span class="color-cycle">TerraTrivia Results</span>

<div class="container" in:slide>
	<h1>Fantastic!</h1>
	<p>
		You were the {ConvertPositionToString(position)} person to solved the puzzle.
	</p>

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
