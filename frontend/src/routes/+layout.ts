import { browser, building } from '$app/environment';
import { getUserGameResults } from '$lib/api/common';
import { getOrCreateUserId } from '$lib/utils/user-id';
import type { LayoutLoad } from './$types';

// prevents server side rendering
// export const ssr = false;

export const prerender = true;

export const load: LayoutLoad = async ({ fetch }) => {
	if (!browser) {
		return {
			gameResults: {
				daily_slash: false,
				connections: false,
				guess_the_npc: false,
				hangman: false,
				terratrivia: false
			},
			userId: building ? 'seo-crawler' : null
		};
	}

	const userId = getOrCreateUserId();
	const gameResults = await getUserGameResults(fetch, userId);

	return { gameResults, userId };
};
