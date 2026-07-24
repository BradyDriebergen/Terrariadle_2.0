<script lang="ts">
	import { page } from '$app/state';
	import { resolve } from '$app/paths';
	import { fly } from 'svelte/transition';
	import { session } from '$lib/store/session.svelte';

	let isDailySlashPage: boolean = $derived(page.url.pathname === '/daily-slash');
	let isConnectionsPage: boolean = $derived(page.url.pathname === '/connections');
	let isGuessTheNpcPage: boolean = $derived(page.url.pathname === '/guess-the-npc');
	let isHangmanPage: boolean = $derived(page.url.pathname === '/hangman');
	let isTerraTriviaPage: boolean = $derived(page.url.pathname === '/terratrivia');
</script>

<div class="landing-page" in:fly={{ y: 50, duration: 500 }}>
	<p>
		Welcome to Terrariadle! <br />
		Test your Terraria knowledge through daily puzzle games.
	</p>
	<br />

	<ul>
		<li aria-current={isDailySlashPage}>
			{#if session.dailySlashStatus}
				<img src="/emojis/red_x.png" alt="Game complete!" />
			{/if}
			<a href={resolve('/daily-slash')}>Daily Slash</a>
		</li>
		<li aria-current={isConnectionsPage}>
			{#if session.connectionStatus}
				<img src="/emojis/red_x.png" alt="Game complete!" />
			{/if}
			<a href={resolve('/connections')}>Connections</a>
		</li>
		<li aria-current={isGuessTheNpcPage}>
			{#if session.guessTheNpcStatus}
				<img src="/emojis/red_x.png" alt="Game complete!" />
			{/if}
			<a href={resolve('/guess-the-npc')}>Guess The NPC</a>
		</li>
		<li aria-current={isHangmanPage}>
			{#if session.hangmanStatus}
				<img src="/emojis/red_x.png" alt="Game complete!" />
			{/if}
			<a href={resolve('/hangman')}>Hangman</a>
		</li>
		<li aria-current={isTerraTriviaPage}>
			{#if session.terratriviaStatus}
				<img src="/emojis/red_x.png" alt="Game complete!" />
			{/if}
			<a href={resolve('/terratrivia')}>TerraTrivia</a>
		</li>
	</ul>
</div>

<style>
	p {
		font-size: 15px;
		color: lightgray;
	}

	ul {
		list-style: none;
		padding: 0;
		margin: 0;
	}

	li {
		margin-right: 1rem;
		margin-bottom: 1rem;
	}

	a {
		text-decoration: none;
		color: grey;
		font-size: 23px;
		display: inline-block;
		transition:
			transform 0.3s ease,
			color 0.3s ease;
	}

	a:hover {
		transform: scale(1.1) rotate(-3deg); /* grows and tilts left */
		color: yellow;
	}

	img {
		position: absolute;
		pointer-events: none;
	}

	@media (max-width: 700px) {
		.landing-page {
			margin-top: -60px;
		}

		p {
			font-size: 17px;
		}
	}
</style>
