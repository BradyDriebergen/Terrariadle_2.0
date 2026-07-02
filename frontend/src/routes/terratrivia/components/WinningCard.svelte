<script lang="ts">
	import { subscribeToPlayerCount } from '$lib/api/common';
	import { getTerraTriviaWinningData } from '$lib/api/terratrivia';
	import RemainingTime from '$lib/components/RemainingTime.svelte';
	import { userIdStore } from '$lib/store/session';
	import { ConvertPositionToString } from '$lib/utils/posToString';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import { slide } from 'svelte/transition';

	// let {
	// 	attempts = 0
	// }: {
	// 	attempts: number;
	// } = $props();

	let position: number = $state(0);
	let playerCount: number = $state(0);

	onMount(async () => {
		const userId = get(userIdStore);
		const winningData = await getTerraTriviaWinningData(userId);

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
