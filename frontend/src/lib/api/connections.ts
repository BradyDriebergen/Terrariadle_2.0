import type { ConnectionsCheckResult } from "$lib/types/connections";
import { error } from "@sveltejs/kit";

export async function checkCategoryGuess(options: string[]): Promise<ConnectionsCheckResult> {
    const userId = localStorage.getItem('user_id');

    if (!userId) {
        throw new Error('Session not found. Try refreshing the page.');
    }

    if (options.length !== 4) {
        throw new Error('Connections guess must have a length of 4');
    }

    const res = await fetch('/api/connections/check-guess', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ user_id: userId, guess: options })
    });

    const body = await res.json();
    if (!res.ok) {
        error(res.status, body.error ?? 'An error occurred when checking guess');
    }

    return body as ConnectionsCheckResult;
}