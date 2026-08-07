import type {
	ConnectionsCheckResult,
	ConnectionsRevealData,
	ConnectionsSession,
	ConnectionsWinningData
} from '$lib/types/connections';
import { ApiError } from '$lib/types/error';
import { parseJsonError, parseJsonSafe } from './utils';

export async function initializeConnectionsGame(
	fetchFn: typeof fetch,
	userId: string
): Promise<ConnectionsSession> {
	const res = await fetchFn(`/api/connections/initialize-game?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to initialize game');
	}

	return parseJsonSafe<ConnectionsSession>(res);
}

export async function checkCategoryGuess(
	options: string[],
	userId: string
): Promise<ConnectionsCheckResult> {
	if (options.length !== 4) {
		throw new ApiError(400, 'Connections guess must have a length of 4');
	}

	const res = await fetch('/api/connections/check-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: options })
	});

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'An error occurred when checking guess');
	}

	return parseJsonSafe<ConnectionsCheckResult>(res);
}

export async function revealConnectionsAnswers(userId: string): Promise<ConnectionsRevealData> {
	const res = await fetch('/api/connections/reveal-answers', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId })
	});

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'An error occurred when revealing answers');
	}

	return parseJsonSafe<ConnectionsRevealData>(res);
}

export async function getConnectionsWinningData(userId: string): Promise<ConnectionsWinningData> {
	const res = await fetch(`/api/connections/winning-data?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to load winning data');
	}

	return parseJsonSafe<ConnectionsWinningData>(res);
}
