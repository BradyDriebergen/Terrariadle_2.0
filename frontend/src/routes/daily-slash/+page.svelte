<script lang="ts">
	import GuessList from './components/GuessList.svelte';
	import GuessPanel from './components/GuessPanel.svelte';
	import PrevWeaponGroup from './components/PrevWeaponGroup.svelte';

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
		<p>You won!</p>
	{/if}

	{#if guesses.length < 1}
		<PrevWeaponGroup previousWeapon={data.previousWeapon} />
	{:else}
		<GuessList guesses={guesses} checks={checks}/>
	{/if}
</div>