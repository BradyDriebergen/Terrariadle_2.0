<script lang="ts">
	import RemainingTime from "$lib/components/RemainingTime.svelte";
	import { ConvertPositionToString } from "$lib/utils/posToString";
	import { typewriter } from "$lib/utils/transitions";
	import { scale } from "svelte/transition";

    let { count, pos, npc, names } = $props();
</script>

<div class="winning-container" in:scale>
    <h1>Spot On!</h1>

    <div class="npc-housing">
        <div class="candle-light" style="bottom: {(npc.name === 'Traveling Merchant') ? "18" : "35"}px"></div>
        <img class="candle" style="bottom: {(npc.name === 'Traveling Merchant') ? '18' : '35'}px" src="/furniture/Candle.png" alt="" />
        <img src={"/furniture/" + npc.name.replace(" ", "") + "Table.png"} alt=""/>
        <img src={"/furniture/" + npc.name.replace(" ", "") + "Chair.png"} alt=""/>
        <img class="npc-image" src={'/npcs/' + npc.path} alt=''/>
    </div>

	{#if pos !== 0}
		<p transition:typewriter={{ speed: 1 }}>
			You were the {ConvertPositionToString(pos)} person to guess today's NPC!
		</p>
	{:else}
		<br />
	{/if}
	<p>{count} people guessed todays weapon</p>

    <div class="bonus-container">
        <h2 class="bonus-title">Bonus Round!</h2>
        <p>Out of the following names, what name can the {npc.name} have?</p>
        <div class="bonus-options">
            {#each names as name}
                <button class="bonus-button">{name}</button>
            {/each}
        </div>
    </div>

	<RemainingTime />
</div>

<style>
    .winning-container {
        border: 20px solid transparent;
        border-image: url('/daily-slash/borders/Wood.png');
        border-image-slice: 17;
        border-image-repeat: round;
        border-radius: 5px;

        background-image: url('/daily-slash/backgrounds/WoodWall.png');
        background-repeat: repeat;
        background-size: 20%;

        width: fit-content;
        margin: auto;
        margin-top: 10px;
        padding: 20px;
        padding-bottom: 0;
    }

    h1 {
        margin-top: 5px;
    }

    .npc-housing {
        position: relative;
    }

    .candle-light {
        position: absolute;
        width: 40px;
        height: 40px;
        background-color: rgb(255, 204, 62);
        border-radius: 50%;
        filter: blur(18px);
        left: 50%;
        margin-left: -59px;
        margin-bottom: 5px;
    }

    .candle {
        position: absolute;
        margin-left: 19px;
        margin-bottom: 5px;
    }

    .npc-image {
        padding-left: 20px;
    }

    .bonus-container {
        width: 350px;
        margin: 50px auto 20px auto;
        padding: 10px;
        background: rgb(31, 47, 82);

        border-radius: 15px;
		border: 2px solid black;
    }

    .bonus-title {
        background-color: var(--color-lightblue);
		width: fit-content;
		margin: auto;
		margin-top: -40px;
		margin-bottom: 10px;
		padding: 10px 20px;

		border-radius: 15px;
		border: 2px solid black;
    }

    .bonus-options {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        grid-auto-rows: 1fr;

        width: 100%;
        height: 100%;
        margin: 0 auto;
        gap: 5px;
    }

    .bonus-button {
        background-color: var(--color-button);
		position: relative;
        height: 50px;
		font-size: 16px;
		text-align: center;

		border-radius: 5px;
		border: thin solid black;
		transition: background-color 0.2s ease;
    }

    .bonus-button:hover {
        cursor: pointer;
		background-color: var(--color-lightblue);
    }
</style>