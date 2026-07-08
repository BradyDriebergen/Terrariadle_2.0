<script lang="ts">
	import { page } from '$app/state';
	import { resolve } from '$app/paths';
	import { fly } from 'svelte/transition';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	let isDailySlashPage: boolean = $derived(page.url.pathname === '/daily-slash');
	let isConnectionsPage: boolean = $derived(page.url.pathname === '/connections');
	let isGuessTheNpcPage: boolean = $derived(page.url.pathname === '/guess-the-npc');
	let isHangmanPage: boolean = $derived(page.url.pathname === '/hangman');
	let isTerraTriviaPage: boolean = $derived(page.url.pathname === '/terratrivia');
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
				class:strikethrough={data.gameResults?.daily_slash}
			>
				Daily Slash
			</a>
		</li>
		<li aria-current={isConnectionsPage}>
			<a 
				href={resolve('/connections')} 
				class:strikethrough={data.gameResults?.connections}
			>
				Connections
			</a>
		</li>
		<li aria-current={isGuessTheNpcPage}>
			<a 
				href={resolve('/guess-the-npc')} 
				class:strikethrough={data.gameResults?.guess_the_npc}
			>
				Guess The NPC
			</a>
		</li>
		<li aria-current={isHangmanPage}>
			<a 
				href={resolve('/hangman')} 
				class:strikethrough={data.gameResults?.hangman}
			>
				Hangman
			</a>
		</li>
		<li aria-current={isTerraTriviaPage}>
			<a 
				href={resolve('/terratrivia')} 
				class:strikethrough={data.gameResults?.terratrivia}
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
