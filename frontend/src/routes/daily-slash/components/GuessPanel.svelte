<script lang="ts">
	import Dropdown from '$lib/components/Dropdown.svelte';
    import hintLock from '$lib/assets/LockedHint.png';
	import LoadingBar from './LoadingBar.svelte';

	let { guesses, submitGuess, weaponList, won } = $props();
	let guessCount = $derived(guesses?.length ?? 0);

	let hints = $state(['', '', '']);
    let showHints = $state([false, false, false])

    async function revealHint(num: number) {
        if (hints[num - 1]) {
            showHints[num - 1] = !showHints[num - 1];
            return;
        }

        const res = await fetch(`http://localhost:3000/api/daily-slash/hint/${num}`);
        hints[num - 1] = await res.json();
        showHints[num - 1] = true;
    }
</script>

<div class="guess-panel">
	<h2>Guess Today's Weapon</h2>

    <div class="loadingBar">
        <LoadingBar guesses={guessCount} won={won} />
    </div>

	<div class="hint-buttons">
		<button
			disabled={guessCount < 4 && !won}
			onclick={() => revealHint(1)}
		>
            {#if guessCount < 4 && !won} 
                <img class="lock" src={hintLock} alt="Locked hint"/> 
            {/if}
			<span>{showHints[0] ? hints[0] : 'Mode Obtained'}</span>
		</button>
		<button
			disabled={guessCount < 7 && !won}
			onclick={() => revealHint(2)}
		>
            {#if guessCount < 7 && !won} 
                <img class="lock" src={hintLock} alt="Locked hint"/> 
            {/if}
            <span>{showHints[1] ? hints[1] : 'Weapon Type'}</span>
		</button>
		<button
			disabled={guessCount < 12 && !won}
			onclick={() => revealHint(3)}
		>
			{#if showHints[2]}
                <img class="hint-3" src={`/weapons/${hints[2]}`} alt={hints[2]} />
            {:else}
                {#if guessCount < 12 && !won} 
                    <img class="lock" src={hintLock} alt="Locked hint"/> 
                {/if}
                <span>Image Clue</span>
            {/if}
		</button>
	</div>

	{#if !won}
		<div class="dropdown">
			<Dropdown
				selectItem={(weaponid: number) => {
					submitGuess(weaponid);
				}}
				itemList={weaponList}
				itemName="weapon"
			/>
		</div>
	{/if}
</div>

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
