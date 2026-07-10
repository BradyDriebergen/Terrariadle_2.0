import { ApiError } from '$lib/types/error';
import type { HangmanCheckResult, HangmanSession, HangmanWinningData } from '$lib/types/hangman';
import { parseJsonSafe } from './utils';

export async function initializeHangmanGame(
	fetchFn: typeof fetch,
	userId: string
): Promise<HangmanSession> {
	const res = await fetchFn(`/api/hangman/initialize-game?user_id=${userId}`);
	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unable to initialize game');
	}

	return body as HangmanSession;
}

export async function checkEnemyGuess(userId: string, letter: string): Promise<HangmanCheckResult> {
	const res = await fetch('/api/hangman/check-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: letter })
	});

	const body = await parseJsonSafe(res);
	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'An error occurred when checking guess');
	}

	return body as HangmanCheckResult;
}

export async function getHangmanWinningData(userId: string): Promise<HangmanWinningData> {
	const res = await fetch(`/api/hangman/winning-data?user_id=${userId}`);
	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unable to load winning data');
	}

	return body as HangmanWinningData;
}
