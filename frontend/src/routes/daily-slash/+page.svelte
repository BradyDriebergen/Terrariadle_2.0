<script lang="ts">
	import GuessPanel from './components/GuessPanel.svelte';
	import PrevWeaponGroup from './components/PrevWeaponGroup.svelte';

	let { data } = $props();

  	let guesses = $state(data.guesses);
	let weapons = $state(data.weapons);

	$inspect(guesses);

	async function submitGuess(weaponid: number) {
		fetch('http://localhost:3000/api/daily-slash/check-guess', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ userId: data.userId, guess: weaponid })
		})
		.then(r => r.json())
		.then(data => {
			guesses.push(data);

			weapons = weapons.filter(w => w.id !== weaponid);
		});
	}
</script>

<div>
	<GuessPanel 
		guesses={guesses}
		submitGuess={submitGuess}
		weaponList={weapons} 
	/>

	{#if guesses.length < 1}
		<PrevWeaponGroup previousWeapon={data.previousWeapon} />
	{/if}
</div>
