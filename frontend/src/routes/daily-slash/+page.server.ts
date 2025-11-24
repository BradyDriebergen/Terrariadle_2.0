import type { SimpleWeapon } from '$lib/types/dailySlash.js';
import { error } from '@sveltejs/kit';

export async function load({ fetch }) {
	const res = await fetch('http://localhost:3000/api/daily-slash/search-items');

	if (!res) {
		error(404, 'Unable to fetch weapons');
	}

	const weapons: SimpleWeapon[] = await res.json();

	return { weapons };
}
