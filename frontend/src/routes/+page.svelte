<script>
	import { page } from '$app/state';
	import { resolve } from '$app/paths';
	import { fly } from 'svelte/transition';

	let isDailySlashPage = $derived(page.url.pathname === '/daily-slash');
	let isConnectionsPage = $derived(page.url.pathname === '/connections');
	let isGuessTheNpcPage = $derived(page.url.pathname === '/guess-the-npc');
	let isHangmanPage = $derived(page.url.pathname === '/hangman');
	let isTerraTriviaPage = $derived(page.url.pathname === '/terratrivia');
</script>

<div in:fly={{ y: 50, duration: 500 }}>
	<p>
		A Terraria inspired daily puzzle game. <br />
		Test your Terraria knowledge.
	</p>
	<br />

	<ul>
		<li aria-current={isDailySlashPage}>
			<a 
				href={resolve('/daily-slash')} 
				class:strikethrough={false}
			>
				Daily Slash
			</a>
		</li>
		<li aria-current={isConnectionsPage}>
			<a 
				href={resolve('/connections')} 
				class:strikethrough={false}
			>
				Connections
			</a>
		</li>
		<li aria-current={isGuessTheNpcPage}>
			<a 
				href={resolve('/guess-the-npc')} 
				class:strikethrough={false}
			>
				Guess The NPC
			</a>
		</li>
		<li aria-current={isHangmanPage}>
			<a 
				href={resolve('/hangman')} 
				class:strikethrough={false}
			>
				Hangman
			</a>
		</li>
		<li aria-current={isTerraTriviaPage}>
			<a 
				href={resolve('/terratrivia')} 
				class:strikethrough={false}
			>
				TerraTrivia
			</a>
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

	.strikethrough {
		text-decoration: line-through;
	}
</style>
