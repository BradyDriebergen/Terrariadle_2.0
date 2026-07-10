<script lang="ts">
	import Header from './Header.svelte';
	import Footer from './Footer.svelte';
	import '../app.css';
	import { navigating, page } from '$app/state';
	import { onMount } from 'svelte';
	import type { LayoutData } from './$types';
	import { session } from '$lib/store/session.svelte';

	let { data, children }: { data: LayoutData; children: any } = $props();

	onMount(() => {
		if (data.gameResults) {
			session.dailySlash = data.gameResults.daily_slash;
			session.connection = data.gameResults.connections;
			session.guessTheNpc = data.gameResults.guess_the_npc;
			session.hangman = data.gameResults.hangman;
			session.terratrivia = data.gameResults.terratrivia;
		}
	});

	let title: string = $state(getTitle(page.url.pathname));
	$effect(() => {
		if (navigating.to) {
			title = getTitle(navigating.to.url.pathname);
		}
	});

	function getTitle(path: string): string {
		switch (path) {
			case '/daily-slash':
				return '| Daily Slash';
			case '/connections':
				return '| Connections';
			case '/guess-the-npc':
				return '| Guess the NPC';
			case '/hangman':
				return '| Hangman';
			case '/terratrivia':
				return '| TerraTrivia';
			case '/about':
				return '| About';
			default:
				return '';
		}
	}
</script>

<svelte:head>
	<title>Terrariadle {title}</title>
	<link rel="icon" href="/logos/TabLogo.ico" />
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
