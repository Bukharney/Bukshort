import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vitest/config';

export default defineConfig({
	plugins: [sveltekit()],
	define: {
		'process.env.API_URL': JSON.stringify(process.env.API_URL || '')
	},
	test: {
		include: ['src/**/*.{test,spec}.{js,ts}']
	}
});
