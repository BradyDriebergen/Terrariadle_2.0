<script lang="ts">
	import { page } from '$app/state';
	import { checkNpcGuess } from '$lib/api/guess-the-npc';
	import Dropdown from '$lib/components/Dropdown.svelte';
	import { session } from '$lib/store/session.svelte';
	import type { NpcGuess } from '$lib/types/guess-the-npc';
	import type { DropdownListItem } from '$lib/types/shared';
	import { cubicInOut } from 'svelte/easing';
	import { slide } from 'svelte/transition';

	let {
		guesses = $bindable<NpcGuess[]>([]),
		finished = $bindable<boolean>(false),
		quote = '',
		npcList = []
	}: {
		guesses: NpcGuess[];
		finished: boolean;
		quote: string;
		npcList: DropdownListItem[];
	} = $props();

	// svelte-ignore state_referenced_locally
	let npcs: DropdownListItem[] = $state(npcList);

	async function submitGuess(npcId: number) {
		try {
			const res = await checkNpcGuess(page.data.userId, npcId);
			guesses = [res.guess, ...guesses];
			finished = res.finished;
			npcs = npcs.filter((w) => w.id !== npcId);

			if (finished) {
				session.guessTheNpcStatus = true;
			}
		} catch (e) {
			alert('Games have refreshed! Refresh the page to start guessing.');

			// TODO: handle error here
			console.error(e);
		}
	}
</script>

<div class="container" style="margin-top: {finished ? 15 : 20}px;">
	{#if !finished}
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

	{#if !finished}
		<div class="dropdown" out:slide={{ duration: 700, easing: cubicInOut }}>
			<Dropdown
				selectItem={(npcId: number) => {
					submitGuess(npcId);
				}}
				itemList={npcs}
				itemType="npc"
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

	@media (max-width: 650px) {
		.quote-box {
			width: 80%;
		}
	}
</style>
