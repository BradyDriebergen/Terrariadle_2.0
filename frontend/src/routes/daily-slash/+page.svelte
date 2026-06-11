<script lang="ts">
	import Confetti from '$lib/components/Confetti.svelte';
	import GuessList from './components/GuessList.svelte';
	import UserInput from './components/UserInput.svelte';
	import GameInfo from './components/GameInfo.svelte';
	import WinningCard from './components/WinningCard.svelte';
	import { checkWeaponGuess } from '$lib/api/daily-slash';
	import type { WeaponListItem, WeaponGuess, WeaponPreview } from '$lib/types/daily-slash';

	let { data } = $props();

	$inspect(data);

	let guesses = $state<WeaponGuess[]>([]);
	let prevWeapon = $state<WeaponPreview | null>(null);
	let finished = $state<boolean>(false);

	$effect(() => {
		// Initialize data once pre-fetch is finished
		if (data.gameContext) {
			guesses = data.gameContext.guesses;
			prevWeapon = data.gameContext.previous_weapon;
			finished = data.gameContext.finished;
		}
	});

	let weaponList = $derived<WeaponListItem[]>(
		(data.weaponList ?? []).filter((w) => !data.gameContext?.guessed_ids.includes(w.weapon_id))
	);

	async function submitGuess(weaponId: number) {
		try {
			const result = await checkWeaponGuess(weaponId);

			// guesses.unshift(result.weapon);
			// checks.unshift(result.check);
			// won = result.won;
			// weapons = weapons.filter((w) => w.id !== weaponId);
		} catch (err) {
			// handle UI feedback here, e.g.:
			console.error(err);
		}
	}
</script>

<svelte:document style:overflow-y="hidden" />

<div>
	<UserInput {guesses} {weaponList} {finished} />

	{#if finished}
		<WinningCard weapon={guesses[0]} userId={data.userId} />
	{/if}

	{#if guesses.length < 1}
		<GameInfo {prevWeapon} />
	{:else}
		<GuessList {guesses} checks={guesses} />
	{/if}

	<Confetti won={finished} />
</div>
