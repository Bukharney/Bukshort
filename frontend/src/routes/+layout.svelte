<script>
	import '../app.postcss';
	import { SvelteToast } from '@zerodevx/svelte-toast';
</script>

<slot />
<SvelteToast options={{ reversed: true, intro: { y: 192 } }} />
