import { ApiError } from '$lib/types/error';
import type { UserGameResults } from '$lib/types/shared';
import { parseJsonError, parseJsonSafe } from './utils';

export function subscribeToPlayerCount(mode: string, onCount: (count: number) => void): () => void {
	const es = new EventSource(`/api/guess-count?mode=${mode}`);

	es.onmessage = (event) => {
		const data = JSON.parse(event.data);
		onCount(data.count);
	};

	es.onerror = (err) => {
		console.error('SSE error:', err);
		if (es.readyState === EventSource.CLOSED) {
			console.error('SSE permanently closed');
		}
	};

	return () => es.close();
}

export async function getUserGameResults(
	fetchFn: typeof fetch,
	userId: string
): Promise<UserGameResults> {
	const res = await fetchFn(`/api/finished-games?user_id=${userId}`);

	if (!res.ok) {
		const err = await parseJsonError(res);
		throw new ApiError(res.status, err.error ?? 'Unable to get game results');
	}

	return parseJsonSafe<UserGameResults>(res);
}
