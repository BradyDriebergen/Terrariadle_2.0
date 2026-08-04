<script lang="ts">
	import Confetti from '$lib/components/Confetti.svelte';
	import GuessList from './components/GuessList.svelte';
	import UserInput from './components/UserInput.svelte';
	import GameInfo from './components/GameInfo.svelte';
	import WinningCard from './components/WinningCard.svelte';
	import type { WeaponGuess, Weapon, WeaponPreview } from '$lib/types/daily-slash';
	import type { DropdownListItem } from '$lib/types/shared';
	import type { PageData } from './$types';
	import { scale } from 'svelte/transition';

	let { data }: { data: PageData } = $props();

	let guesses: WeaponGuess[] = $state([]);
	let guessedIds: number[] = $state([]);
	let prevWeapon: WeaponPreview | null = $state(null);
	let finished: boolean = $state(false);

	let correctWeapon: Weapon | null = $derived(finished ? guesses[0].weapon : null);
	let width: number = $state(0);

	let pageLoading: boolean = $state(true);
	$effect(() => {
		// Initialize data once pre-fetch is finished
		if (data.gameContext) {
			guesses = data.gameContext.guesses;
			guessedIds = data.gameContext.guessed_ids;
			prevWeapon = data.gameContext.previous_weapon;
			finished = data.gameContext.finished;

			pageLoading = false;
		}
	});

	let weaponList: DropdownListItem[] = $derived(
		(data.weaponList ?? []).filter((w) => !guessedIds.includes(w.id))
	);
</script>

<svelte:window bind:innerWidth={width} />

<svelte:document style:overflow-y="hidden" />

{#if !pageLoading}
	<div>
		<UserInput bind:guesses bind:finished {weaponList} />

		{#if finished}
			<WinningCard weaponAnswer={correctWeapon!} />
		{/if}

		{#if guesses.length < 1}
			<GameInfo {prevWeapon} />
		{:else}
			{#if width <= 700}
				<p class="mobile-msg" in:scale>← swipe to view →</p>
			{/if}

			<GuessList {guesses} />
		{/if}

		<Confetti {finished} />
	</div>
{:else}
	<p>Loading...</p>
{/if}

<style>
	.mobile-msg {
		color: grey;
		margin: 10px auto -5px;
	}
</style>
