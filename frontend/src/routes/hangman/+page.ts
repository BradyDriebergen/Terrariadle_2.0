import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import { ApiError } from '$lib/types/error';
import { initializeHangmanGame } from '$lib/api/hangman';

export const load: PageLoad = async ({ fetch, parent }) => {
	const { userId } = await parent();
	if (!userId) return { gameContext: null };

	if (!userId) {
		error(401, 'No session found. Try refreshing the page.');
	}

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
