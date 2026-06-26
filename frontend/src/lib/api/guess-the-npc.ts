import { ApiError } from '$lib/types/error';
import { parseJsonSafe } from './shared';
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
	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unable to initialize game');
	}

	return body as GuessTheNpcSession;
}

export async function getSearchableNpcs(fetchFn: typeof fetch) {
	const res = await fetchFn('/api/guess-the-npc/search-items');
	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unable to find npcs');
	}

	return body as DropdownListItem[];
}

export async function checkNpcGuess(npcId: number): Promise<GuessTheNpcCheckResult> {
	const userId = localStorage.getItem('user_id');

	if (!userId) {
		throw new Error('Session not found. Try refreshing the page.');
	}

	const res = await fetch('/api/guess-the-npc/check-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: npcId })
	});

	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'An error occurred when checking guess');
	}

	return body as GuessTheNpcCheckResult;
}

export async function getNpcWinningData(): Promise<GuessTheNpcWinningData> {
	const userId = localStorage.getItem('user_id');

	if (!userId) {
		throw new Error('Session not found. Try refreshing the page.');
	}

	const res = await fetch(`/api/guess-the-npc/winning-data?user_id=${userId}`);
	return (await res.json()) as GuessTheNpcWinningData;
}

export async function checkNpcName(name: string): Promise<GuessTheNpcMiniGameResult> {
	const userId = localStorage.getItem('user_id');

	if (!userId) {
		throw new Error('Session not found. Try refreshing the page.');
	}

	const res = await fetch('/api/guess-the-npc/check-name-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: name })
	});

	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'An error occurred when checking guess');
	}

	return body as GuessTheNpcMiniGameResult;
}
