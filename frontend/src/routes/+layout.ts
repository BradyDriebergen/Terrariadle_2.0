import { getUserGameResults } from '$lib/api/common';
import { getOrCreateUserId } from '$lib/utils/user-id';
import type { LayoutLoad } from './$types';

// prevents server side rendering
export const ssr = false;

export const load: LayoutLoad = async ({ fetch }) => {
	const userId = getOrCreateUserId();
	const gameResults = await getUserGameResults(fetch, userId);

	return { gameResults, userId };
};
