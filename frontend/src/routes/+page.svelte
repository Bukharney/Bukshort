<script lang="ts">
	import { toast } from '@zerodevx/svelte-toast';
	const API_URL = process.env.API_URL;
	let originalURL = '';
	let shortURL = '';
	let err = false;
	let loading = false;

	async function shortenHandler() {
		loading = true;
		await fetch(API_URL + 'shorten', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				url: originalURL
			})
		})
			.then((res) => res.json())
			.then((data) => {
				shortURL = data.url;
			});
		loading = false;
	}

	function isValidURL(string: string) {
		try {
			new URL(string);
		} catch (_) {
			return false;
		}

		return true;
	}

	function copyToClipboard() {
		navigator.clipboard.writeText(API_URL + shortURL);
		toast.push('Copied to clipboard!', {
			theme: {
				'--toastBackground': '#4F46E5',
				'--toastProgressBackground': '#C7D2FE',
				'--toastProgressFill': '#4F46E5',
				'--toastColor': '#fff'
			}
		});
	}
</script>

{#if loading}
	<script>
		document.title = 'Loading...';
	</script>
{:else}
	<script>
		document.title = 'BukShort';
	</script>
{/if}

<div class=" bg-gray-50 h-screen">
	<header class="">
		<nav
			class="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8"
			aria-label="Global"
		>
			<div class="flex lg:flex-1">
				<a class="-m-1.5 p-1.5" href="#top">
					<span class="sr-only">Your Company</span>
					<img class="h-8 w-auto" src="/letter-b.png" alt="" />
				</a>
			</div>
			<div class="flex items-center">
				<a
					href="https://github.com/Bukharney/shorterURL"
					class="bg-white inline rounded-md shadow-md hover:scale-105 py-1 px-2 text-sm font-medium text-gray-700"
				>
					<img
						src="https://pngimg.com/uploads/github/github_PNG32.png"
						alt="github"
						class="h-8 w-8 inline"
					/>
					GitHub
				</a>
			</div>
		</nav>
	</header>

	<body>
		<div class="relative">
			<div class="max-w-7xl mx-12 xl:mx-auto">
				<div
					class="bg-white relative isolate overflow-hidden px-6 py-6 my-6 shadow-xl rounded-3xl sm:px-16 lg:flex lg:px-24"
				>
					<div class="relative h-full w-full text-white">
						<div class="grid-cols-flow grid-rows-flow grid gap-2 py-2 text-black">
							<label class=" font-bold" for="url">URL</label>
							<input
								class="rounded-md {err
									? 'border-red-600 focus-visible:outline-red-600'
									: 'focus-visible:outline-gray-600'}  border bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-md focus-visible:outline"
								id="url"
								type="text"
								placeholder="https://example.com"
								bind:value={originalURL}
								on:change={() => {
									if (originalURL !== '' && isValidURL(originalURL)) {
										err = false;
									} else {
										err = true;
									}
								}}
							/>
							<p class="text-red-600 text-sm">{err ? 'Invalid URL' : ''}</p>
							<button
								class=" {loading
									? 'cursor-wait'
									: 'cursor-pointer hover:border-indigo-600 hover:bg-indigo-600 hover:text-white'} rounded-md border bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-lg"
								on:click={() => {
									if (originalURL !== '' && isValidURL(originalURL)) {
										err = false;
									} else {
										err = true;
									}
									if (!err && !loading) {
										shortenHandler();
									}
								}}>Shorten</button
							>
						</div>
					</div>
				</div>
				{#if shortURL !== '' && !loading}
					<div
						class="bg-white relative isolate overflow-hidden px-6 py-6 my-6 shadow-xl rounded-3xl sm:px-16 lg:flex lg:px-24"
					>
						<div class="relative h-full w-full text-white">
							<div class="grid-cols-flow grid-rows-flow grid gap-2 py-2">
								<div
									class="pointer-events-none rounded-md border bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-md focus-visible:outline"
								>
									{API_URL + shortURL}
								</div>
								<button
									class="rounded-md border bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-md hover:border-indigo-600 hover:bg-indigo-600 hover:text-white"
									on:click={copyToClipboard}>Copy</button
								>
								<button
									class="rounded-md border bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-md hover:border-indigo-600 hover:bg-indigo-600 hover:text-white"
									on:click={() => {
										window.location.href = API_URL + shortURL;
									}}>Go to this link</button
								>
							</div>
						</div>
					</div>
				{:else if loading}
					<div
						class="bg-white relative isolate overflow-hidden px-6 py-6 my-6 shadow-xl rounded-3xl sm:px-16 lg:flex lg:px-24"
					>
						<div class="flex items-center justify-center h-32 w-full">
							<div
								class="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-current border-r-transparent align-[-0.125em] motion-reduce:animate-[spin_1.5s_linear_infinite]"
								role="status"
							>
								<span
									class="!absolute !-m-px !h-px !w-px !overflow-hidden !whitespace-nowrap !border-0 !p-0 ![clip:rect(0,0,0,0)]"
									>Loading...</span
								>
							</div>
						</div>
					</div>
				{/if}
			</div>
		</div>
	</body>
</div>

<style>
	:root {
		--toastContainerTop: auto;
		--toastContainerRight: auto;
		--toastContainerBottom: 1rem;
		--toastContainerLeft: 1rem;
	}
</style>
