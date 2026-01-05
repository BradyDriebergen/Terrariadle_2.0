<script lang="ts">
	import { colors, type Rarity } from '$lib/types/dailySlash';
	import { flip } from 'svelte/animate';
	import { fly } from 'svelte/transition';
	let { guesses, checks } = $props();

	const guessCorrect = 'background-color: var(--color-green);';
	const guessPartial = 'background-color: var(--color-yellow);';
	const guessWrong = 'background-color: var(--color-red);';

	function checkedObtained(value: number): string {
		if (value === 2) {
			return guessCorrect;
		} else if (value === 1) {
			return guessPartial;
		}
		return guessWrong;
	}

	function getRarityColor(rarity: string) {
		return 'color: ' + colors[rarity as Rarity];
	}
</script>

<div in:fly={{ x: 800, duration: 1000 }}>
	<div class="header">
		<span title="Image of Weapon">Weapon</span>
		<span title="Melee, Ranged, Magic, Summon, Throwing">Damage Type</span>
		<span title="Weapon Damage">Damage</span>
		<span title="Snail, Extremely Slow, Very Slow, Slow, Average, Fast, Very Fast, Insanely Fast"
			>Use Time</span
		>
		<span title="White, Blue, Green, Orange, Light_Red, Pink, Light_Purple, Lime, Yellow, Cyan, Red"
			>Rarity</span
		>
		<span title="Auto or Manual">Operation</span>
		<span title="Yes or No">Material</span>
		<span title="Crafting, Chest, Buy, Drop, Fishing, Background Object">Obtained</span>
	</div>
	{#each guesses as guess, i (guess.id)}
		<div class="row" animate:flip>
			<div in:fly={{ x: 800, duration: 2000 }}>
				<img src={`/weapons/${guess.info['image-path']}`} alt={`${guess.name} image`} />
			</div>

			<span
				style={checks[i].DamageType ? guessCorrect : guessWrong}
				in:fly={{ x: 700, duration: 2000 }}
			>
				{guess.info['damage-type']}
			</span>

			<!-- DAMAGE: triangle behind text -->
			<span
				class="arrow-cell"
				class:arrow-up={checks[i].Damage === 2}
				class:arrow-down={checks[i].Damage === 0}
				style={checks[i].Damage === 1 ? guessCorrect : guessWrong}
				in:fly={{ x: 600, duration: 2000 }}
			>
				{guess.info.damage}
			</span>

			<!-- USE TIME: triangle behind text -->
			<span
				class="arrow-cell"
				class:arrow-up={checks[i].UseTime === 2}
				class:arrow-down={checks[i].UseTime === 0}
				style={checks[i].UseTime === 1 ? guessCorrect : guessWrong}
				in:fly={{ x: 500, duration: 2000 }}
			>
				{guess.info['use-time']}
			</span>

			<!-- RARITY: triangle behind text -->
			<span
				class="arrow-cell"
				class:arrow-up={checks[i].Rarity === 2}
				class:arrow-down={checks[i].Rarity === 0}
				style={(checks[i].Rarity === 1 ? guessCorrect : guessWrong) +
					getRarityColor(guess.info.rarity as string)}
				in:fly={{ x: 400, duration: 2000 }}
			>
				{guess.info.rarity}
			</span>

			<span
				style={checks[i].Operation ? guessCorrect : guessWrong}
				in:fly={{ x: 300, duration: 2000 }}
			>
				{guess.info.operation}
			</span>
			<span
				style={checks[i].Material ? guessCorrect : guessWrong}
				in:fly={{ x: 200, duration: 2000 }}
			>
				{guess.info.material}
			</span>

			<div
				style="flex-direction: column; {checkedObtained(checks[i].Obtained)}"
				in:fly={{ x: 100, duration: 2000 }}
			>
				{#each guess.info.obtained as item}
					<span>{item}</span>
				{/each}
			</div>
		</div>
	{/each}
</div>

<style>
	.header {
		display: flex;
		margin: auto;
		margin-top: 10px;
		width: fit-content;
		padding: 5px 10px;
		gap: 12px;
		border-color: transparent;
		border-bottom: #fff;
		border-style: solid;
	}

	.header span {
		font-size: 13px;
		width: 80px;
	}

	.row {
		display: flex;
		justify-content: center;
		align-items: center;
		margin: 5px auto;
		width: fit-content;
		gap: 8px;
	}

	.row div,
	.row span {
		background-color: rgba(0, 0, 0, 0.4);
		display: flex;
		justify-content: center;
		align-items: center;
		width: 80px;
		height: 80px;
		border-radius: 15px;
		border: 2px solid black;

		/* allow layering inside each cell */
		position: relative;
		z-index: 0;
	}

	.row div span {
		background: none;
		border: none;
		padding: 0;
		gap: 0;
		position: static;
		z-index: auto;
	}

	.row div img {
		width: 50px;
		height: 50px;
		object-fit: contain;
	}

	/* ===== Arrow background layer ===== */

	.arrow-cell::after {
		content: '';
		position: absolute;
		width: 0;
		height: 0;
		z-index: -1; /* BELOW text, ABOVE background */

		/* default: no visible arrow until arrow-up / arrow-down applies */
		border-left: 0 solid transparent;
		border-right: 0 solid transparent;
		border-top: 0 solid transparent;
		border-bottom: 0 solid transparent;
	}

	/* Up-pointing triangle */
	.arrow-cell.arrow-up::after {
		border-left: 30px solid transparent;
		border-right: 30px solid transparent;
		border-bottom: 55px solid rgba(0, 0, 0, 0.2);
	}

	/* Down-pointing triangle */
	.arrow-cell.arrow-down::after {
		border-left: 30px solid transparent;
		border-right: 30px solid transparent;
		border-top: 55px solid rgba(0, 0, 0, 0.2);
	}
</style>
