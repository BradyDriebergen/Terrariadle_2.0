<script lang="ts">
	import Dropdown from '$lib/components/Dropdown.svelte';

	let { npcList, quote, won } = $props();

	function submitGuess(npcId: number) {
		console.log(npcId);
	}
</script>

<div class="container">
	<h2>Guess the NPC</h2>
	<p>Which NPC says this quote?</p>
	<div class="quote-box">
		<p>{quote}</p>

		<div class="quote-options">
			<span>Shop</span>
			<span>Close</span>
			<span>Happiness</span>
		</div>
	</div>

	<div class="dropdown">
		<Dropdown
			selectItem={(npcId: number) => {
				submitGuess(npcId);
			}}
			itemList={npcList}
			itemName="npc"
		/>
	</div>
</div>

<style>
	.container {
		background-color: var(--color-backgroundblue);
		width: fit-content;
		text-align: center;
		margin: auto;
		margin-top: 40px;
		padding: 15px;

		border-radius: 15px;
		border: thin solid black;
	}

	h2 {
		background-color: var(--color-lightblue);
		width: fit-content;
		margin: auto;
		margin-top: -40px;
		margin-bottom: 10px;
		padding: 10px 20px;

		border-radius: 15px;
		border: 2px solid black;
	}

	.quote-box {
		background-color: var(--color-button);
		width: 500px;
		padding: 20px 30px;
		margin: auto;

		border-radius: 15px;
		border: thin solid black;
	}

	.quote-box p {
		text-align: left;
		margin-top: 0;
		font-size: 19px;
	}

	.quote-options {
		margin-top: 30px;
		text-align: left;
	}

	.quote-options span {
		margin-right: 30px;
		color: yellow;
		cursor: default;
	}

	.dropdown {
		margin-top: 20px;
	}
</style>
