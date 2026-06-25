<script lang="ts">
	import GuessList from './components/GuessList.svelte';
	import UserInput from './components/UserInput.svelte';
	import WinningComponent from './components/WinningComponent.svelte';
	import Confetti from '$lib/components/Confetti.svelte';
	import type { PageData } from './$types';
	import type { NpcGuess } from '$lib/types/guess-the-npc';
	import type { DropdownListItem } from '$lib/types/shared';

	let { data }: { data: PageData } = $props();

	let guesses: NpcGuess[] = $state([]);
	let finished: boolean = $state(false);

	$effect(() => {
		// Initialize data once pre-fetch is finished
		if (data.gameContext) {
			guesses = data.gameContext.guesses;
			finished = data.gameContext.finished;
		}
	});

	let quote: string = $derived(data.gameContext?.quote ?? '')
	let npcList: DropdownListItem[] = $derived(
		(data.npcList ?? []).filter((w) => !data.gameContext?.guessed_ids.includes(w.id))
	);
</script>

{#if data.gameContext}
	<UserInput bind:guesses bind:finished {npcList} {quote} />

	{#if finished}
		<WinningComponent npc={guesses[0]} />
	{/if}

	<GuessList {guesses} {finished} />

	<Confetti {finished} />
{:else}
	<p>Loading...</p>
{/if}

<style>
</style>