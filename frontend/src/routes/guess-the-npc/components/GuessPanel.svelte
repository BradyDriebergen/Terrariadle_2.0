<script lang="ts">
	import Dropdown from '$lib/components/Dropdown.svelte';
	import { cubicInOut } from 'svelte/easing';
	import { slide } from 'svelte/transition';

	let { submitGuess, npcList, quote, won } = $props();
</script>

{#if won}
	<span class="color-cycle">Guess The NPC Results</span>
{/if}

<div class="container" style="margin-top: {won ? 15 : 40}px;">
	{#if !won}
		<h2 out:slide={{ duration: 700, easing: cubicInOut }}>Guess the NPC</h2>
		<p out:slide={{ duration: 700, easing: cubicInOut }}>Which NPC says this quote?</p>
	{/if}
	<div class="quote-box">
		<p>{quote}</p>

		<div class="quote-options">
			<span>Shop</span>
			<span>Close</span>
			<span>Happiness</span>
		</div>
	</div>

	{#if !won}
		<div class="dropdown" out:slide={{ duration: 700, easing: cubicInOut }}>
			<Dropdown
				selectItem={(npcId: number) => {
					submitGuess(npcId);
				}}
				itemList={npcList}
				itemName="npc"
			/>
		</div>
	{/if}
</div>

<style>
	.container {
		background-color: var(--color-backgroundblue);
		width: fit-content;
		text-align: center;
		margin: auto;
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

	.color-cycle {
		font-size: 22px;
		animation: colors 6s linear infinite;
	}

	@keyframes colors {
		0% {
			color: rgb(255, 0, 0);
		}
		20% {
			color: rgb(255, 166, 0);
		}
		40% {
			color: rgb(255, 255, 0);
		}
		60% {
			color: rgb(0, 255, 0);
		}
		80% {
			color: rgb(0, 162, 255);
		}
		100% {
			color: rgb(255, 0, 0);
		}
	}
</style>
