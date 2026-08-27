import type { RequestHandler } from './$types';

export const POST: RequestHandler = async ({ request, fetch }) => {
	const backendHost = process.env.API_URL || 'http://backend:8080/';
	const target = new URL('/shorten', backendHost.endsWith('/') ? backendHost : backendHost + '/').toString();

	try {
		const body = await request.text();
		const res = await fetch(target, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body
		});

		const data = await res.text();
		return new Response(data, {
			status: res.status,
			headers: {
				'Content-Type': 'application/json'
			}
		});
	} catch (err: any) {
		return new Response(JSON.stringify({ error: 'Backend unreachable: ' + err.message }), {
			status: 502,
			headers: { 'Content-Type': 'application/json' }
		});
	}
};
