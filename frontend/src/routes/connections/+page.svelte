<script lang="ts">
    import { send, receive } from '$lib/utils/transitions';
	import { flip } from 'svelte/animate';

    let data = $state([
        "The Horseman's Blade",
        "Do You Want to Slay a Snowman?",
        "book",
        "key",
        "pen",
        "mug",
        "sock",
        "coin",
        "lamp",
        "shoe",
        "towel",
        "spoon",
        "box",
        "clock",
        "desk",
        "ball"
    ]);

    let start

    let guesses: string[] = $state([]);
    let selectedStrings: string[] = $state([]);

    function selectCategory(cat: string) {
        if (selectedStrings.includes(cat)) {
            selectedStrings = selectedStrings.filter(s => s !== cat);
        } else {
            selectedStrings.push(cat);
        }
    }

    function guess() {
        data = data.filter(s => !selectedStrings.includes(s))
        guesses.push(...selectedStrings);
        selectedStrings = [];
    }
</script>

<div class="grid">
    {#each guesses as category (category)}
        <button class="category" in:receive={{ key: category }}>
            <span>{category}</span>
        </button>
    {/each}

    {#each data as category (category)}
        <button
            type="button"
            class="category"
            class:Selected={selectedStrings.includes(category)}
            onclick={() => selectCategory(category)}
            disabled={selectedStrings.length === 4 && !selectedStrings.includes(category)}
            out:send={{ key: category }}
            animate:flip={{ duration: 220, easing: t => t }} 
        >
            <span>{category}</span>
        </button>
    {/each}
</div>
<br/>
<button disabled={selectedStrings.length !== 4} onclick={guess}>guess</button>

<style>
    .grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        grid-auto-rows: 1fr;

        width: 500px;
        height: 400px;
        margin: 0 auto;
        gap: 10px;
    }

    .category {
        background-color: var(--color-button);
        width: 100%;
        height: 100%;
        display: grid;
        place-items: center;

        border-radius: 15px;
        border: 2px solid black;
        
        padding: 0;
        font-family: inherit;
        font-size: inherit;
        transition: background-color 0.1s ease;
    }

    .category:hover,
    .category.Selected:hover {
        background-color: var(--color-lightblue);
        cursor: pointer;
    }

    .category:disabled:hover {
        background-color: var(--color-button);
        cursor: default;
    }

    .category.Selected,
    .category.Selected:hover {
        background-color: rgba(0, 255, 255, 0.8);
    }

    .category span {
        padding: 5px;
        font-size: 20px;
    }
</style>
