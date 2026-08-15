import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import { ApiError } from '$lib/types/error';
import { initializeHangmanGame } from '$lib/api/hangman';
import { building } from '$app/environment';

export const load: PageLoad = async ({ fetch, parent }) => {
	const { userId } = await parent();

	if (building || !userId)
		return {
			gameContext: {
				attempts: 6,
				finished: false,
				phrase: ['_', '_', '_', ' ', '_', '_', '_', '_', '_'],
				guesses: []
			}
		};

	try {
		const gameContext = await initializeHangmanGame(fetch, userId);
		return { gameContext };
	} catch (e) {
		if (e instanceof ApiError) {
			error(e.status, e.message);
		}
		error(500, 'Unexpected error initializing game');
	}
};
