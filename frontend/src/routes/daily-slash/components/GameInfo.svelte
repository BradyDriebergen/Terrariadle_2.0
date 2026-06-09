<script lang="ts">
	import { subscribeToPlayerCount } from '$lib/api/common';
	import type { WeaponPreview } from '$lib/types/daily-slash';
	import { colors, type Rarity } from '$lib/types/dailySlash';
	import { onMount } from 'svelte';

	let { 
		prevWeapon
	} : {
		prevWeapon: WeaponPreview | null;
	} = $props();

	let playerCount = $state<number>(0);

	// Streams live player count to the user
	onMount(() => {
        const cleanup = subscribeToPlayerCount('daily_slash', (count) => {
            playerCount = count;
        });

        return cleanup;
    });
</script>

<div>
	<p>{playerCount} people have guessed today's weapon</p>
	<p>Yesterday's weapon was:</p>

	<img
		style="border-color: {colors[prevWeapon?.rarity as Rarity]}"
		src={`/weapons/${prevWeapon?.path}`}
		alt="Previous weapon"
	/>
	<p style="color: {colors[prevWeapon?.rarity as Rarity]}">{prevWeapon?.name}</p>

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
