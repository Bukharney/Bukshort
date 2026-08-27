<script lang="ts">
	import { toast } from '@zerodevx/svelte-toast';
	import QRCode from 'qrcode';

	const API_URL = process.env.API_URL || '';
	let originalURL = '';
	let shortURL = '';
	let err = false;
	let loading = false;
	let copied = false;
	let showQR = false;
	let qrDataUrl = '';

	async function shortenHandler() {
		if (!originalURL || !isValidURL(originalURL)) {
			err = true;
			return;
		}

		err = false;
		loading = true;
		copied = false;
		showQR = false;
		qrDataUrl = '';

		try {
			const endpoint = API_URL ? (API_URL.endsWith('/') ? API_URL + 'shorten' : API_URL + '/shorten') : '/shorten';
			const res = await fetch(endpoint, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					url: originalURL
				})
			});
			const data = await res.json();
			if (data && data.url) {
				shortURL = data.url;
				generateQRCode(getFullURL(data.url));
			} else {
				toast.push('Failed to shorten URL. Please try again.', {
					theme: {
						'--toastBackground': '#EF4444',
						'--toastColor': '#fff'
					}
				});
			}
		} catch (e) {
			toast.push('Network error. Please check your backend connection.', {
				theme: {
					'--toastBackground': '#EF4444',
					'--toastColor': '#fff'
				}
			});
		} finally {
			loading = false;
		}
	}

	function isValidURL(string: string) {
		try {
			const url = new URL(string);
			return url.protocol === 'http:' || url.protocol === 'https:';
		} catch (_) {
			return false;
		}
	}

	function getFullURL(path: string) {
		if (!path) return '';
		if (path.startsWith('http://') || path.startsWith('https://')) return path;
		return API_URL ? `${API_URL}${path}` : `${window.location.origin}/${path}`;
	}

	async function generateQRCode(text: string) {
		try {
			qrDataUrl = await QRCode.toDataURL(text, {
				width: 240,
				margin: 2,
				color: {
					dark: '#0F172A',
					light: '#FFFFFF'
				}
			});
		} catch (err) {
			console.error('Failed to generate QR code', err);
		}
	}

	function copyToClipboard() {
		const full = getFullURL(shortURL);
		navigator.clipboard.writeText(full);
		copied = true;
		toast.push('Copied to clipboard!', {
			theme: {
				'--toastBackground': '#18181B',
				'--toastProgressBackground': '#3F3F46',
				'--toastProgressFill': '#6366F1',
				'--toastColor': '#FFFFFF'
			}
		});

		setTimeout(() => {
			copied = false;
		}, 2500);
	}

	function handleKeydown(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			shortenHandler();
		}
	}
</script>

<svelte:head>
	<title>{loading ? 'Shortening...' : 'BukShort — Minimal URL Shortener'}</title>
</svelte:head>

<div
	class="min-h-screen bg-mesh-pattern flex flex-col justify-between selection:bg-indigo-100 selection:text-indigo-900"
