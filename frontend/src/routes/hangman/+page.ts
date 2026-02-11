import { error } from '@sveltejs/kit';

export async function load({ fetch, parent }) {
	const { userId } = await parent();

	const initData = await fetch(`http://localhost:3000/api/hangman/initialize-game/${userId}`);
	if (!initData) {
		error(404, 'Unable to fetch initializing data');
	}
	const initDataJson = await initData.json();

	const phrase = initDataJson.phrase;
	const guessedLetters = initDataJson.guessedLetters;
	const attempts = initDataJson.attempts;
	const finished = initDataJson.finished;

	return { phrase, guessedLetters, attempts, finished };
}
