<script>
	import { page } from '$app/state';
	import { Spring } from 'svelte/motion';

	let size = new Spring(page.url.pathname === '/' ? 400 : 200);

	$effect(() => {
		if (page.url.pathname === '/') {
			size.target = 400;
		} else {
			size.target = 200;
		}
	});
</script>

<header>
	<div>
		<a href="/">
			<img 
				class="game-logo" 
				src="/logos/TerrariadleLogo.png" 
				alt="Terrariadle" 
				width={size.current}
			/>
		</a>
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
