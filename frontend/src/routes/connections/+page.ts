import type { ConnectionsSession } from '$lib/types/connections.js';
import { error } from '@sveltejs/kit';

export async function load({ fetch, parent }) {
	const { userId } = await parent();

	const contextRes = await fetch(`/api/connections/initialize-game/?user_id=${userId}`);
	if (!contextRes.ok) {
		const body = await contextRes.json();
		error(contextRes.status, body.error ?? 'Unable to initialize game');
	}

	const gameContext = (await contextRes.json()) as ConnectionsSession;

	return { gameContext };
}

