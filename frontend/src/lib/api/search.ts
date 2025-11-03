const API_BASE = 'http://localhost:3000'; // Or from $env/static/public

export async function searchApi(gameMode: string, query: string, signal?: AbortSignal) {
	if (!query || query.trim().length <= 0) return [];

	const res = await fetch(`${API_BASE}/api/${gameMode}/search?q=${encodeURIComponent(query)}`, {
		signal,
		headers: { Accept: 'application/json' }
	});

	if (!res.ok) throw new Error(`HTTP ${res.status}`);
	return res.json();
}