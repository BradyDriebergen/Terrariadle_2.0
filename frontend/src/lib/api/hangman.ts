import { ApiError } from '$lib/types/error';
import type { HangmanCheckResult, HangmanSession, HangmanWinningData } from '$lib/types/hangman';
import { parseJsonError, parseJsonSafe } from './utils';

export async function initializeHangmanGame(
	fetchFn: typeof fetch,
	userId: string
): Promise<HangmanSession> {
	const res = await fetchFn(`/api/hangman/initialize-game?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to initialize game');
	}

	return parseJsonSafe<HangmanSession>(res);
}

export async function checkEnemyGuess(userId: string, letter: string): Promise<HangmanCheckResult> {
	const res = await fetch('/api/hangman/check-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: letter })
	});

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'An error occurred when checking guess');
	}

	return parseJsonSafe<HangmanCheckResult>(res);
}

export async function getHangmanWinningData(userId: string): Promise<HangmanWinningData> {
	const res = await fetch(`/api/hangman/winning-data?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to load winning data');
	}

	return parseJsonSafe<HangmanWinningData>(res);
}
