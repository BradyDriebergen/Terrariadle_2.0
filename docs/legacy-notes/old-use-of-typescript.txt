import type { SimpleWeapon, Weapon } from '$lib/types/dailySlash.js';
import { error } from '@sveltejs/kit';

export async function load({ fetch, parent }) {
	const { userId } = await parent();

	const initData = await fetch(`http://localhost:3000/api/daily-slash/initialize-game/${userId}`);
	if (!initData) {
		error(404, 'Unable to fetch initializing data');
	}
	const initDataJson = await initData.json();

	const previousWeapon = initDataJson.previousWeapon;
	const guesses: Weapon[] = initDataJson.guesses ?? [];
	const checks = initDataJson.checks;
	const won = initDataJson.won;

	let weapons: SimpleWeapon[] = [];
	if (!won) {
		const weaponFetch = await fetch('http://localhost:3000/api/daily-slash/search-items');
		if (!weaponFetch) {
			error(404, 'Unable to fetch weapons');
		}
		weapons = await weaponFetch.json();

		const guessIds: number[] = guesses?.map((g) => g.id) ?? [];
		weapons = weapons.filter((w) => !guessIds.includes(w.id));
	}

	return { weapons, previousWeapon, guesses, checks, won };
}
