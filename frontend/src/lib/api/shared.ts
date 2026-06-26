export function getOrCreateUserId(): string {
	let userId = localStorage.getItem('user_id');

	if (!userId) {
		userId = crypto.randomUUID();
		localStorage.setItem('user_id', userId);
	}

	return userId;
}

export async function parseJsonSafe(res: Response): Promise<any> {
	try {
		return await res.json();
	} catch {
		return null;
	}
}
