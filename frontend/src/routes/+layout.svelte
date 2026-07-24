<script lang="ts">
	import Header from './Header.svelte';
	import Footer from './Footer.svelte';
	import '../app.css';
	import { page } from '$app/state';
	import { onMount, type Snippet } from 'svelte';
	import type { LayoutData } from './$types';
	import { session } from '$lib/store/session.svelte';

	let { data, children }: { data: LayoutData; children: Snippet } = $props();

	onMount(() => {
		if (data.gameResults) {
			session.dailySlashStatus = data.gameResults.daily_slash;
			session.connectionStatus = data.gameResults.connections;
			session.guessTheNpcStatus = data.gameResults.guess_the_npc;
			session.hangmanStatus = data.gameResults.hangman;
			session.terratriviaStatus = data.gameResults.terratrivia;
		}
	});

	const SITE_URL = 'https://terrariadle.com';
	const routeMeta: Record<string, { title: string; description: string }> = {
		'/': {
			title: '',
			description:
				'Welcome to Terrariadle! Test your Terraria knowledge through daily puzzle games.'
		},
		'/daily-slash': {
			title: '| Daily Slash',
			description: 'Try to guess the Terraria weapon using hints from previous guesses!'
		},
		'/connections': {
			title: '| Connections',
			description: 'Find groups of Terraria-themed options that share something in common!'
		},
		'/guess-the-npc': {
			title: '| Guess the NPC',
			description: 'Guess the Terraria NPC from a quote!'
		},
		'/hangman': {
			title: '| Hangman',
			description: 'Guess the name of the Terraria enemy one letter at a time!'
		},
		'/terratrivia': {
			title: '| TerraTrivia',
			description: 'Build Terraria-based terms using letter bundles'
		},
		'/about': {
			title: '| About',
			description: 'Learn more about Terrariadle!'
		}
	};

	let meta = $derived(routeMeta[page.url.pathname] ?? routeMeta['/']);
	let canonicalUrl = $derived(`${SITE_URL}${page.url.pathname}`);
</script>

<svelte:head>
	<title>Terrariadle {meta.title}</title>
	<link rel="icon" href="/logos/TabLogo.ico" />
	<meta name="description" content={meta.description} />
	<link rel="canonical" href={canonicalUrl} />

	<!-- Open Graph — Discord, Slack, Facebook, iMessage all read these -->
	<meta property="og:type" content="website" />
	<meta property="og:site_name" content="Terrariadle" />
	<meta property="og:title" content={`Terrariadle ${meta.title}`} />
	<meta property="og:description" content={meta.description} />
	<meta property="og:url" content={canonicalUrl} />
	<meta property="og:image" content={`${SITE_URL}/og-image.png`} />
	<meta property="og:image:width" content="1200" />
	<meta property="og:image:height" content="630" />

	<!-- Twitter/X Card -->
	<meta name="twitter:card" content="summary_large_image" />
	<meta name="twitter:title" content={`Terrariadle ${meta.title}`} />
	<meta name="twitter:description" content={meta.description} />
	<meta name="twitter:image" content={`${SITE_URL}/og-image.png`} />
</svelte:head>

<div class="app">
	<Header />

	<main>
		{@render children()}
	</main>

	<Footer />
</div>

<style>
	main {
		padding: 20px;
	}
</style>
