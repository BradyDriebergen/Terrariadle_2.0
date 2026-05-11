<!--
  A simple Svelte/TypeScript implementation of a 7 Little Words‑style puzzle.

  Each puzzle consists of seven clues and twenty groups of letters.  Players
  solve the clues by selecting letter groups to assemble the answer words.
  When a correct word is submitted, the corresponding letter groups disappear
  from the board and the clue is marked as solved.  Incorrect guesses are
  rejected and the selection is cleared.

  This component is meant as a standalone page that you can drop into a
  Svelte project.  It uses Svelte’s reactivity and simple styling to
  implement the core mechanics of the game.  Feel free to adapt the
  puzzle data or improve the UI as desired.
-->

<script lang="ts">
  // Define the structure of a single clue.  Each clue has the text of the
  // clue itself, the correct answer and a flag indicating whether it has
  // already been solved.
  interface Clue {
    clue: string;
    answer: string;
    solved: boolean;
  }

  // Define the structure of a letter group.  Each group has a unique ID,
  // the letters it contains and a flag indicating whether it has been used
  // to solve a word.  We also track whether it is currently selected.
  interface LetterGroup {
    id: number;
    letters: string;
    used: boolean;
  }

  // A small utility function to shuffle the letter groups.  This is
  // optional—if you’d like to present the groups in a fixed order, simply
  // remove the call to shuffle() below.
  function shuffle<T>(array: T[]): T[] {
    // Use a copy so we don't mutate the original array.
    const arr = array.slice();
    for (let i = arr.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [arr[i], arr[j]] = [arr[j], arr[i]];
    }
    return arr;
  }

  // Define the puzzle data.  The clues and answers here come from a sample
  // 7 Little Words puzzle.  You can replace these with any seven clue/
  // answer pairs and corresponding twenty letter groups.
  let clues: Clue[] = [
    { clue: 'Paying close attention', answer: 'ATTENTIVE', solved: false },
    { clue: 'Most bothersome', answer: 'PESKIEST', solved: false },
    { clue: 'Gain knowledge', answer: 'LEARN', solved: false },
    { clue: 'Big cats (NFL team)', answer: 'PANTHERS', solved: false },
    { clue: 'Expresses greetings', answer: 'REGARDS', solved: false },
    { clue: 'Asian desert', answer: 'GOBI', solved: false },
    { clue: 'Lukewarm', answer: 'TEPID', solved: false }
  ];

  // Build the twenty letter groups used in the sample puzzle.  Each group
  // corresponds to part of one of the answers above.  For example, the
  // answer "ATTENTIVE" can be made from "AT", "TE", "NTI" and "VE".  The
  // grouping follows the original game’s use of 2–4 letter segments.
  let letterGroups: LetterGroup[] = shuffle([
    { id: 1, letters: 'AT', used: false },
    { id: 2, letters: 'TE', used: false },
    { id: 3, letters: 'NTI', used: false },
    { id: 4, letters: 'VE', used: false },
    { id: 5, letters: 'PE', used: false },
    { id: 6, letters: 'SK', used: false },
    { id: 7, letters: 'IE', used: false },
    { id: 8, letters: 'ST', used: false },
    { id: 9, letters: 'LE', used: false },
    { id: 10, letters: 'ARN', used: false },
    { id: 11, letters: 'PAN', used: false },
    { id: 12, letters: 'THE', used: false },
    { id: 13, letters: 'RS', used: false },
    { id: 14, letters: 'RE', used: false },
    { id: 15, letters: 'GAR', used: false },
    { id: 16, letters: 'DS', used: false },
    { id: 17, letters: 'GO', used: false },
    { id: 18, letters: 'BI', used: false },
    { id: 19, letters: 'TEP', used: false },
    { id: 20, letters: 'ID', used: false }
  ]);

  // Track which groups have been selected for the current guess.  When the
  // selection changes, currentGuess is recalculated to reflect the
  // concatenation of the selected groups’ letters.
  let selectedGroups: LetterGroup[] = [];
  let currentGuess: string = '';
  let message: string = '';

  // When a letter group is clicked, add it to the selection.  We ignore
  // groups that have already been used to solve a word or that are already
  // in the current selection.
  function selectGroup(group: LetterGroup) {
    if (group.used) return;
    // Prevent selecting the same group twice in one guess.
    if (selectedGroups.includes(group)) return;
    selectedGroups = [...selectedGroups, group];
    updateCurrentGuess();
    message = '';
  }

  // Remove the last selected letter group from the guess.  This allows the
  // player to backtrack without clearing the entire selection.
  function undo() {
    if (selectedGroups.length > 0) {
      selectedGroups = selectedGroups.slice(0, -1);
      updateCurrentGuess();
      message = '';
    }
  }

  // Clear all selected letter groups.
  function clearSelection() {
    selectedGroups = [];
    updateCurrentGuess();
    message = '';
  }

  // Update the currentGuess string based on the selected letter groups.
  function updateCurrentGuess() {
    currentGuess = selectedGroups.map((g) => g.letters).join('');
  }

  // Submit the current guess.  If the guess matches one of the unsolved
  // answers exactly, mark that clue solved and mark the used letter groups
  // as used.  Otherwise inform the player that the guess is incorrect.
  function submitGuess() {
    if (!currentGuess) return;
    // Normalize to uppercase for comparison.
    const guess = currentGuess.toUpperCase();
    const clueIndex = clues.findIndex(
      (c) => !c.solved && c.answer.toUpperCase() === guess
    );
    if (clueIndex !== -1) {
      // Correct guess
      clues[clueIndex].solved = true;
      // Mark each selected letter group as used.
      for (const g of selectedGroups) {
        g.used = true;
      }
      clearSelection();
      message = 'Correct!';
    } else {
      // Incorrect guess
      message = 'Incorrect guess. Try again.';
      clearSelection();
    }
  }
