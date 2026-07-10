export async function parseJsonSafe(res: Response): Promise<any> {
	try {
		return await res.json();
	} catch {
		return null;
	}
}
