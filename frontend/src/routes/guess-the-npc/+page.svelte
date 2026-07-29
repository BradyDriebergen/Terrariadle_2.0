<script lang="ts">
	import GuessList from './components/GuessList.svelte';
	import UserInput from './components/UserInput.svelte';
	import WinningCard from './components/WinningCard.svelte';
	import Confetti from '$lib/components/Confetti.svelte';
	import type { PageData } from './$types';
	import type { NpcGuess } from '$lib/types/guess-the-npc';
	import type { DropdownListItem } from '$lib/types/shared';

	let { data }: { data: PageData } = $props();

	let guesses: NpcGuess[] = $state([]);
	let guessedIds: number[] = $state([]);
	let finished: boolean = $state(false);

	$effect(() => {
		// Initialize data once pre-fetch is finished
		if (data.gameContext) {
			guesses = data.gameContext.guesses;
			guessedIds = data.gameContext.guessed_ids;
			finished = data.gameContext.finished;
		}
	});

	let quote: string = $derived(data.gameContext?.quote ?? '');
	let npcList: DropdownListItem[] = $derived(
		(data.npcList ?? []).filter((w) => !guessedIds.includes(w.id))
	);
</script>

{#if data.gameContext}
	<UserInput bind:guesses bind:finished {npcList} {quote} />

	{#if finished}
		<WinningCard npc={guesses[0]} />
	{/if}

	<GuessList {guesses} {finished} />

	<Confetti {finished} />
{:else}
	<p>Loading...</p>
{/if}
