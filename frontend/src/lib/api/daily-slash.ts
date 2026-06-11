import type { WeaponGuess } from '$lib/types/daily-slash';
import { error } from '@sveltejs/kit';

export async function getWeaponHint(num: number): Promise<string> {
	if (num < 1 || num > 3) {
		throw new Error('Invalid hint number input');
	}

	const res = await fetch(`/api/daily-slash/hint/?hint=${num}`);
	return (await res.json()) as string;
}

export async function checkWeaponGuess(weaponId: number): Promise<WeaponGuess> {
	const userId = localStorage.getItem('userId');

	if (!userId) {
		throw new Error('Session not found. Try refreshing the page.');
	}

	const res = await fetch('/api/daily-slash/check-guess', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ user_id: userId, guess: weaponId })
	});

	const body = await res.json();
	if (!res.ok) {
		error(res.status, body.error ?? 'An error occurred when checking guess');
	}

	return body as WeaponGuess;
}
