<script lang="ts">
	import { flip } from "svelte/animate";
	import { scale } from "svelte/transition";

    let { guesses, won } = $props();

    function shake(node: Element, { delay = 80, duration = 500, x = 6 } = {}) {
    return {
        delay,
        duration,
        css: (t: number) => {
            // t goes from 0 → 1 on intro
            const progress = Math.sin(t * Math.PI * 8) * (1 - t)

            return `transform: translateX(${progress * x}px);`
        }
    }
}
</script>

{#if guesses.length !== 0}
    <div in:scale>
        <div class="header">
            <span>Guessed Npcs</span>
        </div>

        {#each guesses as guess, i (guess.id)}
            <div 
                class="item" 
                style="background-color: {won && i === 0 ? 'var(--color-green)' : 'var(--color-red)'}"
                in:shake
                animate:flip={{ duration: 200 }}
            >
                <img src={'/npcs/' + guess.path} alt=""/>
                <span>{guess.name}</span>
            </div>
        {/each}
    </div>
{/if}

<style>
    .header {
		display: flex;
		margin: auto;
		margin-top: 10px;
		width: 200px;
        justify-content: center;
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

    .item {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 10px;
        width: 200px;
        height: 50px;
        background-color: var(--color-red);
        margin: 10px auto;
        padding: 10px;

        border-radius: 15px;
		border: 2px solid black;
    }
</style>