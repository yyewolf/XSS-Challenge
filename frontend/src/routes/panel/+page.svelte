<script>
	import api from '$lib/api.js';
	import { onMount } from 'svelte';

    let flag = "";

	onMount(() => {
		api
			.get('/api/flag')
			.then((res) => {
				flag = res.data.flag;
			})
			.catch((err) => {
				//
			});
	});
</script>

<div class="flex flex-col justify-center items-center min-h-screen">
	<h1 class="text-3xl font-bold text-white mb-4">{flag}</h1>
</div>
