<script lang="ts">
	import Confetti from '$lib/components/Confetti.svelte';
	import GuessList from './components/GuessList.svelte';
	import UserInput from './components/UserInput.svelte';
	import GameInfo from './components/GameInfo.svelte';
	import WinningCard from './components/WinningCard.svelte';
	import { finished } from 'stream';

	let { data } = $props();

	let gameContext = $derived(data.gameContext)
	let weaponList = $derived(data.weaponList)

	async function submitGuess(weaponid: number) {
		fetch('http://localhost:3000/api/daily-slash/check-guess', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ userId: data.userId, guess: weaponid })
		})
			.then((r) => r.json())
			.then((data) => {
				guesses.unshift(data.guess);
				checks.unshift(data.check);
				won = data.won;

				weapons = weapons.filter((w) => w.id !== weaponid);
			});
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
