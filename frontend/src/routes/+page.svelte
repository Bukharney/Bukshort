<script lang="ts">
	let originalURL = '';
	let shortURL = '';
	let err = false;
	let loading = false;

	async function shortenHandler() {
		loading = true;
		await fetch('https://shorterurl.bukharney.tech/shorten', {
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
</script>

{#if loading}
	<script>
		document.title = 'Loading...';
	</script>
{:else}
	<script>
		document.title = 'ShorterURL';
	</script>
{/if}

<header class="bg-inherit">
	<nav class="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8" aria-label="Global">
		<div class="flex lg:flex-1">
			<a class="-m-1.5 p-1.5" href="#top">
				<span class="sr-only">Your Company</span>
				<img class="h-8 w-auto" src="/letter-b.png" alt="" />
			</a>
		</div>
		<div class="flex items-center">
			<a
				href="https://github.com/Bukharney/shorterURL"
				class="inline rounded-md hover:shadow-md py-1 px-2 border border-gray-300 text-sm font-medium text-gray-700"
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
	<div class="relative bg-white">
		<div class="max-w-7xl mx-12 xl:mx-auto">
			<div
				class="relative isolate overflow-hidden px-6 py-6 my-6 shadow-xl rounded-3xl sm:px-16 lg:flex lg:px-24"
			>
				<div class="relative h-full w-full text-white">
					<div class="grid-cols-flow grid-rows-flow grid gap-2 py-2 text-black">
						<label for="url">URL</label>
						<input
							class="rounded-md {err
								? 'border-red-600 '
								: 'border-black '}  border bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-sm hover:bg-gray-100 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
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
								: 'cursor-pointer'} rounded-md border bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-sm hover:bg-gray-100 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
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
					class="relative isolate overflow-hidden px-6 py-6 my-6 shadow-xl rounded-3xl sm:px-16 lg:flex lg:px-24"
				>
					<div class="relative h-full w-full text-white">
						<div class="grid-cols-flow grid-rows-flow grid gap-2 py-2">
							<div
								class=" pointer-events-none text-gray-900 block w-full rounded-md border-0 py-1.5 pl-7 pr-20 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
							>
								{'https://shorterurl.bukharney.tech/' + shortURL}
							</div>
							<button
								class="rounded-md border bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-sm hover:bg-gray-100 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
								on:click={() => {
									navigator.clipboard.writeText('https://shorterurl.bukharney.tech/' + shortURL);
								}}>Copy</button
							>
							<button
								class="rounded-md border bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-sm hover:border-red-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
								on:click={() => {
									window.location.href = 'https://shorterurl.bukharney.tech/' + shortURL;
								}}>Go to this link</button
							>
						</div>
					</div>
				</div>
			{:else if loading}
				<div
					class="relative isolate overflow-hidden px-6 py-6 my-6 shadow-xl rounded-3xl sm:px-16 lg:flex lg:px-24"
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

<style>
</style>
