import { ApiError } from "$lib/types/error";
import type { UserGameResults } from "$lib/types/shared";
import { parseJsonSafe } from "./shared";

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
	const body = await parseJsonSafe(res);

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unable to get game results');
	}

	return body as UserGameResults;
}