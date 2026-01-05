import type { SimpleNpc } from '$lib/types/guess-the-npc.js';
import { error } from '@sveltejs/kit';

export async function load({ fetch, parent }) {
	const { userId } = await parent();

	const npcFetch = await fetch('http://localhost:3000/api/guess-the-npc/search-items');
	if (!npcFetch) {
		error(404, 'Unable to fetch weapons');
	}
	let npcs: SimpleNpc[] = await npcFetch.json();

	const initData = await fetch(`http://localhost:3000/api/guess-the-npc/initialize-game/${userId}`);
	if (!initData) {
		error(404, 'Unable to fetch initializing data');
	}
	const initDataJson = await initData.json();

	const quote = initDataJson.quote;
	const guesses: SimpleNpc[] = initDataJson.guesses ?? [];
	const won = initDataJson.won;

	const guessIds: number[] = guesses?.map((g) => g.id) ?? [];
	npcs = npcs.filter((n) => !guessIds.includes(n.id));

	return { npcs, quote, guesses, won };
}
