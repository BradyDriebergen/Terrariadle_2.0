<script>
	import { resolve } from '$app/paths';
</script>

<header>
	<a href={resolve('/')}>
		<img class="game-logo" src="/logos/TerrariadleLogo.png" alt="Terrariadle" />
	</a>
</header>

<style>
	.game-logo {
		width: 400px;
		max-width: 80%;
		height: auto;
		transition: transform 0.2s ease;
	}

	.game-logo:hover {
		transform: scale(1.025);
	}
</style>
