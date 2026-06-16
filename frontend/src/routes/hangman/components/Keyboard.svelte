<script lang="ts">
	import { cubicInOut } from 'svelte/easing';
	import { slide } from 'svelte/transition';

	let { onKeyPressed, enemyLetters, guessedLetters } = $props();

	const rows: string[][] = [
		['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'],
		['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L'],
		['Z', 'X', 'C', 'V', 'B', 'N', 'M']
	];
</script>

<div class="keyboard" out:slide={{ duration: 700, easing: cubicInOut }}>
	{#each rows as row, rowIndex (rowIndex)}
		<div class="keyboard-row">
			{#each row as letter (letter)}
				<button
					class="keyboard-key"
					class:correct={enemyLetters.includes(letter)}
					class:incorrect={!enemyLetters.includes(letter)}
					disabled={guessedLetters.includes(letter)}
					onclick={() => onKeyPressed(letter)}
				>
					{letter}
				</button>
			{/each}
		</div>
	{/each}
</div>

<style>
	.keyboard {
		display: flex;
		flex-direction: column;
		gap: 8px;
		width: 100%;
	}

	.keyboard-row {
		display: flex;
		justify-content: center;
		gap: 5px;
	}

	.keyboard-key {
		display: flex;
		align-items: center;
		justify-content: center;
		background: var(--color-backgroundblue);
		width: 50px;
		height: 62px;
		font-size: 17px;
		border: 2px solid #000000;
		border-radius: 5px;
		transition: background-color 0.2s ease;
	}

	.keyboard-key:hover:not(:disabled) {
		background: var(--color-lightblue);
		cursor: pointer;
	}

	.keyboard-key:disabled {
		cursor: not-allowed;
	}

	.keyboard-key:disabled.correct {
		background: var(--color-green);
	}

	.keyboard-key:disabled.incorrect {
		background: var(--color-red);
	}
</style>
