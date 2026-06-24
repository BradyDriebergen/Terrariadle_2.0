// +layout.ts
import { browser } from '$app/environment';
import { getOrCreateUserId, getUserId } from '$lib/api/shared';

export function load() {
	if (!browser) return { userId: null };
	return { userId: getOrCreateUserId() };
}
