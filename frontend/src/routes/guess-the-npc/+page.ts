import { ApiError } from '$lib/types/error';
import type { DropdownListItem } from '$lib/types/shared';
import { error } from '@sveltejs/kit';
import { getSearchableNpcs, initializeNpcGame } from '$lib/api/guess-the-npc.js';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, parent }) => {
	const { userId } = await parent();
	if (!userId) return { gameContext: {}, npcList: [] };

	try {
		const gameContext = await initializeNpcGame(fetch, userId);

		let npcList: DropdownListItem[] = [];

		if (!gameContext.finished) {
			npcList = await getSearchableNpcs(fetch);
		}

		return { gameContext, npcList };
	} catch (e) {
		if (e instanceof ApiError) {
			error(e.status, e.message);
		}
		error(500, 'Unexpected error initializing game');
	}
};
