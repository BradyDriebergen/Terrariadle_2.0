<script lang="ts">
	import GuessList from './components/GuessList.svelte';
	import GuessPanel from './components/GuessPanel.svelte';
	import WinningComponent from './components/WinningComponent.svelte';
	import Confetti from '$lib/components/Confetti.svelte';

	let { data } = $props();

	let npcs = $state(data.npcs);
	let guesses = $state(data.guesses);
	let won = $state(data.won);

	async function submitGuess(npcId: number) {
		fetch('http://localhost:3000/api/guess-the-npc/check-guess', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ userId: data.userId, guess: npcId })
		})
			.then((r) => r.json())
			.then((data) => {
				guesses.unshift(data.guess);
				won = data.won;

				npcs = npcs.filter((n) => n.id !== npcId);
			});
	}
</script>

<GuessPanel 
	submitGuess={submitGuess}
	npcList={npcs} 
	quote={data.quote} 
	won={won}
/>

{#if won}
	<WinningComponent
		npc={guesses[0]}
		userId={data.userId}
	/>
{/if}

<GuessList
	guesses={guesses}
	won={won}
/>

<Confetti {won} />

<style>
</style>
