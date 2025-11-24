export function debounceWithHooks<T extends (...args: any[]) => void>(
	fn: T,
	wait = 300,
	hooks?: { onStart?: () => void; onCancel?: () => void }
) {
	let t: ReturnType<typeof setTimeout> | null = null;

	const debounced = ((...args: Parameters<T>) => {
		hooks?.onStart?.();
		if (t) clearTimeout(t);
		t = setTimeout(() => {
			fn(...args);
			t = null;
		}, wait);
	}) as T & { cancel: () => void };

	debounced.cancel = () => {
		if (t) clearTimeout(t);
		t = null;
		hooks?.onCancel?.();
	};

	return debounced;
}
