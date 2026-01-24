<script lang="ts">
	import RemainingTime from "$lib/components/RemainingTime.svelte";

    let name = $state("Duke Fishron");
    let count = $state(5);
    let attempts = $state(1);
</script>

<div 
    class="container" 
    class:fail={attempts == 0}
    class:success={attempts > 0}
>
    {#if attempts == 6}
        <h2 class="title">Outstanding!</h2>
        <span class="subtitle">You guessed the enemy perfectly!</span>
        <br />
        <p>The enemy of the day was:</p>
    {:else if attempts > 0}
        <h2 class="title">Bingo!</h2>
        <p>The enemy of the day was:</p>
    {:else}
        <h2 class="title">Game Over!</h2>
        <p>The correct enemy was:</p>
    {/if}

    <h1>{name}</h1>

    <div class="enemy-image">
        <img src="/enemies/Duke Fishron.png" alt=""/>
    </div>

    <p>{count} people guessed todays enemy</p>
    <RemainingTime />
</div>

<style>
    .container {
		width: fit-content;
        min-width: 300px;
		margin: auto;
		margin-bottom: 20px;
		padding: 0 20px;
	}

    .fail {
        border: 20px solid transparent;
        border-image: url('/daily-slash/borders/HellstoneBrick.png');
        border-image-slice: 17;
        border-image-repeat: round;
        border-radius: 5px;

        background-image: url('/daily-slash/backgrounds/HellstoneWall.png');
        background-repeat: repeat;
        background-size: 20%;
    }

    .success {
        background-color: var(--color-button);
        border-radius: 10px;
		border: 2px solid black;
    }

    .title {
        font-size: 40px;
        margin: 20px auto 0 auto;
    }

    .enemy-image {
        border: 20px solid transparent;
        border-image: url('/hangman/EnemyBorder.png');
        border-image-slice: 16;
        border-image-repeat: round;
        
        background-image: url('/hangman/EnemyBackground.png');
        background-repeat: repeat;
        background-size: 70%;
        background-clip: padding-box;

        padding: 10px;
        margin: auto;
        width: fit-content;
    }

    .enemy-image img {
        max-width: 100px;
        max-height: 100px;
        width: 100px;
        height: auto;
        object-fit: contain;
    }
</style>