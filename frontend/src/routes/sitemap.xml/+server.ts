import type { RequestHandler } from './$types';

export const prerender = true;

const routes = [
	'',
	'/daily-slash',
	'/connections',
	'/guess-the-npc',
	'/hangman',
	'/terratrivia',
	'/about'
];

export const GET: RequestHandler = async () => {
	const body = `
        <?xml version="1.0" encoding="UTF-8"?>
        <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
        ${routes.map((r) => `  <url><loc>https://terrariadle.com${r}</loc></url>`).join('\n')}
        </urlset>`;

	return new Response(body, {
		headers: { 'Content-Type': 'application/xml' }
	});
};
