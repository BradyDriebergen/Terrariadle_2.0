<script lang="ts">
	import RemainingTime from "$lib/components/RemainingTime.svelte";
    import { backgrounds, borders, colors, type Rarity } from "$lib/types/dailySlash";
	import { ConvertPositionToString } from "$lib/utils/posToString";
	import { typewriter } from "$lib/utils/transitions";
	import { onMount } from "svelte";
	import { scale } from "svelte/transition";
    let { weapon, userId } = $props();

    let pos = $state(0);
    let count = $state(0);

    onMount(async () => {
        const winningData = await fetch(`http://localhost:3000/api/daily-slash/winning-data/${userId}`);
        const winningDataJson = await winningData.json();

        pos = winningDataJson.pos;
        count = winningDataJson.count;
    })
</script>

<div 
    class="wrapper" 
    style="
        border-image-source: {borders[weapon.info.rarity as Rarity]}; 
        background: {backgrounds[weapon.info.rarity as Rarity]}"
    in:scale
>
    <h1>You Got It!</h1>
    {#if pos !== 0}
        <p transition:typewriter={{ speed: 1 }}>You were the {ConvertPositionToString(pos)} person to guess today's weapon!</p>
    {:else}
        <br />
    {/if}
    <img 
		style="border-color: {colors[weapon.info.rarity as Rarity]}" 
		src={`/weapons/${weapon.info['image-path']}`} 
		alt='Previous weapon'
        in:scale
	/>
    <h3 class="weapon-name" style="color: {colors[weapon.info.rarity as Rarity]}">{weapon.name}</h3>
    <p>{count} people guessed todays weapon</p>
    <RemainingTime />
</div>

<style>
    .wrapper {
        border: 20px solid transparent;
        border-radius: 5px;
        border-image-slice: 17;
        border-image-repeat: round;

        background-repeat: repeat;
        background-size: 20%;

        display: flex;
        justify-content: center;
        flex-direction: column;
        align-items: center;
        text-align: center;
        padding: 0 20px;
        width: 260px;
        margin: 15px auto;
        box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
    }

    img {
        background-color: var(--color-button);
		padding: 18px;
		width: 45px;
		height: 45px;
		object-fit: contain;

		border-radius: 15px;
		border: 2px solid;
    }

    p {
        margin-top: 0;
    }
</style>