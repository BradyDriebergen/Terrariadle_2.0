<script lang="ts">
	import Dropdown from "$lib/components/Dropdown.svelte";
	import { getContext } from "svelte";

    let weaponList = getContext('weapons');
</script>

<div class="guess-panel">
    <h2>Guess Today's Weapon</h2>
    <div class="hint-buttons">
        <button>Mode Obtained</button>
        <button>Weapon Type</button>
        <button>Image Clue</button>
    </div>

    <div class="dropdown">
        <Dropdown 
            selectItem={(weaponid: number) => {console.log(weaponid)}}
            {weaponList}
        />
    </div>
</div>

<style>
    .guess-panel {
        background-color: var(--color-backgroundblue);
        width: fit-content;
        text-align: center;
        margin: auto;
        margin-top: 40px;
        padding: 15px;
        opacity: 90%;

        border-radius: 15px;
        border: thin solid black;
    }

    h2 {
        background-color: var(--color-lightblue);
        width: fit-content;
        margin: auto;
        margin-top: -40px;
        margin-bottom: 10px;
        padding: 10px 20px;

        border-radius: 15px;
        border: 2px solid black;
    }

    .hint-buttons {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 10px;
    }

    .hint-buttons button {
        background-color: var(--color-button);
        width: 100px;
        height: 100px;
        font-size: 16px;
        cursor: pointer;
        text-align: center;

        border-radius: 15px;
        border: thin solid black;
        transition: background-color 0.2s ease;
    }

    .hint-buttons button:hover {
        background-color: var(--color-lightblue);
    }

    .dropdown {
        margin-top: 15px;
        margin-bottom: 5px;
    }
</style>
