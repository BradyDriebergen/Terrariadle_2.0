// +layout.ts
import { browser } from '$app/environment';
import { getUserGameResults } from '$lib/api/common';
import { getOrCreateUserId } from '$lib/api/shared';
import { ApiError } from '$lib/types/error';
import type { LayoutLoad } from './$types';
import { error } from '@sveltejs/kit';

export const load: LayoutLoad = async ({ fetch }) => {
	if (!browser) return { userId: null, gameResults: null };

	const userId = getOrCreateUserId();

	try {
		const gameResults = await getUserGameResults(fetch, userId);
		return { userId, gameResults };
	} catch (e) {
		if (e instanceof ApiError) {
			error(e.status, e.message);
		}
		error(500, 'Unexpected error initializing game');
	}
}
