import { ApiError, type ApiErrorBody } from '$lib/types/error';
import { parseJsonError, parseJsonSafe } from './utils';
import type {
	GuessTheNpcCheckResult,
	GuessTheNpcMiniGameResult,
	GuessTheNpcSession,
	GuessTheNpcWinningData
} from '$lib/types/guess-the-npc';
import type { DropdownListItem } from '$lib/types/shared';

export async function initializeNpcGame(
	fetchFn: typeof fetch,
	userId: string
): Promise<GuessTheNpcSession> {
	const res = await fetchFn(`/api/guess-the-npc/initialize-game?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to initialize game');
	}

	return parseJsonSafe<GuessTheNpcSession>(res);
}

export async function getSearchableNpcs(fetchFn: typeof fetch) {
	const res = await fetchFn('/api/guess-the-npc/search-items');

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to find npcs');
	}

	return parseJsonSafe<DropdownListItem[]>(res);
}

export async function checkNpcGuess(
	userId: string,
	npcId: number
): Promise<GuessTheNpcCheckResult> {
	const res = await fetch('/api/guess-the-npc/check-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: npcId })
	});

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'An error occurred when checking guess');
	}

	return parseJsonSafe<GuessTheNpcCheckResult>(res);
}

export async function getNpcWinningData(userId: string): Promise<GuessTheNpcWinningData> {
	const res = await fetch(`/api/guess-the-npc/winning-data?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'An error occurred when getting winning data');
	}

	return parseJsonSafe<GuessTheNpcWinningData>(res);
}

export async function checkNpcName(
	userId: string,
	name: string
): Promise<GuessTheNpcMiniGameResult> {
	const res = await fetch('/api/guess-the-npc/check-name-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: name })
	});

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'An error occurred when checking guess');
	}

	return parseJsonSafe<GuessTheNpcMiniGameResult>(res);
}
