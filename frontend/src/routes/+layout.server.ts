// Load method that creates a user id cookie for the user
export function load({ cookies }) {
	let userId = cookies.get('user_id');

	if (!userId) {
		userId = crypto.randomUUID();
		cookies.set('user_id', userId, {
			path: '/',
			httpOnly: true,
			sameSite: 'strict',
			maxAge: 60 * 60 * 24 * 365 // 1 year
		});
	}

	return {
		userId
	};
}
