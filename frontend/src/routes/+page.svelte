<script lang="ts">
	let originalURL = '';
	let shortURL = '';
	let err = false;

	function shortenHandler() {
		fetch('https://shorterurl.bukharney.tech/shorten', {
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
				console.log(data);
				shortURL = data.url;
			});
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

<div class="bg-white">
	<div class="mx-auto max-w-7xl py-24 sm:px-6 sm:py-32 lg:px-8">
		<div
			class="relative isolate overflow-hidden bg-gray-900 px-6 py-6 shadow-2xl sm:rounded-3xl sm:px-16 lg:flex lg:px-24"
		>
			<svg
				viewBox="0 0 1024 1024"
				class="absolute left-1/2 top-1/2 -z-10 h-[64rem] w-[64rem] -translate-y-1/2 [mask-image:radial-gradient(closest-side,white,transparent)] sm:left-full sm:-ml-80 lg:left-1/2 lg:ml-0 lg:-translate-x-1/2 lg:translate-y-0"
				aria-hidden="true"
			>
				<circle
					cx="512"
					cy="512"
					r="512"
					fill="url(#759c1415-0410-454c-8f7c-9a820de03641)"
					fill-opacity="0.7"
				/>
				<defs>
					<radialGradient id="759c1415-0410-454c-8f7c-9a820de03641">
						<stop stop-color="#7775D6" />
						<stop offset="1" stop-color="#E935C1" />
					</radialGradient>
				</defs>
			</svg>
			<div class="relative h-full w-full text-white">
				<div class="grid-cols-flow grid-rows-flow grid gap-2 py-2">
					<label for="url">URL</label>
					<input
						class="rounded-md {err
							? 'border-red-600 border-2'
							: ''} bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-sm hover:bg-gray-100 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
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
						class="rounded-md bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-sm hover:bg-gray-100 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
						on:click={() => {
							if (originalURL !== '' && isValidURL(originalURL)) {
								err = false;
							} else {
								err = true;
							}
						}}
						on:click={err ? () => {} : shortenHandler}
						disabled={err}>Shorten</button
					>
				</div>

				{#if shortURL !== ''}
					<div class="grid-cols-flow grid-rows-flow grid gap-2 py-2">
						<div
							class="text-white block w-full rounded-md border-0 py-1.5 pl-7 pr-20 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
						>
							{'https://shorterurl.bukharney.tech/' + shortURL}
						</div>
						<button
							class="rounded-md bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-sm hover:bg-gray-100 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
							on:click={() => {
								window.location.href = 'https://shorterurl.bukharney.tech/' + shortURL;
							}}>Go to this link</button
						>
						<button
							class="rounded-md bg-white px-3.5 py-2.5 text-sm font-semibold text-gray-900 shadow-sm hover:bg-gray-100 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-white"
							on:click={() => {
								navigator.clipboard.writeText('https://shorterurl.bukharney.tech/' + shortURL);
							}}>Copy</button
						>
					</div>
				{/if}
			</div>
		</div>
	</div>
</div>

<style>
</style>
