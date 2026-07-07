import { error } from '@sveltejs/kit';
import { ApiError } from '$lib/types/error';
import { initializeTerraTriviaGame } from '$lib/api/terratrivia';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, parent }) => {
	const { userId } = await parent();
	if (!userId) return { gameContext: null };

	if (!userId) {
		error(401, 'No session found. Try refreshing the page.');
	}

	try {
		const gameContext = await initializeTerraTriviaGame(fetch, userId);
		return { gameContext };
	} catch (e) {
		if (e instanceof ApiError) {
			error(e.status, e.message);
		}
		error(500, 'Unexpected error initializing game');
	}
};
