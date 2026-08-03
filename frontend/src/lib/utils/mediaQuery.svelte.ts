export function useMediaQuery(query: string) {
	let matches = $state(false);

	$effect(() => {
		const mql = window.matchMedia(query);
		matches = mql.matches;

		const handler = (e: MediaQueryListEvent) => (matches = e.matches);
		mql.addEventListener('change', handler);

		return () => mql.removeEventListener('change', handler);
	});

	return {
		get current(): boolean {
			return matches;
		}
	};
}
