<script lang="ts">
	import Confetti from "svelte-confetti";
	import { scale } from "svelte/transition";

    let { failedGuesses } = $props();
</script>

<div class="container">
{#if failedGuesses < 6}
    <div style="padding-top: 40px;">
        <img
            src="/hangman/TerrariaNoose.png"
            alt=""
            style="position: absolute; margin-left: -20px;"
        />

        <img
            src="/hangman/Guide/Right_Arm.png"
            alt=""
            class="guide-part"
            style="opacity: {failedGuesses >= 1 ? '100%' : '30%'};"
        />

        <img
            src="/hangman/Guide/Torso.png"
            alt=""
            class="guide-part"
            style="opacity: {failedGuesses >= 2 ? '100%' : '30%'};"
        />

        <img
            src="/hangman/Guide/Left_Arm.png"
            alt=""
            class="guide-part"
            style="opacity: {failedGuesses >= 3 ? '100%' : '30%'};"
        />

        <img
            src="/hangman/Guide/Right_Leg.png"
            alt=""
            class="guide-part"
            style="opacity: {failedGuesses >= 4 ? '100%' : '30%'};"
        />

        <img
            src="/hangman/Guide/Left_Leg.png"
            alt=""
            class="guide-part"
            style="opacity: {failedGuesses >= 5 ? '100%' : '30%'};"
        />

        <img
            src="/hangman/Guide/Head.png"
            alt=""
            class="guide-part"
            style="opacity: {failedGuesses >= 6 ? '100%' : '30%'}"
        />
    </div>
{:else}
    <div class="guide-burst">
        <Confetti
            y={[-1, 1]}
            x={[-1, 1]}
            size={30}
            amount={10}
            duration={2000}
            noGravity
            colorArray={[
                'url(/hangman/Guide/Flesh1.png)',
                'url(/hangman/Guide/Flesh2.png)',
                'url(/hangman/Guide/Head.png)',
                'url(/hangman/Guide/Right_Arm.png)'
            ]}
        />
    </div>
    <img class="wof" src="/hangman/Wall_of_Flesh.gif" alt="Wall of Flesh" in:scale/>
{/if}
</div>

<style>
    .container {
        height: 220px;
    }

    .guide-part {
        position: absolute;
        width: 35px;
        height: auto;
        margin-top: 35px;
        margin-left: -12px;
    }

    .guide-burst {
		position: absolute;
		margin-top: 45px;
        margin-left: 48%;
		width: 200px;
		height: 200px;
		pointer-events: none;
    }

    .wof {
        height: 180px;
        padding-top: 20px;
    }
</style>