import { error } from '@sveltejs/kit';

export async function load({ fetch, parent }) {
	const { userId } = await parent();

	const initData = await fetch(`http://localhost:3000/api/connections/initialize-game/${userId}`);
	if (!initData) {
		error(404, 'Unable to fetch initializing data');
	}
	const initDataJson = await initData.json();

	const attempts = initDataJson.attempts;
	const options = initDataJson.options;
	const guesses = initDataJson.guesses;
	const finished = initDataJson.finished;

	return { attempts, options, guesses, finished };
}
