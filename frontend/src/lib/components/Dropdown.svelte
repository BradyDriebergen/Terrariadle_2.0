<script lang="ts">
    import myImage from '$lib/assets/Abigails_Flower.png';

    const items: string[] = ['option a', 'option b', 'option c'];

    let input = $state('');
</script>

<div class="wrapper">
    <input 
        type="text" 
        placeholder="Type any weapon to guess..." 
        bind:value={input}
        onkeydown={e => {
            if (e.key !== 'Enter') return;

            console.log(e.currentTarget);
        }}
    />

    {#if input}
        <div class="dropdown" >
            {#each items as item}
                <button class="item" onclick={() => input = item}>
                    <img src={myImage} alt={item}/>
                    <span>{item}</span>
                </button>
            {/each}
        </div>
    {/if}
</div>

<style>
    .wrapper {
        position: relative;
        display: inline-block;
    }

    input {
        background-color: var(--color-button);
        padding: 10px;
        width: 220px;
        height: 20px;
        font-size: 15px;
        text-align: left;

        border-radius: 5px;
        border: thin solid black;
        text-shadow: none;
        text-shadow: 0px 0px 3px #000, 0px 0px 2px #000;
    }

    .dropdown {
        position: absolute;
        width: 240px;
        top: 100%;
        left: 0;
        background: var(--color-button);
        max-height: 240px;
        overflow-y: auto;
        border-radius: 8px;
        z-index: 50;
        border: thin solid black;
    }
    .item {
        display: flex;
        gap: 8px;
        padding: 6px;
        align-items: center;
        cursor: pointer;
        border-radius: 6px;

        background: none;
        width: 100%;
        border: none;
    }
    .item:hover{
        background: var(--color-lightblue);
    }
    .item span {
        font-size: 14px;
    }

    .item img {
        width: 35px;
        height: 35px;
        object-fit: contain;
    }
    
</style>