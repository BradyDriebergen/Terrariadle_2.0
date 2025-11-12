// Accept SvelteKit's fetch so SSR/client/adapter all work the same.
export async function initializeGame(
  fetch: typeof globalThis.fetch,
  gameMode: string,
  userId: string
) {
  const res = await fetch(`/api/${gameMode}/initialize-game/${encodeURIComponent(userId)}`, {
    headers: { accept: 'application/json' }
  });

  if (!res.ok) {
    // Optional: surface server error text for easier debugging
    const body = await res.text().catch(() => '');
    throw new Error(`Initialize game failed: ${res.status} ${body}`);
  }

  return res.json(); // <-- returns parsed JSON
}