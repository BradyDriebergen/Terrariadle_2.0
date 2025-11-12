import { redirect, error } from '@sveltejs/kit';
import { initializeGame } from '$lib/api/initializeGame';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ cookies, fetch }) => {
  const userId = cookies.get('user_id');
  if (!userId) {
    // pick one:
    throw redirect(302, '/login');
    // or: throw error(401, 'User not logged in');
  }

  // initializeGame already returns JSON, so no second .json()
  const userData = await initializeGame(fetch, 'daily-slash', userId);

  return { userData };
};