<script lang="ts">
	import Confetti from '$lib/components/Confetti.svelte';
	import GuessList from './components/GuessList.svelte';
	import GuessPanel from './components/GuessPanel.svelte';
	import PrevWeaponGroup from './components/PrevWeaponGroup.svelte';
	import WinningPanel from './components/WinningPanel.svelte';

	let { data } = $props();

  	let guesses = $state(data.guesses);
	let checks = $state(data.checks);
	let weapons = $state(data.weapons);
	let won = $state(data.won);

	async function submitGuess(weaponid: number) {
		fetch('http://localhost:3000/api/daily-slash/check-guess', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ userId: data.userId, guess: weaponid })
		})
		.then(r => r.json())
		.then(data => {
			guesses.unshift(data.guess);
			checks.unshift(data.check);
			won = data.won;

			weapons = weapons.filter(w => w.id !== weaponid);
		});
	}
</script>

<svelte:document style:overflow-y="hidden" />

<div>
	<GuessPanel 
		guesses={guesses}
		submitGuess={submitGuess}
		weaponList={weapons} 
		won={won}
	/>

	{#if won}
		<WinningPanel weapon={guesses[0]} userId={data.userId}/>
	{/if}

	{#if guesses.length < 1}
		<PrevWeaponGroup previousWeapon={data.previousWeapon} />
	{:else}
		<GuessList guesses={guesses} checks={checks}/>
	{/if}

	<Confetti won={won} />
</div>