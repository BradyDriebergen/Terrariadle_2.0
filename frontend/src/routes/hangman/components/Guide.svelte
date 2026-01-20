<script>
    let { failedGuesses } = $props();
</script>

<div style="margin-bottom: 190px;">
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

<style>
    .guide-part {
        position: absolute;
        width: 35px;
        height: auto;
        margin-top: 35px;
        margin-left: -12px;
    }
</style>