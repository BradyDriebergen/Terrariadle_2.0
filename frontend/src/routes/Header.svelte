<script>
	import { page } from '$app/state';
	import { resolve } from '$app/paths';
	import { Spring } from 'svelte/motion';

	let isHome = $derived(page.url.pathname === '/')
	let size = new Spring(page.url.pathname === '/' ? 400 : 200);
	let translation = new Spring(page.url.pathname === '/' ? 0 : -37)

	$effect(() => {
		if (isHome) {
			size.target = 400;
			translation.target = 0;
		} else {
			size.target = 200;
			translation.target = -37;
		}
	});
</script>

<header>
	<div>
		<a href={resolve('/')}>
			<img
				class="game-logo"
				src="/logos/TerrariadleLogo.png"
				alt="Terrariadle"
				width={size.current}
				style={`margin-left: ${translation.current}%;`}
			/>
		</a>
		{#if !isHome}
			<a href={resolve('/daily-slash')}>Daily Slash</a>
			<a href={resolve('/connections')}>Connections</a>
			<a href={resolve('/guess-the-npc')}>Guess the NPC</a>
			<a href={resolve('/hangman')}>Hangman</a>
			<a href={resolve('/terratrivia')}>TerraTrivia</a>
		{/if}
	</div>
</header>

<style>
	div {
		margin: auto;
	}

	.game-logo {
		max-width: 80%;
		height: auto;
		transition: transform 0.2s ease;
	}

	.game-logo:hover {
		transform: scale(1.025);
	}
</style>
