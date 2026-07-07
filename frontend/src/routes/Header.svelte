<script>
	import { page } from '$app/state';
	import { resolve } from '$app/paths';
	import { Spring } from 'svelte/motion';
	import { crossfade, scale } from 'svelte/transition';
	import { cubicOut } from 'svelte/easing';

	let isHome = $derived(page.url.pathname === '/')
	let headerHeight = new Spring(page.url.pathname === '/' ? 140 : 100)

	$inspect(page.url.pathname)

	let dailySlashLogoLink = '/logos/daily_slash_logo.png';
	let connectionsLogoLink = '/logos/connections_logo.png';
	let guessTheNpcLogoLink = '/logos/guess_the_npc_logo.png';
	let hangmanLogoLink = '/logos/hangman_logo.png';
	let terraTriviaLogoLink = '/logos/terratrivia_logo.png';

	const [send, receive] = crossfade({
		duration: 400,
		easing: cubicOut
	});

	$effect(() => {
		if (isHome) {
			headerHeight.target = 140;
		} else {
			headerHeight.target = 100;
		}
	});
</script>

<header style={`height: ${headerHeight.current}px;`}>
	{#if isHome}
		<a
			href={resolve('/')}
			class="logo-slot centered"
			in:receive={{ key: 'block' }}
			out:send={{ key: 'block' }}
		>
			<img class="game-logo" src="/logos/TerrariadleLogo.png" alt="Terrariadle"/>
		</a>
	{:else}
		<a
			href={resolve('/')}
			class="logo-slot corner"
			in:receive={{ key: 'block' }}
			out:send={{ key: 'block' }}
		>
			<img class="game-logo" src="/logos/TerrariadleLogo.png" alt="Terrariadle" />
		</a>
		<div class="links" in:scale out:scale>
			<a 
				class:Selected={page.url.pathname === '/daily-slash'} 
				href={resolve('/daily-slash')}
			>
				<img src={dailySlashLogoLink} alt="Daily slash" />
			</a>
			<a 
				class:Selected={page.url.pathname === '/connections'} 
				href={resolve('/connections')}
			>
				<img src={connectionsLogoLink} alt="Connections" />
			</a>
			<a 
				class:Selected={page.url.pathname === '/guess-the-npc'} 
				href={resolve('/guess-the-npc')}
			>
				<img src={guessTheNpcLogoLink} alt="Guess the NPC" />
			</a>
			<a 
				class:Selected={page.url.pathname === '/hangman'} 
				href={resolve('/hangman')}
			>
				<img src={hangmanLogoLink} alt="Hangman" />
			</a>
			<a 
				class:Selected={page.url.pathname === '/terratrivia'} 
				href={resolve('/terratrivia')}
			>
				<img src={terraTriviaLogoLink} alt="TerraTrivia" />
			</a>
		</div>
	{/if}
</header>

<style>
	header {
		position: relative;
		display: grid;
		grid-template-columns: 1fr auto 1fr;
		align-items: center;
	}

	.logo-slot {
		position: absolute;
		grid-row: 1;
		grid-column: 1;
		display: block;
	}

	.centered {
		justify-self: center;
		width: 400px;
	}

	.corner {
		justify-self: start;
		width: 200px;
	}

	.game-logo {
		margin: auto;
		max-width: 80%;
		display: block;
		transition: transform 0.2s ease;
	}

	.game-logo:hover {
		transform: scale(1.025);
	}

	.links {
		grid-column: 2;
		grid-row: 1;
		justify-self: center;
		display: flex;
		gap: 12px;
	}

	.links a {
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: var(--color-backgroundblue);
		border-radius: 8px;
		border: 2px solid black;
		padding: 6px;
	}

	.links a:hover {
		background: var(--color-lightblue);
	}

	.links a.Selected,
	.links a.Selected:hover {
		background: rgba(139, 31, 31, 0.862);
	}

	.links img {
		width: 40px;
	}
</style>