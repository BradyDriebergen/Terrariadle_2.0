<script lang="ts">

  let { onKeyPressed, letters } = $props();

  let selectedLetters: string[] = $state([]);

  const rows: string[][] = [
    ['Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P'],
    ['A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L'],
    ['Z', 'X', 'C', 'V', 'B', 'N', 'M']
  ];

  function handleLetterClick(letter: string) {
    if (!selectedLetters.includes(letter)) {
      selectedLetters.push(letter);
    }
  }
</script>

<div class="keyboard">
    {#each rows as row, rowIndex (rowIndex)}
        <div class="keyboard-row">
            {#each row as letter (letter)}
                <button
                    class="keyboard-key"
                    class:correct={letters.includes(letter)}
                    disabled={selectedLetters.includes(letter)}
                    onclick={() => {
                        handleLetterClick(letter);
                        onKeyPressed(letter);
                    }}
                >
                    {letter}
                </button>
            {/each}
        </div>
    {/each}
</div>

<style>
    .keyboard {
        display: flex;
        flex-direction: column;
        gap: 8px;
        width: 100%;
    }

    .keyboard-row {
        display: flex;
        justify-content: center;
        gap: 5px;
    }

    .keyboard-key {
        display: flex;
        align-items: center;
        justify-content: center;
        background: var(--color-backgroundblue);
        width: 50px;
        height: 62px;
        font-size: 17px;
        border: 2px solid #000000;
        border-radius: 5px;
        transition: background-color 0.2s ease;
    }

    .keyboard-key:hover:not(:disabled) {
        background: var(--color-lightblue);
        cursor: pointer;
    }

    .keyboard-key:disabled {
        background: var(--color-red);
        cursor: not-allowed;
    }

    .keyboard-key:disabled.correct {
        background: var(--color-green);
    }
</style>
