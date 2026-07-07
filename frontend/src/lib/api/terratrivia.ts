import { ApiError } from '$lib/types/error';
import type {
	TerraTriviaCheckResult,
	TerraTriviaSession,
	TerraTriviaWinningData
} from '$lib/types/terratrivia';
import { parseJsonSafe } from './shared';

export async function initializeTerraTriviaGame(
	fetchFn: typeof fetch,
	userId: string
): Promise<TerraTriviaSession> {
	const res = await fetchFn(`/api/terratrivia/initialize-game?user_id=${userId}`);
	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unable to initialize game');
	}

	return body as TerraTriviaSession;
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

	const body = await parseJsonSafe(res);
	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'An error occurred when checking guess');
	}

	return body as TerraTriviaCheckResult;
}

export async function getTerraTriviaWinningData(userId: string): Promise<TerraTriviaWinningData> {
	const res = await fetch(`/api/terratrivia/winning-data?user_id=${userId}`);
	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unable to load winning data');
	}

	return body as TerraTriviaWinningData;
}
