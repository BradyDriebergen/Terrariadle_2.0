import type { PageLoad } from './$types';
import { initializeConnectionsGame } from '$lib/api/connections';
import { ApiError } from '$lib/types/error';
import { error } from '@sveltejs/kit';
import { building } from '$app/environment';

export const load: PageLoad = async ({ fetch, parent }) => {
	const { userId } = await parent();

	if (building || !userId) {
		return {
			gameContext: {
				attempts: 4,
				finished: false,
				options: [
					'Blood Zombie',
					'Leaf Blower',
					'Drippler',
					'Menogoblin Shark',
					'Nettle Burst',
					'Dungeon Spirit',
					'Spinal Tap',
					'Morning Star',
					'Moth',
					'Seedling',
					'Antlion Swarmer',
					'Thorn Hook',
					'Durendal',
					'Clown',
					'Kaleidoscope',
					'Demon Eye'
				],
				solved_categories: []
			}
		};
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
