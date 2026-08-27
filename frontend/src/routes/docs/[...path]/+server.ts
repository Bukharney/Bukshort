import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ params, fetch }) => {
	const backendHost = process.env.API_URL || 'http://backend:8080/';
	const subpath = params.path || 'index.html';
	const target = new URL(`/docs/${subpath}`, backendHost.endsWith('/') ? backendHost : backendHost + '/').toString();

	try {
		const res = await fetch(target);
		const contentType = res.headers.get('content-type') || 'text/plain';
		const body = await res.arrayBuffer();

		return new Response(body, {
			status: res.status,
			headers: {
				'Content-Type': contentType
			}
		});
	} catch (err: any) {
		return new Response('Swagger docs unreachable', { status: 502 });
	}
};
