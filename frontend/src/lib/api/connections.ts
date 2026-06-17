import type { ConnectionsCheckResult, ConnectionsSession, ConnectionsWinningData } from "$lib/types/connections";
import { ApiError } from "$lib/types/error";
import { getUserId, parseJsonSafe } from "./shared";

export async function initializeConnectionsGame(
  fetchFn: typeof fetch,
  userId: string
): Promise<ConnectionsSession> {
    const res = await fetchFn(`/api/connections/initialize-game/?user_id=${userId}`);

    if (!res.ok) {
        const body = await res.json().catch(() => null);
        throw new ApiError(res.status, body?.error ?? 'Unable to initialize game');
    }

    return res.json() as Promise<ConnectionsSession>;
}

export async function checkCategoryGuess(
    options: string[]
): Promise<ConnectionsCheckResult> {
    const userId = getUserId();

    if (options.length !== 4) {
        throw new ApiError(400, 'Connections guess must have a length of 4');
    }

    const res = await fetch('/api/connections/check-guess', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ user_id: userId, guess: options })
    });

    const body = await parseJsonSafe(res);

    if (!res.ok) {
        throw new ApiError(res.status, body?.error ?? 'An error occurred when checking guess');
    }

    return body as ConnectionsCheckResult;
}

export async function getConnectionsWinningData(): Promise<ConnectionsWinningData> {
    const userId = getUserId();

    const res = await fetch(`/api/connections/winning-data/?user_id=${userId}`);
    const body = await parseJsonSafe(res);

    if (!res.ok) {
        throw new ApiError(res.status, body?.error ?? 'Unable to load winning data');
    }

    return body as ConnectionsWinningData;
}