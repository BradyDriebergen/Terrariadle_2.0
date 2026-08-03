<script lang="ts">
	import { colors, CompareResult, type Rarity, type WeaponGuess } from '$lib/types/daily-slash';
	import { useMediaQuery } from '$lib/utils/mediaQuery.svelte';
	import { flip } from 'svelte/animate';
	import { fly } from 'svelte/transition';

	let {
		guesses
	}: {
		guesses: WeaponGuess[];
	} = $props();

	const guessCorrect = 'background-color: var(--color-green);';
	const guessPartial = 'background-color: var(--color-yellow);';
	const guessWrong = 'background-color: var(--color-red);';

	const weaponTitle = 'Image of Weapon';
	const damageTypeTitle = 'Melee, Ranged, Magic, etc.';
	const damageTitle = "Weapon's damage";
	const useTimeTitle = 'Snail, Slow, Average, Very Fast, etc.';
	const rarityTitle = 'White, Green, Pink, Red, etc.';
	const operationTitle = 'Auto or Manual';
	const materialTitle = 'Yes or No';
	const obtainedTitle = 'Crafting, Chest, Buy, etc.';

	function checkedObtained(value: CompareResult): string {
		switch (value) {
			case CompareResult.Match:
				return guessCorrect;
			case CompareResult.PartialMatch:
				return guessPartial;
			case CompareResult.NoMatch:
				return guessWrong;
			default:
				break;
		}

		return guessWrong;
	}

	function getRarityColor(rarity: Rarity): string {
		return 'color: ' + colors[rarity];
	}

	const isMobile = useMediaQuery('(max-width: 800px)');

	const flyX = (x: number) => (isMobile.current ? x / 4 : x);
	const flyDuration = () => (isMobile.current ? 400 : 2000);
</script>

{#snippet body()}
	<div class="header">
		<span title={weaponTitle}>Weapon</span>
		<span title={damageTypeTitle}>Damage Type</span>
		<span title={damageTitle}>Damage</span>
		<span title={useTimeTitle}>Use Time</span>
		<span title={rarityTitle}>Rarity</span>
		<span title={operationTitle}>Operation</span>
		<span title={materialTitle}>Material</span>
		<span title={obtainedTitle}>Obtained</span>
	</div>
	{#each guesses as guess (guess.weapon.id)}
		<div class="row" animate:flip>
			<!-- Weapon icon -->
			<div in:fly={{ x: flyX(560), duration: flyDuration() }}>
				<img src={`/weapons/${guess.weapon.image_path}`} alt={`${guess.weapon.name} image`} />
			</div>

			<!-- Damage type -->
			<span
				style={guess.checks.damage_type ? guessCorrect : guessWrong}
				in:fly={{ x: flyX(480), duration: flyDuration() }}
			>
				{guess.weapon.damage_type}
			</span>

			<!-- Damage -->
			<span
				class="arrow-cell"
				class:arrow-up={guess.checks.damage === CompareResult.Higher}
				class:arrow-down={guess.checks.damage === CompareResult.Lower}
				style={guess.checks.damage === CompareResult.Match ? guessCorrect : guessWrong}
				in:fly={{ x: flyX(400), duration: flyDuration() }}
			>
				{guess.weapon.damage}
			</span>

			<!-- Use time -->
			<span
				class="arrow-cell"
				class:arrow-up={guess.checks.use_time === CompareResult.Higher}
				class:arrow-down={guess.checks.use_time === CompareResult.Lower}
				style={guess.checks.use_time === CompareResult.Match ? guessCorrect : guessWrong}
				in:fly={{ x: flyX(320), duration: flyDuration() }}
			>
				{guess.weapon.use_time}
			</span>

			<!-- Rarity -->
			<span
				class="arrow-cell"
				class:arrow-up={guess.checks.rarity === CompareResult.Higher}
				class:arrow-down={guess.checks.rarity === CompareResult.Lower}
				style={(guess.checks.rarity === CompareResult.Match ? guessCorrect : guessWrong) +
					getRarityColor(guess.weapon.rarity as Rarity)}
				in:fly={{ x: flyX(240), duration: flyDuration() }}
			>
				{guess.weapon.rarity}
			</span>

			<!-- Operation -->
			<span
				style={guess.checks.operation ? guessCorrect : guessWrong}
				in:fly={{ x: flyX(160), duration: flyDuration() }}
			>
				{guess.weapon.operation}
			</span>

			<!-- Material -->
			<span
				style={guess.checks.material ? guessCorrect : guessWrong}
				in:fly={{ x: flyX(80), duration: flyDuration() }}
			>
				{guess.weapon.material}
			</span>

			<div
				style="flex-direction: column; {checkedObtained(guess.checks.obtained)}"
				in:fly={{ x: flyX(0), duration: flyDuration() }}
			>
				{#each guess.weapon.obtained as item (item)}
					<span>{item}</span>
				{/each}
			</div>
		</div>
	{/each}
{/snippet}

{#if !isMobile.current}
	<div class="container" in:fly={{ x: flyX(800), duration: flyDuration() }}>
		{@render body()}
	</div>
{:else}
	<div class="container">
		{@render body()}
	</div>
{/if}

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

	@media (max-width: 800px) {
		.container {
			max-width: 100%;
			overflow-y: auto;
		}

		.header {
			gap: 11px;
			padding: 5px 0;
		}

		.header span {
			font-size: 12px;
			width: 70px;
		}

		.row div,
		.row span {
			width: 70px;
			height: 70px;
		}
	}
</style>
