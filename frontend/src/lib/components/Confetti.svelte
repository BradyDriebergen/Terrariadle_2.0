<script lang="ts">
	import { Confetti } from 'svelte-confetti';

	let { finished } = $props();

	let showConfetti = $state([false, false, false]);

	$effect(() => {
		if (!finished) {
			showConfetti = [false, false, false];
			return;
		}

		const t1 = setTimeout(() => (showConfetti[0] = true), 500);
		const t2 = setTimeout(() => (showConfetti[1] = true), 1000);
		const t3 = setTimeout(() => (showConfetti[2] = true), 1500);

		return () => {
			clearTimeout(t1);
			clearTimeout(t2);
			clearTimeout(t3);
		};
	});
</script>

{#if finished}
	{#if showConfetti[0]}
		<div
			class="confetti"
			style="top: {Math.random() * 10 + 20}%; left: {Math.random() * 10 + 20}%;"
		>
			<Confetti
				y={[-1, 1]}
				x={[-1, 1]}
				size={20}
				delay={[50, 100]}
				colorArray={[
					'url(/confetti/Confetti1.png)',
					'url(/confetti/Confetti2.png)',
					'url(/confetti/Confetti3.png)',
					'url(/confetti/Confetti4.png)'
				]}
			/>
		</div>
	{/if}
	{#if showConfetti[1]}
		<div
			class="confetti"
			style="top: {Math.random() * 20 + 60}%; left: {Math.random() * 50 + 30}%;"
		>
			<Confetti
				y={[-1, 1]}
				x={[-1, 1]}
				size={20}
				delay={[50, 100]}
				colorArray={[
					'url(/confetti/Confetti1.png)',
					'url(/confetti/Confetti2.png)',
					'url(/confetti/Confetti3.png)',
					'url(/confetti/Confetti4.png)'
				]}
			/>
		</div>
	{/if}
	{#if showConfetti[2]}
		<div
			class="confetti"
			style="top: {Math.random() * 10 + 20}%; left: {Math.random() * 10 + 80}%;"
		>
			<Confetti
				y={[-1, 1]}
				x={[-1, 1]}
				size={20}
				delay={[50, 100]}
				colorArray={[
					'url(/confetti/Confetti1.png)',
					'url(/confetti/Confetti2.png)',
					'url(/confetti/Confetti3.png)',
					'url(/confetti/Confetti4.png)'
				]}
			/>
		</div>
	{/if}
{/if}

<style>
	.confetti {
		position: fixed;
		top: 20px;
		right: 20px;
		width: 200px;
		height: 200px;
		pointer-events: none;
	}
</style>
