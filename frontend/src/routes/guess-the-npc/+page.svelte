<script lang="ts">
	import GuessPanel from './components/GuessPanel.svelte';

	let { data } = $props();
</script>

<GuessPanel 
	npcList={data.npcs} 
	quote={data.quote} 
	won={data.won}
/>

<style>
</style>
