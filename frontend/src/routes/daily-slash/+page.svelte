<script lang="ts">
	import Confetti from '$lib/components/Confetti.svelte';
	import GuessList from './components/GuessList.svelte';
	import UserInput from './components/UserInput.svelte';
	import GameInfo from './components/GameInfo.svelte';
	import WinningCard from './components/WinningCard.svelte';
	import type { WeaponGuess, Weapon, WeaponPreview } from '$lib/types/daily-slash';
	import type { DropdownListItem } from '$lib/types/common';

	let { data } = $props();

	$inspect(data);

	let guesses = $state<WeaponGuess[]>([]);
	let prevWeapon = $state<WeaponPreview | null>(null);
	let finished = $state<boolean>(false);

	let correctWeapon = $derived<Weapon | null>(finished ? guesses[0].weapon : null);

	$effect(() => {
		// Initialize data once pre-fetch is finished
		if (data.gameContext) {
			guesses = data.gameContext.guesses;
			prevWeapon = data.gameContext.previous_weapon;
			finished = data.gameContext.finished;
		}
	});

	let weaponList = $derived<DropdownListItem[]>(
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