</script>

<div class="game-container">
  <div class="clues">
    {#each clues as clue, index}
      <div class="clue-item">
        <div class="clue-header">
          <span class="clue-number">{index + 1}.</span>
          <span class="clue-text">{clue.clue} ({clue.answer.length})</span>
        </div>
        <div class="clue-answer">
          {#if clue.solved}
            <strong>{clue.answer}</strong>
          {:else}
            <!-- Render a series of underscores equal in length to the answer to hint at the word length. -->
            {Array.from({ length: clue.answer.length }).map(() => '_').join(' ')}
          {/if}
        </div>
      </div>
    {/each}
  </div>

  <div class="selected">
    <h3>Current Guess:</h3>
    <div class="guess" aria-live="polite">{currentGuess}</div>
    <div class="selected-buttons">
      <!-- Show the currently selected groups.  Clicking on the last selected group acts as an undo operation. -->
      {#each selectedGroups as group, idx (group.id)}
        <button
          class="selected-group-button"
          on:click={() => idx === selectedGroups.length - 1 && undo()}
          title="Click to undo last selection"
        >
          {group.letters}
        </button>
      {/each}
    </div>
    <div class="actions">
      <button on:click={submitGuess} disabled={selectedGroups.length === 0}
        >Submit</button
      >
      <button on:click={undo} disabled={selectedGroups.length === 0}>Undo</button>
      <button on:click={clearSelection} disabled={selectedGroups.length === 0}
        >Clear</button
      >
    </div>
    {#if message}
      <p class="message">{message}</p>
    {/if}
  </div>

  <div class="letter-groups">
    {#each letterGroups as group (group.id)}
      <button
        class="letter-group-button"
        disabled={group.used}
        on:click={() => selectGroup(group)}
        aria-disabled={group.used}
      >
        {group.letters}
      </button>
    {/each}
  </div>
</div>

<style>
  /* Container layout */
  .game-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    max-width: 640px;
    margin: 1rem auto;
    font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI',
      Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    padding: 0 1rem;
  }

  /* Clue list styling */
  .clues {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
  }

  .clue-item {
    display: flex;
    flex-direction: column;
  }

  .clue-header {
    display: flex;
    gap: 0.5rem;
    align-items: baseline;
  }

  .clue-number {
    font-weight: bold;
  }

  .clue-text {
    font-weight: 500;
  }

  .clue-answer {
    margin-left: 1.5rem;
    font-size: 1.2rem;
    font-family: monospace;
  }

  /* Selected guess area */
  .selected {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
    border-top: 1px solid #ddd;
    padding-top: 1rem;
  }

  .guess {
    font-size: 1.5rem;
    min-height: 2rem;
  }

  .selected-buttons {
    display: flex;
    gap: 0.5rem;
    flex-wrap: wrap;
  }

  .selected-group-button {
    padding: 0.25rem 0.5rem;
    font-size: 1rem;
    border: 1px solid #888;
    border-radius: 4px;
    background-color: #e8f4fd;
    cursor: pointer;
  }

  .selected-group-button:hover {
    background-color: #d7eafb;
  }

  .actions {
    display: flex;
    gap: 0.5rem;
  }

  .actions button {
    padding: 0.4rem 0.8rem;
    font-size: 1rem;
    border: 1px solid #ccc;
    border-radius: 4px;
    background-color: #f5f5f5;
    cursor: pointer;
  }

  .actions button[disabled],
  .letter-group-button[disabled] {
    opacity: 0.5;
    cursor: not-allowed;
  }

  .actions button:hover:not([disabled]) {
    background-color: #e0e0e0;
  }

  .message {
    margin-top: 0.5rem;
    font-weight: bold;
    color: #b00020;
  }

  /* Letter groups styling */
  .letter-groups {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
    justify-content: center;
    border-top: 1px solid #ddd;
    padding-top: 1rem;
  }

  .letter-group-button {
    padding: 0.5rem 0.8rem;
    font-size: 1rem;
    border: 1px solid #bbb;
    border-radius: 4px;
    background-color: #fff;
    cursor: pointer;
    min-width: 3.5rem;
    text-align: center;
  }

  .letter-group-button:hover:not([disabled]) {
    background-color: #f0f0f0;
  }
</style>