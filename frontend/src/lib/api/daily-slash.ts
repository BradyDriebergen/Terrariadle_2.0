import type {
	DailySlashCheckResult,
	DailySlashSession,
	DailySlashWinningData
} from '$lib/types/daily-slash';
import { ApiError, type ApiErrorBody } from '$lib/types/error';
import { parseJsonError, parseJsonSafe } from './utils';
import type { DropdownListItem } from '$lib/types/shared';

export async function initializeDailySlashGame(
	fetchFn: typeof fetch,
	userId: string
): Promise<DailySlashSession> {
	const res = await fetchFn(`/api/daily-slash/initialize-game?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to initialize game');
	}

	return parseJsonSafe<DailySlashSession>(res);
}

export async function getSearchableWeapons(fetchFn: typeof fetch) {
	const res = await fetchFn('/api/daily-slash/search-items');

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to find weapons');
	}

	return parseJsonSafe<DropdownListItem[]>(res);
}

export async function getWeaponHint(num: number): Promise<string> {
	if (num < 1 || num > 3) {
		throw new Error('Invalid hint number input');
	}

	const res = await fetch(`/api/daily-slash/hint?hint=${num}`);
	return (await res.json()) as string;
}

export async function checkWeaponGuess(
	userId: string,
	weaponId: number
): Promise<DailySlashCheckResult> {
	const res = await fetch('/api/daily-slash/check-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: weaponId })
	});

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'An error occurred when checking guess');
	}

	return parseJsonSafe<DailySlashCheckResult>(res);
}

export async function getDailySlashWinningData(userId: string): Promise<DailySlashWinningData> {
	const res = await fetch(`/api/daily-slash/winning-data?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'An error occurred when getting winning data');
	}

	return parseJsonSafe<DailySlashWinningData>(res);
}
