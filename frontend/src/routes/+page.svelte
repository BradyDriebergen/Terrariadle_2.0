<script>
	import { page } from '$app/state';
</script>

<div>
	<ul>
		<li aria-current={page.url.pathname === '/daily-slash' ? 'page' : undefined}>
			<a href="/daily-slash">Daily Slash</a>
		</li>
		<li aria-current={page.url.pathname === '/connections' ? 'page' : undefined}>
			<a href="/connections">Connections</a>
		</li>
		<li aria-current={page.url.pathname === '/guess-the-npc' ? 'page' : undefined}>
			<a href="/guess-the-npc">Guess The NPC</a>
		</li>
		<li aria-current={page.url.pathname === '/hangman' ? 'page' : undefined}>
			<a href="/hangman">Hangman</a>
		</li>
		<li aria-current={page.url.pathname === '/about' ? 'page' : undefined}>
			<a href="/about">About</a>
		</li>
	</ul>
</div>

<style>
	ul {
		list-style: none;
		color: yellow;
		padding: 0;
		margin: 0;
	}

	li {
		margin-right: 1rem;
		margin-bottom: 1rem;
	}

	a {
		text-decoration: none; /* remove underline */
		color: inherit; /* match the <li> text color */
		display: inline-block; /* allows transform scaling */
		transition:
			transform 0.3s ease,
			color 0.3s ease;
	}

	a:hover {
		transform: scale(1.1) rotate(-3deg); /* grows and tilts left */
	}
</style>
