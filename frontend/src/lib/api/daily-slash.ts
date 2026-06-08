import { error } from "@sveltejs/kit";

export async function checkWeaponGuess(weaponId: number) {
    const userId = localStorage.getItem('userId');

    const res = await fetch('/api/daily-slash/check-guess', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ userId, guess: weaponId })
    });

    if (!res.ok) {
        const body = await res.json();
        error(res.status, body.error ?? 'Unable to initialize game');
    }
    
    return res.json();
}