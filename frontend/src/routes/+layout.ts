import { browser } from '$app/environment';
import { getUserGameResults } from '$lib/api/common';
import { getOrCreateUserId } from '$lib/utils/user-id';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ fetch }) => {
	if (!browser) return { userId: null };

	const userId = getOrCreateUserId();
	const gameResults = await getUserGameResults(fetch, userId);

	return { gameResults, userId };
};
