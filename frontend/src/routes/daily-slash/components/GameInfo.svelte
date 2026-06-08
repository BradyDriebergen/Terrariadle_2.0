<script lang="ts">
	import { colors, type Rarity } from '$lib/types/dailySlash';

	let { previousWeapon } = $props();
</script>

<div>
	<p>5 people have guessed today's weapon</p>
	<p>Yesterday's weapon was:</p>

	<img
		style="border-color: {colors[previousWeapon.rarity as Rarity]}"
		src={`/weapons/${previousWeapon.path}`}
		alt="Previous weapon"
	/>
	<p style="color: {colors[previousWeapon.rarity as Rarity]}">{previousWeapon.name}</p>

	<p>Guess any weapon to begin.</p>
</div>

<style>
	img {
		background-color: var(--color-button);
		padding: 18px;
		width: 45px;
		height: 45px;
		object-fit: contain;

		border-radius: 15px;
		border: 2px solid;
	}
</style>
