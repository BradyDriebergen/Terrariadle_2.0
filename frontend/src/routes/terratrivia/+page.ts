import { error } from '@sveltejs/kit';
import { ApiError } from '$lib/types/error';
import { initializeTerraTriviaGame } from '$lib/api/terratrivia';
import type { PageLoad } from './$types';
import { building } from '$app/environment';

export const load: PageLoad = async ({ fetch, parent }) => {
	const { userId } = await parent();

	if (building)
		return {
			gameContext: {
				finished: false,
				trivia_items: [
					{ id: 1165, clue: 'Increases range of crafting stations', letter_count: 11, answer: '' },
					{ id: 1554, clue: 'Flying hardmode Corruption enemy', letter_count: 9, answer: '' },
					{ id: 986, clue: 'Hermes Boots + Flower Boots', letter_count: 10, answer: '' },
					{ id: 1639, clue: 'Sandstorm enemy', letter_count: 12, answer: '' },
					{ id: 1459, clue: 'Underworld slime', letter_count: 9, answer: '' },
					{ id: 1295, clue: 'Summons the Eyeball Spring pet', letter_count: 9, answer: '' },
					{ id: 331, clue: 'Pick that can mine all blocks', letter_count: 7, answer: '' }
				],
				chunks: [
					'ARTI',
					'LA',
					'SANL',
					'YTUM',
					'COR',
					'PIC',
					'KSAW',
					'RU',
					'SPR',
					'VASL',
					'OAF',
					'OTS',
					'PTOR',
					'FAI',
					'IME',
					'ANGR',
					'EYE',
					'BLER',
					'ING',
					'RYBO'
				]
			}
		};

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
