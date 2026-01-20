<script lang="ts">
	import { cubicInOut } from 'svelte/easing';
	import { slide } from 'svelte/transition';
	import Guide from './components/Guide.svelte';
	import Keyboard from './components/Keyboard.svelte';

    let { data } = $props();

    let failedGuesses: number = $state(0);
    let finished: boolean = $state(false);

    function onKeyPressed(letter: string) {
        console.log(letter);
    }
</script>

{#if !finished}
    <!-- out:slide={{ duration: 700, easing: cubicInOut }} -->
	<div class="title-box">
		<h2>Hangman</h2>
		<p>Guess letters one by one to figure out the enemy before hanging the Guide!</p>
	</div>
{/if}

<Guide {failedGuesses}/>

<button onclick={() => failedGuesses++}>Add</button>

<Keyboard {onKeyPressed} letters={""}/>

<style>
    .title-box {
		background-color: var(--color-backgroundblue);
		width: fit-content;
		text-align: center;
		margin: 30px auto 30px auto;
		padding: 0 15px;

		border-radius: 15px;
		border: thin solid black;
	}

	.title-box h2 {
		background-color: var(--color-lightblue);
		width: fit-content;
		margin: auto;
		margin-top: -30px;
		margin-bottom: 10px;
		padding: 10px 20px;

		border-radius: 15px;
		border: 2px solid black;
	}
</style>