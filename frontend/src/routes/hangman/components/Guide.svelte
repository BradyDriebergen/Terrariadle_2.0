<script lang="ts">
	import Confetti from "svelte-confetti";
	import { scale } from "svelte/transition";

    let { attempts } = $props();
</script>

<div class="container">
{#if attempts > 0}
    <div style="padding-top: 40px;">
        <img
            src="/hangman/TerrariaNoose.png"
            alt=""
            style="position: absolute; margin-left: -20px;"
        />

        <img
            src="/hangman/guide/Right_Arm.png"
            alt=""
            class="guide-part"
            style:opacity={attempts <= 5 ? 1 : 0.3}
        />

        <img
            src="/hangman/guide/Torso.png"
            alt=""
            class="guide-part"
            style:opacity={attempts <= 4 ? 1 : 0.3}
        />

        <img
            src="/hangman/guide/Left_Arm.png"
            alt=""
            class="guide-part"
            style:opacity={attempts <= 3 ? 1 : 0.3}
        />

        <img
            src="/hangman/guide/Right_Leg.png"
            alt=""
            class="guide-part"
            style:opacity={attempts <= 2 ? 1 : 0.3}
        />

        <img
            src="/hangman/guide/Left_Leg.png"
            alt=""
            class="guide-part"
            style:opacity={attempts <= 1 ? 1 : 0.3}
        />

        <img
            src="/hangman/guide/Head.png"
            alt=""
            class="guide-part"
            style:opacity={attempts <= 0 ? 1 : 0.3}
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
                'url(/hangman/guide/Flesh1.png)',
                'url(/hangman/guide/Flesh2.png)',
                'url(/hangman/guide/Head.png)',
                'url(/hangman/guide/Right_Arm.png)'
            ]}
        />
    </div>
    <img class="wof" src="/hangman/Wall_of_Flesh.gif" alt="Wall of Flesh" in:scale/>
{/if}
</div>
<p>Wrong guesses: {6 - attempts} / 6</p>

<style>
    .container {
        height: 190px;
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
    }
</style>