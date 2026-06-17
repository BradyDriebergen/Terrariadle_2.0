import { ApiError } from "$lib/types/error";

export function getUserId(): string {
  const userId = localStorage.getItem('user_id');
  if (!userId) {
    throw new ApiError(401, 'Session not found. Try refreshing the page.');
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