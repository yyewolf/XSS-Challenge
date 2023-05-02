<script lang="ts">
	// The ordering of these imports is critical to your app working properly
	import '@skeletonlabs/skeleton/themes/theme-skeleton.css';
	// If you have source.organizeImports set to true in VSCode, then it will auto change this ordering
	import '@skeletonlabs/skeleton/styles/all.css';
	// Most of your app wide CSS should be put in this file
	import '../app.postcss';

	import api from '$lib/api.js';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	import { AppBar } from '@skeletonlabs/skeleton';

	let connected = false;

	onMount(() => {
		api
			.get('/api/connected')
			.then((res) => {
				connected = true;
			})
			.catch((err) => {
				//
			})
			.finally(() => {
				let path = $page.route.id;
				if (!connected) {
					if (path === '/panel' || path === '/reports' || path === '/readreport') {
						window.location.href = '/';
					}
				}
			});
	});
</script>

<!-- Navbar -->
<AppBar>
	<svelte:fragment slot="lead">
		<a href="/">
			<img src="/android-chrome-192x192.png" alt="logo" class="h-8" />
		</a>
	</svelte:fragment>
	<a class="text-xl font-semibold hover:text-gray-200" href="/">Iks Esse Esse</a>
	<svelte:fragment slot="trail">
		{#if connected}
			<a class="text-xl font-semibold hover:text-gray-200" href="/panel">Panel</a>
			<a class="text-xl font-semibold hover:text-gray-200" href="/reports">Reports</a>
		{:else}
			<a class="text-xl font-semibold hover:text-gray-200" href="/report">Report</a>
			<a class="text-xl font-semibold hover:text-gray-200" href="/login">Login</a>
		{/if}
	</svelte:fragment>
</AppBar>

<slot />
