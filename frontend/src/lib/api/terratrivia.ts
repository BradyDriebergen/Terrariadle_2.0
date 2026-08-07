import { ApiError, type ApiErrorBody } from '$lib/types/error';
import type {
	TerraTriviaCheckResult,
	TerraTriviaSession,
	TerraTriviaWinningData
} from '$lib/types/terratrivia';
import { parseJsonError, parseJsonSafe } from './utils';

export async function initializeTerraTriviaGame(
	fetchFn: typeof fetch,
	userId: string
): Promise<TerraTriviaSession> {
	const res = await fetchFn(`/api/terratrivia/initialize-game?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to initialize game');
	}

	return parseJsonSafe<TerraTriviaSession>(res);
}

export async function checkTriviaQuestionGuess(
	userId: string,
	guess: string
): Promise<TerraTriviaCheckResult> {
	const res = await fetch('/api/terratrivia/check-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess })
	});

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'An error occurred when checking guess');
	}

	return parseJsonSafe<TerraTriviaCheckResult>(res);
}

export async function getTerraTriviaWinningData(userId: string): Promise<TerraTriviaWinningData> {
	const res = await fetch(`/api/terratrivia/winning-data?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to load winning data');
	}

	return parseJsonSafe<TerraTriviaWinningData>(res);
}
