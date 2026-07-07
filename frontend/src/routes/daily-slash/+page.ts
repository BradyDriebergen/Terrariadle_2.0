import type { DropdownListItem } from '$lib/types/shared.js';
import { error } from '@sveltejs/kit';
import { ApiError } from '$lib/types/error';
import { getSearchableWeapons, initializeDailySlashGame } from '$lib/api/daily-slash.js';
import type { PageLoad } from './$types';

export const load: PageLoad = async ({ fetch, parent }) => {
	const { userId } = await parent();
	if (!userId) return { gameContext: null, weaponList: [] };

	try {
		const gameContext = await initializeDailySlashGame(fetch, userId);

		let weaponList: DropdownListItem[] = [];

		if (!gameContext.finished) {
			weaponList = await getSearchableWeapons(fetch);
		}

		return { gameContext, weaponList };
	} catch (e) {
		if (e instanceof ApiError) {
			error(e.status, e.message);
		}
		error(500, 'Unexpected error initializing game');
	}
};
