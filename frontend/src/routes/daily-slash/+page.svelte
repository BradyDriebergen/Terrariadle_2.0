<script lang="ts">
	import Confetti from '$lib/components/Confetti.svelte';
	import GuessList from './components/GuessList.svelte';
	import UserInput from './components/UserInput.svelte';
	import GameInfo from './components/GameInfo.svelte';
	import WinningCard from './components/WinningCard.svelte';
	import type { WeaponGuess, Weapon, WeaponPreview } from '$lib/types/daily-slash';
	import type { DropdownListItem } from '$lib/types/shared';

	let { data } = $props();

	let guesses: WeaponGuess[] = $state([]);
	let prevWeapon: WeaponPreview | null = $state(null);
	let finished: boolean = $state(false);

	let correctWeapon: Weapon | null = $derived(finished ? guesses[0].weapon : null);

	$effect(() => {
		// Initialize data once pre-fetch is finished
		if (data.gameContext) {
			guesses = data.gameContext.guesses;
			prevWeapon = data.gameContext.previous_weapon;
			finished = data.gameContext.finished;
		}
	});

	let weaponList: DropdownListItem[] = $derived(
		(data.weaponList ?? []).filter((w) => !data.gameContext?.guessed_ids.includes(w.id))
	);
</script>

<svelte:document style:overflow-y="hidden" />

<div>
	<UserInput bind:guesses bind:finished {weaponList} />

	{#if finished}
		<WinningCard weaponAnswer={correctWeapon!} />
	{/if}

	{#if guesses.length < 1}
		<GameInfo {prevWeapon} />
	{:else}
		<GuessList {guesses} />
	{/if}

	<Confetti {finished} />
</div>
