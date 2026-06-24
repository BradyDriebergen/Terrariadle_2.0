import type { DropdownListItem } from '$lib/types/shared.js';
import type { DailySlashSession } from '$lib/types/daily-slash.js';
import { error } from '@sveltejs/kit';
import { ApiError } from '$lib/types/error';
import { getSearchableWeapons, initializeDailySlashGame } from '$lib/api/daily-slash.js';

export async function load({ fetch, parent }) {
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
}
