import { ApiError } from '$lib/types/error';
import type { DropdownListItem } from '$lib/types/shared';
import { error } from '@sveltejs/kit';
import { getSearchableNpcs, initializeNpcGame } from '$lib/api/guess-the-npc.js';

export async function load({ fetch, parent }) {
	const { userId } = await parent();
	if (!userId) return { gameContext: null, npcList: [] };

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
}