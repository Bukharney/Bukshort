import type { RequestHandler } from './$types';
import { redirect } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ params, fetch }) => {
	const backendHost = process.env.API_URL || 'http://backend:8080/';
	const target = new URL(`/${params.shortURL}`, backendHost.endsWith('/') ? backendHost : backendHost + '/').toString();

	try {
		const res = await fetch(target, { redirect: 'manual' });
		
		// If backend returns redirect 301 or 302, follow location
		if (res.status === 302 || res.status === 301) {
			const location = res.headers.get('location');
			if (location) {
				throw redirect(res.status, location);
			}
		}

		const data = await res.text();
		return new Response(data, {
			status: res.status,
			headers: {
				'Content-Type': res.headers.get('content-type') || 'text/plain'
			}
		});
	} catch (err: any) {
		// Re-throw SvelteKit redirect exceptions
		if (err && typeof err === 'object' && 'status' in err && 'location' in err) {
			throw err;
		}
		return new Response('URL not found or backend unreachable', { status: 404 });
	}
};
