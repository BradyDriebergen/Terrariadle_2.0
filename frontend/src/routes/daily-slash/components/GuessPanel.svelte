<script lang="ts">
    import { onDestroy, tick } from 'svelte';
    import { debounceWithHooks } from '$lib/utils/debounce';
    import { searchApi } from '$lib/api/search';

    const GAME_MODE = 'daily-slash';

    interface SearchResult {
        weaponId: number;
        name: string;
        path: string;
    }

    // Search variables
    let query = '';
    let results: SearchResult[] = [];
    let loading = false;
    let error: string | null = null;
    let controller: AbortController | null = null;

    // Dropdown & keyboard state
    let showDropdown = false;
    let highlighted = -1;
    let dropdownEl: HTMLDivElement | null = null;

    // Calls search API
    async function doSearch(value: string) {
        try {
            controller?.abort();
            controller = new AbortController();
            loading = true;
            results = await searchApi(GAME_MODE, value, controller.signal);
            console.log(results)
            error = null;
            highlighted = 0;
        } catch (e: any) {
            if (e.name !== 'AbortError') error = e.message;
        } finally {
            loading = false;
        }
    }
    const debouncedSearch = debounceWithHooks(doSearch, 80, {
        onStart: () => (loading = true),
        onCancel: () => (loading = false),
    });


    // Text input event
    function handleInput(e: Event) {
        query = (e.target as HTMLInputElement).value;
        showDropdown = true;
        debouncedSearch(query);
    }

    // Handles keys pressed
    function handleKeydown(e: KeyboardEvent) {
        if (!showDropdown) return;
        if (e.key === 'ArrowDown') {
            highlighted = Math.min(highlighted + 1, results.length - 1);
            e.preventDefault();
        } else if (e.key === 'ArrowUp') {
            highlighted = Math.max(highlighted - 1, 0);
            e.preventDefault();
        } else if (e.key === 'Enter') {
            if (highlighted >= 0 && results[highlighted]) selectItem(results[highlighted]);
            e.preventDefault();
        }
    }

    // Method to search once an option has been selected
    function selectItem(item: SearchResult) {
        query = '';
        results = [];
        highlighted = 0;

        // Call parent to check guess
    }

    $: if (dropdownEl && highlighted >= 0) {
        // wait for DOM to update so the highlighted item exists
        tick().then(() => {
            if (!dropdownEl) return;
            const child = dropdownEl.children[highlighted] as HTMLElement | undefined;
            if (child) child.scrollIntoView({ behavior: 'smooth', block: 'nearest', inline: 'nearest' });
        });
    }

    // Cancels hanging requests on page leave
    onDestroy(() => {
        debouncedSearch.cancel();
        controller?.abort();
    });
</script>

<div class="guess-panel">
    <h2>Guess Today's Weapon</h2>
    <div class="hint-buttons">
        <button>Mode Obtained</button>
        <button>Weapon Type</button>
        <button>Image Clue</button>
    </div>
    <div class="input-section">
        <input 
            type="text" 
            placeholder="Type any weapon to guess..." 
            on:input={handleInput} 
            on:keydown={handleKeydown}
            bind:value={query}
        />
    </div>

    <!-- Custom dropdown -->
    {#if results.length}
        <div class="dropdown" role="listbox" aria-label="Search results" bind:this={dropdownEl}>
            {#each results as item, i}
                <div
                    class="dropdown-item"
                    class:selected={i === highlighted}
                    on:mousedown|preventDefault={() => selectItem(item)}
                    on:mouseenter={() => (highlighted = i)}
                    role="option"
                    aria-selected={i === highlighted}
                    tabindex="-1"
                >
                    <img src={`/weapons/${item.path}`} alt={item.name} />
                    <span>{item.name}</span>
                </div>
            {/each}
        </div>
    {:else if !results.length && query && !loading}
        <div class="dropdown" role="listbox" aria-label="Search results">
            <div class="dropdown-item">
                <span>No results...</span>
            </div>
        </div>
    {/if}
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

    .input-section {
        display: flex;
        justify-content: center;
        align-items: center;
        margin: 20px auto 5px auto;
    }

    .input-section input {
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

    /* Dropdown styles */
    .dropdown {
        background: var(--color-button);
        position: absolute;
        width: 245px;
        margin-left: 13px;
        max-height: 240px;
        overflow-y: auto;
        border-radius: 8px;
        box-shadow: 0 6px 18px rgba(0,0,0,0.3);
        z-index: 50;

        border: thin solid black;
    }

    .dropdown-item {
        display: flex;
        gap: 8px;
        padding: 6px;
        align-items: center;
        cursor: pointer;
        border-radius: 6px;
    }
    .dropdown-item:hover,
    .dropdown-item.selected {
        background: var(--color-lightblue);
    }

    .dropdown-item img {
        width: 35px;
        height: 35px;
        object-fit: contain;
    }
    .dropdown-item span {
        font-size: 14px;
    }
</style>
