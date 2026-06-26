<script lang="ts">
	import Dropdown from '$lib/components/Dropdown.svelte';
	import LoadingBar from './LoadingBar.svelte';
	import { scale, slide } from 'svelte/transition';
	import { cubicInOut } from 'svelte/easing';
	import type { WeaponGuess } from '$lib/types/daily-slash';
	import { checkWeaponGuess, getWeaponHint } from '$lib/api/daily-slash';
	import type { DropdownListItem } from '$lib/types/shared';
	import { get } from 'svelte/store';
	import { userIdStore } from '$lib/store/session';

	let {
		guesses = $bindable<WeaponGuess[]>([]),
		finished = $bindable<boolean>(false),
		weaponList = []
	}: {
		guesses: WeaponGuess[];
		finished: boolean;
		weaponList: DropdownListItem[];
	} = $props();

	// svelte-ignore state_referenced_locally
	let weapons: DropdownListItem[] = $state(weaponList);

	let guessCount: number = $derived(guesses?.length ?? 0);

	type HintState = { text: string; visible: boolean };
	let hints = $state<HintState[]>([
		{ text: '', visible: false },
		{ text: '', visible: false },
		{ text: '', visible: false }
	]);

	let hint1Locked: boolean = $derived(guessCount < 4 && !finished);
	let hint2Locked: boolean = $derived(guessCount < 7 && !finished);
	let hint3Locked: boolean = $derived(guessCount < 12 && !finished);

	async function revealHint(num: number) {
		try {
			if (hints[num - 1].text) {
				hints[num - 1].visible = !hints[num - 1].visible;
				return;
			}

			const res = await getWeaponHint(num);
			hints[num - 1] = { text: res, visible: true };
		} catch (e: any) {
			// handle error here
			console.error(e);
		}
	}

	async function submitGuess(weaponId: number) {
		try {
			const userId = get(userIdStore);
			const res = await checkWeaponGuess(userId, weaponId);
			guesses = [res.guess_result, ...guesses];
			finished = res.finished;
			weapons = weapons.filter((w) => w.id !== weaponId);
		} catch (e) {
			// handle error here
			console.error(e);
		}
	}
</script>

{#if !finished}
	<div class="guess-panel" out:slide={{ duration: 700, easing: cubicInOut }}>
		<h2>Guess Today's Weapon</h2>

		<div class="loadingBar" class:finished>
			<LoadingBar {guessCount} {finished} />
		</div>

		<div class="hint-buttons">
			<!-- Hint 1 -->
			<button disabled={hint1Locked} onclick={() => revealHint(1)}>
				{#if hint1Locked}
					<img
						class="lock"
						out:scale={{ duration: 1000 }}
						src="/daily-slash/LockedHint.png"
						alt="Locked hint 1"
					/>
				{/if}
				<span>{hints[0].visible ? hints[0].text : 'Mode Obtained'}</span>
			</button>

			<!-- Hint 2 -->
			<button disabled={hint2Locked} onclick={() => revealHint(2)}>
				{#if hint2Locked}
					<img
						class="lock"
						out:scale={{ duration: 1000 }}
						src="/daily-slash/LockedHint.png"
						alt="Locked hint 2"
					/>
				{/if}
				<span>{hints[1].visible ? hints[1].text : 'Weapon Type'}</span>
			</button>

			<!-- Hint 3 -->
			<button disabled={hint3Locked} onclick={() => revealHint(3)}>
				{#if !hints[2].visible}
					{#if hint3Locked}
						<img
							class="lock"
							out:scale={{ duration: 1000 }}
							src="/daily-slash/LockedHint.png"
							alt="Locked hint 3"
						/>
					{/if}
					<span>Image Clue</span>
				{:else}
					<img class="hint-3" src={`/weapons/${hints[2].text}`} alt={hints[2].text} />
				{/if}
			</button>
		</div>

		<div class="dropdown">
			<Dropdown
				selectItem={(weaponid: number) => {
					submitGuess(weaponid);
				}}
				itemList={weapons}
				itemType="weapon"
			/>
		</div>
	</div>
{:else}
	<span class="color-cycle">Daily Slash Results</span>
{/if}

<style>
	.guess-panel {
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

	.hint-buttons {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 10px;
	}

	.hint-buttons button {
		background-color: var(--color-button);
		position: relative;
		width: 100px;
		height: 100px;
		font-size: 16px;
		text-align: center;

		border-radius: 15px;
		border: thin solid black;
		transition: background-color 0.2s ease;
	}
	.hint-buttons button .lock {
		position: absolute;
		top: 9px;
		left: 24px;
		width: 50px;
		opacity: 50%;
	}
	.hint-buttons button .hint-3 {
		width: 45px;
		height: 45px;
		object-fit: contain;
		filter: blur(4px);
	}

	.hint-buttons button:not(:disabled):hover {
		cursor: pointer;
		background-color: var(--color-lightblue);
	}

	.loadingBar {
		margin-top: 25px;
		margin-bottom: 10px;
	}

	.dropdown {
		margin-top: 15px;
		margin-bottom: 5px;
	}
</style>
