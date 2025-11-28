import type { PreviousWeapon, SimpleWeapon } from '$lib/types/dailySlash.js';
import { error } from '@sveltejs/kit';

export async function load({ fetch, parent }) {
	const { userId } = await parent();

	const weaponFetch = await fetch('http://localhost:3000/api/daily-slash/search-items');
	if (!weaponFetch) {
		error(404, 'Unable to fetch weapons');
	}
	const weapons: SimpleWeapon[] = await weaponFetch.json();

	const initData = await fetch(`http://localhost:3000/api/daily-slash/initialize-game/${userId}`);
	if (!weaponFetch) {
		error(404, 'Unable to fetch weapons');
	}
	const initDataJson = await initData.json();

	const previousWeapon = initDataJson.previousWeapon;
	const guesses = initDataJson.guesses;

	return { weapons, previousWeapon, guesses };
}
