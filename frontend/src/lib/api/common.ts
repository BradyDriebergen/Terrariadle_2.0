export function subscribeToPlayerCount(mode: string, onCount: (count: number) => void): () => void {
	const es = new EventSource(`/api/guess-count?mode=${mode}`);

	es.onmessage = (event) => {
		const data = JSON.parse(event.data);
		onCount(data.count);
	};

	es.onerror = (err) => {
		console.error('SSE error:', err);
		if (es.readyState === EventSource.CLOSED) {
			console.error('SSE permanently closed');
		}
	};

	return () => es.close();
}
