<script lang="ts">
	import Confetti from '$lib/components/Confetti.svelte';
	import GuessList from './components/GuessList.svelte';
	import UserInput from './components/UserInput.svelte';
	import GameInfo from './components/GameInfo.svelte';
	import WinningCard from './components/WinningCard.svelte';
	import { finished } from 'stream';
	import { checkWeaponGuess } from '$lib/api/daily-slash';

	let { data } = $props();

	let gameContext = $derived(data.gameContext)
	let weaponList = $derived(data.weaponList)

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
	<UserInput guesses={gameContext.guesses} {submitGuess} weaponList={weaponList} won={gameContext.finished} />

	{#if gameContext.finished}
		<WinningCard weapon={gameContext.guesses[0]} userId={data.userId} />
	{/if}

	{#if gameContext.guesses.length < 1}
		<GameInfo previousWeapon={gameContext.previous_weapon} />
	{:else}
		<GuessList guesses={gameContext.guesses} checks={gameContext.guesses} />
	{/if}

	<Confetti won={gameContext.finished} />
</div>
