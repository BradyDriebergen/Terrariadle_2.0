<script lang="ts">
    import { backgrounds, borders, colors, type Rarity } from "$lib/types/dailySlash";
	import { onMount } from "svelte";
    let { weapon, userId } = $props();

    let pos = $state(0);
    let count = $state(0);
    let time = $state(0);

    onMount(async () => {
        const winningData = await fetch(`http://localhost:3000/api/daily-slash/winning-data/${userId}`);
        const winningDataJson = await winningData.json();

        const timeData = await fetch(`http://localhost:3000/api/remaining-time`);
        const timeDataJson = await timeData.json();

        time = timeDataJson.remaining;
        pos = winningDataJson.pos;
        count = winningDataJson.count;
    })

    $effect(() => {
    if (time <= 0) return;

    const id = setInterval(() => {
        if (time <= 1) {
            time = 0;
            clearInterval(id);
        } else {
            time -= 1;
        }
    }, 1000);

        return () => {
            clearInterval(id);
        };
    });

    let remainingTime = $derived.by(() => {
        if (time <= 0) return 'Refresh Now!';

        const hours = Math.floor(time / 3600);
        const minutes = Math.floor((time % 3600) / 60);
        const seconds = time % 60;

        const pad = (n: number) => n.toString().padStart(2, "0");

        return `${pad(hours)}:${pad(minutes)}:${pad(seconds)}`;
    })
</script>

<div 
    class="wrapper" 
    style="
        border-image-source: {borders[weapon.info.rarity as Rarity]}; 
        background: {backgrounds[weapon.info.rarity as Rarity]}"
>
    <h1>You Got It!</h1>
    <p>You were the {pos}th person to guess today's weapon!</p>
    <img 
		style="border-color: {colors[weapon.info.rarity as Rarity]}" 
		src={`/weapons/${weapon.info['image-path']}`} 
		alt='Previous weapon'
	/>
    <h3 class="weapon-name" style="color: {colors[weapon.info.rarity as Rarity]}">{weapon.name}</h3>
    <p>{count} people guessed todays weapon</p>
    <p>Next game starts in: {remainingTime}</p>
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
        width: fit-content;
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