<script>
	import { page } from '$app/state';
	import { resolve } from '$app/paths';
</script>

<div>
	<p>
		A Terraria inspired daily puzzle game. <br />
		Test your Terraria knowledge.
	</p>
	<br />

	<ul>
		<li aria-current={page.url.pathname === '/daily-slash' ? 'page' : undefined}>
			<a href={resolve('/daily-slash')}>Daily Slash</a>
		</li>
		<li aria-current={page.url.pathname === '/connections' ? 'page' : undefined}>
			<a href={resolve('/connections')}>Connections</a>
		</li>
		<li aria-current={page.url.pathname === '/guess-the-npc' ? 'page' : undefined}>
			<a href={resolve('/guess-the-npc')}>Guess The NPC</a>
		</li>
		<li aria-current={page.url.pathname === '/hangman' ? 'page' : undefined}>
			<a href={resolve('/hangman')}>Hangman</a>
		</li>
		<li aria-current={page.url.pathname === '/about' ? 'page' : undefined}>
			<a href={resolve('/about')}>About</a>
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
</style>
