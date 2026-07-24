<script>
	import { page } from '$app/state';
	import { resolve } from '$app/paths';
	import { Spring } from 'svelte/motion';
	import { crossfade, scale } from 'svelte/transition';
	import { cubicOut } from 'svelte/easing';
	import { session } from '$lib/store/session.svelte';

	let isHome = $derived(page.url.pathname === '/' || page.url.pathname === '/about');
	// svelte-ignore state_referenced_locally
	let headerHeight = new Spring(isHome ? 140 : 100);

	// Easter egg for showing old achievement logos
	const isEasterEgg = Math.random() < 1 / 1000;
	const prefix = isEasterEgg ? 'old_' : '';

	let dailySlashLogoLink = `/logos/${prefix}daily_slash_logo.png`;
	let connectionsLogoLink = `/logos/${prefix}connections_logo.png`;
	let guessTheNpcLogoLink = `/logos/${prefix}guess_the_npc_logo.png`;
	let hangmanLogoLink = `/logos/${prefix}hangman_logo.png`;
	let terraTriviaLogoLink = `/logos/${prefix}terratrivia_logo.png`;

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
			<img class="game-logo" src="/logos/TerrariadleLogo.png" alt="Terrariadle" />
			<!-- Remove once the site's been up for a bit -->
			<h2 class="new-msg color-cycle">NEW AND IMPROVED!</h2>
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
			<a class:Selected={page.url.pathname === '/daily-slash'} href={resolve('/daily-slash')}>
				<img
					class:incomplete={!session.dailySlashStatus}
					src={dailySlashLogoLink}
					alt="Daily slash"
				/>
			</a>
			<a class:Selected={page.url.pathname === '/connections'} href={resolve('/connections')}>
				<img
					class:incomplete={!session.connectionStatus}
					src={connectionsLogoLink}
					alt="Connections"
				/>
			</a>
			<a class:Selected={page.url.pathname === '/guess-the-npc'} href={resolve('/guess-the-npc')}>
				<img
					class:incomplete={!session.guessTheNpcStatus}
					src={guessTheNpcLogoLink}
					alt="Guess the NPC"
				/>
			</a>
			<a class:Selected={page.url.pathname === '/hangman'} href={resolve('/hangman')}>
				<img class:incomplete={!session.hangmanStatus} src={hangmanLogoLink} alt="Hangman" />
			</a>
			<a class:Selected={page.url.pathname === '/terratrivia'} href={resolve('/terratrivia')}>
				<img
					class:incomplete={!session.terratriviaStatus}
					src={terraTriviaLogoLink}
					alt="TerraTrivia"
				/>
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

	/* Remove once site has been up a few months */
	.new-msg {
		position: absolute;
		left: 50%;
		top: 100px;
		transform: rotate(-4deg);
		font-size: 22px;
	}

	.logo-slot {
		position: absolute;
		grid-row: 1;
		grid-column: 1;
		display: block;
	}

	.centered {
		justify-self: center;
		max-width: 400px;
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
		transition: background 0.08s;
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

	.incomplete {
		filter: grayscale(100%);
	}

	@media (max-width: 700px) {
		header {
			margin-bottom: 70px;
		}

		.new-msg {
			font-size: 20px;
		}

		.centered {
			margin-top: 20px;
		}

		.corner {
			justify-self: center;
		}

		.links {
			margin-top: 100px;
		}
	}
</style>
