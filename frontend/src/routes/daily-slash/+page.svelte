<script lang="ts">
	import GuessPanel from './components/GuessPanel.svelte';
	import PrevWeaponGroup from './components/PrevWeaponGroup.svelte';
	import { setContext } from 'svelte';

	let { data } = $props();

	setContext('weapons', data.weapons);

  	let guesses = $state(0);
</script>

<div>
  <button onclick={() => guesses++}>Increase Guesses</button>
	<GuessPanel numberOfGuesses={guesses} />
	<PrevWeaponGroup previousWeapon={data.previousWeapon} />
</div>
