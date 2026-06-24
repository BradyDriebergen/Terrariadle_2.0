import type {
	DailySlashCheckResult,
	DailySlashSession,
	DailySlashWinningData
} from '$lib/types/daily-slash';
import { ApiError } from '$lib/types/error';
import { error } from '@sveltejs/kit';
import { parseJsonSafe } from './shared';
import type { DropdownListItem } from '$lib/types/shared';

export async function initializeDailySlashGame(
	fetchFn: typeof fetch,
	userId: string
): Promise<DailySlashSession> {
	const res = await fetchFn(`/api/daily-slash/initialize-game?user_id=${userId}`);
	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unable to initialize game');
	}

	return body as DailySlashSession;
}

export async function getSearchableWeapons(fetchFn: typeof fetch) {
	const res = await fetchFn('/api/daily-slash/search-items');
	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unable to find weapons');
	}

	return body as DropdownListItem[];
}

export async function getWeaponHint(num: number): Promise<string> {
	if (num < 1 || num > 3) {
		throw new Error('Invalid hint number input');
	}

	const res = await fetch(`/api/daily-slash/hint?hint=${num}`);
	return (await res.json()) as string;
}

export async function checkWeaponGuess(weaponId: number): Promise<DailySlashCheckResult> {
	const userId = localStorage.getItem('user_id');

	if (!userId) {
		throw new Error('Session not found. Try refreshing the page.');
	}

	const res = await fetch('/api/daily-slash/check-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: weaponId })
	});

	const body = await res.json();
	if (!res.ok) {
		error(res.status, body.error ?? 'An error occurred when checking guess');
	}

	return body as DailySlashCheckResult;
}

export async function getDailySlashWinningData(): Promise<DailySlashWinningData> {
	const userId = localStorage.getItem('user_id');

	if (!userId) {
		throw new Error('Session not found. Try refreshing the page.');
	}

	const res = await fetch(`/api/daily-slash/winning-data?user_id=${userId}`);
	return (await res.json()) as DailySlashWinningData;
}
