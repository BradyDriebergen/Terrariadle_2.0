export async function parseJsonSafe<T = unknown>(res: Response): Promise<T> {
	try {
		return (await res.json()) as T;
	} catch {
		return {} as T;
	}
}
