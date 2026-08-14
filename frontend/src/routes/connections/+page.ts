import type { PageLoad } from './$types';
import { initializeConnectionsGame } from '$lib/api/connections';
import { ApiError } from '$lib/types/error';
import { error } from '@sveltejs/kit';

export const load: PageLoad = async ({ fetch, parent }) => {
	const { userId } = await parent();
	if (!userId) return { gameContext: {} };

	if (!userId) {
		error(401, 'No session found. Try refreshing the page.');
	}

	try {
		const gameContext = await initializeConnectionsGame(fetch, userId);
		return { gameContext };
	} catch (e) {
		if (e instanceof ApiError) {
			error(e.status, e.message);
		}
		error(500, 'Unexpected error initializing game');
	}
};
