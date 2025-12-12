<script lang="ts">
	import Header from './Header.svelte';
	import Footer from './Footer.svelte';
	import '../app.css';
	import { navigating, page } from '$app/state';

	let { children } = $props();

	let title = $state(getTitle(page.url.pathname));
	$effect(() => {
		if (navigating.to) {
			title = getTitle(navigating.to.url.pathname);
		}
	})

	function getTitle(path: string): string {
		switch(path) {
			case '/daily-slash':
				return '| Daily Slash';
			case '/connections':
				return '| Connections';
			case '/guess-the-npc':
				return '| Guess the NPC';
			case '/hangman':
				return '| Hangman';
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
</style>