>
	<!-- Navbar Header -->
	<header class="w-full border-b border-zinc-200/60 bg-white/70 backdrop-blur-md sticky top-0 z-50">
		<nav class="mx-auto flex max-w-5xl items-center justify-between px-6 py-4">
			<a href="/" class="flex items-center gap-2.5 group">
				<div
					class="h-9 w-9 rounded-xl bg-zinc-900 flex items-center justify-center text-white font-bold text-lg shadow-sm group-hover:scale-105 transition-transform duration-200"
				>
					B
				</div>
				<span class="font-bold text-zinc-900 text-lg tracking-tight">
					Buk<span class="text-indigo-600">Short</span>
				</span>
			</a>

			<a
				href="https://github.com/Bukharney/shorterURL"
				target="_blank"
				rel="noreferrer"
				class="inline-flex items-center gap-2 rounded-full border border-zinc-200 bg-white px-3.5 py-1.5 text-xs font-semibold text-zinc-700 shadow-sm hover:border-zinc-300 hover:bg-zinc-50 hover:text-zinc-900 transition-all"
			>
				<svg class="h-4 w-4 fill-current text-zinc-800" viewBox="0 0 24 24">
					<path
						d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"
					/>
				</svg>
				<span>GitHub</span>
			</a>
		</nav>
	</header>

	<!-- Main Content Hero Container -->
	<main class="flex-1 flex items-center justify-center p-6 py-12 md:py-20">
		<div class="w-full max-w-2xl mx-auto space-y-8">
			<!-- Hero Header Title -->
			<div class="text-center space-y-3">
				<h1 class="text-4xl md:text-5xl font-extrabold text-zinc-900 tracking-tight leading-tight">
					Shorten links, <span
						class="text-transparent bg-clip-text bg-gradient-to-r from-indigo-600 to-violet-600"
						>simplify reach.</span
					>
				</h1>
				<p class="text-zinc-500 text-sm md:text-base max-w-md mx-auto">
					Transform long, complex URLs into clean, memorable links in seconds.
				</p>
			</div>

			<!-- Main Input Card -->
			<div
				class="bg-white/90 backdrop-blur-xl border border-zinc-200/90 rounded-2xl p-4 md:p-6 shadow-xl shadow-zinc-200/60 transition-all"
			>
				<div class="space-y-3">
					<label
						for="url-input"
						class="block text-xs font-bold uppercase tracking-wider text-zinc-500"
					>
						Enter Long URL
					</label>

					<div class="relative flex flex-col sm:flex-row items-stretch gap-2.5">
						<div class="relative flex-1">
							<input
								id="url-input"
								type="url"
								placeholder="https://example.com/very-long-link-to-shorten"
								bind:value={originalURL}
								on:keydown={handleKeydown}
								on:input={() => {
									if (err && isValidURL(originalURL)) err = false;
								}}
								class="w-full rounded-xl border bg-zinc-50/50 px-4 py-3.5 text-sm font-medium text-zinc-900 placeholder-zinc-400 transition-all focus:bg-white focus:outline-none focus:ring-2 ${err
									? 'border-red-400 focus:border-red-500 focus:ring-red-200'
									: 'border-zinc-200 focus:border-indigo-500 focus:ring-indigo-100'}"
							/>
						</div>

						<button
							on:click={shortenHandler}
							disabled={loading}
							class="inline-flex items-center justify-center gap-2 rounded-xl bg-zinc-900 px-6 py-3.5 text-sm font-semibold text-white shadow-md shadow-zinc-900/10 hover:bg-zinc-800 active:scale-[0.98] disabled:opacity-70 disabled:cursor-not-allowed transition-all cursor-pointer whitespace-nowrap min-w-[120px]"
						>
							{#if loading}
								<svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
									<circle
										class="opacity-25"
										cx="12"
										cy="12"
										r="10"
										stroke="currentColor"
										stroke-width="4"
									/>
									<path
										class="opacity-75"
										fill="currentColor"
										d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
									/>
								</svg>
								<span>Shortening...</span>
							{:else}
								<span>Shorten Link</span>
								<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M14 5l7 7m0 0l-7 7m7-7H3"
									/>
								</svg>
							{/if}
						</button>
					</div>

					{#if err}
						<p class="text-xs text-red-500 font-medium flex items-center gap-1">
							<svg class="h-3.5 w-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
								/>
							</svg>
							Please enter a valid URL (e.g. https://example.com)
						</p>
					{/if}
				</div>
			</div>

			<!-- Shortened Result Card -->
			{#if shortURL && !loading}
				<div
					class="bg-white/90 backdrop-blur-xl border border-zinc-200/90 rounded-2xl p-5 md:p-6 shadow-xl shadow-zinc-200/60 space-y-4 animate-fade-in-scale"
				>
					<div class="flex items-center justify-between">
						<span
							class="text-xs font-bold uppercase tracking-wider text-emerald-600 flex items-center gap-1.5"
						>
							<span class="h-2 w-2 rounded-full bg-emerald-500 animate-pulse" />
							Shortened URL Ready
						</span>

						<button
							on:click={() => (showQR = !showQR)}
							class="inline-flex items-center gap-1.5 text-xs font-semibold text-zinc-600 hover:text-indigo-600 transition-colors"
						>
							<svg class="h-4 w-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M12 4v1m6 11h2m-6 0h-2v4m0-11v3m0 0h.01M12 12h4.01M16 20h4M4 12h4m12 0h.01M5 8h2a1 1 0 001-1V5a1 1 0 00-1-1H5a1 1 0 00-1 1v2a1 1 0 001 1zm12 0h2a1 1 0 001-1V5a1 1 0 00-1-1h-2a1 1 0 00-1 1v2a1 1 0 001 1zM5 20h2a1 1 0 001-1v-2a1 1 0 00-1-1H5a1 1 0 00-1 1v2a1 1 0 001 1z"
								/>
							</svg>
							<span>{showQR ? 'Hide QR Code' : 'Show QR Code'}</span>
						</button>
					</div>

					<div class="flex flex-col sm:flex-row items-stretch sm:items-center gap-3">
						<!-- Displayed Short Link -->
						<div
							class="flex-1 bg-zinc-50 border border-zinc-200 rounded-xl px-4 py-3 text-sm font-semibold text-zinc-900 select-all truncate"
						>
							{getFullURL(shortURL)}
						</div>

						<!-- Action Buttons -->
						<div class="flex items-center gap-2">
							<button
								on:click={copyToClipboard}
								class="flex-1 sm:flex-none inline-flex items-center justify-center gap-1.5 rounded-xl border border-zinc-200 bg-white px-4 py-3 text-sm font-semibold text-zinc-800 shadow-sm hover:bg-zinc-50 hover:border-zinc-300 transition-all ${copied
									? '!bg-indigo-50 !border-indigo-200 !text-indigo-700'
									: ''}"
							>
								{#if copied}
									<svg
										class="h-4 w-4 text-indigo-600"
										fill="none"
										stroke="currentColor"
										viewBox="0 0 24 24"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											stroke-width="2.5"
											d="M5 13l4 4L19 7"
										/>
									</svg>
									<span>Copied!</span>
								{:else}
									<svg
										class="h-4 w-4 text-zinc-500"
										fill="none"
										stroke="currentColor"
										viewBox="0 0 24 24"
									>
										<path
											stroke-linecap="round"
											stroke-linejoin="round"
											stroke-width="2"
											d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
										/>
									</svg>
									<span>Copy</span>
								{/if}
							</button>

							<a
								href={getFullURL(shortURL)}
								target="_blank"
								rel="noreferrer"
								class="flex-1 sm:flex-none inline-flex items-center justify-center gap-1.5 rounded-xl bg-zinc-900 px-4 py-3 text-sm font-semibold text-white shadow-sm hover:bg-zinc-800 transition-all"
							>
								<span>Open</span>
								<svg
									class="h-4 w-4 opacity-70"
									fill="none"
									stroke="currentColor"
									viewBox="0 0 24 24"
								>
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
									/>
								</svg>
							</a>
						</div>
					</div>

					<!-- QR Code Collapsible Section -->
					{#if showQR && qrDataUrl}
						<div
							class="pt-3 border-t border-zinc-100 flex flex-col items-center justify-center space-y-3 animate-fade-in-scale"
						>
							<div class="p-3 bg-white border border-zinc-200 rounded-2xl shadow-sm">
								<img src={qrDataUrl} alt="QR Code for Short URL" class="w-44 h-44 rounded-lg" />
							</div>
							<a
								href={qrDataUrl}
								download="bukshort-qr.png"
								class="text-xs font-semibold text-indigo-600 hover:text-indigo-800 hover:underline inline-flex items-center gap-1"
							>
								<svg class="h-3.5 w-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
									<path
										stroke-linecap="round"
										stroke-linejoin="round"
										stroke-width="2"
										d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"
									/>
								</svg>
								Download QR Image
							</a>
						</div>
					{/if}
				</div>
			{/if}
		</div>
	</main>

	<!-- Footer -->
	<footer class="w-full border-t border-zinc-200/60 py-6 text-center text-xs text-zinc-400">
		<p>BukShort &copy; {new Date().getFullYear()} — Built with SvelteKit &amp; Go</p>
	</footer>
</div>

<style>
	:root {
		--toastContainerTop: auto;
		--toastContainerRight: auto;
		--toastContainerBottom: 1.5rem;
		--toastContainerLeft: 50%;
		--toastTransform: translateX(-50%);
	}
</style>
