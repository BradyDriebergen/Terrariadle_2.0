import type { DailySlashSession, WeaponListItem, WeaponPreview } from '$lib/types/daily-slash.js';
import { error } from '@sveltejs/kit';

export async function load({ fetch, parent }) {
	const { userId } = await parent();

	const contextRes = await fetch(`/api/daily-slash/initialize-game/?user_id=${userId}`);
	if (!contextRes.ok) {
		const body = await contextRes.json();
		error(contextRes.status, body.error ?? 'Unable to initialize game');
	}

	const gameContext = await contextRes.json() as DailySlashSession;

	let weaponList: WeaponListItem[] = [];

	if (!gameContext.finished) {
		const weaponsRes = await fetch('/api/daily-slash/search-items');
		if (!weaponsRes.ok) {
			const body = await weaponsRes.json();
			error(weaponsRes.status, body.error ?? 'Unable to fetch weapons');
		}

		weaponList = await weaponsRes.json() as WeaponListItem[];
	}

	return { gameContext, weaponList };
}
